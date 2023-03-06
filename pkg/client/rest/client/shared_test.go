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

package client_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

// Required to force io.ReadAll error
type errReader int

func (errReader) Read(p []byte) (n int, err error) {
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
	notFoundResponse := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><error><code>1004</code></error>`))),
	}

	failUnmarshalResponse := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(bytes.NewReader([]byte(`<?xml version="1.0" encoding="UTF-8" ?><badformat></badformat>`))),
	}

	badBodyResponse := &http.Response{
		StatusCode: 403,
		Body:       io.NopCloser(errReader(0)),
	}

	err := client.HandleResponse(response)
	require.Error(t, err)

	err = client.HandleResponse(emptyStatusResponse)
	require.Equal(t, "server error: failed", err.Error())

	err = client.HandleResponse(emptyCodeResponse)
	require.Equal(t, "server error: status code 403", err.Error())

	err = client.HandleResponse(notFoundResponse)
	require.Equal(t, "server error: not found", err.Error())

	err = client.HandleResponse(badBodyResponse)
	require.Equal(t, "server error: bad body", err.Error())

	err = client.HandleResponse(failUnmarshalResponse)
	require.Equal(t, "expected element type <error> but have <badformat>", err.Error())

}
