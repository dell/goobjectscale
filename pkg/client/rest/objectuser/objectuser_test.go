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

package objectuser_test

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/dell/goobjectscale/pkg/client/api"
	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func TestObjectUser(t *testing.T) {
	var (
		r   *recorder.Recorder
		err error
	)

	r, err = recorder.New("fixtures")
	if err != nil {
		log.Fatal(err)
	}

	r.AddHook(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")

		return nil
	}, recorder.BeforeSaveHook)

	r.SetRealTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec
		},
	})

	c := client.Simple{
		Endpoint: "https://testserver",
		Authenticator: &client.AuthService{
			Gateway:       "https://testgateway",
			SharedSecret:  "OSC234DSF223423",
			PodName:       "objectscale-graphql-7d754f8499-ng4h6",
			Namespace:     "svc-objectscale-domain-c8",
			ObjectScaleID: "IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		},
		HTTPClient:     newRecordedHTTPClient(r),
		OverrideHeader: false,
	}
	clientset := rest.NewClientSet(&c)

	for scenario, fn := range map[string]func(t *testing.T, clientset api.ClientSet){
		"list":         testList,
		"getInfo":      testGetInfo,
		"getSecret":    testGetSecret,
		"createSecret": testCreateSecret,
		"deleteSecret": testDeleteSecret,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testList(t *testing.T, clientset api.ClientSet) {
	data, err := clientset.ObjectUser().List(context.TODO(), nil)
	require.NoError(t, err)
	require.Len(t, data.BlobUser, 1)
	require.Equal(t, data.BlobUser[0].UserID, "zmvodjnrbmjxagvwcxf5cg==")
	require.Equal(t, data.BlobUser[0].Namespace, "small-operator-acceptance")

	_, err = clientset.ObjectUser().List(context.TODO(), map[string]string{"a": "b"})
	require.Error(t, err)
}

var objectUserGetInfoTest = []struct {
	in      string
	out     *model.ObjectUserInfo
	withErr bool
}{
	{
		in: "zmvodjnrbmjxagvwcxf5cg==",
		out: &model.ObjectUserInfo{
			Namespace: "small-operator-acceptance",
			Name:      "zmvodjnrbmjxagvwcxf5cg==",
			Created:   "Mon Nov 25 09:55:38 GMT 2019",
		},
	},
	{
		in:      "non existed UID",
		out:     nil,
		withErr: true,
	},
}

func testGetInfo(t *testing.T, clientset api.ClientSet) {
	for _, tt := range objectUserGetInfoTest {
		t.Run(tt.in, func(t *testing.T) {
			data, err := clientset.ObjectUser().GetInfo(context.TODO(), tt.in, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, data, tt.out)
		})
	}
}

var objectUserGetSecretTest = []struct {
	in      string
	out     *model.ObjectUserSecret
	withErr bool
}{
	{
		in: "zmvodjnrbmjxagvwcxf5cg==",
		out: &model.ObjectUserSecret{
			SecretKey1:    "cmowa2dxeXZrM2U1NWptdHdrdTB6a3B4YnRub3RwZHY=",
			KeyTimestamp1: "2019-11-25 09:56:32.364",
			Link: model.Link{
				HREF: "/object/secret-keys",
				Rel:  "self",
			},
		},
	},
	{
		in:      "non existed UID",
		out:     nil,
		withErr: true,
	},
}

func testGetSecret(t *testing.T, clientset api.ClientSet) {
	for _, tt := range objectUserGetSecretTest {
		t.Run(tt.in, func(t *testing.T) {
			data, err := clientset.ObjectUser().GetSecret(context.TODO(), tt.in, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, data, tt.out)
		})
	}
}

var objectUserCreateSecretTest = []struct {
	in1     string
	in2     model.ObjectUserSecretKeyCreateReq
	out     *model.ObjectUserSecretKeyCreateRes
	withErr bool
}{
	{
		in1: "zmvodjnrbmjxagvwcxf5cg==",
		in2: model.ObjectUserSecretKeyCreateReq{
			SecretKey: "aaaaaa",
		},
		out: &model.ObjectUserSecretKeyCreateRes{
			SecretKey:          "aaaaaa",
			KeyTimeStamp:       "2020-03-12 19:45:00.333",
			KeyExpiryTimestamp: "",
		},
	},
	{
		in1:     "non existed UID",
		out:     nil,
		withErr: true,
	},
}

func testCreateSecret(t *testing.T, clientset api.ClientSet) {
	for _, tt := range objectUserCreateSecretTest {
		t.Run(tt.in1, func(t *testing.T) {
			data, err := clientset.ObjectUser().CreateSecret(context.TODO(), tt.in1, tt.in2, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, data, tt.out)
		})
	}
}

var objectUserDeleteSecretTest = []struct {
	in1     string
	in2     model.ObjectUserSecretKeyDeleteReq
	withErr bool
}{
	{
		in1: "zmvodjnrbmjxagvwcxf5cg==",
		in2: model.ObjectUserSecretKeyDeleteReq{
			SecretKey: "aaaaaa",
		},
	},
	{
		in1: "non existed UID",
		in2: model.ObjectUserSecretKeyDeleteReq{
			SecretKey: "aaaaaa",
		},
		withErr: true,
	},
}

func testDeleteSecret(t *testing.T, clientset api.ClientSet) {
	for _, tt := range objectUserDeleteSecretTest {
		t.Run(tt.in1, func(t *testing.T) {
			err := clientset.ObjectUser().DeleteSecret(context.TODO(), tt.in1, tt.in2, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
