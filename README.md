# Dell ObjectScale Management Go SDK

This project provides a development SDK for the ObjectScale object store management API for Go based applications

## Example

```go
restClient := rest.NewClientSet(client.NewClient("https://ecs-hostname:4443", "https://ecs-hostname:443", "user", "password", &http.Client{}, true))
listParams := map[string]string{}
bucketList := restClient.Buckets().List(listParams)
```
