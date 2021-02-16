package objmt_test

import (
	"crypto/tls"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"
)

var (
	rec        *recorder.Recorder
	httpClient *http.Client
	clientset  *rest.ClientSet
)

func TestMain(m *testing.M) {
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
	httpClient = &http.Client{Transport: rec}
	clientset = rest.NewClientSet(
		"root",
		"ChangeMe",
		"https://testserver:443",
		httpClient,
		false,
	)
	defer func() {
		if err := rec.Stop(); err != nil {
			log.Fatal(err)
		}
	}()
	m.Run()
}

func TestAccountInfo_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetAccountBillingInfo([]string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestAccountSample_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetAccountBillingSample([]string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestBucketInfo_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetBucketBillingInfo("a12345", []string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestBucketSample_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetBucketBillingSample("a12345", []string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestBucketPerf_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetBucketBillingPerf("a12345", []string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestReplicationInfo_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetReplicationInfo("a12345", [][]string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestReplicationSample_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetReplicationSample("a12345", [][]string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestStoreBillingInfo_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetStoreBillingInfo(nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestStoreBillingSample_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetStoreBillingSample(nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestStorePerf_List(t *testing.T) {
	data, err := clientset.ObjectMt().GetStoreReplicationData([]string{}, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
}
