// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package vpool_test

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/dell/goobjectscale/pkg/client/api"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func TestVPool(t *testing.T) {
	var (
		r   *recorder.Recorder
		err error
	)

	r, err = recorder.New("../../../../fixtures/testdata/fixtures")
	if err != nil {
		log.Fatal(err)
	}

	r.AddHook(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")

		return nil
	}, recorder.BeforeSaveHook)

	objectscaleAuthUser := client.AuthUser{
		Gateway:  "https://testgateway:4443",
		Username: "",
		Password: "",
	}

	clientset := rest.NewClientSet(&client.Simple{
		Endpoint:       "https://testgateway:4443",
		Authenticator:  &objectscaleAuthUser,
		OverrideHeader: false,
		HTTPClient:     newRecordedHTTPClient(r),
	})

	for scenario, fn := range map[string]func(t *testing.T, mgmtClient api.ClientSet){
		"list": testListVPool,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func testListVPool(t *testing.T, mgmtClient api.ClientSet) {
	vpools, err := mgmtClient.VPools().List(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, vpools)
}

func TestInvalidJSONResponse(t *testing.T) {
	c := client.Simple{
		Endpoint: "https://testgateway",
		Authenticator: &client.AuthUser{
			Gateway:  "https://testgateway",
			Username: "testuser",
			Password: "testpassword",
		},
		HTTPClient:     NewTestHTTPClient(),
		OverrideHeader: true,
	}
	clientset := rest.NewClientSet(&c)
	list, err := clientset.VPools().List(context.Background())
	assert.Nil(t, list)
	assert.NotNil(t, err)
}

type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip mocks an http request and returns an http response.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func NewTestHTTPClient() *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		switch req.URL.String() {
		case "https://testgateway/login":
			header := make(http.Header)
			header.Set("X-Sds-Auth-Token", "")
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"user":"root"}`))),
				Header:     header,
			}
		case "https://testgateway/vdc/data-service/vpools":
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`invalid-json`))),
			}
		}
		return nil
	})
}
