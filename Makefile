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

.PHONY: actions action-help
actions: ## Run all GitHub Action checks that run on a pull request creation
	@echo "Running all GitHub Action checks for pull request events..."
	@act -l | grep -v ^Stage | grep pull_request | grep -v image_security_scan | awk '{print $$2}' | while read WF; do \
		echo "Running workflow: $${WF}"; \
		act pull_request --no-cache-server --platform ubuntu-latest=ghcr.io/catthehacker/ubuntu:act-latest --job "$${WF}"; \
	done

action-help: ## Echo instructions to run one specific workflow locally
	@echo "GitHub Workflows can be run locally with the following command:"
	@echo "act pull_request --no-cache-server --platform ubuntu-latest=ghcr.io/catthehacker/ubuntu:act-latest --job <jobid>"
	@echo ""
	@echo "Where '<jobid>' is a Job ID returned by the command:"
	@echo "act -l"
	@echo ""
	@echo "NOTE: if act is not installed, it can be downloaded from https://github.com/nektos/act"