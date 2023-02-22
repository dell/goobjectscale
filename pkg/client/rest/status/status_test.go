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

package status_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

func TestStatus(t *testing.T) {
	t.Run("should return build info", func(t *testing.T) {
		var (
			r   *recorder.Recorder
			err error
		)
		r, err = recorder.New("fixtures")
		if err != nil {
			log.Fatal(err)
		}
		r.AddHook(func(i *cassette.Interaction) error {
			delete(i.Request.Headers, "Authorization")
			delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
			return nil
		}, recorder.BeforeSaveHook)
		clientset := rest.NewClientSet(client.NewServiceClient(
			"https://testserver",
			"https://testgateway",
			"svc-objectscale-domain-c8",
			"objectscale-graphql-7d754f8499-ng4h6",
			"OSC234DSF223423",
			"IgQBVjz4mq1M6wmKjHmfDgoNSC56NGPDbLvnkaiuaZKpwHOMFOMGouNld7GXCC690qgw4nRCzj3EkLFgPitA2y8vagG6r3yrUbBdI8FsGRQqW741eiYykf4dTvcwq8P6",
			newRecordedHTTPClient(r),
			false,
		))
		rebuildInfo, err := clientset.Status().GetRebuildStatus("testdevice1",
			"testdevice1-ss-0", "testdomain", "1", nil)
		require.NoError(t, err)
		assert.Equal(t, rebuildInfo.Level, 1)
		assert.Equal(t, rebuildInfo.RemainingBytes, 1024)
		require.Equal(t, rebuildInfo.TotalBytes, 2048)
		_, err = clientset.Status().GetRebuildStatus("testdevice1",
			"testdevice1-ss-1", "testdomain", "1", nil)
		require.Error(t, err)
	})
}

func newRecordedHTTPClient(r *recorder.Recorder) *http.Client {
	return &http.Client{Transport: r}
}
