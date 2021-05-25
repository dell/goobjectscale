package api

import (
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// ClientSet represents a client interface of supported resources
type ClientSet interface {
	// Buckets returns a bucket client interface
	Buckets() BucketsInterface
	ObjectUser() ObjectUserInterface
	Tenants() TenantsInterface
	ObjectMt() ObjmtInterface
	CRR() CRRInterface
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

// TenantsInterface represents an tenant resource client interface.
type TenantsInterface interface {
	// List returns a list of tenants within the ObjectScale object store.
	List(params map[string]string) (*model.TenantList, error)

	// Get returns an account tenant in the ObjectScale object store
	Get(name string, params map[string]string) (*model.Tenant, error)

	// Create creates a tenant and returns it
	Create(payload model.TenantCreate) (*model.Tenant, error)

	// Delete deletes a tenant
	Delete(name string) error

	// Update updates a specific tenant (currently only default bucket block size and alias fields supported)
	Update(payload model.TenantUpdate, name string) error

	// GetQuota gets the quota of a tenant
	GetQuota(name string, params map[string]string) (*model.TenantQuota, error)

	// DeleteQuota deletes the quota of a tenant
	DeleteQuota(name string) error

	// SetQuota sets the quota of a tenant
	SetQuota(name string, payload model.TenantQuotaSet) error
}

// ObjectUserInterface represents an interface for objMT service metrics.
type ObjmtInterface interface {
	// GetAccountBillingInfo returns billing info metrics for defined accounts
	GetAccountBillingInfo(ids []string, params map[string]string) (*model.AccountBillingInfoList, error)

	// GetAccountBillingSample returns billing sample (time-window) metrics for defined accounts
	GetAccountBillingSample(ids []string, params map[string]string) (*model.AccountBillingSampleList, error)

	// GetBucketBillingInfo returns billing info metrics for defined buckets and account
	GetBucketBillingInfo(account string, ids []string, params map[string]string) (*model.BucketBillingInfoList, error)

	// GetBucketBillingSample returns billing sample (time-window) metrics for defined buckets and account
	GetBucketBillingSample(account string, ids []string, params map[string]string) (*model.BucketBillingSampleList, error)

	// GetBucketBillingPerf returns performance metrics for defined buckets and account
	GetBucketBillingPerf(account string, ids []string, params map[string]string) (*model.BucketPerfDataList, error)

	// GetReplicationInfo returns billing info metrics for defined replication pairs and account
	GetReplicationInfo(account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationInfoList, error)

	// GetReplicationSample returns billing sample (time-window) metrics for defined replication pairs and account
	GetReplicationSample(account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationSampleList, error)

	// GetStoreBillingInfo returns billing info metrics for object store
	GetStoreBillingInfo(params map[string]string) (*model.StoreBillingInfoList, error)

	// GetStoreBillingSample returns billing sample (time-window) metrics for object store
	GetStoreBillingSample(params map[string]string) (*model.StoreBillingSampleList, error)

	// GetStoreReplicationData returns CRR metrics for defined object stores
	GetStoreReplicationData(ids []string, params map[string]string) (*model.StoreReplicationDataList, error)
}

// CRRInterface represents an interface for Cross Region Replication (CRR)
type CRRInterface interface {
	// PauseReplication temporarily pauses source and destination object stores' replication communication
	// pauses for the provided milliseconds
	PauseReplication(destObjectScale string, destObjectStore string, durationMills int, param map[string]string) error

	// SuspendReplication suspends source and destination object stores' replication communication
	SuspendReplication(destObjectScale string, destObjectStore string, param map[string]string) error

	// ResumeReplication resumes source and destination object stores' replication communication
	ResumeReplication(destObjectScale string, destObjectStore string, param map[string]string) error

	// ThrottleReplication throttles source and destination object stores' replication communication
	// throttles the provided MB per second
	ThrottleReplication(destObjectScale string, destObjectStore string, mbPerSecond int, param map[string]string) error

	// Get returns the replication configuration regarding pause/resume/suspend/throttle information
	Get(destObjectScale string, destObjectStore string, param map[string]string) (*model.CRR, error)
}
