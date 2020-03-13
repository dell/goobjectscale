package fake

import (
	"errors"
	"fmt"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/api"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// ClientSet is a set of clients for each API section
type ClientSet struct {
	buckets    api.BucketsInterface
	objectUser api.ObjectUserInterface
}

// NewClientSet returns a new client set based on the provided REST client parameters
func NewClientSet(objs ...interface{}) *ClientSet {
	var (
		policy       = make(map[string]string)
		bucketList   []model.Bucket
		blobUsers    []model.BlobUser
		userSecrets  []UserSecret
		userInfoList []UserInfo
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
	}
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

// ObjectUser implements the client API.
func (c *ClientSet) ObjectUser() api.ObjectUserInterface {
	return c.objectUser
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
func (o *ObjectUsers) List(params map[string]string) (*model.ObjectUserList, error) {
	return o.Users, nil
}

// GetSecret returns information about object user secrets.
func (o *ObjectUsers) GetSecret(uid string, params map[string]string) (*model.ObjectUserSecret, error) {
	if _, ok := o.Secrets[uid]; !ok {
		return nil, fmt.Errorf("secret for %s is not found", uid)
	}
	return o.Secrets[uid], nil
}

// CreateSecret will create a specific secret
func (o *ObjectUsers) CreateSecret(uid string, req model.ObjectUserSecretKeyCreateReq, params map[string]string) (*model.ObjectUserSecretKeyCreateRes, error) {
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
		return nil, fmt.Errorf("User %s already has 2 valid keys", uid)
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
func (o *ObjectUsers) DeleteSecret(uid string, req model.ObjectUserSecretKeyDeleteReq, params map[string]string) error {
	if _, ok := o.Secrets[uid]; !ok {
		return fmt.Errorf("User %s not found", uid)
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
		return fmt.Errorf("User %s secret key not found", uid)
	}
}

func clearSecretKey2(key *model.ObjectUserSecret) {
	key.SecretKey2 = ""
	key.KeyExpiryTimestamp2 = ""
	key.KeyExpiryTimestamp2 = ""
}

// GetSecret returns information about object user secrets.
func (o *ObjectUsers) GetInfo(uid string, params map[string]string) (*model.ObjectUserInfo, error) {
	if _, ok := o.InfoList[uid]; !ok {
		return nil, fmt.Errorf("info for %s is not found", uid)
	}
	return o.InfoList[uid], nil
}

// Buckets implements the buckets API
type Buckets struct {
	items  []model.Bucket
	policy map[string]string
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

// GetPolicy implements the buckets API
func (b *Buckets) GetPolicy(bucketName string, param map[string]string) (string, error) {
	if policy, ok := b.policy[fmt.Sprintf("%s/%s", bucketName, param["namespace"])]; ok {
		return policy, nil
	}
	return "", nil
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
