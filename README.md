# :lock: **Important Notice**
Starting with the release of **Container Storage Modules v1.16.0**, this repository will no longer be maintained as an open source project. Future development will continue under a closed source model. This change reflects our commitment to delivering even greater value to our customers by enabling faster innovation and more deeply integrated features with the Dell storage portfolio.<br>
For existing customers using Dell’s Container Storage Modules, you will continue to receive:
* **Ongoing Support & Community Engagement**<br>
       You will continue to receive high-quality support through Dell Support and our community channels. Your experience of engaging with the Dell community remains unchanged.
* **Streamlined Deployment & Updates**<br>
        Deployment and update processes will remain consistent, ensuring a smooth and familiar experience.
* **Access to Documentation & Resources**<br>
       All documentation and related materials will remain publicly accessible, providing transparency and technical guidance.
* **Continued Access to Current Open Source Version**<br>
       The current open-source version will remain available under its existing license for those who rely on it.

Moving to a closed source model allows Dell’s development team to accelerate feature delivery and enhance integration across our Enterprise Kubernetes Storage solutions ultimately providing a more seamless and robust experience.<br>
We deeply appreciate the contributions of the open source community and remain committed to supporting our customers through this transition.<br>

For questions or access requests, please contact the maintainers via [Dell Support](https://www.dell.com/support/kbdoc/en-in/000188046/container-storage-interface-csi-drivers-and-container-storage-modules-csm-how-to-get-support).


> 🚧 **NOTE:** The *goobjectscale* is in a preview stage, and API is subject to change between minor updates. Until `v1.0.0` we do not guarantee backwards compatibility between versions.

# goobjectscale

[![GitHub license](https://img.shields.io/github/license/dell/goobjectscale)](https://github.com/dell/goobjectscale/blob/main/LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/dell/goobjectscale)](https://github.com/dell/goobjectscale/releases/latest)
[![Development Actions](https://github.com/dell/goobjectscale/actions/workflows/development.yaml/badge.svg)](https://github.com/dell/goobjectscale/actions/workflows/development.yaml)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dell/goobjectscale)](https://github.com/dell/goobjectscale/blob/main/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/dell/goobjectscale)](https://goreportcard.com/report/github.com/dell/goobjectscale)
[![Go Reference](https://pkg.go.dev/badge/github.com/dell/goobjectscale.svg)](https://pkg.go.dev/github.com/dell/goobjectscale)

## Table of Contents

* [Code of Conduct](./docs/CODE_OF_CONDUCT.md)
* [Maintainer Guide](./docs/MAINTAINER_GUIDE.md)
* [Committer Guide](./docs/COMMITTER_GUIDE.md)
* [Contributing Guide](./docs/CONTRIBUTING.md)
* [Maintainers](./docs/MAINTAINERS.md)
* [Support](./docs/SUPPORT.md)
* [Security](./docs/SECURITY.md)

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

