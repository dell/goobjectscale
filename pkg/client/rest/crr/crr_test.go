package crr_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCRR(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CRR Spec")
}

var _ = Describe("CRR", func() {

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
			"OSTOKEN-eyJ4NXUiOi",
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
