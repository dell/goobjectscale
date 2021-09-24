package client_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/testutils"
)

func TestRest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest client Spec")
}

var _ = Describe("Rest client", func() {

	var (
		clientset *rest.ClientSet
	)

	Context("with a valid user", func() {

		var (
			err      error
			captures map[string]interface{}
		)

		BeforeEach(func() {
			captures = map[string]interface{}{}
			clientset = rest.NewClientSet(
				"https://host",
				"OSTOKEN-eyJ4NXUiOi",
				newTestHTTPClient(captures),
				false,
			)
			err = clientset.Client().MakeRemoteCall(client.Request{
				Method:      http.MethodGet,
				Path:        "/test",
				ContentType: client.ContentTypeJSON,
			}, nil)
		})

		It("should not error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("should query test route", func() {
			Expect(captures["test"]).To(Equal(1))
		})

		It("should not query an unknown route", func() {
			Expect(captures["notfound"]).To(BeNil())
		})

		It("should return an OK response", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("with an invalid endpoint", func() {
		var (
			err      error
			captures map[string]interface{}
		)

		BeforeEach(func() {
			captures = map[string]interface{}{}
			clientset = rest.NewClientSet(
				":not:a:valid:url",
				"OSTOKEN-eyJ4NXUiOi",
				newTestHTTPClient(captures),
				true,
			)
			err = clientset.Client().MakeRemoteCall(client.Request{
				Method:      http.MethodGet,
				Path:        "",
				ContentType: client.ContentTypeJSON,
			}, nil)
		})

		It("should return an error", func() {
			e := "parse \":not:a:valid:url\": missing protocol scheme"
			Expect(err.Error()).To(Equal(e))
		})

		It("should not make any remote calls", func() {
			Expect(captures["login"]).To(BeNil())
			Expect(captures["test"]).To(BeNil())
			Expect(captures["notfound"]).To(BeNil())
		})
	})

	Context("with an invalid content type", func() {
		var (
			err      error
			captures map[string]interface{}
		)

		BeforeEach(func() {
			captures = map[string]interface{}{}
			clientset = rest.NewClientSet(
				"https://host",
				"OSTOKEN-eyJ4NXUiOi",
				newTestHTTPClient(captures),
				false,
			)
			err = clientset.Client().MakeRemoteCall(client.Request{
				Method:      http.MethodGet,
				Path:        "",
				ContentType: "NotAContentType",
			}, nil)
		})

		It("should return an error", func() {
			e := "invalid content-type"
			Expect(err.Error()).To(Equal(e))
		})

		It("should not make a login call", func() {
			Expect(captures["login"]).To(BeNil())
		})

		It("should not make the test call", func() {
			Expect(captures["test"]).To(BeNil())
		})

		It("should not make an unfound call", func() {
			Expect(captures["notfound"]).To(BeNil())
		})
	})
})

func newTestHTTPClient(captures map[string]interface{}) *http.Client {
	return testutils.NewTestClient(func(req *http.Request) *http.Response {
		header := make(http.Header)
		switch req.URL.String() {
		case "https://host/test":
			testutils.IncrCapture(captures, "test")
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("OK"))),
				Header:     header,
			}
		default:
			testutils.IncrCapture(captures, "notfound")
			return &http.Response{
				StatusCode: 404,
				Header:     make(http.Header),
			}
		}
	})
}
