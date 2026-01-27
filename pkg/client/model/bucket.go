// Copyright © 2022 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package model

import (
	"time"
)

const (
	S3 = "S3"
)

type ObjectBucketParam struct {
	Name              string `json:"name,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	Vpool             string `json:"vpool,omitempty"`
	HeadType          string `json:"head_type,omitempty"`
	FsAccessEnabled   bool   `json:"filesystem_enabled,omitempty"`
	EncryptionEnabled bool   `json:"is_encryption_enabled,omitempty"`
	IsStaleAllowed    bool   `json:"is_stale_allowed,omitempty"`
	Retention         *int   `json:"retention,omitempty"`
	NotificationSize  *int   `json:"notificationSize,omitempty"`
	BlockSize         *int   `json:"blockSize,omitempty"`
}

type BucketList struct {
	Buckets []Bucket `json:"object_bucket"`
}

type Bucket struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Link struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"link"`
	Namespace                       string      `json:"namespace"`
	Vpool                           string      `json:"vpool"`
	Locked                          bool        `json:"locked"`
	FsAccessEnabled                 bool        `json:"fs_access_enabled"`
	Softquota                       string      `json:"softquota"`
	Created                         time.Time   `json:"created"`
	IsStaleAllowed                  bool        `json:"is_stale_allowed"`
	IsTsoReadOnly                   bool        `json:"is_tso_read_only"`
	IsObjectLockEnabled             bool        `json:"is_object_lock_enabled"`
	DefaultObjectLockRetentionMode  interface{} `json:"default_object_lock_retention_mode"`
	DefaultObjectLockRetentionYears interface{} `json:"default_object_lock_retention_years"`
	DefaultObjectLockRetentionDays  interface{} `json:"default_object_lock_retention_days"`
	DefaultRetention                int         `json:"default_retention"`
	BlockSize                       int         `json:"block_size"`
	AutoCommitPeriod                int         `json:"auto_commit_period"`
	NotificationSize                int         `json:"notification_size"`
	BlockSizeInCount                int         `json:"blockSizeInCount"`
	NotificationSizeInCount         int         `json:"notificationSizeInCount"`
	IsEncryptionEnabled             string      `json:"is_encryption_enabled"`
	Tag                             []struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"Tag"`
	Retention                         int    `json:"retention"`
	DefaultGroup                      string `json:"default_group"`
	DefaultGroupFileReadPermission    bool   `json:"default_group_file_read_permission"`
	DefaultGroupFileWritePermission   bool   `json:"default_group_file_write_permission"`
	DefaultGroupFileExecutePermission bool   `json:"default_group_file_execute_permission"`
	DefaultGroupDirReadPermission     bool   `json:"default_group_dir_read_permission"`
	DefaultGroupDirWritePermission    bool   `json:"default_group_dir_write_permission"`
	DefaultGroupDirExecutePermission  bool   `json:"default_group_dir_execute_permission"`
	MinMaxGovernor                    struct {
		EnforceRetention         bool        `json:"enforce_retention"`
		MinimumFixedRetention    interface{} `json:"minimum_fixed_retention"`
		MaximumFixedRetention    interface{} `json:"maximum_fixed_retention"`
		MinimumVariableRetention interface{} `json:"minimum_variable_retention"`
		MaximumVariableRetention interface{} `json:"maximum_variable_retention"`
	} `json:"min_max_governor"`
	AuditDeleteExpiration              int         `json:"audit_delete_expiration"`
	EnableAdvancedMetadataSearch       bool        `json:"enableAdvancedMetadataSearch"`
	AdvancedMetadataSearchTargetName   interface{} `json:"advancedMetadataSearchTargetName"`
	AdvancedMetadataSearchTargetStream interface{} `json:"advancedMetadataSearchTargetStream"`
	Owner                              string      `json:"owner"`
	APIType                            string      `json:"api_type"`
	SearchMetadata                     struct {
		IsEnabled bool `json:"isEnabled"`
		Metadata  []struct {
			Datatype string `json:"datatype"`
			Name     string `json:"name"`
			Type     string `json:"type"`
		} `json:"metadata"`
		MaxKeys        int  `json:"maxKeys"`
		MetadataTokens bool `json:"metadata_tokens"`
	} `json:"search_metadata"`
}

type BucketOwner struct {
	Namespace string `json:"namespace"`
	NewOwner  string `json:"new_owner"`
}
