package client_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/testutils"
)

func TestRest(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"validUser":          testValidUser,
		"InvalidEndpoint":    testInvalidEndpoint,
		"InvalidContentType": testInvalidContentType,
		"FailedAuth":         testFailedAuth,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testValidUser(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, false),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "/test",
		ContentType: client.ContentTypeJSON,
	}, nil)

	require.NoError(t, err)
	assert.Equal(t, captures["login"], 1)
	assert.Equal(t, captures["test"], 1)
	assert.Nil(t, captures["notfound"])
}

func testInvalidEndpoint(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		":not:a:valid:url",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, false),
		true,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: client.ContentTypeJSON,
	}, nil)
	e := "parse \":not:a:valid:url\": missing protocol scheme"
	require.Error(t, err)
	assert.Equal(t, err.Error(), e)
	assert.Nil(t, captures["login"])
	assert.Nil(t, captures["test"])
	assert.Nil(t, captures["notfound"])
}

func testInvalidContentType(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, false),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "",
		ContentType: "NotAContentType",
	}, nil)

	assert.Equal(t, err.Error(), "invalid content-type")
	assert.Nil(t, captures["login"])
	assert.Nil(t, captures["test"])
	assert.Nil(t, captures["notfound"])
}

func testFailedAuth(t *testing.T) {
	captures := map[string]interface{}{}
	clientset := rest.NewClientSet(client.NewClient(
		"https://testserver",
		"https://testgateway",
		"testuser",
		"testpassword",
		newTestHTTPClient(captures, true),
		false,
	))
	err := clientset.Client().MakeRemoteCall(client.Request{
		Method:      http.MethodGet,
		Path:        "/test",
		ContentType: client.ContentTypeJSON,
	}, nil)

	assert.Equal(t, err.Error(), "auth failure")
	assert.Equal(t, captures["login"], 1)
	assert.Nil(t, captures["test"])
	assert.Nil(t, captures["notfound"])
}

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
		case "https://testserver/test":
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
