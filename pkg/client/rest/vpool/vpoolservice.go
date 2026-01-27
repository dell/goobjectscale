// Copyright © 2022 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package vpool

import (
	"context"
	"net/http"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

const (
	DataServiceVPoolsPath = "/vdc/data-service/vpools"
)

type Service struct {
	Client client.RemoteCaller
}

func (b Service) List(ctx context.Context) ([]model.DataServiceVPool, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        DataServiceVPoolsPath,
		ContentType: client.ContentTypeJSON,
	}
	vPools := &model.DataServiceVPoolList{}

	err := b.Client.MakeRemoteCall(ctx, req, vPools)
	if err != nil {
		return nil, err
	}

	return vPools.DataServiceVPools, nil
}
