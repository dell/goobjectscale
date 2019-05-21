package buckets_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dnaeon/go-vcr/recorder"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
)

func TestBuckets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buckets Spec")
}

var _ = Describe("Buckets", func() {

	var (
		r *recorder.Recorder
		clientset *rest.ClientSet
		err error
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
		)
	})

	AfterSuite(func() {
		err = r.Stop()
		if err != nil {
			log.Fatal(err)
		}
	})

	Context("#List", func() {
		var (
			buckets *model.BucketList
			err     error
		)

		Context("with no params", func() {
			BeforeEach(func() {
				buckets, err = clientset.Buckets().List(map[string]string{})
			})

			It("shouldn't error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have buckets", func() {
				Expect(len(buckets.Items)).To(Equal(19))
			})
		})
	})
})

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
