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

package crr_test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/dell/goobjectscale/pkg/client/api"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func TestCRR(t *testing.T) {
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
		"pause":      testPause,
		"suspend":    testSuspend,
		"resume":     testResume,
		"get":        testGet,
		"throttle":   testThrottle,
		"unthrottle": testUnthrottle,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testPause(t *testing.T, clientset api.ClientSet) {
	err := clientset.CRR().PauseReplication(context.TODO(), "test-objectscale", "test-objectstore", map[string]string{"pauseEndMills": "3000"})
	require.NoError(t, err)
}

func testSuspend(t *testing.T, clientset api.ClientSet) {
	err := clientset.CRR().SuspendReplication(context.TODO(), "test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
}

func testResume(t *testing.T, clientset api.ClientSet) {
	err := clientset.CRR().ResumeReplication(context.TODO(), "test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
}

func testThrottle(t *testing.T, clientset api.ClientSet) {
	err := clientset.CRR().ThrottleReplication(context.TODO(), "test-objectscale", "test-objectstore", map[string]string{"throttleMBPerSecond": "3000"})
	require.NoError(t, err)
}

func testUnthrottle(t *testing.T, clientset api.ClientSet) {
	err := clientset.CRR().UnthrottleReplication(context.TODO(), "test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
}

func testGet(t *testing.T, clientset api.ClientSet) {
	crr, err := clientset.CRR().Get(context.TODO(), "test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, crr.DestObjectStore, "test-objectstore")

	_, err = clientset.CRR().Get(context.TODO(), "bad-objectscale", "bad-objectstore", map[string]string{})
	require.Error(t, err)
}
