# Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#      http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
