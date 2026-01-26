// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package api

import (
	"context"

	"github.com/dell/goobjectscale/pkg/client/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name ClientSet
type ClientSet interface {
	Buckets() BucketServiceInterface
	VPools() VPoolServiceInterface
}

//go:generate go run github.com/vektra/mockery/v2@latest --name BucketServiceInterface
type BucketServiceInterface interface {
	Create(ctx context.Context, createParam *model.ObjectBucketParam) (*model.Bucket, error)
	Get(ctx context.Context, bucketName string, param map[string]string) (*model.Bucket, error)
	List(ctx context.Context, param map[string]string) (*model.BucketList, error)
	Delete(ctx context.Context, bucketName string, param map[string]string) error
	UpdatePolicy(ctx context.Context, bucketName string, policy string, param map[string]string) error
	GetPolicy(ctx context.Context, bucketName string, param map[string]string) (string, error)
	DeletePolicy(ctx context.Context, bucketName string, param map[string]string) error
}

//go:generate go run github.com/vektra/mockery/v2@latest --name VPoolServiceInterface
type VPoolServiceInterface interface {
	List(context.Context) ([]model.DataServiceVPool, error)
}
