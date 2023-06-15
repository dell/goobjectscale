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

package buckets_test

import (
	"context"
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

func TestBuckets(t *testing.T) {
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
		"list":         testList,
		"get":          testGet,
		"create":       testCreate,
		"delete":       testDelete,
		"getQuota":     testGetQuota,
		"getPolicy":    testGetPolicy,
		"updatePolicy": testUpdatePolicy,
		"deletePolicy": testDeletePolicy,
		"UpdateQuota":  testUpdateQuota,
		"deleteQuota":  testDeleteQuota,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testList(t *testing.T, clientset *rest.ClientSet) {
	buckets, err := clientset.Buckets().List(context.TODO(), map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, len(buckets.Items), 18)
	_, err = clientset.Buckets().List(context.TODO(), map[string]string{"a": "b"})
	require.Error(t, err)
}

func testGet(t *testing.T, clientset *rest.ClientSet) {
	bucket, err := clientset.Buckets().Get(context.TODO(), "Files", map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, bucket.Name, "Files")
	_, err = clientset.Buckets().Get(context.TODO(), "Flies", map[string]string{})
	require.Error(t, err)
}

func testCreate(t *testing.T, clientset *rest.ClientSet) {
	createBucket := model.Bucket{
		Name:             "testbucket1",
		ReplicationGroup: "urn:storageos:ReplicationGroupInfo:104b3728-fba1-41b3-8055-4592348f1d24:global",
		Namespace:        "130820808912778549",
	}
	bucket, err := clientset.Buckets().Create(context.TODO(), createBucket)

	require.NoError(t, err)
	assert.Equal(t, bucket.Name, "testbucket1")

	//Delete bucket after
	delErr := clientset.Buckets().Delete(context.TODO(), "testbucket1", "130820808912778549", true)
	if delErr != nil {
		panic(delErr)
	}
}

func testDelete(t *testing.T, clientset *rest.ClientSet) {
	createBucket := model.Bucket{
		Name:             "testbucket1",
		ReplicationGroup: "urn:storageos:ReplicationGroupInfo:104b3728-fba1-41b3-8055-4592348f1d24:global",
		Namespace:        "130820808912778549",
	}
	_, createErr := clientset.Buckets().Create(context.TODO(), createBucket)
	if createErr != nil {
		panic(createErr)
	}
	err := clientset.Buckets().Delete(context.TODO(), "testbucket1", "130820808912778549", true)
	require.NoError(t, err)
	err = clientset.Buckets().Delete(context.TODO(), "unknownbucket", "130820808912778549", true)
	require.Error(t, err)
}

func testGetQuota(t *testing.T, clientset *rest.ClientSet) {
	quota, err := clientset.Buckets().GetQuota(context.TODO(), "testbucket1", "130820808912778549")
	require.NoError(t, err)
	assert.Equal(t, int(quota.BlockSize), -1)
	_, err = clientset.Buckets().GetQuota(context.TODO(), "testbucket", "130820808912778549")
	require.Error(t, err)
}

func testGetPolicy(t *testing.T, clientset *rest.ClientSet) {
	_, err := clientset.Buckets().GetPolicy(context.TODO(), "testbucket1", map[string]string{})
	require.NoError(t, err)
	_, err = clientset.Buckets().GetPolicy(context.TODO(), "unknownbucket", map[string]string{})
	require.Error(t, err)
}

func testUpdatePolicy(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.Buckets().UpdatePolicy(context.TODO(), "testbucket1", "testpolicy", map[string]string{})
	require.Nil(t, err)
}

func testDeletePolicy(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.Buckets().DeletePolicy(context.TODO(), "testbucket1", map[string]string{})
	require.Nil(t, err)
}

func testUpdateQuota(t *testing.T, clientset *rest.ClientSet) {
	bucketQuotaUpdate := model.BucketQuotaUpdate{}
	bucketQuota := model.BucketQuota{
		BucketName: "testbucket1",
	}
	bucketQuotaUpdate.BucketQuota = bucketQuota
	err := clientset.Buckets().UpdateQuota(context.TODO(), bucketQuotaUpdate)
	require.Nil(t, err)

}

func testDeleteQuota(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.Buckets().DeleteQuota(context.TODO(), "testbucket1", "130820808912778549")
	require.Nil(t, err)
}
