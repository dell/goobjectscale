package serviceprovider_test

import (
	"encoding/xml"
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

func TestServiceProvider(sp *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(sp, "ServiceProvider Spec")
}

var _ = Describe("ServiceProvider", func() {

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

	Context("#Get", func() {
		Context("with no params", func() {
			var (
				serviceProvider *model.ServiceProvider
				err             error
			)

			It("should have the requested ServiceProvider", func() {
				serviceProvider, err = clientset.ServiceProvider().Get(map[string]string{})
				Expect(err).ToNot(HaveOccurred())
				Expect(serviceProvider.UniqueID).To(Equal("48eef7e9-aadc-3a01-813b-54b06abb8d80"))
			})
		})
	})

	Context("#Create", func() {
		Context("with no params", func() {
			var (
				serviceProvider *model.ServiceProvider
				err             error
			)

			BeforeEach(func() {
				payload := model.ServiceProviderCreate{
					XMLName: xml.Name{},
					ServiceProvider: model.ServiceProvider{
						CreateTime:   "2020-04-16T19:42:02Z",
						DNS:          "127.0.0.1",
						Etag:         "0000000000000000",
						JavaKeystore: "/u3+7QAAAAIAAAABAAAAAQAE...",
						KeyAlias:     "saml",
						KeyPassword:  "pass123",
						UniqueID:     "48eef7e9-aadc-3a01-813b-54b06abb8d80",
						UUID:         "25e0e516-a30f-49b0-bcbc-b2fb9ee96c90",
					},
				}
				serviceProvider, err = clientset.ServiceProvider().Create(payload)
			})

			It("should have the requested ServiceProvider", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(serviceProvider.UniqueID).To(Equal("48eef7e9-aadc-3a01-813b-54b06abb8d80"))
			})
		})
	})

	Context("#Delete", func() {
		Context("with no params", func() {
			var (
				err error
			)

			BeforeEach(func() {
				err = clientset.ServiceProvider().Delete()
			})

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
