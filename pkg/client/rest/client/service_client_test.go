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

	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func TestServiceRest(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"NoErrorsNoObject":   testServiceNoErrorsNoObject,
		"BadRequest":         testServiceBadRequest,
		"EmptyBody":          testServiceEmptyBody,
		"WithBody":           testServiceWithBody,
		"InvalidEndpoint":    testServiceInvalidEndpoint,
		"InvalidContentType": testServiceInvalidContentType,
		"OverrideHeader":     testServiceOverrideHeader,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}
func testServiceInvalidEndpoint(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		":not:a:valid:url",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
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
func testServiceInvalidContentType(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: "NotAContentType",
	}, nil)

	require.Equal(t, err.Error(), "invalid content-type")
}

func testServiceBadRequest(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
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

func testServiceEmptyBody(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
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

func testServiceWithBody(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
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

func testServiceNoErrorsNoObject(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
		false,
	))

	//Content-type JSON
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(req, nil)
	require.Nil(t, err)

	//Content-type XML
	req = client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/xml",
		ContentType: client.ContentTypeXML,
	}
	err = clientset.Client().MakeRemoteCall(req, nil)
	require.Nil(t, err)
}

func testServiceOverrideHeader(t *testing.T) {
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newTestServiceHTTPClient(),
		true,
	))

	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ok/json",
		ContentType: client.ContentTypeJSON,
	}
	err := clientset.Client().MakeRemoteCall(req, nil)
	require.Nil(t, err)
}
func newTestServiceHTTPClient() *http.Client {
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
		if req.URL.String() == "https://testgateway/mgmt/serviceLogin" {
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
