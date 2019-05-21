package model

import "encoding/xml"

// Bucket is an object storage bucket
type Bucket struct {
	// XMLName is the name of the xml tag used XML marshalling
	XMLName xml.Name `xml:"object_bucket"`

	// APIType is the object API type used by the bucket
	APIType string `json:"api_type" xml:"api_type"`

	// AuditDeleteExpiration is the amount of time to retain deletion audit
	// entries
	AuditDeleteExpiration int `json:"audit_delete_expiration" xml:"audit_delete_expiration"`

	// Created is the date and time that the bucket was created
	Created string `json:"created" xml:"created"`

	// ID is the id of the bucket scoped to the cluster instance
	ID string `json:"id" xml:"id"`

	// Name is the name of the cluster instance
	Name string `json:"name" xml:"name"`

	// EncryptionEnabled displays if this bucket is configured for encryption
	// at rest
	EncryptionEnabled bool `json:"is_encryption_enabled" xml:"is_encryption_enabled"`

	// SoftQuota is the warning quota level for the bucket
	SoftQuota string `json:"softquota" xml:"softquota"`

	// FSEnabled indicates if the bucket has file-system support enabled
	FSEnabled bool `json:"fs_access_enabled" xml:"fs_access_enabled"`

	// Locked indicates if the bucket is locked
	Locked bool `json:"locked" xml:"locked"`

	// VPool is the replication group id of the bucket
	VPool string `json:"vpool" xml:"vpool"`

	// Namespace is the namespace of the bucket
	Namespace string `json:"namespace" xml:"namespace"`

	// Owner is the s3 object user owner of the bucket
	Owner string `json:"owner" xml:"owner"`

	// StaleAllowed indicates if access to the bucket is allowed during an
	// outage
	StaleAllowed bool `json:"is_stale_allowed" xml:"is_stale_allowed"`

	// TSOReadOnly indicates if access to the bucket is allowed during an
	// outage
	TSOReadOnly bool `json:"is_tso_read_only" xml:"is_tso_read_only"`

	// DefaultRetention is the default retention period for objects in bucket
	DefaultRetention int64 `json:"default_retention,omit_empty" xml:"default_retention,omit_empty"`

	// BlockSize is the bucket size at which new object creations will be
	// blocked
	BlockSize int64 `json:"block_size,omit_empty" xml:"block_size,omit_empty"`

	// NotificationSize is the bucket size at which the users will be notified
	NotificationSize int64 `json:"notification_size,omit_empty" xml:"notification_size,omit_empty"`

	// Tags is a list of arbitrary metadata keys and values applied to the
	// bucket
	Tags TagSet `json:"TagSet" xml:"TagSet"`

	// Retention is the default retention value for the bucket
	Retention int64 `json:"retention" xml:"retention"`

	// DefaultGroupFileReadPermission is a flag indicating the Read permission
	// for default group
	DefaultGroupFileReadPermission bool `json:"default_group_file_read_permission" xml:"default_group_file_read_permission"`

	// DefaultGroupFileWritePermission is a flag indicating the Execute permission
	// for default group
	DefaultGroupFileExecutePermission bool `json:"default_group_file_execute_permission" xml:"default_group_file_execute_permission"`

	// DefaultGroupFileExecutePermission is a flag indicating the Write permission
	// for default group
	DefaultGroupFileWritePermission bool `json:"default_group_file_write_permission" xml:"default_group_file_write_permission"`

	// DefaultGroupDirReadPermission is a flag indicating the Read permission
	// for default group
	DefaultGroupDirReadPermission bool `json:"default_group_dir_read_permission" xml:"default_group_dir_read_permission"`

	// DefaultGroupDirWritePermission is a flag indicating the Execute permission
	// for default group
	DefaultGroupDirExecutePermission bool `json:"default_group_dir_execute_permission" xml:"default_group_dir_execute_permission"`

	// DefaultGroupDirExecutePermission is a flag indicating the Write permission
	// for default group
	DefaultGroupDirWritePermission bool `json:"default_group_dir_write_permission" xml:"default_group_dir_write_permission"`

	// DefaultGroup is the bucket's default group
	DefaultGroup string `json:"default_group,omitempty" xml:"default_group,omitempty"`

	// SearchMetadata is the custom metadata for enabled for querying on the
	// bucket
	SearchMetadata `json:"search_metadata" xml:"search_metadata"`

	// MinMaxGovenor enforces minimum and maximum retention for bucket objects
	MinMaxGovenor `json:"min_max_govenor" xml:"min_max_govenor"`

}

// BucketList is a list of object storage buckets
type BucketList struct {
	// XMLName is the name of the xml tag used XML marshalling
	XMLName xml.Name `json:"object_buckets" xml:"object_buckets"`

	// Items is the list of buckets in the list
	Items []Bucket `json:"object_bucket" xml:"object_bucket"`

	// MaxBuckets is the maximum number of buckets requested in the listing
	MaxBuckets int `json:"max_buckets,omit_empty" xml:"max_buckets,omit_empty"`

	// NextMarker is a reference object to receive the next set of buckets
	NextMarker string `json:"next_marker,omit_empty" xml:"next_marker,omit_empty"`

	// Filter is a string query used to limit the returned buckets in the
	// listing
	Filter string `json:"Filter,omit_empty" xml:"Filter,omit_empty"`

	// NextPageLink is a hyperlink to the next page in the bucket listing
	NextPageLink string `json:"next_page_link,omit_empty" xml:"next_page_link,omit_empty"`
}

// MinMaxGovenor enforces minimum and maximum retention for bucket objects
type MinMaxGovenor struct {

	// EnforceRetention indicates if retention should be enforced for this
	// min-max-govenor
	EnforceRetention bool `json:"enforce_retention" xml:"enforce_retention"`

	// MinimumFixedRetention  is the minimum fixed retention for objects within
	// a bucket
	MinimumFixedRetention int64 `json:"minimum_fixed_retention" xml:"minimum_fixed_retention"`

	// MinimumVariableRetention  is the minimum variable retention for objects
	// within a bucket
	MinimumVariableRetention int64 `json:"minimum_variable_retention" xml:"minimum_variable_retention"`

	// MaximumFixedRetention  is the maximum fixed retention for objects within
	// a bucket
	MaximumFixedRetention int64 `json:"maximum_fixed_retention" xml:"maximum_fixed_retention"`

	// MaximumVariableRetention  is the maximum variable retention for objects
	// within a bucket
	MaximumVariableRetention int64 `json:"maximum_variable_retention" xml:"maximum_variable_retention"`

	// Link is the hyperlink to this resource
	Link `json:"link" xml:"link"`

	// Inactive indicates if the bucket has been placed into an inactive state,
	// typically prior to deletion
	Inactive bool `json:"inactive" xml:"inactive"`

	// Global indicates if the resource is global
	Global bool `json:"global" xml:"global"`

	// Remote indicates if the resource is remote to the current API instance
	Remote bool `json:"remote" xml:"remote"`

	// Internal indicates if the resource is an internal resource
	Internal bool `json:"internal" xml:"internal"`

	// VDCLink is a link from a bucket to a VDC
	VDCLink `json:"vdc" xml:"vdc"`
}

// VDCLink is a link from a bucket to a VDC
type VDCLink struct {
	// ID is the identifier for the VDC
	ID string `json:"id" xml:"id"`

	// Link is a hyperlink to the VDC
	Link `json:"link" xml:"link"`
}
