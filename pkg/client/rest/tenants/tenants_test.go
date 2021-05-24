package tenants_test

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

func TestTenants(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tenants Spec")
}

var _ = Describe("Tenants", func() {

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
			tenants *model.TenantList
			err     error
		)

		Context("with no params", func() {
			BeforeEach(func() {
				tenants, err = clientset.Tenants().List(map[string]string{})
			})

			It("shouldn't error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have tenants", func() {
				Expect(len(tenants.Items)).To(Equal(1))
			})
		})
	})

	Context("#Get", func() {
		Context("with no params", func() {
			var (
				tenant *model.Tenant
				err    error
			)

			It("should have the requested tenant", func() {
				tenant, err = clientset.Tenants().Get("10d9817c-3696-4625-854e-82b21d8c0795", map[string]string{})
				Expect(err).ToNot(HaveOccurred())
				Expect(tenant.ID).To(Equal("10d9817c-3696-4625-854e-82b21d8c0795"))
			})
		})
	})

	Context("#Create", func() {
		Context("with no params", func() {
			var (
				tenant *model.Tenant
				err    error
			)

			BeforeEach(func() {
				payload := model.TenantCreate{
					XMLName:           xml.Name{},
					Alias:             "",
					AccountID:         "test-account",
					EncryptionEnabled: false,
					ComplianceEnabled: false,
					BucketBlockSize:   0,
				}
				tenant, err = clientset.Tenants().Create(payload)
			})

			It("should have the requested tenant", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(tenant.ID).To(Equal("test-account"))
			})
		})
	})

	Context("#Delete", func() {
		Context("with no params", func() {
			var (
				err    error
			)

			BeforeEach(func() {
				err = clientset.Tenants().Delete("test-account")
			})

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})


	Context("#Get Quota", func() {
		Context("with no params", func() {
			var (
				err    error
			)

			BeforeEach(func() {
				payload := model.TenantQuotaSet{
					XMLName:                 xml.Name{},
					BlockSize:               "5",
					NotificationSize:        "6",
					BlockSizeInCount:        "7",
					NotificationSizeInCount: "8",
				}
				err = clientset.Tenants().SetQuota("test-account", payload)
			})

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Context("#Get Quota", func() {
		Context("with no params", func() {
			var (
				err    error
				quota *model.TenantQuota
			)

			BeforeEach(func() {
				quota, err = clientset.Tenants().GetQuota("test-account", nil)
			})

			It("should not error", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(quota.BlockSize).To(Equal("5"))
			})
		})
	})

	Context("#Delete Quota", func() {
		Context("with no params", func() {
			var (
				err    error
			)

			BeforeEach(func() {
				err = clientset.Tenants().DeleteQuota("test-account")
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
