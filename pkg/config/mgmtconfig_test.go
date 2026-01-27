// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package config

import (
	"reflect"
	"testing"
)

func TestMgmtConfigReadEnv(t *testing.T) {
	tests := []struct {
		name        string
		envVars     map[string]string
		wantErr     bool
		expectedCfg *MgmtConfig
	}{
		{
			name: "success - all env vars set",
			envVars: map[string]string{
				"MGMT_ENDPOINT_URL": "https://example.com",
				"MGMT_TLS_VERIFY":   "true",
				"MGMT_USERNAME":     "user",
				"MGMT_PASSWORD":     "pass",
			},
			wantErr: false,
			expectedCfg: &MgmtConfig{
				EndpointURL: "https://example.com",
				TLSVerify:   true,
				Username:    "user",
				Password:    "pass",
			},
		},
		{
			name: "success - some env vars set",
			envVars: map[string]string{
				"MGMT_ENDPOINT_URL": "https://example.com",
			},
			wantErr: false,
			expectedCfg: &MgmtConfig{
				EndpointURL: "https://example.com",
			},
		},
		{
			name: "failure - invalid env var value",
			envVars: map[string]string{
				"MGMT_TLS_VERIFY": "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set env vars for this test
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}
			// Reset env vars after this test
			defer func() {
				for key := range tt.envVars {
					t.Setenv(key, "")
				}
			}()

			cfg := &MgmtConfig{}
			err := cfg.ReadEnv()

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(cfg, tt.expectedCfg) {
				t.Errorf("ReadEnv() cfg = %+v, expected %+v", cfg, tt.expectedCfg)
			}
		})
	}
}
