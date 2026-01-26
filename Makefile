# Copyright © 2023 - 2025 Dell Inc. or its subsidiaries. All Rights Reserved.
#
# This software contains the intellectual property of Dell Inc.
# or is licensed to Dell Inc. from third parties. Use of this software
# and the intellectual property contained therein is expressly limited to the
# terms and conditions of the License Agreement under which it is provided by or
# on behalf of Dell Inc. or its subsidiaries.

.PHONY: all
all: clean codegen lint test

.PHONY: help
help:	##show help
	@fgrep --no-filename "##" $(MAKEFILE_LIST) | fgrep --invert-match fgrep | sed --expression='s/\\$$//' | sed --expression='s/##//'

.PHONY: clean
clean:	##clean directory
	rm --force --recursive pkg/client/api/mocks
	go clean

########################################################################
##                                 GO                                 ##
########################################################################

.PHONY: codegen
codegen: clean	##regenerate files
	go generate ./...

########################################################################
##                              TESTING                               ##
########################################################################

.PHONY: test
test: unit-test	##run unit tests

.PHONY: unit-test
unit-test:	##run tests
	( go clean -cache; CGO_ENABLED=0 go test -v -coverprofile=c.out ./...)

.PHONY: unit-test-race
unit-test-race:	##run unit tests with race condition reporting
	( go clean -cache; CGO_ENABLED=1 go test -race -v -coverprofile=c.out ./...)

########################################################################
##                              TOOLING                               ##
########################################################################

.PHONY: lint
lint:	##run golangci-lint over the repository
	golangci-lint run

.PHOHY: generate-fixtures
generate-fixtures:
	cd fixtures && rm -f testdata/fixtures.yaml && go run generate_fixtures.go
