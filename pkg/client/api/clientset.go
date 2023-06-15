// Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"

	"github.com/dell/goobjectscale/pkg/client/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name (ClientSet)|(([A-Z][A-Za-z]+)Interface)

// ClientSet represents a client interface of supported resources
type ClientSet interface {
	// Buckets returns a bucket client interface
	Buckets() BucketsInterface
	ObjectUser() ObjectUserInterface
	Tenants() TenantsInterface
	ObjectMt() ObjmtInterface
	AlertPolicies() AlertPoliciesInterface
	CRR() CRRInterface
	Status() StatusInterface
	FederatedObjectStores() FederatedObjectStoresInterface
}

// StatusInterface represents status resource client interface
type StatusInterface interface {
	// GetRebuildStatus returns rebuild status of an ObjectScale object store
	GetRebuildStatus(ctx context.Context, objStoreName, ssPodName, ssPodNameSpace, level string, params map[string]string) (*model.RebuildInfo, error)
}

// FederatedObjectStoresInterface represents a replication store client interface.
type FederatedObjectStoresInterface interface {
	// List returns a list of federated object stores
	List(ctx context.Context, params map[string]string) (*model.FederatedObjectStoreList, error)
}

// BucketsInterface represents a bucket resource client interface.
type BucketsInterface interface {
	// List returns a list of buckets within the ObjectScale object store
	List(ctx context.Context, params map[string]string) (*model.BucketList, error)

	// GetPolicy returns current policy for a bucket as json string.
	GetPolicy(ctx context.Context, bucketName string, param map[string]string) (string, error)

	// UpdatePolicy adds/replaces new policy to the existing bucket.
	UpdatePolicy(ctx context.Context, bucketName string, policy string, param map[string]string) error

	// DeletePolicy removes a policy from an existing bucket.
	DeletePolicy(ctx context.Context, bucketName string, param map[string]string) error

	// Get returns a bucket in the ObjectScale object store
	Get(ctx context.Context, name string, params map[string]string) (*model.Bucket, error)

	// Create creates a new bucket in the ObjectScale object store
	Create(ctx context.Context, createParam model.Bucket) (*model.Bucket, error)

	// Delete deletes bucket from the ObjectScale object store
	Delete(ctx context.Context, name string, namespace string, emptyBucket bool) error

	// GetQuota Gets the quota for the given bucket and namespace.
	GetQuota(ctx context.Context, bucketName string, namespace string) (*model.BucketQuotaInfo, error)

	// UpdateQuota Updates the quota for the specified bucket.
	UpdateQuota(ctx context.Context, bucketQuota model.BucketQuotaUpdate) error

	// DeleteQuota Deletes the quota setting for the given bucket and namespace.
	DeleteQuota(ctx context.Context, bucketName string, namespace string) error
}

// ObjectUserInterface represents an object user resource client interface.
type ObjectUserInterface interface {
	// List returns a list of object users within the ObjectScale object store.
	List(ctx context.Context, params map[string]string) (*model.ObjectUserList, error)

	// GetInfo returns information about an object user within the ObjectScale object store.
	GetInfo(ctx context.Context, uid string, params map[string]string) (*model.ObjectUserInfo, error)

	// GetSecret returns information about object user secrets.
	GetSecret(ctx context.Context, uid string, params map[string]string) (*model.ObjectUserSecret, error)

	// CreateSecret creates a secret for an object user within the Objectscale object store
	CreateSecret(ctx context.Context, uid string, req model.ObjectUserSecretKeyCreateReq, params map[string]string) (*model.ObjectUserSecretKeyCreateRes, error)

	// DeleteSecret delete a secret for an object user within the Objectscale object store
	DeleteSecret(ctx context.Context, uid string, req model.ObjectUserSecretKeyDeleteReq, params map[string]string) error
}

// AlertPoliciesInterface represents a alert policy resource client interface.
type AlertPoliciesInterface interface {
	// List returns a list of alert policies within the ObjectScale object store.
	List(ctx context.Context, params map[string]string) (*model.AlertPolicies, error)

	// Get returns the Alert Policy
	Get(ctx context.Context, policyName string) (*model.AlertPolicy, error)

	// Create creates an Alert Policy and returns it
	Create(ctx context.Context, payload model.AlertPolicy) (*model.AlertPolicy, error)

	// Delete deletes an Alert Policy
	Delete(ctx context.Context, policyName string) error

	// Update updates an Alert Policy and returns it
	Update(ctx context.Context, payload model.AlertPolicy, policyName string) (*model.AlertPolicy, error)
}

