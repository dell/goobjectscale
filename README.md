# Dell EMC ObjectScale Management Go SDK

This project provides a development SDK for the ObjectScale object store management API for Go based applications

## Example

```go
restClient := rest.NewClientSet("username", "password", "https://ecs-hostname:4443", &http.Client{})
listParams := map[string]string{}
bucketList := restClient.Buckets().List(listParams)
```
