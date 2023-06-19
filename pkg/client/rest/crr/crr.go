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

package crr

import (
	"context"
	"net/http"
	"path"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

// CRR is a REST implementation of the CRR interface.
type CRR struct {
	Client client.RemoteCaller
}

// PauseReplication implements the CRR interface.
func (c *CRR) PauseReplication(ctx context.Context, destObjectScale string, destObjectStore string, params map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore, "pause"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	return c.Client.MakeRemoteCall(ctx, req, nil)
}

// SuspendReplication implements the CRR interface.
func (c *CRR) SuspendReplication(ctx context.Context, destObjectScale string, destObjectStore string, params map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore, "suspend"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	return c.Client.MakeRemoteCall(ctx, req, nil)
}

// ResumeReplication implements the CRR interface.
func (c *CRR) ResumeReplication(ctx context.Context, destObjectScale string, destObjectStore string, params map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore, "resume"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	return c.Client.MakeRemoteCall(ctx, req, nil)
}

// UnthrottleReplication implements the CRR interface.
func (c *CRR) UnthrottleReplication(ctx context.Context, destObjectScale string, destObjectStore string, params map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore, "unthrottle"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	return c.Client.MakeRemoteCall(ctx, req, nil)
}

// ThrottleReplication implements the CRR interface.
func (c *CRR) ThrottleReplication(ctx context.Context, destObjectScale string, destObjectStore string, params map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore, "throttle"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	return c.Client.MakeRemoteCall(ctx, req, nil)
}

// Get implements the CRR interface.
func (c *CRR) Get(ctx context.Context, destObjectScale string, destObjectStore string, params map[string]string) (*model.CRR, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	config := &model.CRR{}

	err := c.Client.MakeRemoteCall(ctx, req, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
