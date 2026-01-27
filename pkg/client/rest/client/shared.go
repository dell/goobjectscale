// Copyright © 2023 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package client

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dell/goobjectscale/pkg/client/model"
)

// RemoteCaller interface is used to create backend calls.
// into represents type, _into_ which data will be unmarshalled.
// Naming follows Effective Go naming convention https://go.dev/doc/effective_go#interface-names
type RemoteCaller interface {
	MakeRemoteCall(ctx context.Context, r Request, into interface{}) error
}

// HandleResponse handles custom behavior based on server response.
func HandleResponse(resp *http.Response) error {
	if resp.StatusCode >= http.StatusBadRequest {
		switch resp.Body {
		case nil:
			switch {
			case resp.Status != "":
				return fmt.Errorf("server error: %s", strings.ToLower(resp.Status))
			case resp.StatusCode != 0:
				return fmt.Errorf("server error: status code %d", resp.StatusCode)
			}
		default:
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			apiError := &model.Error{}

			err = xml.Unmarshal(body, apiError)
			if err != nil {
				return err
			}

			return fmt.Errorf("server error: %w", apiError)
		}
	}

	return nil
}
