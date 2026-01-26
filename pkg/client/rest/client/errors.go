// Copyright © 2023 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package client

import "errors"

var (
	// ErrAuthorization is returned when the client is unable to authenticate with the server.
	ErrAuthorization = errors.New("authorization")

	// ErrContentType is returned when the client or server responds with an unknown content type header.
	ErrContentType = errors.New("content type")
)
