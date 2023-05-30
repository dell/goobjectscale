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

package alertpolicies_test

import (
	"encoding/xml"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func TestAlertpolicies(t *testing.T) {
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
	for scenario, fn := range map[string]func(t *testing.T, clientset *rest.ClientSet){
		"list":   testList,
		"get":    testGet,
		"create": testCreate,
		"update": testUpdate,
		"delete": testDelete,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testList(t *testing.T, clientset *rest.ClientSet) {
	alertPolicies, err := clientset.AlertPolicies().List(map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, len(alertPolicies.Items), 3)
	_, err = clientset.AlertPolicies().List(map[string]string{"a": "b"})
	require.Error(t, err)
}

func testGet(t *testing.T, clientset *rest.ClientSet) {
	alertPolicy, err := clientset.AlertPolicies().Get("testPolicy")
	require.NoError(t, err)
	assert.Equal(t, alertPolicy.PolicyName, "testPolicy")
	_, err = clientset.AlertPolicies().Get("nonexistentPolicy")
	require.Error(t, err)
}

func testCreate(t *testing.T, clientset *rest.ClientSet) {
	payload := model.AlertPolicy{
		XMLName:    xml.Name{},
		PolicyName: "testPolicy",
	}
	alertPolicy, err := clientset.AlertPolicies().Create(payload)
	require.NoError(t, err)
	assert.Equal(t, alertPolicy.PolicyName, "testPolicy")
}

func testUpdate(t *testing.T, clientset *rest.ClientSet) {
	payload := model.AlertPolicy{
		PolicyName: "testPolicy",
	}
	_, err := clientset.AlertPolicies().Update(payload, "testPolicy")
	require.NoError(t, err)
	// Updating nonexistent policy
	_, err = clientset.AlertPolicies().Update(payload, "nonexistentPolicy")
	require.Error(t, err)
}

func testDelete(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.AlertPolicies().Delete("testPolicy")
	require.NoError(t, err)
	err = clientset.AlertPolicies().Delete("")
	require.Error(t, err)
}
