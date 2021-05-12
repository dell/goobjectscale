package model

import "encoding/xml"


type CRR struct {
	XMLName xml.Name `xml:"ReplicationAdminConfiguration"`

	DestObjectScale string `xml:"DestinationObjectScale"`

	DestObjectStore string `xml:"DestinationObjectStore"`

	PauseStartMills int `xml:"PauseStartMills"`

	PauseEndMills int `xml:"PauseEndMills"`

	SuspendStartMills int `xml:"SuspendStartMills"`

	ThrottleBandwidth int `xml:"ThrottleBandwidth"`
}
