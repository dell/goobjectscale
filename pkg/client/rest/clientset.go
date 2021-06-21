package rest

import (
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/alertpolicies"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/crr"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/buckets"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/objectuser"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/objmt"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/tenants"
)

// ClientSet is a set of clients for each API section
type ClientSet struct {
	client        *client.Client
	buckets       api.BucketsInterface
	objectUser    api.ObjectUserInterface
	alertPolicies api.AlertPoliciesInterface
	tenants       api.TenantsInterface
	objmt         api.ObjmtInterface
	crr           api.CRRInterface
}

// Returns a new client set based on the provided REST client parameters
func NewClientSet(u string, p string, e string, h *http.Client, overrideHdr bool) *ClientSet {
	c := &client.Client{
		Username:       u,
		Password:       p,
		Endpoint:       e,
		HTTPClient:     h,
		OverrideHeader: overrideHdr,
	}
	return &ClientSet{
		client:        c,
		buckets:       &buckets.Buckets{Client: c},
		objectUser:    &objectuser.ObjectUser{Client: c},
		alertPolicies: &alertpolicies.AlertPolicies{Client: c},
		tenants:       &tenants.Tenants{Client: c},
		objmt:         &objmt.Objmt{Client: c},
		crr:           &crr.CRR{Client: c},
	}
}

// Client returns the REST client used in the clientset
func (c *ClientSet) Client() *client.Client {
	return c.client
}

// Buckets implements the client API
func (c *ClientSet) Buckets() api.BucketsInterface {
	return c.buckets
}

// ObjectUser implements the client API
func (c *ClientSet) ObjectUser() api.ObjectUserInterface {
	return c.objectUser
}

// AlertPolicies implements the client API
func (c *ClientSet) AlertPolicies() api.AlertPoliciesInterface {
	return c.alertPolicies
}

// Tenants implements the client API
func (c *ClientSet) Tenants() api.TenantsInterface {
	return c.tenants
}

// ObjectMt implements the client API for objMT metrics
func (c *ClientSet) ObjectMt() api.ObjmtInterface {
	return c.objmt
}

// ObjectMt implements the client API for objMT metrics
func (c *ClientSet) CRR() api.CRRInterface {
	return c.crr
}
