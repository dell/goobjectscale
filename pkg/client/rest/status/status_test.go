package status_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func TestStatus(t *testing.T) {
	t.Run("should return build info", func(t *testing.T) {
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
		rebuildInfo, err := clientset.Status().GetRebuildStatus("testdevice1",
			"testdevice1-ss-0", "testdomain", "1", nil)
		require.NoError(t, err)
		assert.Equal(t, rebuildInfo.Level, 1)
		assert.Equal(t, rebuildInfo.RemainingBytes, 1024)
		assert.Equal(t, rebuildInfo.TotalBytes, 2048)
	})
}

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
