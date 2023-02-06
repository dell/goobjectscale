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

package buckets_test

import (
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
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newRecordedHTTPClient(r),
		false,
	))
	for scenario, fn := range map[string]func(t *testing.T, clientset *rest.ClientSet){
		"list":   testList,
		"get":    testGet,
		"create": testCreate,
		"delete": testDelete,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testList(t *testing.T, clientset *rest.ClientSet) {
	buckets, err := clientset.Buckets().List(map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, len(buckets.Items), 18)
}

func testGet(t *testing.T, clientset *rest.ClientSet) {
	bucket, err := clientset.Buckets().Get("Files", map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, bucket.Name, "Files")
}

func testCreate(t *testing.T, clientset *rest.ClientSet) {
	createBucket := model.Bucket{
		Name:             "testbucket1",
		ReplicationGroup: "urn:storageos:ReplicationGroupInfo:104b3728-fba1-41b3-8055-4592348f1d24:global",
		Namespace:        "130820808912778549",
	}
	bucket, err := clientset.Buckets().Create(createBucket)

	require.NoError(t, err)
	assert.Equal(t, bucket.Name, "testbucket1")

	//Delete bucket after
	delErr := clientset.Buckets().Delete("testbucket1", "130820808912778549")
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
	_, createErr := clientset.Buckets().Create(createBucket)
	if createErr != nil {
		panic(createErr)
	}
	err := clientset.Buckets().Delete("testbucket1", "130820808912778549")
	require.NoError(t, err)

}
