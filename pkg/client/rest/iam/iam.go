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

// Example function using IAM client to create new user in ObjectScale.
//
//	func ExampleCreateIAMUser(userName string) {
//		x509Client := *http.DefaultClient
//		objClient = client.AuthUser{Gateway: "https://testgateway", Username: "username", Password: "password"}
//		iamSession, _ := session.NewSession(&aws.Config{ 	// github.com/aws/aws-sdk-go/service/session
//				Endpoint:                      "https://testgateway/iam",
//				Region:                        "us-west-2",
//				CredentialsChainVerboseErrors: aws.Bool(true),
//				HTTPClient: x509Client,
//				},
//			})
//		iamClient = iam.New(iamSession) // github.com/aws/aws-sdk-go/service/iam
//		InjectTokenToIAMClient(iamClient, &objClient, x509Client)
//		InjectAccountIDToIAMClient(iamClient, "osaid185e2bf9e8ae35f")
//		user, err := iamClient.CreateUser(&iam.CreateUserInput{
//			UserName: userName,
//		})
//	}
package iam

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"

	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

const (
	SDSHandlerName       = "X-Sds-Handler"
	SDSHeaderName        = "X-Sds-Auth-Token"
	AccountIDHandlerName = "X-Emc-Handler"
	AccountIDHeaderName  = "X-Emc-Namespace"
)

// InjectTokenToIAMClient configure IAM client to connect with Objectscale
func InjectTokenToIAMClient(clientIam *iam.IAM, clientObjectscale client.Authenticator, httpClient http.Client) {
	clientIam.Handlers.Sign.RemoveByName(v4.SignRequestHandler.Name)

	handler := request.NamedHandler{
		Name: SDSHandlerName,
		Fn: func(r *request.Request) {
			if !clientObjectscale.IsAuthenticated() {
				err := clientObjectscale.Login(&httpClient)
				if err != nil {
					r.Error = err // no return intentional
				}
			}

			token := clientObjectscale.Token()
			r.HTTPRequest.Header.Add(SDSHeaderName, token)
		},
	}

	swapped := clientIam.Handlers.Sign.SwapNamed(handler)
	if !swapped {
		clientIam.Handlers.Sign.PushFrontNamed(handler)
	}
}

// InjectAccountIDToIAMClient configure IAM client to connect with Objectscale Accont
func InjectAccountIDToIAMClient(clientIam *iam.IAM, AccountID string) {
	clientIam.Handlers.Sign.RemoveByName(v4.SignRequestHandler.Name)
	handler := request.NamedHandler{
		Name: AccountIDHandlerName,
		Fn: func(r *request.Request) {
			r.HTTPRequest.Header.Add(AccountIDHeaderName, AccountID)
		},
	}

	swapped := clientIam.Handlers.Sign.SwapNamed(handler)
	if !swapped {
		clientIam.Handlers.Sign.PushFrontNamed(handler)
	}
}
