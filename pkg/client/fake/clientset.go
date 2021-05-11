package fake

import (
	"errors"
	"fmt"
	"time"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// ClientSet is a set of clients for each API section
type ClientSet struct {
	buckets    api.BucketsInterface
	objectUser api.ObjectUserInterface
	tenants    api.TenantsInterface
	objectMt   api.ObjmtInterface
	crr        api.CRRInterface
}

// NewClientSet returns a new client set based on the provided REST client parameters
func NewClientSet(objs ...interface{}) *ClientSet {
	var (
		policy                      = make(map[string]string)
		bucketList                  []model.Bucket
		blobUsers                   []model.BlobUser
		userSecrets                 []UserSecret
		userInfoList                []UserInfo
		tenantList                  []model.Tenant
		accountBillingInfoList      *model.AccountBillingInfoList
		accountBillingSampleList    *model.AccountBillingSampleList
		bucketBillingInfoList       *model.BucketBillingInfoList
		bucketBillingSampleList     *model.BucketBillingSampleList
		bucketBillingPerfList       *model.BucketPerfDataList
		bucketReplicationInfoList   *model.BucketReplicationInfoList
		bucketReplicationSampleList *model.BucketReplicationSampleList
		storeBillingInfoList        *model.StoreBillingInfoList
		storeBillingSampleList      *model.StoreBillingSampleList
		storeReplicationDataList    *model.StoreReplicationDataList
		crr                         *model.CRR
	)
	for _, o := range objs {
		switch object := o.(type) {
		case *model.Bucket:
			bucketList = append(bucketList, *object)
		case *BucketPolicy:
			policy[fmt.Sprintf("%s/%s", object.BucketName, object.Namespace)] = object.Policy
		case *model.BlobUser:
			blobUsers = append(blobUsers, *object)
		case *UserSecret:
			userSecrets = append(userSecrets, *object)
		case *UserInfo:
			userInfoList = append(userInfoList, *object)
		case *model.Tenant:
			tenantList = append(tenantList, *object)
		case *model.AccountBillingInfoList:
			accountBillingInfoList = object
		case *model.AccountBillingSampleList:
			accountBillingSampleList = object
		case *model.BucketBillingInfoList:
			bucketBillingInfoList = object
		case *model.BucketBillingSampleList:
			bucketBillingSampleList = object
		case *model.BucketPerfDataList:
			bucketBillingPerfList = object
		case *model.BucketReplicationInfoList:
			bucketReplicationInfoList = object
		case *model.BucketReplicationSampleList:
			bucketReplicationSampleList = object
		case *model.StoreBillingInfoList:
			storeBillingInfoList = object
		case *model.StoreBillingSampleList:
			storeBillingSampleList = object
		case *model.StoreReplicationDataList:
			storeReplicationDataList = object
		case *model.CRR:
			crr = object
		default:
			panic(fmt.Sprintf("Fake client set doesn't support %T type", o))
		}
	}

	return &ClientSet{
		buckets: &Buckets{
			items:  bucketList,
			policy: policy,
		},
		objectUser: NewObjectUsers(blobUsers, userSecrets, userInfoList),
		tenants: &Tenants{
			items: tenantList,
		},
		objectMt: &Objmt{
			accountBillingInfoList:      accountBillingInfoList,
			accountBillingSampleList:    accountBillingSampleList,
			bucketBillingInfoList:       bucketBillingInfoList,
			bucketBillingSampleList:     bucketBillingSampleList,
			bucketBillingPerfList:       bucketBillingPerfList,
			bucketReplicationInfoList:   bucketReplicationInfoList,
			bucketReplicationSampleList: bucketReplicationSampleList,
			storeBillingInfoList:        storeBillingInfoList,
			storeBillingSampleList:      storeBillingSampleList,
			storeReplicationDataList:    storeReplicationDataList,
		},
		crr: &CRR{
			Config: crr,
		},
	}
}

// CRR implements the client API
func (c *ClientSet) CRR() api.CRRInterface {
	return c.crr
}

// Buckets implements the client API
func (c *ClientSet) Buckets() api.BucketsInterface {
	return c.buckets
}

type BucketPolicy struct {
	BucketName string
	Policy     string
	Namespace  string
}

// Tenants implements the client API.
func (c *ClientSet) Tenants() api.TenantsInterface {
	return c.tenants
}

// ObjectUser implements the client API.
func (c *ClientSet) ObjectUser() api.ObjectUserInterface {
	return c.objectUser
}

// ObjectMt implements the client API for objMT metrics
func (c *ClientSet) ObjectMt() api.ObjmtInterface {
	return c.objectMt
}

// UserSecret make easiest passing secret about users.
type UserSecret struct {
	UID    string
	Secret *model.ObjectUserSecret
}

// UserInfo make easiest passing info about users.
type UserInfo struct {
	UID  string
	Info *model.ObjectUserInfo
}

// ObjectUsers contains information about object users to be used in fake client set.
type ObjectUsers struct {
	Users    *model.ObjectUserList
	Secrets  map[string]*model.ObjectUserSecret
	InfoList map[string]*model.ObjectUserInfo
}

// NewObjectUsers returns initialized ObjectUsers.
func NewObjectUsers(blobUsers []model.BlobUser, userSecrets []UserSecret, userInfoList []UserInfo) *ObjectUsers {
	mappedUserSecrets := map[string]*model.ObjectUserSecret{}
	mappedUserInfoList := map[string]*model.ObjectUserInfo{}
	for _, s := range userSecrets {
		mappedUserSecrets[s.UID] = s.Secret
	}
	for _, i := range userInfoList {
		mappedUserInfoList[i.UID] = i.Info
	}
	return &ObjectUsers{
		&model.ObjectUserList{
			BlobUser: blobUsers,
		},
		mappedUserSecrets,
		mappedUserInfoList,
	}
}

// List returns a list of object users.
func (o *ObjectUsers) List(_ map[string]string) (*model.ObjectUserList, error) {
	return o.Users, nil
}

// GetSecret returns information about object user secrets.
func (o *ObjectUsers) GetSecret(uid string, _ map[string]string) (*model.ObjectUserSecret, error) {
	if _, ok := o.Secrets[uid]; !ok {
		return nil, fmt.Errorf("secret for %s is not found", uid)
	}
	return o.Secrets[uid], nil
}

// CreateSecret will create a specific secret
func (o *ObjectUsers) CreateSecret(uid string, req model.ObjectUserSecretKeyCreateReq, _ map[string]string) (*model.ObjectUserSecretKeyCreateRes, error) {
	if _, ok := o.Secrets[uid]; !ok {
		o.Secrets[uid] = &model.ObjectUserSecret{
			SecretKey1: req.SecretKey,
		}
		return &model.ObjectUserSecretKeyCreateRes{
			SecretKey: req.SecretKey,
		}, nil
	}

	switch {
	case o.Secrets[uid].SecretKey1 != "" && o.Secrets[uid].SecretKey2 != "":
		return nil, fmt.Errorf("user %s already has 2 valid keys", uid)
	case o.Secrets[uid].SecretKey1 != "":
		o.Secrets[uid].SecretKey2 = req.SecretKey
		return &model.ObjectUserSecretKeyCreateRes{
			SecretKey: req.SecretKey,
		}, nil
	default:
		o.Secrets[uid].SecretKey1 = req.SecretKey
		return &model.ObjectUserSecretKeyCreateRes{
			SecretKey: req.SecretKey,
		}, nil
	}
}

// DeleteSecret will delete a specific secret
func (o *ObjectUsers) DeleteSecret(uid string, req model.ObjectUserSecretKeyDeleteReq, _ map[string]string) error {
	if _, ok := o.Secrets[uid]; !ok {
		return fmt.Errorf("user %s not found", uid)
	}
	switch req.SecretKey {
	case o.Secrets[uid].SecretKey1:
		o.Secrets[uid].SecretKey1 = o.Secrets[uid].SecretKey2
		clearSecretKey2(o.Secrets[uid])
		return nil
	case o.Secrets[uid].SecretKey2:
		clearSecretKey2(o.Secrets[uid])
		return nil
	default:
		return fmt.Errorf("user %s secret key not found", uid)
	}
}

func clearSecretKey2(key *model.ObjectUserSecret) {
	key.SecretKey2 = ""
	key.KeyExpiryTimestamp2 = ""
	key.KeyExpiryTimestamp2 = ""
}

// GetSecret returns information about object user secrets.
func (o *ObjectUsers) GetInfo(uid string, _ map[string]string) (*model.ObjectUserInfo, error) {
	if _, ok := o.InfoList[uid]; !ok {
		return nil, fmt.Errorf("info for %s is not found", uid)
	}
	return o.InfoList[uid], nil
}

// Tenants implements the tenants API
type Tenants struct {
	items []model.Tenant
}

// Get implements the tenants API
func (t *Tenants) Get(id string, _ map[string]string) (*model.Tenant, error) {
	for _, tenant := range t.items {
		if tenant.ID == id {
			return &tenant, nil
		}
	}
	return nil, errors.New("not found")
}

// Get implements the tenants API
func (t *Tenants) List(_ map[string]string) (*model.TenantList, error) {
	return &model.TenantList{Items: t.items}, nil
}

// Buckets implements the buckets API
type Buckets struct {
	items  []model.Bucket
	policy map[string]string
}

// List implements the buckets API
func (b *Buckets) List(_ map[string]string) (*model.BucketList, error) {
	return &model.BucketList{Items: b.items}, nil
}

// Get implements the buckets API
func (b *Buckets) Get(name string, _ map[string]string) (*model.Bucket, error) {
	for _, bucket := range b.items {
		if bucket.Name == name {
			return &bucket, nil
		}
	}
	return nil, errors.New("not found")
}

// GetPolicy implements the buckets API
func (b *Buckets) GetPolicy(bucketName string, param map[string]string) (string, error) {
	if policy, ok := b.policy[fmt.Sprintf("%s/%s", bucketName, param["namespace"])]; ok {
		return policy, nil
	}
	return "", nil
}

// DeletePolicy implements the buckets API
func (b *Buckets) DeletePolicy(bucketName string, param map[string]string) error {
	found := false
	for _, bucket := range b.items {
		if bucket.Name == bucketName {
			found = true
			break
		}
	}
	if found {
		delete(b.policy, fmt.Sprintf("%s/%s", bucketName, param["namespace"]))
		return nil
	} else {
		return errors.New("bucket not found")
	}
}

// UpdatePolicy implements the buckets API
func (b *Buckets) UpdatePolicy(bucketName string, policy string, param map[string]string) error {
	found := false
	for _, bucket := range b.items {
		if bucket.Name == bucketName {
			found = true
			break
		}
	}
	if found {
		b.policy[fmt.Sprintf("%s/%s", bucketName, param["namespace"])] = policy
		return nil
	} else {
		return errors.New("bucket not found")
	}
}

// Create implements the buckets API
func (b *Buckets) Create(createParam model.Bucket) (*model.Bucket, error) {
	for _, existingBucket := range b.items {
		if existingBucket.Namespace == createParam.Namespace && existingBucket.Name == createParam.Name {
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

// GetQuota gets the quota for the given bucket and namespace.
func (b *Buckets) GetQuota(bucketName string, _ string) (*model.BucketQuotaInfo, error) {
	for _, bucket := range b.items {
		if bucket.Name == bucketName {
			return &model.BucketQuotaInfo{
				BucketQuota: model.BucketQuota{
					BucketName:       bucket.Name,
					Namespace:        bucket.Namespace,
					NotificationSize: bucket.NotificationSize,
					BlockSize:        bucket.BlockSize,
				},
			}, nil
		}
	}

	return nil, errors.New("not found")
}

// UpdateQuota updates the quota for the specified bucket.
func (b *Buckets) UpdateQuota(bucketQuota model.BucketQuotaUpdate) error {
	for i := 0; i < len(b.items); i++ {
		if b.items[i].Name == bucketQuota.BucketName {
			b.items[i].BlockSize = bucketQuota.BlockSize
			b.items[i].NotificationSize = bucketQuota.NotificationSize
			return nil
		}
	}
	return errors.New("not found")
}

// DeleteQuota deletes the quota setting for the given bucket and namespace.
func (b *Buckets) DeleteQuota(bucketName string, _ string) error {
	for i := 0; i < len(b.items); i++ {
		if b.items[i].Name == bucketName {
			b.items[i].BlockSize = -1
			b.items[i].NotificationSize = -1
			return nil
		}
	}
	return errors.New("not found")
}

// Objmt is a fake (mocked) implementation of the Objmt interface
type Objmt struct {
	accountBillingInfoList      *model.AccountBillingInfoList
	accountBillingSampleList    *model.AccountBillingSampleList
	bucketBillingInfoList       *model.BucketBillingInfoList
	bucketBillingSampleList     *model.BucketBillingSampleList
	bucketBillingPerfList       *model.BucketPerfDataList
	bucketReplicationInfoList   *model.BucketReplicationInfoList
	bucketReplicationSampleList *model.BucketReplicationSampleList
	storeBillingInfoList        *model.StoreBillingInfoList
	storeBillingSampleList      *model.StoreBillingSampleList
	storeReplicationDataList    *model.StoreReplicationDataList
}

// GetStoreBillingInfo returns billing info metrics for object store
func (mt *Objmt) GetStoreBillingInfo(_ map[string]string) (*model.StoreBillingInfoList, error) {
	return mt.storeBillingInfoList, nil
}

// GetStoreBillingSample returns billing sample (time-window) metrics for object store
func (mt *Objmt) GetStoreBillingSample(_ map[string]string) (*model.StoreBillingSampleList, error) {
	return mt.storeBillingSampleList, nil
}

// GetStoreReplicationData returns CRR metrics for defined object stores
func (mt *Objmt) GetStoreReplicationData(_ []string, _ map[string]string) (*model.StoreReplicationDataList, error) {
	return mt.storeReplicationDataList, nil
}

// GetAccountBillingInfo returns billing info metrics for defined accounts
func (mt *Objmt) GetAccountBillingInfo(_ []string, _ map[string]string) (*model.AccountBillingInfoList, error) {
	return mt.accountBillingInfoList, nil
}

// GetAccountBillingSample returns billing sample (time-window) metrics for defined accounts
func (mt *Objmt) GetAccountBillingSample(_ []string, _ map[string]string) (*model.AccountBillingSampleList, error) {
	return mt.accountBillingSampleList, nil
}

// GetBucketBillingInfo returns billing info metrics for defined buckets and account
func (mt *Objmt) GetBucketBillingInfo(_ string, _ []string, _ map[string]string) (*model.BucketBillingInfoList, error) {
	return mt.bucketBillingInfoList, nil
}

// GetBucketBillingSample returns billing sample (time-window) metrics for defined buckets and account
func (mt *Objmt) GetBucketBillingSample(_ string, _ []string, _ map[string]string) (*model.BucketBillingSampleList, error) {
	return mt.bucketBillingSampleList, nil
}

// GetBucketBillingPerf returns performance metrics for defined buckets and account
func (mt *Objmt) GetBucketBillingPerf(_ string, _ []string, _ map[string]string) (*model.BucketPerfDataList, error) {
	return mt.bucketBillingPerfList, nil
}

// GetReplicationInfo returns billing info metrics for defined replication pairs and account
func (mt *Objmt) GetReplicationInfo(_ string, _ [][]string, _ map[string]string) (*model.BucketReplicationInfoList, error) {
	return mt.bucketReplicationInfoList, nil
}

// GetReplicationSample returns billing sample (time-window) metrics for defined replication pairs and account
func (mt *Objmt) GetReplicationSample(_ string, _ [][]string, _ map[string]string) (*model.BucketReplicationSampleList, error) {
	return mt.bucketReplicationSampleList, nil
}

// CRR implements the crr API
type CRR struct {
	Config *model.CRR
}

// PauseReplication implements the CRR API
func (c *CRR) PauseReplication(destObjectScale string, destObjectStore string, durationMills int, _ map[string]string) error {
	c.Config.DestObjectScale = destObjectScale
	c.Config.DestObjectStore = destObjectStore
	c.Config.PauseStartMills = int(time.Millisecond)
	c.Config.PauseEndMills = c.Config.PauseStartMills + durationMills
	return nil
}

// SuspendReplication implements the CRR API
func (c *CRR) SuspendReplication(destObjectScale string, destObjectStore string, _ map[string]string) error {
	c.Config.DestObjectScale = destObjectScale
	c.Config.DestObjectStore = destObjectStore
	c.Config.SuspendStartMills = int(time.Millisecond)
	return nil
}

// ResumeReplication implements the CRR API
func (c *CRR) ResumeReplication(destObjectScale string, destObjectStore string, _ map[string]string) error {
	c.Config.DestObjectScale = destObjectScale
	c.Config.DestObjectStore = destObjectStore
	c.Config.SuspendStartMills = 0
	return nil
}

// ThrottleReplication implements the CRR API
func (c *CRR) ThrottleReplication(destObjectScale string, destObjectStore string, mbPerSecond int, _ map[string]string) error {
	c.Config.DestObjectScale = destObjectScale
	c.Config.DestObjectStore = destObjectStore
	c.Config.ThrottleBandwidth = mbPerSecond
	return nil
}

// Get implements the CRR API
func (c *CRR) Get(destObjectScale string, destObjectStore string, _ map[string]string) (*model.CRR, error) {
	c.Config.DestObjectScale = destObjectScale
	c.Config.DestObjectStore = destObjectStore
	return c.Config, nil
}
