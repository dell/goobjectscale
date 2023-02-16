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

package objmt_test

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"

	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func TestObjmt(t *testing.T) {
	var rec *recorder.Recorder
	var err error
	rec, err = recorder.New("objmt_fixtures")
	if err != nil {
		log.Fatal(err)
	}
	rec.AddHook(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
		return nil
	}, recorder.BeforeSaveHook)
	rec.SetRealTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	rec.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		if r.Body == nil {
			return cassette.DefaultMatcher(r, i)
		}
		var b bytes.Buffer
		if _, err := b.ReadFrom(r.Body); err != nil {
			return false
		}
		r.Body = io.NopCloser(&b)
		return cassette.DefaultMatcher(r, i) && (b.String() == "" || b.String() == i.Body)
	})
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newRecordedHTTPClient(rec),
		false,
	))
	for scenario, fn := range map[string]func(t *testing.T, clientset *rest.ClientSet){
		"accountInfo":        testAccountInfoList,
		"accountSample":      testAccountSampleList,
		"bucketInfo":         testBucketInfoList,
		"bucketSample":       testBucketSampleList,
		"bucketPerf":         testBucketPerfList,
		"replicationInfo":    testReplicationInfoList,
		"replicationSample":  testReplicationSampleList,
		"storeBillingInfo":   testStoreBillingInfoList,
		"storeBillingSample": testStoreBillingSampleList,
		"storePerf":          testStorePerfList,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testAccountInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetAccountBillingInfo([]string{"aaa", "bbb", "ccc"}, nil)
	//Success
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetAccountBillingInfo([]string{}, nil)
	require.Error(t, err)
}

func testAccountSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetAccountBillingSample([]string{"aaa", "bbb", "ccc"}, nil)
	//Success
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetAccountBillingSample([]string{}, nil)
	require.Error(t, err)
}

func testBucketInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetBucketBillingInfo("a12345", []string{"aaa", "bbb", "ccc"}, nil)
	//Success
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetBucketBillingInfo("a12345", []string{}, nil)
	require.Error(t, err)
}

func testBucketSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetBucketBillingSample("a12345", []string{"aaa", "bbb", "ccc"}, nil)
	//Success
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetBucketBillingSample("a12345", []string{}, nil)
	require.Error(t, err)
}

func testBucketPerfList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetBucketBillingPerf("a12345", []string{"aaa", "bbb", "ccc"}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetBucketBillingPerf("a12345", []string{}, nil)
	require.Error(t, err)
}

func testReplicationInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetReplicationInfo("a12345", [][]string{{"a", "b"}, {"c", "d"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetReplicationInfo("a12345", [][]string{}, nil)
	require.Error(t, err)
}

func testReplicationSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetReplicationSample("a12345", [][]string{{"a", "b"}, {"c", "d"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetReplicationSample("a12345", [][]string{}, nil)
	require.Error(t, err)
}

func testStoreBillingInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetStoreBillingInfo(nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	_, err = clientset.ObjectMt().GetStoreBillingInfo(map[string]string{"a": "b"})
	require.Error(t, err)
}

func testStoreBillingSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetStoreBillingSample(nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	_, err = clientset.ObjectMt().GetStoreBillingSample(map[string]string{"a": "b"})
	require.Error(t, err)
}

func testStorePerfList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetStoreReplicationData([]string{"aaa", "bbb", "ccc"}, nil)
	//Success
	require.NoError(t, err)
	require.NotNil(t, data)
	//Fail
	_, err = clientset.ObjectMt().GetStoreReplicationData([]string{}, nil)
	require.Error(t, err)
}
