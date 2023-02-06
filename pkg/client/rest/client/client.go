//
//
//  Copyright © 2021 - 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//       http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//
//

package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

type Client struct {
	// Endpoint is the URL of the management API
	Endpoint string `json:"endpoint"`

	// Gateway is the auth endpoint
	Gateway string `json:"gateway"`

	// Username used to authenticate management user
	Username string `json:"username"`

	// Password used to authenticate management user
	Password string `json:"password"`

	HTTPClient  *http.Client
	token       string
	authRetries int

	// Should X-EMC-Override header be added into the request
	OverrideHeader bool
}

var _ RemoteCaller = (*Client)(nil)

func NewClient(endpoint, gateway, username, password string, h *http.Client, overrideHdr bool) *Client {
	return &Client{
		Endpoint:       endpoint,
		Gateway:        gateway,
		Username:       username,
		Password:       password,
		HTTPClient:     h,
		OverrideHeader: overrideHdr,
	}
}

func (c *Client) login() error {
	u, err := url.Parse(c.Gateway)
	if err != nil {
		return err
	}
	u.Path = "/mgmt/login"
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if err = HandleResponse(resp); err != nil {
		return err
	}
	c.token = resp.Header.Get("X-SDS-AUTH-TOKEN")
	if c.token == "" {
		return fmt.Errorf("server error: login failed")
	} else {
		c.authRetries = 0
	}
	return nil
}

func (c *Client) isLoggedIn() bool {
	return c.token != ""
}

// MakeRemoteCall executes an API request against the client endpoint, returning
// the object body of the response into a response object
// NOTE: this is WET not DRY, as the same code is copied for ServiceClient
func (c *Client) MakeRemoteCall(r Request, into interface{}) error {
	var (
		obj []byte
		err error
		q   = url.Values{}
	)
	switch r.ContentType {
	case ContentTypeXML:
		obj, err = xml.Marshal(r.Body)
	case ContentTypeJSON:
		if raw, ok := r.Body.(json.RawMessage); ok {
			obj, err = raw.MarshalJSON()
		} else {
			obj, err = json.Marshal(r.Body)
		}
	default:
		return errors.New("invalid content-type")
	}
	if err != nil {
		return err
	}
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return err
	}
	u.Path = r.Path
	if r.Params != nil {
		for key, value := range r.Params {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(r.Method, u.String(), bytes.NewBuffer(obj))
	if err != nil {
		return err
	}
	if !c.isLoggedIn() {
		if err = c.login(); err != nil {
			return err
		}
	}
	req.Header.Add("Accept", r.ContentType)
	req.Header.Add("Content-Type", r.ContentType)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("X-SDS-AUTH-TOKEN", c.token)
	if c.OverrideHeader {
		req.Header.Add("X-EMC-Override", "true")
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}

	switch {
	case resp.StatusCode == http.StatusUnauthorized:
		if c.authRetries < AuthRetriesMax {
			c.authRetries += 1
			c.token = ""
			return c.MakeRemoteCall(r, into)
		}
		return errors.New(strings.ToLower(resp.Status))
	case resp.StatusCode > 399:
		ecsError := &model.Error{}
		switch r.ContentType {
		case ContentTypeXML:
			if err = xml.Unmarshal(body, ecsError); err != nil {
				return err
			}
		case ContentTypeJSON:
			if err = json.Unmarshal(body, ecsError); err != nil {
				return err
			}
		}
		return fmt.Errorf("%s: %s", ecsError.Description, ecsError.Details)
	case into == nil:
		// No errors found, and no response object defined, so just return
		// without error
		return nil
	default:
		if len(body) == 0 {
			return nil
		}
		switch r.ContentType {
		case ContentTypeXML:
			if err = xml.Unmarshal(body, into); err != nil {
				return err
			}
		case ContentTypeJSON:
			if err = json.Unmarshal(body, into); err != nil {
				return err
			}
		}
	}
	return nil
}
