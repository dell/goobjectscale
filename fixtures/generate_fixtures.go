// Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// This software contains the intellectual property of Dell Inc.
// or is licensed to Dell Inc. from third parties. Use of this software
// and the intellectual property contained therein is expressly limited to the
// terms and conditions of the License Agreement under which it is provided by or
// on behalf of Dell Inc. or its subsidiaries.

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
	"gopkg.in/yaml.v3"
)

const (
	EndpointURL = ""
	Username    = ""
	Password    = ""
	bucket1Name = "bucket1"
	bucket2Name = "bucket2"
	namespace   = "ns1"
	user        = "test_user_1"
	objectUser  = "object_user1"
)

func NewRecorder(name string) (*recorder.Recorder, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // #nosec G402
		},
	}

	r, err := recorder.New(filepath.Join("testdata", strings.ReplaceAll(name, "/", "_")),
		recorder.WithRealTransport(transport),
		recorder.WithHook(func(i *cassette.Interaction) error {
			i.Request.Host = ""
			i.Request.RemoteAddr = ""
			i.Response.Duration = 0
			i.Request.URL = replaceIPWithHostname(i.Request.URL)
			i.Request.Headers = nil
			return nil
		}, recorder.BeforeSaveHook))
	if err != nil {
		log.Fatal(err)
	}
	return r, nil
}

func main() {
	recorder, err := NewRecorder("fixtures")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recorder.Stop(); err != nil {
			log.Fatal(err)
		}
		formatFile()
	}()

	httpClient := &http.Client{
		Transport: recorder,
	}

	objectscaleAuthUser := client.AuthUser{
		Gateway:  EndpointURL,
		Username: Username,
		Password: Password,
	}

	clientset := rest.NewClientSet(&client.Simple{
		Endpoint:       EndpointURL,
		Authenticator:  &objectscaleAuthUser,
		OverrideHeader: false,
		HTTPClient:     httpClient,
	})

	// Bucket APIs
	_, err = clientset.Buckets().Get(context.Background(), bucket1Name, map[string]string{"namespace": namespace})
	if err != nil {
		log.Fatal(err)
	}

	_, err = clientset.Buckets().Create(context.Background(), &model.ObjectBucketParam{
		Name:      bucket2Name,
		Namespace: namespace,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = clientset.Buckets().List(context.Background(), map[string]string{"namespace": namespace})
	if err != nil {
		log.Fatal(err)
	}

	// update policy
	bucketPolicyJSON := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": ["*"],
				"Resource": [
					"arn:aws:s3:::bucket2/*"
				],
				"Principal": {
					"AWS": "urn:ecs:iam::ns1:user/iamuser"
				}
			}
		]
	}`
	err = clientset.Buckets().UpdatePolicy(context.Background(), bucket2Name, bucketPolicyJSON, map[string]string{"namespace": namespace})
	if err != nil {
		log.Fatal(err)
	}

	bucketPolicy, err := clientset.Buckets().GetPolicy(context.Background(), bucket2Name, map[string]string{"namespace": namespace})
	if err != nil {
		log.Fatal(err)
	}

	if bucketPolicy == "" {
		log.Fatal("policy not found, but should be there")
	}

	err = clientset.Buckets().DeletePolicy(context.Background(), bucket2Name, map[string]string{"namespace": namespace})
	if err != nil {
		log.Fatal(err)
	}

	// VPool Service
	_, err = clientset.VPools().List(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Clean up
	err = clientset.Buckets().Delete(context.Background(), bucket2Name, map[string]string{"namespace": namespace})
	if err != nil {
		log.Fatal(err)
	}
}

func formatFile() {
	// Path to your YAML file
	filePath := "testdata/fixtures.yaml"

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal into a generic map
	var content map[string]interface{}
	if err := yaml.Unmarshal(data, &content); err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	file, err := os.Create("testdata/fixtures.yaml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	defer encoder.Close()

	if err := encoder.Encode(content); err != nil {
		fmt.Println("Error encoding YAML:", err)
		return
	}
}

func replaceIPWithHostname(str string) string {
	// Regular expression pattern to match IP addresses
	ipPattern := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	// Replace IP addresses with a fake hostname
	return ipPattern.ReplaceAllStringFunc(str, func(ip string) string {
		// Check if the matched string is a valid IP address
		if net.ParseIP(ip) != nil {
			return "testgateway"
		}
		return ip
	})
}
