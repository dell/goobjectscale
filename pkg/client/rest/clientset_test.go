// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package rest

import (
	"net/http"
	"testing"

	"github.com/dell/goobjectscale/pkg/client/rest/client"
	"github.com/stretchr/testify/assert"
)

func TestNewClientSet(t *testing.T) {
	tests := []struct {
		name   string
		client client.RemoteCaller
	}{
		{
			name: "success - valid client",
			client: &client.Simple{
				Endpoint:       "https://testserver",
				Authenticator:  &client.AuthUser{},
				HTTPClient:     &http.Client{},
				OverrideHeader: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientSet := NewClientSet(tt.client)
			assert.NotNil(t, clientSet)
			assert.NotNil(t, clientSet.Client())
			assert.NotNil(t, clientSet.Buckets())
			assert.NotNil(t, clientSet.VPools())
		})
	}
}
