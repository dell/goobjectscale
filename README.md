# goobjectscale

![GitHub license](https://img.shields.io/github/license/dell/goobjectscale)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/dell/goobjectscale)
![GitHub branch checks state](https://img.shields.io/github/checks-status/dell/goobjectscale/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dell/goobjectscale)
[![Go Report Card](https://goreportcard.com/badge/github.com/dell/goobjectscale)](https://goreportcard.com/report/github.com/dell/goobjectscale)
[![Go Reference](https://pkg.go.dev/badge/github.com/dell/goobjectscale.svg)](https://pkg.go.dev/github.com/dell/goobjectscale)

> 🚧 **NOTE:** The *goobjectscale* is in a preview stage, and API is subject to change between minor updates. Until `v1.0.0` we do not guarantee backwards compatibility between versions.

## Overview

_goobjectscale_ is a Go package that provides a client for the Dell ObjectScale Managment HTTP API and helpers for Dell ObjectScale Identity and Access Management (IAM) HTTP API (compatible with Amazon Web Servicess IAM).

## Examples

The tests provide working examples for how to use the package, but here are a few code snippets to further illustrate the basic ideas.

### Initialize a new client


```go
import (
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

// First, provide user credentials for your ObjectScale.
objectscaleAuthUser := objectscaleClient.AuthUser{
	Gateway:  "https://gateway.example.com:443", // Look at FAQ how to obtain it.
	Username: "example-user",
	Password: "example-password",
}

// Next create valid HTTP transport, including TLS config.
transport := &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true // Insecure only for demo purpose.
	},
}

// Finally, create REST clientset
clientset := rest.NewClientSet(&objectscaleClient.Simple{
	Endpoint:       "https://objectstore.example.com:4443", // Look at FAQ how to obtain it.
	Authenticator:  &user,
	HTTPClient:     &http.Client{Transport: transport},
	OverrideHeader: false,
})
```

### Get existing bucket

```go
// Create new parameters map, that will be provided to the Get call.
parameters := map[string]string{
	"namespace": "osaia3382ab190a7a3df", // "namespace" is a required parameter (look at FAQ how to obtain it).
}

// NOTE: Create clientset beforehand.

// Get existing bucket.
bucket, err := clientset.Buckets().Get("example-bucket", parameters)
```

### Create new bucket

```go
import "github.com/dell/goobjectscale/pkg/client/model"

// NOTE: Create clientset beforehand.

// Create new bucket using clientset.
bucket, err = clientset.Buckets().Create(&model.Bucket{
	Name: "example-bucket",             // Name is a required field.
	Namespace: "osaia3382ab190a7a3df",  // Namespace is a required field.
})
```

### Delete the bucket

```go
import "github.com/dell/goobjectscale/pkg/client/model"

// NOTE: Create clientset beforehand.

// Delete bucket requires bucket name, namespace (look at FAQ how to obtain it) and emptyBucket parameters.
// EmptyBucket indicates if the bucket should be deleted, if it has any objects.
err := clientset.Buckets().Delete("example-bucket", "osaia3382ab190a7a3df", false)
```

### Initialize IAM client

```go
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/session"
)

// Let's use default HTTP Client.
x509Client := *http.DefaultClient

// First, provide user credentials for your ObjectScale.
objectscaleAuthUser := &objectscaleClient.AuthUser{
	Gateway:  "https://gateway.example.com:443", // Look at FAQ how to obtain it.
	Username: "example-user",
	Password: "example-password",
}

// Create new session with custom endpoint.
iamSession, err := session.NewSession(&aws.Config{
	Endpoint:                      "https://gateway.example.com:443", // Look at FAQ how to obtain it.
	Region:                        "us-west-1",
	CredentialsChainVerboseErrors: aws.Bool(true),
	HTTPClient: x509Client,
})

// Create new IAM client using the session above.
iamClient = iam.New(iamSession)

// Before using IAM client, we need to do some additional setup.
// First we need to inject ObjectScale access token from objectscaleAuthUser structure.
InjectTokenToIAMClient(iamClient, objectscaleAuthUser, x509Client)

// Next we need to inject Account ID / Namespace (look at FAQ how to obtain it).
InjectAccountIDToIAMClient(iamClient, "osaia3382ab190a7a3df")

// Finally we can use the IAM client to do API calls to ObjectScale.
user, err := iamClient.CreateUser(&iam.CreateUserInput{
	UserName: userName,
})
```

## Contributions

**First:** if you're unsure or afraid of anything, just ask or submit the issue or pull request anyway. You won't be yelled at for giving your best effort. The worst that can happen is that you'll be politely asked to change something.

We appreciate any sort of contributions, and don't want a wall of rules to get in the way of that.

## Licensing

Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

## FAQ

### What is the Namespace? What is the AccountID?

Those two terms used around codebase are unfortunately referring to the same thing. The easiest way to obtain it is to look in Objectscale Portal.

1. Login into ObjectScale Portal;
2. Select *Accounts* tab in the panel on the left side of your screen;
3. You should now see list of accounts. Select one of the values from column called *Account ID*.

### How to obtain ObjectScale Gateway endpoint URL?

The easiest way to obtain ObjectScale Gateway endpoint URL is to look in ObjectScale Portal.

1. Login into ObjectScale Portal;
2. From the menu on left side of the screen select *Administration* tab;
3. After unfolding *Administration* tab enter *ObjectScale* page;
4. Select *Federation* tab;
5. In the table you will see one or more values, unroll selected one;
6. In the table, you will now see *External Endpoint* value associated with *objectscale-gateway-internal*.
7. The endpoint must be of the following format: `https://<IP-ADDRESS>:443` or `https://<EXTERNAL-HOSTNAME>`

### How to obtain ObjectScale Objectstore endpoint URL?

The easiest way to obtain ObjectScale Objectstore endpoint URL is to look in ObjectScale Portal.


1. Login into ObjectScale Portal;
2. From the menu on left side of the screen select *Administration* tab;
3. After unfolding *Administration* tab enter *ObjectScale* page;
4. Select one of the object stores visible in the table, and click its name;
5. You should see *Summary* of that object store.
6. In the *Management Service details* section, you will see value under *IP address* column.
7. The endpoint must be of the following format: `https://<IP-ADDRESS>:4443` or `https://<EXTERNAL-HOSTNAME>`

## Support

For any issues, questions or feedback, please follow our [support process](https://github.com/dell/csm/blob/main/docs/SUPPORT.md)
