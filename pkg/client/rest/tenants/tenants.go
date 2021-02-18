package tenants

import (
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

// Tenants is a REST implementation of the Tenants interface
type Tenants struct {
	Client *client.Client
}

// List implements the tenants interface
func (t *Tenants) List(params map[string]string) (*model.TenantList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/object/tenants",
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	tenantList := &model.TenantList{}
	err := t.Client.MakeRemoteCall(req, tenantList)
	if err != nil {
		return nil, err
	}
	return tenantList, nil
}
