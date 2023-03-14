//
//
//  Copyright © 2021 - 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
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

package model_test

import (
	"errors"
	"testing"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Error":      testErrorError,
		"StatusCode": testErrorStatusCode,
		"Is":         testErrorIs,
	} {
		t.Run(scenario, fn)
	}
}

func testErrorError(t *testing.T) {
	testCases := []struct {
		name     string
		err      *model.Error
		expected string
	}{
		{
			name: "fully filled error",
			err: &model.Error{
				Code:        404,
				Description: "description",
				Details:     "details",
			},
			expected: "description: details",
		},
		{
			name: "error without code",
			err: &model.Error{
				Description: "description",
				Details:     "details",
			},
			expected: "description: details",
		},
		{
			name: "error without code and description",
			err: &model.Error{
				Details: "details",
			},
			expected: "Unknown: details",
		},
		{
			name: "error without code and details",
			err: &model.Error{
				Description: "description",
			},
			expected: "description",
		},
		{
			name:     "empty error",
			err:      &model.Error{},
			expected: "Unknown",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.err.Error())
		})
	}
}

func testErrorStatusCode(t *testing.T) {
	testCases := []struct {
		name     string
		err      *model.Error
		expected int64
	}{
		{
			name: "code is present",
			err: &model.Error{
				Code: 404,
			},
			expected: 404,
		},
		{
			name:     "code is not present",
			err:      &model.Error{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.err.StatusCode())
		})
	}
}

func testErrorIs(t *testing.T) {
	testCases := []struct {
		name     string
		err      *model.Error
		target   error
		expected bool
	}{
		{
			name: "target and error are the same",
			err: &model.Error{
				Code: 404,
			},
			target: &model.Error{
				Code: 404,
			},
			expected: true,
		},
		{
			name: "target and error are not the same",
			err: &model.Error{
				Code: 404,
			},
			target: &model.Error{
				Code: 505,
			},
			expected: false,
		},
		{
			name: "target is nil",
			err: &model.Error{
				Code: 404,
			},
			target:   nil,
			expected: false,
		},
		{
			name: "different type, message equal to description",
			err: &model.Error{
				Description: "not found",
				Code:        404,
			},
			target:   errors.New("not found"),
			expected: true,
		},
		{
			name: "different type, message equal description and details",
			err: &model.Error{
				Description: "not found",
				Details:     "something is missing",
				Code:        404,
			},
			target:   errors.New("not found: something is missing"),
			expected: true,
		},
		{
			name: "different type, message missing details",
			err: &model.Error{
				Description: "not found",
				Details:     "something is missing",
				Code:        404,
			},
			target:   errors.New("not found"),
			expected: false,
		},
		{
			name: "different type different message",
			err: &model.Error{
				Code: 404,
			},
			target:   errors.New("wrong error"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, errors.Is(tc.err, tc.target))
		})
	}
}
