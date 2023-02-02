package alertpolicies_test

import (
	"encoding/xml"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}

func TestAlertpolicies(t *testing.T) {
	var (
		r   *recorder.Recorder
		err error
	)
	r, err = recorder.New("fixtures")
	if err != nil {
		log.Fatal(err)
	}

	r.AddHook(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
		return nil
	}, recorder.BeforeSaveHook)
	clientset := rest.NewClientSet(client.NewServiceClient(
		"https://testserver",
		"https://testgateway",
		"svc-objectscale-domain-c8",
		"objectscale-graphql-7d754f8499-ng4h6",
		"OSC234DSF223423",
		"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
		newRecordedHTTPClient(r),
		false,
	))
	for scenario, fn := range map[string]func(t *testing.T, clientset *rest.ClientSet){
		"list":   testList,
		"get":    testGet,
		"create": testCreate,
		"update": testUpdate,
		"delete": testDelete,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, clientset)
		})
	}
}

func testList(t *testing.T, clientset *rest.ClientSet) {
	alertPolicies, err := clientset.AlertPolicies().List(map[string]string{})
	require.NoError(t, err)
	assert.Equal(t, len(alertPolicies.Items), 3)
}

func testGet(t *testing.T, clientset *rest.ClientSet) {
	alertPolicy, err := clientset.AlertPolicies().Get("testPolicy")
	require.NoError(t, err)
	assert.Equal(t, alertPolicy.PolicyName, "testPolicy")
}

func testCreate(t *testing.T, clientset *rest.ClientSet) {
	payload := model.AlertPolicy{
		XMLName:    xml.Name{},
		PolicyName: "testPolicy",
	}
	alertPolicy, err := clientset.AlertPolicies().Create(payload)
	require.NoError(t, err)
	assert.Equal(t, alertPolicy.PolicyName, "testPolicy")
}

func testUpdate(t *testing.T, clientset *rest.ClientSet) {
	payload := model.AlertPolicy{
		PolicyName: "testPolicy",
	}
	_, err := clientset.AlertPolicies().Update(payload, "testPolicy")
	require.NoError(t, err)
}

func testDelete(t *testing.T, clientset *rest.ClientSet) {
	err := clientset.AlertPolicies().Delete("testPolicy")
	require.NoError(t, err)
}
