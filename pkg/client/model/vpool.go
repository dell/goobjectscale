// Copyright © 2022 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package model

type DataServiceVPoolList struct {
	DataServiceVPools []DataServiceVPool `json:"data_service_vpool"`
}

type DataServiceVPool struct {
	Global               interface{}     `json:"global"`
	Remote               interface{}     `json:"remote"`
	Vdc                  interface{}     `json:"vdc"`
	VarrayMappings       []VarrayMapping `json:"varrayMappings"`
	Name                 string          `json:"name"`
	ID                   string          `json:"id"`
	Link                 interface{}     `json:"link"`
	CreationTime         int64           `json:"creation_time"`
	Inactive             bool            `json:"inactive"`
	Internal             interface{}     `json:"internal"`
	Description          string          `json:"description"`
	IsAllowAllNamespaces bool            `json:"isAllowAllNamespaces"`
	EnableRebalancing    bool            `json:"enable_rebalancing"`
	UseReplicationTarget bool            `json:"useReplicationTarget"`
	IsFullRep            bool            `json:"isFullRep"`
}

type VarrayMapping struct {
	Name                string `json:"name"`
	Value               string `json:"value"`
	IsReplicationTarget bool   `json:"is_replication_target"`
}
