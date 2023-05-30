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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

func testLogin(t *testing.T, auth client.Authenticator) {
	fixtureFailedServiceauth := client.AuthService{
		Gateway:       ":not:a:valid:url",
		SharedSecret:  "",
		PodName:       "objectscale-graphql-7d754f8499-ng4h6",
		Namespace:     "svc-objectscale-domain-c8",
		ObjectScaleID: "IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
	}
	badAuth := &fixtureFailedServiceauth
	err := badAuth.Login(NewTestHTTPClient())
	require.Error(t, err)
	fixtureFailedServiceauth.Gateway = "bad:gate:way"
	badAuth = &fixtureFailedServiceauth
	err = badAuth.Login(NewTestHTTPClient())
	require.Error(t, err)

}
