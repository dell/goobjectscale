
> 🚧 **NOTE:** The *goobjectscale* is in a preview stage, and API is subject to change between minor updates. Until `v1.0.0` we do not guarantee backwards compatibility between versions.

# goobjectscale

[![GitHub license](https://img.shields.io/github/license/dell/goobjectscale)](https://github.com/dell/goobjectscale/blob/main/LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/dell/goobjectscale)](https://github.com/dell/goobjectscale/releases/latest)
[![GitHub branch checks state](https://img.shields.io/github/checks-status/dell/goobjectscale/main)](https://github.com/dell/goobjectscale/actions)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dell/goobjectscale)](https://github.com/dell/goobjectscale/blob/main/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/dell/goobjectscale)](https://goreportcard.com/report/github.com/dell/goobjectscale)
[![Go Reference](https://pkg.go.dev/badge/github.com/dell/goobjectscale.svg)](https://pkg.go.dev/github.com/dell/goobjectscale)

## Description
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
objectscaleAuthUser := client.AuthUser{
	Gateway:  "https://gateway.example.com:443", // See FAQ on how to get it.
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
	Endpoint:       "https://objectstore.example.com:4443", // See FAQ on how to get it.
	Authenticator:  &user,
	HTTPClient:     &http.Client{Transport: transport},
	OverrideHeader: false,
})
```

### Get existing bucket

```go
// Create new parameters map, that will be provided to the Get call.
parameters := map[string]string{
	"namespace": "osaia3382ab190a7a3df", // "namespace" is a required parameter (see FAQ on how to get it).
}

// NOTE: Create clientset beforehand.

ctx := context.TODO() // Only for demo purpose.

// Get existing bucket.
bucket, err := clientset.Buckets().Get(ctx, "example-bucket", parameters)
```

### Create new bucket

```go
import "github.com/dell/goobjectscale/pkg/client/model"

// NOTE: Create clientset beforehand.

ctx := context.TODO() // Only for demo purpose.

// Create new bucket using clientset.
bucket, err = clientset.Buckets().Create(ctx, &model.Bucket{
	Name: "example-bucket",             // Name is a required field.
	Namespace: "osaia3382ab190a7a3df",  // Namespace is a required field.
})
```

### Delete the bucket

```go
import "github.com/dell/goobjectscale/pkg/client/model"

// NOTE: Create clientset beforehand.

ctx := context.TODO() // Only for demo purpose.

// Delete bucket requires bucket name, namespace (see FAQ on how to get it) and emptyBucket parameters.
// EmptyBucket indicates if the bucket should be deleted, if it has any objects.
err := clientset.Buckets().Delete(ctx, "example-bucket", "osaia3382ab190a7a3df", false)
```

### Initialize IAM client

```go
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/session"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

// Let's use default HTTP Client.
x509Client := *http.DefaultClient

// First, provide user credentials for your ObjectScale.
objectscaleAuthUser := &client.AuthUser{
	Gateway:  "https://gateway.example.com:443", // See FAQ on how to get it.
	Username: "example-user",
	Password: "example-password",
}

// Create new session with custom endpoint.
iamSession, err := session.NewSession(&aws.Config{
	Endpoint:                      "https://gateway.example.com:443", // See FAQ on how to get it.
	Region:                        "us-west-1",
	CredentialsChainVerboseErrors: aws.Bool(true),
	HTTPClient: x509Client,
})

// Create new IAM client using the session above.
iamClient = iam.New(iamSession)

// Before using IAM client, we need to do some additional setup.
// First we need to inject ObjectScale access token from objectscaleAuthUser structure.
InjectTokenToIAMClient(iamClient, objectscaleAuthUser, x509Client)

// Next we need to inject Account ID / Namespace (see FAQ on how to get it).
InjectAccountIDToIAMClient(iamClient, "osaia3382ab190a7a3df")

// Finally we can use the IAM client to do API calls to ObjectScale.
user, err := iamClient.CreateUser(&iam.CreateUserInput{
	UserName: userName,
})
```

## FAQ

### What is the Namespace? What is the AccountID?

Those two terms used around codebase are unfortunately referring to the same thing. The easiest way to obtain it is to look in Objectscale Portal.

1. Log in to the ObjectScale Portal;
2. Select *Accounts* tab in the panel on the left side of your screen;
3. You should now see list of accounts. Select one of the values from column called *Account ID*.

### How to obtain ObjectScale Gateway endpoint URL?

The easiest way to obtain ObjectScale Gateway endpoint URL is to look in ObjectScale Portal.

1. Log in to the ObjectScale Portal;
2. From the menu on left side of the screen select *Administration* tab;
3. After unfolding *Administration* tab enter *ObjectScale* page;
4. Select *Federation* tab;
5. In the table you will see one or more values, unroll selected one;
6. In the table, you will now see *External Endpoint* value associated with *objectscale-gateway-internal*.
7. The endpoint must be of the following format: `https://<IP-ADDRESS>:4443` or `https://<EXTERNAL-HOSTNAME>`

### How to obtain ObjectScale Objectstore endpoint URL?

The easiest way to obtain ObjectScale Objectstore endpoint URL is to look in ObjectScale Portal.


1. Log in to the ObjectScale Portal;
2. From the menu on left side of the screen select *Administration* tab;
3. After unfolding *Administration* tab enter *ObjectScale* page;
4. Select one of the object stores visible in the table, and click its name;
5. You should see *Summary* of that object store.
6. In the *Management Service details* section, you will see value under *IP address* column.
7. The endpoint must be of the following format: `https://<IP-ADDRESS>:4443` or `https://<EXTERNAL-HOSTNAME>`
