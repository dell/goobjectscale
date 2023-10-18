// Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/go-logr/logr"
)

var _ RemoteCaller = (*Simple)(nil) // interface guard

// Allowed content types.
const (
	ContentTypeXML  = "application/xml"
	ContentTypeJSON = "application/json"
)

// Simple is the simple client definition.
type Simple struct {
	// Endpoint is the URL of the management API
	Endpoint string `json:"endpoint"`

	// Authenticator!=nil means Authenticator.Login will be called to
	// obtain login credentials.
	Authenticator Authenticator

	// Should X-EMC-Override header be added into the request
	OverrideHeader bool

	HTTPClient *http.Client

	log logr.Logger
}

// SetLogger sets the logger in the Simple client.
func (s *Simple) SetLogger(log logr.Logger) {
	s.log = log
}

// MakeRemoteCall executes an API request against the client endpoint, returning
// the object body of the response into a response object.
func (s *Simple) MakeRemoteCall(ctx context.Context, r Request, into interface{}) error {
	err := r.Validate(s.Endpoint)
	if err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	// Do performs a single http request.
	Do := func(ctx context.Context) error {
		req, err := s.buildHTTPRequest(ctx, r)
		if err != nil {
			return err
		}

		s.log.V(8).Info("Request prepared.", //nolint:gomnd
			"Header", req.Header,
			"URL", req.URL,
		)

		resp, err := s.HTTPClient.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		s.log.V(8).Info("Response obtained.", //nolint:gomnd
			"ContentLength", resp.ContentLength,
			"Header", req.Header,
			"URL", req.URL,
			"StatusCode", resp.StatusCode,
			"Status", resp.Status,
			"URL", resp.Header,
		)

		// not checking error, as this is a best-effort call, and we do not care about the error here.
		body, _ := io.ReadAll(resp.Body)
		s.log.V(10).Info("Response obtained (extra info).", //nolint:gomnd
			"Body", string(body),
		)

		err = validateResponse(r, resp)
		if err != nil {
			return err
		}

		if into != nil {
			if err := unmarshal(r, resp, into); err != nil {
				return err
			}
		}

		return nil
	}

	// If Authenticator is nil then just perform a single request; otherwise
	// perform AuthRetriesMax requests but only if returned error is an authorization
	// error.
	if s.Authenticator == nil {
		return Do(ctx)
	}

	if !s.Authenticator.IsAuthenticated() {
		if err := s.Authenticator.Login(ctx, s.HTTPClient); err != nil {
			return fmt.Errorf("%w: login: %s", ErrAuthorization, err.Error())
		}
	}

	for tries := 0; tries < AuthRetriesMax; tries++ {
		err := Do(ctx)

		switch {
		case errors.Is(err, ErrAuthorization):
			if err = s.Authenticator.Login(ctx, s.HTTPClient); err != nil {
				// TODO Depending on how the error is constructed we could potentially
				//      leak credentials here.  Must be careful.
				return fmt.Errorf("%w: retry login", err)
			}

			continue
		default:
			return err
		}
	}

	return fmt.Errorf("%w: exhausted authentication tries", ErrAuthorization)
}

func (s *Simple) buildHTTPRequest(ctx context.Context, r Request) (*http.Request, error) {
	req, err := r.HTTPWithContext(ctx, s.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("simple client: %w", err)
	}

	req.Header.Add("Accept", r.ContentType)
	req.Header.Add("Content-Type", r.ContentType)
	req.Header.Add("Accept", "application/xml")

	if s.Authenticator != nil {
		req.Header.Add("X-SDS-AUTH-TOKEN", s.Authenticator.Token())
	}

	if s.OverrideHeader {
		req.Header.Add("X-EMC-Override", "true")
	}

	return req, nil
}

func validateResponse(r Request, resp *http.Response) error {
	switch {
	case resp.StatusCode == http.StatusUnauthorized:
		return ErrAuthorization
	case resp.StatusCode >= http.StatusBadRequest:
		var ecsError model.Error
		if err := unmarshal(r, resp, &ecsError); err != nil {
			return err
		}

		return ecsError
	}

	return nil
}

// Unmarshal unmarshals resp.Body into v with respect to returned Content-Type or
// original requested Content-Type.
func unmarshal(r Request, resp *http.Response, v interface{}) error {
	// Use content type sent by server first but fall back to original request content type.
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = r.ContentType
	}

	// Reading all of resp.Body into memory is inefficient, and it's better to send
	// it directly into decoders.  However if body is empty the decoders will
	// receive an EOF.  They can also receive EOF for malformed responses.
	// We need to differentiate between EOF from empty responses and malformed
	// responsed.
	var cw CountWriter
	body := io.TeeReader(resp.Body, &cw)

	// HandleError handles a decoding error.
	// If incoming error is EOF and cw.N == 0 then it's EOF due to empty response
	// and error is discarded if-and-only-if v is non-nil; i.e. it's an error
	// if we got empty body but expected a response.
	HandleError := func(err error) error {
		// v == nil check disabled because some test fixtures are missing proper response and/or some API
		// calls are written strangely -- Tenants.SetQuota for example.
		if (errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)) && cw.N == 0 /* && v == nil TODO Causes problems -- see note */ {
			return nil
		}

		return err
	}

	switch contentType {
	case ContentTypeJSON:
		decoder := json.NewDecoder(body)
		if err := HandleError(decoder.Decode(v)); err != nil {
			return fmt.Errorf("response: json: %w", err)
		}
	case ContentTypeXML:
		decoder := xml.NewDecoder(body)
		if err := HandleError(decoder.Decode(v)); err != nil {
			return fmt.Errorf("response: xml: %w", err)
		}
	default:
		return fmt.Errorf("response: %s: %w", contentType, ErrContentType)
	}

	return nil
}
