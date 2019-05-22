package buckets

import (
	"fmt"
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

var _ api.BucketsInterface = (*Buckets)(nil)

// Buckets is a REST implementation of the Buckets interface
type Buckets struct {
	Client *client.Client
}

// Get implements the buckets interface
func (b *Buckets) Get(name string, params map[string]string) (*model.Bucket, error) {
	req := client.Request{
		Method: http.MethodGet,
		Path: fmt.Sprintf("/object/bucket/%s/info", name),
		ContentType: client.ContentTypeXML,
		Params: params,
	}
	bucket := &model.BucketInfo{}
	err := b.Client.MakeRemoteCall(req, bucket)
	if err != nil {
		return nil, err
	}
	return &bucket.Bucket, nil
}

// List implements the buckets interface
func (b *Buckets) List(params map[string]string) (*model.BucketList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/object/bucket",
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	bucketList := &model.BucketList{}
	err := b.Client.MakeRemoteCall(req, bucketList)
	if err != nil {
		return nil, err
	}
	return bucketList, nil
}
