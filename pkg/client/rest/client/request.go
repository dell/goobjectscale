// Copyright © 2023 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// Request is an ObjectScale API request wrapper.
type Request struct {
	// Method is the method of REST API request
	Method string

	// Path is the path of REST API request
	Path string

	// Body is the body of REST API request
	Body interface{}

	// ContentType is the body of REST API request
	ContentType string

	// Params are the parameters of the REST API request
	Params map[string]string
}

// HTTPWithContext converts the Request data into an http.Request object.
func (r Request) HTTPWithContext(ctx context.Context, uri string) (*http.Request, error) {
	var (
		obj []byte
		err error
		q   = url.Values{}
	)

	switch r.ContentType {
	case ContentTypeXML:
		obj, err = xml.Marshal(r.Body)
		if err != nil {
			return nil, fmt.Errorf("marshal xml: %w", err)
		}
	case ContentTypeJSON:
		if raw, ok := r.Body.(json.RawMessage); ok {
			obj, err = raw.MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("marshal raw json: %w", err)
			}
		} else {
			obj, err = json.Marshal(r.Body)
			if err != nil {
				return nil, fmt.Errorf("marshal json: %w", err)
			}
		}
	default:
		return nil, fmt.Errorf("request: %s: %w", r.ContentType, ErrContentType)
	}

	u, _ := url.Parse(uri)
	u.Path = r.Path

	for key, value := range r.Params {
		q.Add(key, value)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, r.Method, u.String(), bytes.NewBuffer(obj))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	return req, nil
}

// Validate performs lightweight validation and can be used to detect misconfigured
// requests.
func (r Request) Validate(uri string) error {
	switch r.ContentType {
	case ContentTypeJSON, ContentTypeXML:
	default:
		return fmt.Errorf("request: %s: %w", r.ContentType, ErrContentType)
	}

	_, err := url.Parse(uri)
	if err != nil {
		return fmt.Errorf("parse uri: %s: %w", uri, err)
	}

	return nil
}
