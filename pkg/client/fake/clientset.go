package fake

import (
	"errors"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// ClientSet is a set of clients for each API section
type ClientSet struct {
	buckets api.BucketsInterface
}

// Returns a new client set based on the provided REST client parameters
func NewClientSet(objs ...interface{}) *ClientSet {
	var bucketList []model.Bucket
	for _, o := range objs {
		switch object := o.(type) {
		case *model.Bucket:
			bucketList = append(bucketList, *object)
		}
	}
	return &ClientSet{
		buckets: &Buckets{items: bucketList},
	}
}

// Buckets implements the client API
func (c *ClientSet) Buckets() api.BucketsInterface {
	return c.buckets
}

// Buckets implements the buckets API
type Buckets struct {
	items []model.Bucket
}

// List implements the buckets API
func (b *Buckets) List(params map[string]string) (*model.BucketList, error) {
	return &model.BucketList{Items: b.items}, nil
}

// Get implements the buckets API
func (b *Buckets) Get(name string, params map[string]string) (*model.Bucket, error) {
	for _, bucket := range b.items {
		if bucket.Name == name {
			return &bucket, nil
		}
	}
	return nil, errors.New("not found")
}

// Create implements the buckets API
func (b *Buckets) Create(createParam model.Bucket) (*model.Bucket, error) {
	for _, existingBucket := range b.items {
		if existingBucket.Namespace == createParam.Namespace && existingBucket.Name == createParam.Name	 {
			return nil, errors.New("duplicate found")
		}
	}
	b.items = append(b.items, createParam)
	return &createParam, nil
}

// Delete implements the buckets API
func (b *Buckets) Delete(name string, namespace string) error {
	for i, existingBucket := range b.items {
		if existingBucket.Name == name && existingBucket.Namespace == namespace {
			b.items = append(b.items[:i], b.items[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
