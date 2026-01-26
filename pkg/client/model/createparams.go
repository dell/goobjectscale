// Copyright © 2022 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package model

import (
	"strconv"

	"github.com/pkg/errors"
)

type CreateBucketRequestParams struct {
	ReplicationGroup          string `json:"replicationGroup,omitempty"`
	EncryptionEnabled         bool   `json:"encryptionEnabled,omitempty"`
	FilesystemEnabled         bool   `json:"filesystemEnabled,omitempty"`
	AccessDuringOutageEnabled bool   `json:"accessDuringOutageEnabled,omitempty"`
	QuotaLimit                *int   `json:"quotaLimit,omitempty"`
	QuotaWarn                 *int   `json:"quotaWarn,omitempty"`
	DefaultRetention          *int   `json:"defaultRetention,omitempty"`
	Expiration                *int   `json:"expiration,omitempty"`
}

func (p *CreateBucketRequestParams) ParseFrom(parameters map[string]string) error {
	if val, ok := parameters["replicationGroup"]; ok {
		p.ReplicationGroup = val
	}

	encryptionEnabled, err := getBoolValueByKeyName(parameters, "encryptionEnabled")
	if err != nil {
		return err
	}
	p.EncryptionEnabled = encryptionEnabled

	filesystemEnabled, err := getBoolValueByKeyName(parameters, "filesystemEnabled")
	if err != nil {
		return err
	}
	p.FilesystemEnabled = filesystemEnabled

	accessDuringOutageEnabled, err := getBoolValueByKeyName(parameters, "accessDuringOutageEnabled")
	if err != nil {
		return err
	}
	p.AccessDuringOutageEnabled = accessDuringOutageEnabled

	quotaLimit, err := getIntValueByKeyName(parameters, "quotaLimit")
	if err != nil {
		return err
	}
	p.QuotaLimit = quotaLimit

	quotaWarn, err := getIntValueByKeyName(parameters, "quotaWarn")
	if err != nil {
		return err
	}
	p.QuotaWarn = quotaWarn

	defaultRetention, err := getIntValueByKeyName(parameters, "defaultRetention")
	if err != nil {
		return err
	}
	p.DefaultRetention = defaultRetention

	expiration, err := getIntValueByKeyName(parameters, "expiration")
	if err != nil {
		return err
	}
	p.Expiration = expiration

	return nil
}

func getBoolValueByKeyName(parameters map[string]string, key string) (bool, error) {
	if val, ok := parameters[key]; ok {
		result, err := strconv.ParseBool(val)
		if err != nil {
			return false, errors.Wrapf(err, "Failed to extract %s parameter for bucket creation", key)
		}
		return result, nil
	}
	return false, nil
}

func getIntValueByKeyName(parameters map[string]string, key string) (*int, error) {
	if val, ok := parameters[key]; ok {
		result, err := strconv.Atoi(val)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to extract %s parameter for bucket creation", key)
		}
		return &result, nil
	}
	return nil, nil
}