// TenantsInterface represents an tenant resource client interface.
type TenantsInterface interface {
	// List returns a list of tenants within the ObjectScale object store.
	List(ctx context.Context, params map[string]string) (*model.TenantList, error)

	// Get returns an account tenant in the ObjectScale object store
	Get(ctx context.Context, name string, params map[string]string) (*model.Tenant, error)

	// Create creates a tenant and returns it
	Create(ctx context.Context, payload model.TenantCreate) (*model.Tenant, error)

	// Delete deletes a tenant
	Delete(ctx context.Context, name string) error

	// Update updates a specific tenant (currently only default bucket block size and alias fields supported)
	Update(ctx context.Context, payload model.TenantUpdate, name string) error

	// GetQuota gets the quota of a tenant
	GetQuota(ctx context.Context, name string, params map[string]string) (*model.TenantQuota, error)

	// DeleteQuota deletes the quota of a tenant
	DeleteQuota(ctx context.Context, name string) error

	// SetQuota sets the quota of a tenant
	SetQuota(ctx context.Context, name string, payload model.TenantQuotaSet) error
}

// ObjmtInterface represents an interface for objMT service metrics.
type ObjmtInterface interface {
	// GetAccountBillingInfo returns billing info metrics for defined accounts
	GetAccountBillingInfo(ctx context.Context, ids []string, params map[string]string) (*model.AccountBillingInfoList, error)

	// GetAccountBillingSample returns billing sample (time-window) metrics for defined accounts
	GetAccountBillingSample(ctx context.Context, ids []string, params map[string]string) (*model.AccountBillingSampleList, error)

	// GetBucketBillingInfo returns billing info metrics for defined buckets and account
	GetBucketBillingInfo(ctx context.Context, account string, ids []string, params map[string]string) (*model.BucketBillingInfoList, error)

	// GetBucketBillingSample returns billing sample (time-window) metrics for defined buckets and account
	GetBucketBillingSample(ctx context.Context, account string, ids []string, params map[string]string) (*model.BucketBillingSampleList, error)

	// GetBucketBillingPerf returns performance metrics for defined buckets and account
	GetBucketBillingPerf(ctx context.Context, account string, ids []string, params map[string]string) (*model.BucketPerfDataList, error)

	// GetReplicationInfo returns billing info metrics for defined replication pairs and account
	GetReplicationInfo(ctx context.Context, account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationInfoList, error)

	// GetReplicationSample returns billing sample (time-window) metrics for defined replication pairs and account
	GetReplicationSample(ctx context.Context, account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationSampleList, error)

	// GetStoreBillingInfo returns billing info metrics for object store
	GetStoreBillingInfo(ctx context.Context, params map[string]string) (*model.StoreBillingInfoList, error)

	// GetStoreBillingSample returns billing sample (time-window) metrics for object store
	GetStoreBillingSample(ctx context.Context, params map[string]string) (*model.StoreBillingSampleList, error)

	// GetStoreReplicationData returns CRR metrics for defined object stores
	GetStoreReplicationData(ctx context.Context, ids []string, params map[string]string) (*model.StoreReplicationDataList, error)
}

// CRRInterface represents an interface for Cross Region Replication (CRR).
type CRRInterface interface {
	// PauseReplication temporarily pauses source and destination object stores' replication communication
	// pauses for the provided future epoch time in milliseconds
	PauseReplication(ctx context.Context, destObjectScale string, destObjectStore string, param map[string]string) error

	// SuspendReplication suspends source and destination object stores' replication communication
	SuspendReplication(ctx context.Context, destObjectScale string, destObjectStore string, param map[string]string) error

	// ResumeReplication resumes source and destination object stores' replication communication
	ResumeReplication(ctx context.Context, destObjectScale string, destObjectStore string, param map[string]string) error

	// UnthrottleReplication resumes resumes replication sans any configured throttle cap
	UnthrottleReplication(ctx context.Context, destObjectScale string, destObjectStore string, param map[string]string) error

	// ThrottleReplication throttles source and destination object stores' replication communication
	// throttles the provided MB per second
	ThrottleReplication(ctx context.Context, destObjectScale string, destObjectStore string, param map[string]string) error

	// Get returns the replication configuration regarding pause/resume/suspend/throttle information
	Get(ctx context.Context, destObjectScale string, destObjectStore string, param map[string]string) (*model.CRR, error)
}
