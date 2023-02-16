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
	"errors"
	"fmt"

	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

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
		"NoErrorsNoObject":   testNoErrorsNoObject,
		"BadRequest":         testBadRequest,
		"EmptyBody":          testEmptyBody,
		"WithBody":           testWithBody,
		"InvalidEndpoint":    testInvalidEndpoint,
		"InvalidContentType": testInvalidContentType,
		"FailedAuth":         testFailedLogin,
		"HandleResponse":     testHandleResponse,
		"NoGateway":          testNoGateway,
		"OverriderHeader":    testOverrideHeader,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
func testInvalidEndpoint(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		":not:a:valid:url",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		true,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: client.ContentTypeJSON,
	}, nil)
	e := "parse \":not:a:valid:url\": missing protocol scheme"
	require.Error(t, err)
	require.Equal(t, err.Error(), e)
}
func testInvalidContentType(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: "NotAContentType",
	}, nil)

	require.Equal(t, err.Error(), "invalid content-type")
}

func testFailedLogin(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser1",
		"testpassword1",
		newTestHTTPClient(),
		false,
	))

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/failedLogin",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(req, nil)
	require.Error(t, err)
}

func testBadRequest(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		false,
	))
	//Content-type JSON
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/badRequest",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(req, nil)
	require.Error(t, err)
	//Content-type XML
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/badRequest",
		ContentType: client.ContentTypeXML,
	}
	nerr := clientset.Client().MakeRemoteCall(req, nil)
	require.Error(t, nerr)
}

func testEmptyBody(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		false,
	))
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/emptyBody",
		ContentType: client.ContentTypeJSON,
	}
	bucketList := &model.BucketList{}
	err := clientset.Client().MakeRemoteCall(req, bucketList)
	require.Nil(t, err)
}

func testWithBody(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		false,
	))
	//JSON
	//Success
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}
	ecsError := &model.Error{}
	err := clientset.Client().MakeRemoteCall(req, ecsError)
	require.Nil(t, err)
	//Error
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/xml",
		ContentType: client.ContentTypeJSON,
	}
	err = clientset.Client().MakeRemoteCall(req, ecsError)
	require.Equal(t, err.Error(), "invalid character '<' looking for beginning of value")
	//XML
	//Success
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/xml",
		ContentType: client.ContentTypeXML,
	}
	err = clientset.Client().MakeRemoteCall(req, ecsError)
	require.Nil(t, err)
	//Error
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeXML,
	}
	err = clientset.Client().MakeRemoteCall(req, ecsError)
	require.Equal(t, err.Error(), "EOF")
}

func testNoErrorsNoObject(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		false,
	))

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}

	//Content-type JSON
	err := clientset.Client().MakeRemoteCall(req, nil)
	require.Nil(t, err)

	//Content-type XML
	err = clientset.Client().MakeRemoteCall(req, nil)
	require.Nil(t, err)
}

func testNoGateway(t *testing.T) {
	clientset := rest.NewClientSet(&client.Client{
		Endpoint:       "https://testserver",
		Username:       "testuser",
		Password:       "testpassword",
		HTTPClient:     newTestHTTPClient(),
		OverrideHeader: false,
	})
	ecsError := &model.Error{}
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/badRequest",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(req, ecsError)
	require.Error(t, err)

}

func testOverrideHeader(t *testing.T) {
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(),
		true,
	))

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}

	//Content-type JSON
	err := clientset.Client().MakeRemoteCall(req, nil)
	require.Nil(t, err)
}

func newTestHTTPClient() *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		header := make(http.Header)
		if req.URL.String() == "https://testserver/ok/json" {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"OK"}`))),
				Header:     header,
			}
		}
		if req.URL.String() == "https://testserver/ok/xml" {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><error><Description>"OK"</Description></error>`))),
				Header:     header,
			}
		}
		if req.URL.String() == "https://testserver/failedLogin" {
			return &http.Response{
				StatusCode: 401,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"Unauthorized"}`))),
				Header:     header,
			}
		}
		if req.URL.String() == "https://testserver/badRequest" {
			return &http.Response{
				StatusCode: 403,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"Forbidden"}`))),
				Header:     header,
			}
		}
		if req.URL.String() == "https://testserver/emptyBody" {
			return &http.Response{
				StatusCode: 200,
				Header:     header,
			}
		}
		if req.URL.String() == "https://testgateway/mgmt/login" {
			reqAuth := fmt.Sprint(req.Header["Authorization"])
			defaultAuthCreds := "[Basic dGVzdHVzZXI6dGVzdHBhc3N3b3Jk]"
			header.Set("X-Sds-Auth-Token", "")
			if reqAuth == defaultAuthCreds {
				header.Set("X-Sds-Auth-Token", "TESTTOKEN")
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"OK"}`))),
				Header:     header,
			}
		}
		return nil
	})
}

// Required to force io.ReadAll error
type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("server error: bad body")
}

func testHandleResponse(t *testing.T) {
	response := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><error><description>"OK"</description></error>`))),
	}
	emptyStatusResponse := &http.Response{
		StatusCode: 403,
		Body:       nil,
		Status:     "failed",
	}
	emptyCodeResponse := &http.Response{
		StatusCode: 403,
		Body:       nil,
	}
	notFoundResponse := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><error><code>1004</code></error>`))),
	}

	badBodyResponse := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(errReader(0)),
	}

	err := client.HandleResponse(response)
	require.Error(t, err)

	err = client.HandleResponse(emptyStatusResponse)
	require.Equal(t, err.Error(), "server error: failed")

	err = client.HandleResponse(emptyCodeResponse)
	require.Equal(t, err.Error(), "server error: status code 403")

	err = client.HandleResponse(notFoundResponse)
	require.Equal(t, err.Error(), "server error: not found")

	err = client.HandleResponse(badBodyResponse)
	require.Equal(t, err.Error(), "server error: bad body")

}
