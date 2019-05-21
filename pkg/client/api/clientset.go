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
	// List returns a list of buckets in the ObjectScale instance buckets
	List(params map[string]string) (*model.BucketList, error)
}