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

package client_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

// IncrCapture increments the integer value of a capture field in the map
func IncrCapture(captures map[string]interface{}, name string) {
	currentVal, ok := captures[name].(int)
	if !ok || captures[name] == nil {
		captures[name] = 0
	}
	captures[name] = currentVal + 1
}

// RoundTripFunc is a transport mock that makes a fake HTTP response locally
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip mocks an http request and returns an http response
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestRest(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"validUser":          testValidUser,
		"InvalidEndpoint":    testInvalidEndpoint,
		"InvalidContentType": testInvalidContentType,
		"FailedAuth":         testFailedAuth,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testValidUser(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, false),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "/test",
		ContentType: client.ContentTypeJSON,
	}, nil)

	require.NoError(t, err)
	assert.Equal(t, captures["login"], 1)
	assert.Equal(t, captures["test"], 1)
	assert.Nil(t, captures["notfound"])
}

func testInvalidEndpoint(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		":not:a:valid:url",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, false),
		true,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: client.ContentTypeJSON,
	}, nil)
	e := "parse \":not:a:valid:url\": missing protocol scheme"
	require.Error(t, err)
	assert.Equal(t, err.Error(), e)
	assert.Nil(t, captures["login"])
	assert.Nil(t, captures["test"])
	assert.Nil(t, captures["notfound"])
}

func testInvalidContentType(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, false),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: "NotAContentType",
	}, nil)

	assert.Equal(t, err.Error(), "invalid content-type")
	assert.Nil(t, captures["login"])
	assert.Nil(t, captures["test"])
	assert.Nil(t, captures["notfound"])
}

func testFailedAuth(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, true),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "/test",
		ContentType: client.ContentTypeJSON,
	}, nil)

	assert.Equal(t, err.Error(), "auth failure")
	assert.Equal(t, captures["login"], 1)
	assert.Nil(t, captures["test"])
	assert.Nil(t, captures["notfound"])
}

func newTestHTTPClient(captures map[string]interface{}, authFailure bool) *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		header := make(http.Header)
		switch req.URL.String() {
		case "https://testgateway/mgmt/login":
			IncrCapture(captures, "login")
			switch authFailure {
			case true:
				return &http.Response{
					StatusCode: 401,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("auth failure"))),
					Header:     header,
				}
			default:
				header.Set("X-Sds-Auth-Token", "TESTTOKEN")
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("OK"))),
					Header:     header,
				}
			}
		case "https://testserver/test":
			IncrCapture(captures, "test")
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("OK"))),
				Header:     header,
			}
		default:
			IncrCapture(captures, "notfound")
			return &http.Response{
				StatusCode: 404,
				Header:     make(http.Header),
			}
		}
	})
}
