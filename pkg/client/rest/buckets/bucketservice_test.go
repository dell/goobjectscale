// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package buckets_test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/dell/goobjectscale/pkg/client/api"
	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func TestBuckets(t *testing.T) {
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
		Username: "root",
		Password: "Password123!",
	}

	clientset := rest.NewClientSet(&client.Simple{
		Endpoint:       "https://testgateway:4443",
		Authenticator:  &objectscaleAuthUser,
		OverrideHeader: false,
		HTTPClient:     newRecordedHTTPClient(r),
	})

	for scenario, fn := range map[string]func(t *testing.T, mgmtClient api.ClientSet){
		"get":          testGet,
		"create":       testCreate,
		"delete":       testDelete,
		"list":         testListBuckets,
		"updatePolicy": testUpdatePolicy,
		"getPolicy":    testGetPolicy,
		"deletePolicy": testDeletePolicy,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testGet(t *testing.T, mgmtClient api.ClientSet) {
	bucket, err := mgmtClient.Buckets().Get(context.Background(), "bucket1", map[string]string{"namespace": "ns1"})
	require.NoError(t, err)
	assert.Equal(t, bucket.Name, "bucket1")

	// get unknown bucket
	_, err = mgmtClient.Buckets().Get(context.Background(), "bucket3", map[string]string{"namespace": "ns1"})
	require.Error(t, err)
}

func testListBuckets(t *testing.T, mgmtClient api.ClientSet) {
	buckets, err := mgmtClient.Buckets().List(context.Background(), map[string]string{"namespace": "ns1"})
	require.NoError(t, err)
	assert.Equal(t, 5, len(buckets.Buckets))
}

func testCreate(t *testing.T, mgmtClient api.ClientSet) {
	createBucket := &model.ObjectBucketParam{
		Name:      "bucket2",
		Namespace: "ns1",
	}
	bucket, err := mgmtClient.Buckets().Create(context.Background(), createBucket)
	require.NoError(t, err)
	assert.Equal(t, bucket.Name, "bucket2")

	// create unknown bucket
	createBucket = &model.ObjectBucketParam{
		Name:      "bucket3",
		Namespace: "ns1",
	}
	_, err = mgmtClient.Buckets().Create(context.Background(), createBucket)
	require.Error(t, err)
}

func testDelete(t *testing.T, mgmtClient api.ClientSet) {
	err := mgmtClient.Buckets().Delete(context.Background(), "bucket2", map[string]string{"namespace": "ns1"})
	require.NoError(t, err)

	// delete unknown bucket
	err = mgmtClient.Buckets().Delete(context.Background(), "unknownbucket", map[string]string{"namespace": "ns1"})
	require.Error(t, err)
}

func testUpdatePolicy(t *testing.T, mgmtClient api.ClientSet) {
	bucketPolicyJSON := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["*"],"Resource":["arn:aws:s3:::*"]}]}`
	err := mgmtClient.Buckets().UpdatePolicy(context.Background(), "bucket2", bucketPolicyJSON, map[string]string{"namespace": "ns1"})
	require.NoError(t, err)
}

func testDeletePolicy(t *testing.T, mgmtClient api.ClientSet) {
	err := mgmtClient.Buckets().DeletePolicy(context.Background(), "bucket2", map[string]string{"namespace": "ns1"})
	require.NoError(t, err)

	err = mgmtClient.Buckets().DeletePolicy(context.Background(), "bucketunknown", map[string]string{"namespace": "ns1"})
	require.Error(t, err)
}

func testGetPolicy(t *testing.T, mgmtClient api.ClientSet) {
	policy, err := mgmtClient.Buckets().GetPolicy(context.Background(), "bucket2", map[string]string{"namespace": "ns1"})
	require.NoError(t, err)
	assert.NotNil(t, policy)
}
