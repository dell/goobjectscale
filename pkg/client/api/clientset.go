package api

import (
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// ClientSet represents a client interface of supported resources
type ClientSet interface {
	// Buckets returns a bucket client interface
	Buckets() BucketsInterface
	ObjectUser() ObjectUserInterface
}

// BucketsInterfaces represents a bucket resource client interface
type BucketsInterface interface {
	// List returns a list of buckets within the ObjectScale object store
	List(params map[string]string) (*model.BucketList, error)

	// GetPolicy returns current policy for a bucket as json string.
	GetPolicy(bucketName string, param map[string]string) (string, error)

	// UpdatePolicy adds/replaces new policy to the existing bucket.
	UpdatePolicy(bucketName string, policy string, param map[string]string) error

	// DeletePolicy removes a policy from an existing bucket.
	DeletePolicy(bucketName string, param map[string]string) error

	// Get returns a bucket in the ObjectScale object store
	Get(name string, params map[string]string) (*model.Bucket, error)

	// Create creates a new bucket in the ObjectScale object store
	Create(createParam model.Bucket) (*model.Bucket, error)

	// Delete deletes bucket from the ObjectScale object store
	Delete(name string, namespace string) error

	// Gets the quota for the given bucket and namespace.
	GetQuota(bucketName string, namespace string) (*model.BucketQuotaInfo, error)

	// Updates the quota for the specified bucket.
	UpdateQuota(bucketQuota model.BucketQuotaUpdate) error

	// Deletes the quota setting for the given bucket and namespace.
	DeleteQuota(bucketName string, namespace string) error
}

// ObjectUserInterface represents an object user resource client interface.
type ObjectUserInterface interface {
	// List returns a list of object users within the ObjectScale object store.
	List(params map[string]string) (*model.ObjectUserList, error)

	// GetInfo returns information about an object user within the ObjectScale object store.
	GetInfo(uid string, params map[string]string) (*model.ObjectUserInfo, error)

	// GetSecret returns information about object user secrets.
	GetSecret(uid string, params map[string]string) (*model.ObjectUserSecret, error)

	// CreateSecret creates a secret for an object user within the Objectscale object store
	CreateSecret(uid string, req model.ObjectUserSecretKeyCreateReq, params map[string]string) (*model.ObjectUserSecretKeyCreateRes, error)

	// DeleteSecret delete a secret for an object user within the Objectscale object store
	DeleteSecret(uid string, req model.ObjectUserSecretKeyDeleteReq, params map[string]string) error
}
