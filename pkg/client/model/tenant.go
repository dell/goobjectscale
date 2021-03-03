package model

import "encoding/xml"

// TenantInfo is an object store tenant with an alternate XML tag name
type TenantInfo struct {
	// XMLName is the name of the xml tag used XML marshalling
	XMLName xml.Name `xml:"tenant_info"`

	Tenant
}

// Tenant is an object store tenant
type Tenant struct {
	// XMLName is the name of the xml tag used XML marshalling
	XMLName xml.Name `xml:"tenant"`

	// ID is the id of the tenant scoped to the cluster instance
	ID string `json:"id,omitempty" xml:"id,omitempty"`

	// EncryptionEnabled displays if this tenant is configured for encryption at rest
	EncryptionEnabled bool `json:"is_encryption_enabled,omitempty" xml:"is_encryption_enabled,omitempty"`

	// ComplianceEnabled displays if this tenant is configured for compliance retention
	ComplianceEnabled bool `json:"is_compliance_enabled,omitempty" xml:"is_compliance_enabled,omitempty"`

	// ReplicationGroup is the default replication group id of the tenant
	ReplicationGroup string `json:"default_data_services_vpool,omitempty" xml:"default_data_services_vpool,omitempty"`

	// BucketBlockSize is the default bucket size at which new object creations will be blocked
	BucketBlockSize int64 `json:"default_bucket_block_size,omit_empty,omitempty" xml:"default_bucket_block_size,omit_empty,omitempty"`
}

// TenantList is a list of object store tenants
type TenantList struct {
	// XMLName is the name of the xml tag used XML marshalling
	XMLName xml.Name `json:"tenants" xml:"tenants"`

	// Items is the list of tenants in the list
	Items []Tenant `json:"tenant" xml:"tenant"`
}