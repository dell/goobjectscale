package serviceprovider

import (
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

// ServiceProvider is a REST implementation of the ServiceProvider interface
type ServiceProvider struct {
	Client *client.Client
}

var _ api.ServiceProviderInterface = &ServiceProvider{}

// Get implements the ServiceProvider interface
func (sp *ServiceProvider) Get(params map[string]string) (*model.ServiceProvider, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/ecs-service-provider",
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	serviceprovider := &model.ServiceProvider{}
	err := sp.Client.MakeRemoteCall(req, serviceprovider)
	if err != nil {
		return nil, err
	}
	return serviceprovider, nil
}

// Create implements the ServiceProvider interface
func (sp *ServiceProvider) Create(payload model.ServiceProviderCreate) (*model.ServiceProvider, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        "/ecs-service-provider",
		ContentType: client.ContentTypeXML,
		Body:        payload,
	}
	serviceprovider := &model.ServiceProvider{}
	err := sp.Client.MakeRemoteCall(req, serviceprovider)
	if err != nil {
		return nil, err
	}
	return serviceprovider, nil
}

// Delete implements the ServiceProvider interface
func (sp *ServiceProvider) Delete() error {
	req := client.Request{
		Method:      http.MethodDelete,
		Path:        "/ecs-service-provider",
		ContentType: client.ContentTypeXML,
	}
	serviceprovider := &model.ServiceProvider{}
	err := sp.Client.MakeRemoteCall(req, serviceprovider)
	if err != nil {
		return err
	}
	return nil
}
