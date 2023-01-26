package alertpolicies_test

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
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func TestTenants(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AlertPolicies Spec")
}

var _ = Describe("AlertPolicies", func() {

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

	AfterSuite(func() {
		err = r.Stop()
		if err != nil {
			log.Fatal(err)
		}
	})

	Context("#List", func() {
		var (
			alertPolicies *model.AlertPolicies
			err           error
		)

		Context("with no params", func() {
			It("should have alertPolicies", func() {
				alertPolicies, err = clientset.AlertPolicies().List(map[string]string{})
				Expect(err).ToNot(HaveOccurred())
				Expect(len(alertPolicies.Items)).To(Equal(3))
			})
		})
	})

	Context("#Get", func() {
		Context("with no params", func() {
			var (
				alertPolicy *model.AlertPolicy
				err         error
			)

			It("should have the requested alertPolicy", func() {
				alertPolicy, err = clientset.AlertPolicies().Get("testPolicy")
				Expect(err).ToNot(HaveOccurred())
				Expect(alertPolicy.PolicyName).To(Equal("testPolicy"))
			})
		})
	})

	Context("#Create", func() {
		Context("with no params", func() {
			var (
				alertPolicy *model.AlertPolicy
				err         error
			)

			BeforeEach(func() {
				payload := model.AlertPolicy{
					XMLName:    xml.Name{},
					PolicyName: "testPolicy",
				}
				alertPolicy, err = clientset.AlertPolicies().Create(payload)
			})

			It("should have the requested alertPolicy", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(alertPolicy.PolicyName).To(Equal("testPolicy"))
			})
		})
	})

	Context("#Delete", func() {
		Context("with no params", func() {
			var (
				err error
			)

			BeforeEach(func() {
				err = clientset.AlertPolicies().Delete("testPolicy")
			})

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Context("#Update", func() {
		Context("with no params", func() {
			var (
				err error
			)

			BeforeEach(func() {
				payload := model.AlertPolicy{
					PolicyName: "testPolicy",
				}
				_, err = clientset.AlertPolicies().Update(payload, "testPolicy")
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
