// Copyright © 2023 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package client

import (
	"io"
)

// CountWriter is an io.Writer that counts bytes written without actually writing anything.
type CountWriter struct {
	N int
}

var _ io.Writer = (*CountWriter)(nil) // interface guard

// Write adds len(p) to the current value of N and returns len(p) without
// writing anything.
func (w *CountWriter) Write(p []byte) (int, error) {
	w.N += len(p)
	return len(p), nil
}
