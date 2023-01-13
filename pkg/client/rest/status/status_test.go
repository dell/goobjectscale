package status_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
)

func TestStatus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Status Spec")
}

var _ = Describe("Status", func() {

	var (
		r         *recorder.Recorder
		clientset *rest.ClientSet
		err       error
	)

	BeforeSuite(func() {
		r, err = recorder.New("fixtures")
		if err != nil {
			log.Fatal(err)
		}
		r.AddFilter(func(i *cassette.Interaction) error {
			delete(i.Request.Headers, "Authorization")
			delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
			return nil
		})
		clientset = rest.NewClientSet(
			"https://testserver",
			"https://testgateway",
			"svc-objectscale-domain-c8",
			"objectscale-graphql-7d754f8499-ng4h6",
			"OSC234DSF223423",
			"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
			newRecordedHTTPClient(r),
			false,
		)
	})

	AfterSuite(func() {
		err = r.Stop()
		if err != nil {
			log.Fatal(err)
		}
	})

	Context("#Get", func() {

		Context("with no params", func() {
			var (
				rebuildInfo *model.RebuildInfo
				err         error
			)
			BeforeEach(func() {
				rebuildInfo, err = clientset.Status().GetRebuildStatus("testdevice1",
					"testdevice1-ss-0", "testdomain", "1", nil)
			})

			It("should return the rebulidInfo", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(rebuildInfo.Level).To(Equal(1))
				Expect(rebuildInfo.RemainingBytes).To(Equal(1024))
				Expect(rebuildInfo.TotalBytes).To(Equal(2048))
			})

		})
	})

})

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
