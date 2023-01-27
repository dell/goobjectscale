package crr_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func TestCRR(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CRR Spec")
}

var (
	r         *recorder.Recorder
	clientset *rest.ClientSet
	err       error
)

var _ = BeforeSuite(func() {
	r, err = recorder.New("fixtures")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
		return nil
	})
	clientset = rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newRecordedHTTPClient(r),
		false,
	))
})

var _ = AfterSuite(func() {
	err = r.Stop()
	if err != nil {
		log.Fatal(err)
	}
})

var _ = Describe("CRR", func() {

	Context("#Pause", func() {
		var (
			err error
		)
		BeforeEach(func() {
			err = clientset.CRR().PauseReplication("test-objectscale", "test-objectstore", map[string]string{"pauseEndMills": "3000"})
		})

		It("shouldn't error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("#Suspend", func() {
		var (
			err error
		)
		BeforeEach(func() {
			err = clientset.CRR().SuspendReplication("test-objectscale", "test-objectstore", map[string]string{})
		})

		It("shouldn't error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("#Resume", func() {
		var (
			err error
		)
		BeforeEach(func() {
			err = clientset.CRR().ResumeReplication("test-objectscale", "test-objectstore", map[string]string{})
		})

		It("shouldn't error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("#Unthrottle", func() {
		var (
			err error
		)
		BeforeEach(func() {
			err = clientset.CRR().UnthrottleReplication("test-objectscale", "test-objectstore", map[string]string{})
		})

		It("shouldn't error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("#Throttle", func() {
		var (
			err error
		)
		BeforeEach(func() {
			err = clientset.CRR().ThrottleReplication("test-objectscale", "test-objectstore", map[string]string{"throttleMBPerSecond": "3000"})
		})

		It("shouldn't error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("#Get", func() {
		var (
			crr *model.CRR
			err error
		)
		BeforeEach(func() {
			crr, err = clientset.CRR().Get("test-objectscale", "test-objectstore", map[string]string{})
		})

		It("should return the bucket successfully", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(crr.DestObjectStore).To(Equal("test-objectstore"))
		})
	})

})

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
