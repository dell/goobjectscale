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

package iam_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	awsIAM "github.com/aws/aws-sdk-go/service/iam"
	objscaleIAM "github.com/dell/goobjectscale/pkg/client/rest/iam"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

var (
	iamClient *awsIAM.IAM
	endpoint  = "https://testgateway/iam"
	region    = "us-west-2"
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

// newTestClient returns *http.Client with Transport replaced to avoid making real calls
func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestIam(t *testing.T) {
	tests := map[string]func(*testing.T, *awsIAM.IAM){
		"testInjectTokenToIAMClient":     testInjectTokenToIAMClient,
		"testInjectAccountIDToIAMClient": testInjectAccountIDToIAMClient,
	}

	iamSession, _ := session.NewSession(&aws.Config{
		Endpoint:                      &endpoint,
		Region:                        &region,
		CredentialsChainVerboseErrors: aws.Bool(true),
		HTTPClient:                    &http.Client{},
	})

	iamClient = awsIAM.New(iamSession)

	t.Run("IAM Test", func(t *testing.T) {
		for scenario, fn := range tests {
			t.Run(scenario, func(t *testing.T) {
				fn(t, iamClient)
			})
		}
	})
}

func testInjectTokenToIAMClient(t *testing.T, iamClient *awsIAM.IAM) {
	TestClient := newTestHTTPClient()

	objscaleIAM.InjectTokenToIAMClient(iamClient, &FixtureServiceauth, *TestClient)
	r := &request.Request{
		HTTPRequest: &http.Request{
			Header: http.Header{},
		},
	}
	iamClient.Handlers.Sign.Run(r)
	checkHeader(t, *r, objscaleIAM.SDSHeaderName, "TESTTOKEN")

	objscaleIAM.InjectTokenToIAMClient(iamClient, &FixtureUserAuth, *TestClient)
	r = &request.Request{
		HTTPRequest: &http.Request{
			Header: http.Header{},
		},
	}
	iamClient.Handlers.Sign.Run(r)
	checkHeader(t, *r, objscaleIAM.SDSHeaderName, "TESTTOKEN")
}

func testInjectAccountIDToIAMClient(t *testing.T, iamClient *awsIAM.IAM) {
	objscaleIAM.InjectAccountIDToIAMClient(iamClient, "xxx")
	r := &request.Request{
		HTTPRequest: &http.Request{
			Header: http.Header{},
		},
	}
	iamClient.Handlers.Sign.Run(r)
	checkHeader(t, *r, objscaleIAM.AccountIDHeaderName, "xxx")

	objscaleIAM.InjectAccountIDToIAMClient(iamClient, "yyy")
	r = &request.Request{
		HTTPRequest: &http.Request{
			Header: http.Header{},
		},
	}
	iamClient.Handlers.Sign.Run(r)
	checkHeader(t, *r, objscaleIAM.AccountIDHeaderName, "yyy")

}

func checkHeader(t *testing.T, request request.Request, expectedHeader string, expectedHeaderValue string) {
	sdsHeader, ok := request.HTTPRequest.Header[expectedHeader]
	if !ok {
		t.Errorf("expect %v header", expectedHeader)
	}
	actualHeaderValue := sdsHeader[0]
	if actualHeaderValue != expectedHeaderValue {
		t.Fatalf("expected: %v actual: %v token in %v header", expectedHeaderValue, actualHeaderValue, expectedHeader)
	}
}

func newTestHTTPClient() *http.Client {
	return newTestClient(func(req *http.Request) *http.Response {
		header := make(http.Header)
		switch req.URL.String() {
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
