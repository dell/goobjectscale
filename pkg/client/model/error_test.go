// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package model

import (
	"testing"

	"github.com/pkg/errors"
)

func TestError_Is(t *testing.T) {
	tests := []struct {
		name     string
		errorObj Error
		target   error
		expected bool
	}{
		{
			name: "success - matching error codes",
			errorObj: Error{
				Code: 500,
			},
			target: &Error{
				Code: 500,
			},
			expected: true,
		},
		{
			name: "failure - non-matching error codes",
			errorObj: Error{
				Code: 500,
			},
			target: &Error{
				Code: 404,
			},
			expected: false,
		},
		{
			name: "success - matching error messages",
			errorObj: Error{
				Description: "internal server error",
			},
			target: &Error{
				Description: "internal server error",
			},
			expected: true,
		},
		{
			name: "success - target error does not implement statusCoder interface",
			errorObj: Error{
				Description: "internal server error",
			},
			target:   errors.New("internal server error"),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.errorObj.Is(tt.target)

			if result != tt.expected {
				t.Errorf("Is() result = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	tests := []struct {
		name        string
		errorObj    Error
		expectedErr string
	}{
		{
			name: "success - error with description",
			errorObj: Error{
				Description: "internal server error",
			},
			expectedErr: "internal server error",
		},
		{
			name: "success - error with description and details",
			errorObj: Error{
				Description: "internal server error",
				Details:     "more details",
			},
			expectedErr: "internal server error: more details",
		},
		{
			name: "success - error with no description",
			errorObj: Error{
				Details: "more details",
			},
			expectedErr: "Unknown: more details",
		},
		{
			name:        "success - error with no description or details",
			errorObj:    Error{},
			expectedErr: "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errStr := tt.errorObj.Error()

			if errStr != tt.expectedErr {
				t.Errorf("Error() error = %v, expected %v", errStr, tt.expectedErr)
			}
		})
	}
}
