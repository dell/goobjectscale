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
			items: bucketList,
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

// Get implements the buckets API
func (b *Buckets) GetPolicy(bucketName string, param map[string]string) (string, error) {
	return b.policy[fmt.Sprintf("%s/%s", bucketName, param["namespace"])], nil
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
