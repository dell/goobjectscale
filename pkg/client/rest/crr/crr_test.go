package crr_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
		return nil
	})
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

func testPause(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.CRR().PauseReplication("test-objectscale", "test-objectstore", map[string]string{"pauseEndMills": "3000"})
	require.NoError(t, err)
}

func testSuspend(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.CRR().SuspendReplication("test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
}

func testResume(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.CRR().ResumeReplication("test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
}
func testThrottle(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.CRR().ThrottleReplication("test-objectscale", "test-objectstore", map[string]string{"throttleMBPerSecond": "3000"})
	require.NoError(t, err)
}
func testUnthrottle(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.CRR().UnthrottleReplication("test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
}

func testGet(t *testing.T, clientset *rest.ClientSet) {
	crr, err := clientset.CRR().Get("test-objectscale", "test-objectstore", map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, crr.DestObjectStore, "test-objectstore")
}
