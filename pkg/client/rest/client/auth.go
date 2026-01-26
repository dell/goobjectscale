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
	"fmt"
	"net/http"
	"net/url"
)

// AuthRetriesMax is the maximum number of times the client will attempt to
// login before returning an error.
const AuthRetriesMax = 3

// Authenticator can perform a Login to the gateway.
//
//go:generate go run github.com/vektra/mockery/v2@latest --name Authenticator
type Authenticator interface {
	// IsAuthenticated returns true if the authenticated has been established.  This
	// does not mean the next request is guaranteed to succeed as authentication can
	// become expired.
	IsAuthenticated() bool

	// Login obtains fresh authentication token(s) from the server.
	Login(context.Context, *http.Client) error

	// Token returns the current authentication token.
	Token() string
}

var _ Authenticator = (*AuthUser)(nil) // interface guard

// AuthUser is an out-of-cluster or username+password based Authenticator.
type AuthUser struct {
	// Gateway is the auth endpoint
	Gateway string `json:"gateway"`

	// Username used to authenticate management user
	Username string `json:"username"`

	// Password used to authenticate management user
	Password string `json:"password"`

	token string
}

// IsAuthenticated returns true if the authenticated has been established.  This
// does not mean the next request is guaranteed to succeed as authentication can
// become expired.
func (auth *AuthUser) IsAuthenticated() bool {
	return auth.token != ""
}

// Login obtains fresh authentication token(s) from the server.
func (auth *AuthUser) Login(ctx context.Context, ht *http.Client) error {
	basicAuth := func(r *http.Request) { r.SetBasicAuth(auth.Username, auth.Password) }

	u, err := url.Parse(auth.Gateway)
	if err != nil {
		return err
	}

	u.Path = "/login"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	basicAuth(req)

	resp, err := ht.Do(req)
	if err != nil {
		return err
	}

	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	defer resp.Body.Close()

	if err = HandleResponse(resp); err != nil {
		return err
	}

	auth.token = resp.Header.Get("X-SDS-AUTH-TOKEN")
	if auth.token == "" {
		return fmt.Errorf("server error: login failed")
	}

	return nil
}

// Token returns the current authentication token.
func (auth *AuthUser) Token() string {
	return auth.token
}
