// Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

func testServiceLogin(t *testing.T, _ client.Authenticator) {
	fixtureFailedServiceAuth := client.AuthService{
		Gateway:       ":not:a:valid:url",
		SharedSecret:  "",
		PodName:       "objectscale-graphql-7d754f8499-ng4h6",
		Namespace:     "svc-objectscale-domain-c8",
		ObjectScaleID: "IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
	}

	badAuth := &fixtureFailedServiceAuth
	err := badAuth.Login(context.TODO(), NewTestHTTPClient())
	require.Error(t, err)

	fixtureFailedServiceAuth.Gateway = "bad:gate:way"

	badAuth = &fixtureFailedServiceAuth
	err = badAuth.Login(context.TODO(), NewTestHTTPClient())
	require.Error(t, err)
}

func testUserLogin(t *testing.T, _ client.Authenticator) {
	fixtureFailedUserAuth := client.AuthUser{
		Gateway:  ":not:a:valid:url",
		Username: "testuser",
		Password: "testpassword",
	}

	badAuth := &fixtureFailedUserAuth
	err := badAuth.Login(context.TODO(), NewTestHTTPClient())
	require.Error(t, err)

	fixtureFailedUserAuth.Gateway = "bad:gate:way"

	badAuth = &fixtureFailedUserAuth
	err = badAuth.Login(context.TODO(), NewTestHTTPClient())
	require.Error(t, err)
}
