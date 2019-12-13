// Package objectuser contains API to work with Object user resource of object scale.
package objectuser

import (
	"fmt"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"net/http"
)

// ObjectUser is a REST implementation of the object user interface
type ObjectUser struct {
	Client *client.Client
}

// GetInfo returns information about an object user within the ObjectScale object store.
func (o *ObjectUser) GetInfo(uid string, params map[string]string) (*model.ObjectUserInfo, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("object/users/%s/info", uid),
		ContentType: client.ContentTypeJSON,
		Params:      params,
	}
	ou := &model.ObjectUserInfo{}
	err := o.Client.MakeRemoteCall(req, ou)
	if err != nil {
		return nil, err
	}
	return ou, nil
}

// GetSecret returns information about object user secrets.
func (o *ObjectUser) GetSecret(uid string, params map[string]string) (*model.ObjectUserSecret, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("/object/user-secret-keys/%s", uid),
		ContentType: client.ContentTypeJSON,
		Params:      params,
	}
	ou := &model.ObjectUserSecret{}
	err := o.Client.MakeRemoteCall(req, ou)
	if err != nil {
		return nil, err
	}
	return ou, nil
}

// List returns a list of object users within the ObjectScale object store.
func (o *ObjectUser) List(params map[string]string) (*model.ObjectUserList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "object/users",
		ContentType: client.ContentTypeJSON,
		Params:      params,
	}
	ouList := &model.ObjectUserList{}
	err := o.Client.MakeRemoteCall(req, ouList)
	if err != nil {
		return nil, err
	}
	return ouList, nil
}
