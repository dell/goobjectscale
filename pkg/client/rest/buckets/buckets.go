package buckets

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

var _ api.BucketsInterface = (*Buckets)(nil)

type Buckets struct {
	Client *client.Client
}

func (b *Buckets) List(params map[string]string) (*model.BucketList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/object/bucket",
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	resp, err := b.Client.MakeRemoteCall(req)
	if err != nil {
		return nil, err
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	retval := &model.BucketList{}
	if err = xml.Unmarshal(body, retval); err !=  nil {
		return nil, err
	}
	return retval, nil
}
