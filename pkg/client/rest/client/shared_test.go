// Copyright © 2023 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package client_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
	"github.com/stretchr/testify/require"
)

// Required to force io.ReadAll error.
type errReader int

func (errReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("server error: bad body")
}

func TestHandleResponse(t *testing.T) {
	response := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><error><description>"OK"</description></error>`))),
	}
	emptyStatusResponse := &http.Response{
		StatusCode: 403,
		Body:       nil,
		Status:     "failed",
	}
	emptyCodeResponse := &http.Response{
		StatusCode: 403,
		Body:       nil,
	}
	badBodyResponse := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(errReader(0)),
	}
	okError := &model.Error{
		Code:        0,
		Description: "OK",
	}

	err := client.HandleResponse(response)
	require.ErrorIs(t, err, okError)

	err = client.HandleResponse(emptyStatusResponse)
	require.Equal(t, "server error: failed", err.Error())

	err = client.HandleResponse(emptyCodeResponse)
	require.Equal(t, "server error: status code 403", err.Error())

	err = client.HandleResponse(badBodyResponse)
	require.Equal(t, "server error: bad body", err.Error())
}
