.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: prepare-lint
prepare-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.21.0
