//
//
//  Copyright © 2021 - 2022 Dell Inc. or its subsidiaries. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//       http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//
//

package model

import "encoding/xml"

type Error struct {
	// XMLName is the name of the xml tag used XML marshalling
	XMLName xml.Name `xml:"error"`

	// Code is the error code returned by the management API
	Code int64 `xml:"code" json:"code"`

	// Description is a human readable description of the error
	Description string `xml:"description" json:"description"`

	// Details are additional error information
	Details string `xml:"details" json:"details"`

	// Retryable signifies if a request returning an error of this type should
	// be retried
	Retryable bool `xml:"retryable" json:"retryable"`
}
