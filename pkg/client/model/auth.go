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

package model

// RKELoginRequest is a model of login request body, posted to /mgmt/auth/login on RKE platform.
type RKELoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RKELoginResponse is a model of login response body, received from /mgmt/auth/login on RKE platform.
type RKELoginResponse struct {
	AccessToken      string `json:"access_token"`
	AccessExpiresIn  int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}
