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
				"root",
				"changme",
				"https://host",
				"https://testgateway",
				newTestHTTPClient(captures, false),
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

		It("should login", func() {
			Expect(captures["login"]).To(Equal(1))
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
				"root",
				"changme",
				":not:a:valid:url",
				"https://testgateway",
				newTestHTTPClient(captures, false),
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
				"root",
				"changme",
				"https://host",
				"https://testgateway",
				newTestHTTPClient(captures, false),
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

	Context("with an auth failure", func() {
		var (
			err      error
			captures map[string]interface{}
		)

		BeforeEach(func() {
			captures = map[string]interface{}{}
			clientset = rest.NewClientSet(
				"root",
				"changme",
				"https://host",
				"https://testgateway",
				newTestHTTPClient(captures, true),
				false,
			)
			err = clientset.Client().MakeRemoteCall(client.Request{
				Method:      http.MethodGet,
				Path:        "/test",
				ContentType: client.ContentTypeJSON,
			}, nil)
		})

		It("shouldn't error", func() {
			Expect(err.Error()).To(Equal("auth failure"))
		})

		It("should attempt 3 logins", func() {
			Expect(captures["login"]).To(Equal(1))
		})

		It("should not make the test call", func() {
			Expect(captures["test"]).To(BeNil())
		})

		It("should not make an unfound call", func() {
			Expect(captures["notfound"]).To(BeNil())
		})
	})
})

func newTestHTTPClient(captures map[string]interface{}, authFailure bool) *http.Client {
	return testutils.NewTestClient(func(req *http.Request) *http.Response {
		header := make(http.Header)
		switch req.URL.String() {
		case "https://testgateway/mgmt/login":
			testutils.IncrCapture(captures, "login")
			switch authFailure {
			case true:
				return &http.Response{
					StatusCode: 401,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("auth failure"))),
					Header:     header,
				}
			default:
				header.Set("X-Sds-Auth-Token", "TESTTOKEN")
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("OK"))),
					Header:     header,
				}
			}
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
