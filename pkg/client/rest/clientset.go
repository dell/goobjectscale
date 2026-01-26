// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package rest

import (
	"github.com/dell/goobjectscale/pkg/client/api"
	"github.com/dell/goobjectscale/pkg/client/rest/buckets"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
	"github.com/dell/goobjectscale/pkg/client/rest/vpool"
)

// ClientSet is a set of clients for each API section.
type ClientSet struct {
	client  client.RemoteCaller
	buckets api.BucketServiceInterface
	vpool   api.VPoolServiceInterface
}

var _ api.ClientSet = (*ClientSet)(nil)

// NewClientSet returns a new client set based on the provided REST client parameters.
func NewClientSet(c client.RemoteCaller) *ClientSet {
	return &ClientSet{
		client:  c,
		buckets: &buckets.Buckets{Client: c},
		vpool:   &vpool.Service{Client: c},
	}
}

// Client returns the REST client used in the ClientSet.
func (c *ClientSet) Client() client.RemoteCaller {
	return c.client
}

// Buckets implements the client API.
func (c *ClientSet) Buckets() api.BucketServiceInterface {
	return c.buckets
}

func (c *ClientSet) VPools() api.VPoolServiceInterface {
	return c.vpool
}
