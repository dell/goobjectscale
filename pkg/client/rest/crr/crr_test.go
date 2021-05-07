package crr_test

import (
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"log"
	"net/http"
	"testing"

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
			"root",
			"ChangeMe",
			"https://testserver",
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
			err = clientset.CRR().PauseReplication("test-objectscale", "test-objectstore", 3000, map[string]string{})
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

	Context("#Throttle", func() {
		var (
			err error
		)
		BeforeEach(func() {
			err = clientset.CRR().ThrottleReplication("test-objectscale", "test-objectstore", 3000, map[string]string{})
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

		It("shouldn't error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return the bucket", func() {
			Expect(crr).ToNot(BeNil())
		})
	})

})

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
