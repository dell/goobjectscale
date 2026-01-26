// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package model

import (
	"errors"
	"fmt"
	"strings"
)

type Error struct {
	Code        int64  `json:"code"`
	Description string `json:"description"`
	Details     string `json:"details"`
	Retryable   bool   `json:"retryable"`
}

var (
	// Error codes specific to ObjectScale
	ErrParameterNotFound = Error{Code: 1004}
	ErrEntityNotFound    = Error{Code: 2000}
)

var _ error = Error{}

func (err Error) StatusCode() int64 {
	return err.Code
}

func (err Error) Is(target error) bool {
	// create intermediate interface
	type statusCoder interface {
		StatusCode() int64
	}

	var statusCoderErr statusCoder
	// validate if target implements statusCoder interface,
	// and compare the status codes
	if errors.As(target, &statusCoderErr) {
		return err.Code == statusCoderErr.StatusCode()
	}

	// if someone is already relying on error message comparisons, then don't break it
	return strings.EqualFold(err.Error(), target.Error())
}

func (err Error) Error() string {
	if err.Description == "" {
		err.Description = "Unknown"
	}

	if err.Details != "" {
		return fmt.Sprintf("%s: %s", err.Description, err.Details)
	}

	return err.Description
}
