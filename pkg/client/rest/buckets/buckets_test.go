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
				Expect(len(buckets.Items)).To(Equal(18))
			})
		})
	})

	Context("#Get", func() {
		var (
			bucket *model.Bucket
			err    error
		)

		Context("with no params", func() {
			BeforeEach(func() {
				bucket, err = clientset.Buckets().Get("Files", map[string]string{})
			})

			It("shouldn't error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the bucket", func() {
				Expect(bucket.Name).To(Equal("Files"))
			})
		})
	})

	Context("#Create", func() {
		var (
			bucket *model.Bucket
			err    error
		)

		Context("without one existing", func() {
			BeforeEach(func() {
				createBucket := model.Bucket{
					Name:             "testbucket1",
					ReplicationGroup: "urn:storageos:ReplicationGroupInfo:104b3728-fba1-41b3-8055-4592348f1d24:global",
					Namespace:        "130820808912778549",
				}
				bucket, err = clientset.Buckets().Create(createBucket)
			})

			AfterEach(func() {
				delErr := clientset.Buckets().Delete("testbucket1", "130820808912778549")
				if delErr != nil {
					panic(delErr)
				}
			})

			It("shouldn't error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the bucket", func() {
				Expect(bucket.Name).To(Equal("testbucket1"))
			})
		})
	})

	Context("#Delete", func() {
		var err error

		Context("with no params", func() {
			BeforeEach(func() {
				createBucket := model.Bucket{
					Name:             "testbucket1",
					ReplicationGroup: "urn:storageos:ReplicationGroupInfo:104b3728-fba1-41b3-8055-4592348f1d24:global",
					Namespace:        "130820808912778549",
				}
				_, createErr := clientset.Buckets().Create(createBucket)
				if createErr != nil {
					panic(createErr)
				}
				err = clientset.Buckets().Delete("testbucket1", "130820808912778549")
			})

			It("shouldn't error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

})

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
