// Copyright © 2022 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package buckets

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

const (
	BucketServicePath = "/object/bucket"
	NamespaceParam    = "namespace"
	Info              = "info"
	ACL               = "acl"
)

// Buckets is a REST implementation of the Buckets interface.
type Buckets struct {
	Client client.RemoteCaller
}

// BucketService handles communication with bucket related API of mgmt api
// type BucketService mgmtAPIService

func (b *Buckets) GetPolicy(ctx context.Context, bucketName string, params map[string]string) (string, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("/object/bucket/%s/policy", bucketName),
		ContentType: client.ContentTypeJSON,
		Params:      params,
	}
	var bucketPolicy json.RawMessage

	err := b.Client.MakeRemoteCall(ctx, req, &bucketPolicy)
	if err != nil {
		return "", err
	}

	return string(bucketPolicy), nil
}

func (b *Buckets) UpdatePolicy(ctx context.Context, bucketName string, policy string, param map[string]string) error {
	req := client.Request{
		Method:      http.MethodPut,
		Path:        fmt.Sprintf("/object/bucket/%s/policy", bucketName),
		ContentType: client.ContentTypeJSON,
		Body:        json.RawMessage(policy),
		Params:      param,
	}

	return b.Client.MakeRemoteCall(ctx, req, nil)
}

func (b *Buckets) DeletePolicy(ctx context.Context, bucketName string, param map[string]string) error {
	req := client.Request{
		Method:      http.MethodDelete,
		Path:        fmt.Sprintf("/object/bucket/%s/policy", bucketName),
		ContentType: client.ContentTypeJSON,
		Params:      param,
	}

	return b.Client.MakeRemoteCall(ctx, req, nil)
}

func (b *Buckets) Create(ctx context.Context, createParam *model.ObjectBucketParam) (*model.Bucket, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        BucketServicePath,
		ContentType: client.ContentTypeJSON,
		Body:        &createParam,
	}
	bucket := &model.Bucket{}

	err := b.Client.MakeRemoteCall(ctx, req, bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (b *Buckets) Get(ctx context.Context, name string, param map[string]string) (*model.Bucket, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        path.Join("object", "bucket", name, "info"),
		ContentType: client.ContentTypeJSON,
		Params:      param,
	}
	bucket := &model.Bucket{}

	err := b.Client.MakeRemoteCall(ctx, req, bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (b *Buckets) Delete(ctx context.Context, name string, param map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join(BucketServicePath, name, "deactivate"),
		ContentType: client.ContentTypeJSON,
		Params:      param,
	}
	bucket := &model.Bucket{}

	err := b.Client.MakeRemoteCall(ctx, req, bucket)
	if err != nil {
		return err
	}

	return nil
}

func (b *Buckets) List(ctx context.Context, param map[string]string) (*model.BucketList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        BucketServicePath,
		Body:        nil,
		Params:      param,
		ContentType: client.ContentTypeJSON,
	}
	buckets := &model.BucketList{}

	err := b.Client.MakeRemoteCall(ctx, req, buckets)
	if err != nil {
		return nil, err
	}

	return buckets, nil
}
