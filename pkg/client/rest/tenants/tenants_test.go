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

package tenants_test

import (
	"encoding/xml"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
func TestTenants(t *testing.T) {
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
		"list":        testList,
		"get":         testGet,
		"create":      testCreate,
		"update":      testUpdate,
		"delete":      testDelete,
		"getQuota":    testGetQuota,
		"setQuota":    testSetQuota,
		"deleteQuota": testDeleteQuota,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testList(t *testing.T, clientset *rest.ClientSet) {
	tenants, err := clientset.Tenants().List(map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, len(tenants.Items), 1)
	_, err = clientset.Tenants().List(map[string]string{"a": "b"})
	require.Error(t, err)
}
func testGet(t *testing.T, clientset *rest.ClientSet) {
	tenant, err := clientset.Tenants().Get("10d9817c-3696-4625-854e-82b21d8c0795", map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, "10d9817c-3696-4625-854e-82b21d8c0795", tenant.ID)
	_, err = clientset.Tenants().Get("0", map[string]string{})
	require.Error(t, err)

}
func testCreate(t *testing.T, clientset *rest.ClientSet) {
	payload := model.TenantCreate{
		XMLName:           xml.Name{},
		Alias:             "",
		AccountID:         "test-account",
		EncryptionEnabled: false,
		ComplianceEnabled: false,
		BucketBlockSize:   0,
	}
	tenant, err := clientset.Tenants().Create(payload)
	require.NoError(t, err)
	assert.Equal(t, "test-account", tenant.ID)
}
func testUpdate(t *testing.T, clientset *rest.ClientSet) {
	payload := model.TenantUpdate{
		Alias:           "test-alias",
		BucketBlockSize: 100,
	}
	err := clientset.Tenants().Update(payload, "test-account")
	require.NoError(t, err)
	err = clientset.Tenants().Update(payload, "noaccount")
	require.Error(t, err)
}
func testDelete(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.Tenants().Delete("test-account")
	require.NoError(t, err)
	err = clientset.Tenants().Delete("noaccount")
	require.Error(t, err)
}
func testGetQuota(t *testing.T, clientset *rest.ClientSet) {
	quota, err := clientset.Tenants().GetQuota("test-account", nil)
	require.NoError(t, err)
	assert.Equal(t, "5", quota.BlockSize)
}
func testSetQuota(t *testing.T, clientset *rest.ClientSet) {
	payload := model.TenantQuotaSet{
		XMLName:                 xml.Name{},
		BlockSize:               "5",
		NotificationSize:        "6",
		BlockSizeInCount:        "7",
		NotificationSizeInCount: "8",
	}
	err := clientset.Tenants().SetQuota("test-account", payload)
	require.NoError(t, err)
}
func testDeleteQuota(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.Tenants().DeleteQuota("test-account")
	require.NoError(t, err)
}
