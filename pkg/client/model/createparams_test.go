// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package model

import (
	"reflect"
	"testing"
)

func TestCreateBucketRequestParamsParseFrom(t *testing.T) {
	tests := []struct {
		name        string
		parameters  map[string]string
		wantErr     bool
		expectedCfg *CreateBucketRequestParams
	}{
		{
			name: "success - all parameters set",
			parameters: map[string]string{
				"namespace":                 "my-namespace",
				"replicationGroup":          "my-replication-group",
				"encryptionEnabled":         "true",
				"filesystemEnabled":         "true",
				"accessDuringOutageEnabled": "true",
				"quotaLimit":                "100",
				"quotaWarn":                 "50",
				"defaultRetention":          "30",
				"expiration":                "60",
			},
			wantErr: false,
			expectedCfg: &CreateBucketRequestParams{
				ReplicationGroup:          "my-replication-group",
				EncryptionEnabled:         true,
				FilesystemEnabled:         true,
				AccessDuringOutageEnabled: true,
				QuotaLimit:                func(i int) *int { return &i }(100),
				QuotaWarn:                 func(i int) *int { return &i }(50),
				DefaultRetention:          func(i int) *int { return &i }(30),
				Expiration:                func(i int) *int { return &i }(60),
			},
		},
		{
			name: "success - some parameters set",
			parameters: map[string]string{
				"replicationGroup": "my-replication-group",
			},
			wantErr: false,
			expectedCfg: &CreateBucketRequestParams{
				ReplicationGroup: "my-replication-group",
			},
		},
		{
			name: "failure - invalid boolean parameter value for encryptionEnabled",
			parameters: map[string]string{
				"encryptionEnabled": "invalid",
			},
			wantErr: true,
		},
		{
			name: "failure - invalid boolean parameter value for filesystemEnabled",
			parameters: map[string]string{
				"filesystemEnabled": "invalid",
			},
			wantErr: true,
		},
		{
			name: "failure - invalid boolean parameter value for accessDuringOutageEnabled",
			parameters: map[string]string{
				"accessDuringOutageEnabled": "invalid",
			},
			wantErr: true,
		},
		{
			name: "failure - invalid integer parameter value for quotaLimit",
			parameters: map[string]string{
				"quotaLimit": "invalid",
			},
			wantErr: true,
		},
		{
			name: "failure - invalid integer parameter value for quotaWarn",
			parameters: map[string]string{
				"quotaWarn": "invalid",
			},
			wantErr: true,
		},
		{
			name: "failure - invalid integer parameter value for defaultRetention",
			parameters: map[string]string{
				"defaultRetention": "invalid",
			},
			wantErr: true,
		},
		{
			name: "failure - invalid integer parameter value for expiration",
			parameters: map[string]string{
				"expiration": "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &CreateBucketRequestParams{}
			err := cfg.ParseFrom(tt.parameters)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(cfg, tt.expectedCfg) {
				t.Errorf("ParseFrom() cfg = %+v, expected %+v", cfg, tt.expectedCfg)
			}
		})
	}
}
