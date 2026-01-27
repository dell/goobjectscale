// Copyright © 2022 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package config

import "github.com/kelseyhightower/envconfig"

type MgmtConfig struct {
	EndpointURL string `envconfig:"MGMT_ENDPOINT_URL"`
	TLSVerify   bool   `envconfig:"MGMT_TLS_VERIFY"`
	Username    string `envconfig:"MGMT_USERNAME"`
	Password    string `envconfig:"MGMT_PASSWORD"`
}

func (c *MgmtConfig) ReadEnv() error {
	return envconfig.Process("MGMT", c)
}
