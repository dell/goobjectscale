package buckets

import (
	"net/http"
	"path"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

// Buckets is a REST implementation of the Buckets interface
type Buckets struct {
	Client *client.Client
}

// Get implements the buckets interface
func (b *Buckets) Get(name string, params map[string]string) (*model.Bucket, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path: path.Join("object", "bucket", name, "info"),
		ContentType: client.ContentTypeXML,
		Params:      params,
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

// Create implements the buckets interface
func (b *Buckets) Create(createParam *model.Bucket) (*model.Bucket, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        "/object/bucket",
		ContentType: client.ContentTypeXML,
		Body:        &model.BucketCreate{Bucket: *createParam},
	}
	bucket := &model.Bucket{}
	err := b.Client.MakeRemoteCall(req, bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

// Delete implements the buckets interface
func (b *Buckets) Delete(name string, namespace string) error {
	req := client.Request{
		Method: http.MethodPost,
		Path: path.Join("object", "bucket", name, "deactivate"),
		Params: map[string]string{"namespace": namespace},
		ContentType: client.ContentTypeJSON,
	}
	err := b.Client.MakeRemoteCall(req, nil)
	if err != nil {
		return err
	}
	return nil
}
