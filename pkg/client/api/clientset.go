package api

import (
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// ClientSet represents a client interface of supported resources
type ClientSet interface {
	// Buckets returns a bucket client interface
	Buckets() BucketsInterface
}

// BucketsInterfaces represents a bucket resource client interface
type BucketsInterface interface {
	// List returns a list of buckets within the ObjectScale object store
	List(params map[string]string) (*model.BucketList, error)

	// Get returns a bucket in the ObjectScale object store
	Get(name string, params map[string]string) (*model.Bucket, error)

	// Create creates a new bucket in the ObjectScale object store
	Create(createParam *model.Bucket) (*model.Bucket, error)

	// Delete deletes bucket from the ObjectScale object store
	Delete(name string, namespace string) error
}