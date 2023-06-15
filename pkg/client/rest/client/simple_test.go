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

package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"

	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

var FixtureUserAuth = client.AuthUser{
	Gateway:  "https://testgateway",
	Username: "testuser",
	Password: "testpassword",
}

var FixtureServiceauth = client.AuthService{
	Gateway:       "https://testgateway",
	SharedSecret:  "OSC234DSF223423",
	PodName:       "objectscale-graphql-7d754f8499-ng4h6",
	Namespace:     "svc-objectscale-domain-c8",
	ObjectScaleID: "IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
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

func TestSimple(t *testing.T) {
	tests := map[string]func(*testing.T, client.Authenticator){
		"NoErrorsNoObject":   testNoErrorsNoObject,
		"BadRequest":         testBadRequest,
		"EmptyBody":          testEmptyBody,
		"WithBody":           testWithBody,
		"InvalidEndpoint":    testInvalidEndpoint,
		"InvalidContentType": testInvalidContentType,
		"FailedLogin":        testFailedLogin,
		"NoGateway":          testNoGateway,
		"OverriderHeader":    testOverrideHeader,
		"FailedAuth":         testFailedAuth,
		"TestHttp":           testHTTP,
		"TestLogin":          testLogin,
	}

	t.Run("service-auth", func(t *testing.T) {
		for scenario, fn := range tests {
			t.Run(scenario, func(t *testing.T) {
				service := FixtureServiceauth // shallow copy
				fn(t, &service)
			})
		}
	})

	t.Run("user-auth", func(t *testing.T) {
		for scenario, fn := range tests {
			t.Run(scenario, func(t *testing.T) {
				user := FixtureUserAuth // shallow copy
				fn(t, &user)
			})
		}
	})
}

func testInvalidEndpoint(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       ":not:a:valid:url",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: true,
	}
	clientset := rest.NewClientSet(&c)
	err := clientset.Client().MakeRemoteCall(context.TODO(), client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: client.ContentTypeJSON,
	}, nil)
	errUrl := &url.Error{}
	require.Error(t, err)
	assert.ErrorAs(t, err, &errUrl)
}

func testInvalidContentType(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)
	err := clientset.Client().MakeRemoteCall(context.TODO(),
		client.Request{
			Method:      http.MethodGet,
			Path:        "",
			ContentType: "NotAContentType",
		}, nil)
	require.ErrorIs(t, err, client.ErrContentType)
}
func testFailedAuth(t *testing.T, auth client.Authenticator) {
	badUser := client.AuthUser{
		Gateway:  "https://testgateway",
		Username: "testuser1",
		Password: "testpassword1",
	}
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  &badUser,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Error(t, err)
}

func testFailedLogin(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/failedLogin",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Error(t, err)
	assert.Equal(t, "authorization: exhausted authentication tries", err.Error())
	c.Authenticator = nil
	clientset = rest.NewClientSet(&c)
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Error(t, err)
}

func testBadRequest(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)
	//Content-type JSON
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/badRequest",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Error(t, err)
	//Content-type XML
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/badRequest",
		ContentType: client.ContentTypeXML,
	}
	nerr := clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Error(t, nerr)
}

func testEmptyBody(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/emptyBody",
		ContentType: client.ContentTypeJSON,
	}
	bucketList := &model.BucketList{}
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, bucketList)
	require.Nil(t, err)
}

func testWithBody(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)
	//JSON
	//Success
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}
	ecsError := &model.Error{}
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.Nil(t, err)
	//Error
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/xml",
		ContentType: client.ContentTypeJSON,
	}
	errJSONSyntax := &json.SyntaxError{}
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.ErrorAs(t, err, &errJSONSyntax)
	require.Contains(t, err.Error(), "invalid character '<' looking for beginning of value")
	// force error in request.go
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
		Body:        math.Inf(-1),
	}
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.Error(t, err)
	//XML
	//Success
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/xml",
		ContentType: client.ContentTypeXML,
	}
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.Nil(t, err)
	//Error
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeXML,
	}
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.ErrorIs(t, err, io.EOF)
	require.Contains(t, err.Error(), "EOF")
	// force error in request.go
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeXML,
		Body:        map[string]string{},
	}
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.Error(t, err)
}

func testNoErrorsNoObject(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}

	//Content-type JSON
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Nil(t, err)

	//Content-type XML
	err = clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Nil(t, err)
}

func testNoGateway(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)
	ecsError := &model.Error{}
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/badRequest",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, ecsError)
	require.Error(t, err)

}

func testOverrideHeader(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: true,
	}
	clientset := rest.NewClientSet(&c)

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}

	//Content-type JSON
	err := clientset.Client().MakeRemoteCall(context.TODO(), req, nil)
	require.Nil(t, err)
}

func testHTTP(t *testing.T, auth client.Authenticator) {
	c := client.Simple{
		Endpoint:       "https://testserver",
		Authenticator:  auth,
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: false,
	}
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: "NotAContentType",
		Params:      map[string]string{"param1": "param"},
	}
	_, err := req.HTTP(c.Endpoint)
	require.Error(t, err)
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/test",
		ContentType: client.ContentTypeJSON,
		Params:      map[string]string{"param1": "param"},
	}
	_, err = req.HTTP(c.Endpoint)
	require.NoError(t, err)

	req = client.Request{
		Method:      "not a method",
		Path:        "/test",
		ContentType: client.ContentTypeJSON,
		Params:      map[string]string{"param1": "param"},
	}
	_, err = req.HTTP(c.Endpoint)
	require.Error(t, err)
	require.Equal(t, "create request: net/http: invalid method \"not a method\"", err.Error())
}

func NewTestHTTPClient() *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		header := make(http.Header)
		switch req.URL.String() {
		case "https://testserver/ok/json":
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"OK"}`))),
				Header:     header,
			}

		case "https://testserver/ok/xml":
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><error><Description>"OK"</Description></error>`))),
				Header:     header,
			}

		case "https://testserver/failedLogin":
			return &http.Response{
				StatusCode: 401,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"Unauthorized"}`))),
				Header:     header,
			}

		case "https://testserver/badRequest":
			return &http.Response{
				StatusCode: 403,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"Forbidden"}`))),
				Header:     header,
			}

		case "https://testserver/emptyBody":
			return &http.Response{
				StatusCode: 200,
				Header:     header,
			}

		case "https://testgateway/mgmt/login":
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

		case "https://testgateway/mgmt/serviceLogin":
			header.Set("X-SDS-AUTH-TOKEN", "TESTTOKEN")
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"Description":"OK"}`))),
				Header:     header,
			}

		}
		return nil
	})
}
