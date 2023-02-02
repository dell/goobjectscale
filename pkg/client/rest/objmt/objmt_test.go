package objmt_test

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"
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
	rec.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
		return nil
	})
	rec.SetTransport(&http.Transport{
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
		r.Body = ioutil.NopCloser(&b)
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
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testAccountSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetAccountBillingSample([]string{"aaa", "bbb", "ccc"}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testBucketInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetBucketBillingInfo("a12345", []string{"aaa", "bbb", "ccc"}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testBucketSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetBucketBillingSample("a12345", []string{"aaa", "bbb", "ccc"}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testBucketPerfList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetBucketBillingPerf("a12345", []string{"aaa", "bbb", "ccc"}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testReplicationInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetReplicationInfo("a12345", [][]string{{"a", "b"}, {"c", "d"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testReplicationSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetReplicationSample("a12345", [][]string{{"a", "b"}, {"c", "d"}}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testStoreBillingInfoList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetStoreBillingInfo(nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testStoreBillingSampleList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetStoreBillingSample(nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func testStorePerfList(t *testing.T, clientset *rest.ClientSet) {
	data, err := clientset.ObjectMt().GetStoreReplicationData([]string{"aaa", "bbb", "ccc"}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}
