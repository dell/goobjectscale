package rest

import (
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/buckets"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/objectuser"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/tenants"
)

// ClientSet is a set of clients for each API section
type ClientSet struct {
	client     *client.Client
	buckets    api.BucketsInterface
	objectUser api.ObjectUserInterface
	tenants    api.TenantsInterface
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
		client:     c,
		buckets:    &buckets.Buckets{Client: c},
		objectUser: &objectuser.ObjectUser{Client: c},
		tenants:    &tenants.Tenants{Client: c},
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

// Tenants implements the client API
func (c *ClientSet) Tenants() api.TenantsInterface {
	return c.tenants
}
