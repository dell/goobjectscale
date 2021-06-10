package model

import "encoding/xml"

// ServiceProvider is an IAM ServiceProvider
type ServiceProvider struct {
	// XMLName is the name of the xml tag used for XML marshalling
	XMLName      xml.Name `xml:"service_provider"`
	CreateTime   string   `xml:"create_time"`
	DNS          string   `xml:"dns"`
	Etag         string   `xml:"etag"`
	JavaKeystore string   `xml:"java_keystore"`
	KeyAlias     string   `xml:"key_alias"`
	KeyPassword  string   `xml:"key_password"`
	UniqueID     string   `xml:"unique_id"`
	UUID         string   `xml:"uuid"`
}

// ServiceProviderCreate is the ServiceProvider creation input
type ServiceProviderCreate struct {
	// XMLName is the name of the xml tag used for XML marshalling
	XMLName         xml.Name        `xml:"service_provider_create"`
	ServiceProvider ServiceProvider `xml:"service_provider"`
}
