<!--
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# How to Contribute

Become one of the contributors to this project! We thrive to build a welcoming and open community for anyone who wants to use the project or contribute to it. There are just a few small guidelines you need to follow. To help us create a safe and positive community experience for all, we require all participants to adhere to the [Code of Conduct](CODE_OF_CONDUCT.md).

## Table of Contents

* [Become a contributor](#Become-a-contributor)
* [Submitting issues](#Submitting-issues)
* [Triage issues](#Triage-issues)
* [Your first contribution](#Your-first-contribution)
* [Branching](#Branching)
* [Signing your commits](#Signing-your-commits)
* [Pull requests](#Pull-requests)
* [Code reviews](#Code-reviews)
* [TODOs in the code](#TODOs-in-the-code)

## Become a contributor

You can contribute to this project in several ways. Here are some examples:

* Contribute to the documentation and codebase.
* Report and triage bugs.
* Create feature requests
* Fix bugs and implement features
* Write technical documentation and blog posts, for users and contributors.
* Help others by answering questions about this project.

## Submitting issues

All issues related to the repository, should be submitted [here][issues]. Issues will be triaged and labels will be used to indicate the type of issue. This section outlines the types of issues that can be submitted.

### Report bugs

We aim to track and document everything related to product via the Issues page. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.

Before submitting a new issue, make sure someone hasn't already reported the problem. Look through the [existing issues][issues] for similar issues.

Report a bug by submitting a [bug report][new-bug-report]. Make sure that you provide as much information as possible on how to reproduce the bug.

When opening a Bug please include the following information to help with debugging:

1. Version of relevant software: this software, Kubernetes, Dell Storage Platform, Helm, etc.
2. Details of the issue explaining the problem: what, when, where
3. The expected outcome that was not met (if any)
4. Supporting troubleshooting information. __Note: Do not provide private company information that could compromise your company's security.__

An Issue __must__ be created before submitting any pull request. Any pull request that is created should be linked to an Issue.

### Feature request

If you have an idea of how to improve this project, submit a [feature request][new-feature-request].

### Answering questions

If you have a question and you can't find the answer in the documentation or issues, the next step is to submit a [question][new-question].

We'd love your help answering questions being asked by other users.

## Triage issues

Triage helps ensure that issues resolve quickly by:

* Ensuring the issue's intent and purpose is conveyed precisely. This is necessary because it can be difficult for an issue to explain how an end user experiences a problem and what actions they took.
* Giving a contributor the information they need before they commit to resolving an issue.
* Lowering the issue count by preventing duplicate issues.
* Streamlining the development process by preventing duplicate discussions.

If you don't have the knowledge or time to code, consider helping with _issue triage_. The community will thank you for saving them time by spending some of yours.

Read more about the ways you can [Triage issues](ISSUE_TRIAGE.md).

## Your first contribution

Unsure where to begin contributing? Start by browsing issues labeled `beginner friendly` or `help wanted`.

* [Beginner-friendly][beginner-friendly] issues are generally straightforward to complete.
* [Help wanted][help-wanted] issues are problems we would like the community to help us with regardless of complexity.

When you're ready to contribute, it's time to create a pull request.

## Branching

This repository follows a scaled trunk branching strategy where short-lived branches are created off of the main branch. When coding is complete, the branch is merged back into main after being approved in a pull request code review.

### Steps for branching and contributing

1. Fork the repository.
2. Create a branch off of the main branch.
3. Make your changes and commit them to your branch.
4. If other code changes have merged into the upstream main branch, perform a rebase of those changes into your branch.
5. Open a [pull request][pulls] between your branch and the upstream main branch.
6. Once your pull request has merged, your branch can be deleted.


## Signing your commits

We require that developers sign off their commits to certify that they have permission to contribute the code in a pull request. This way of certifying is commonly known as the [Developer Certificate of Origin (DCO)](https://developercertificate.org/). We encourage all contributors to read the DCO text before signing a commit and making contributions.

GitHub will prevent a pull request from being merged if there are any unsigned commits.

### Signing a commit

GPG (GNU Privacy Guard) will be used to sign commits.  Follow the instructions [here](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/signing-commits) to create a GPG key and configure your GitHub account to use that key.

Make sure you have your user name and e-mail set.  This will be required for your signed commit to be properly verified.  Check the following references:

* Setting up your github user name [reference](https://help.github.com/articles/setting-your-username-in-git/)
* Setting up your e-mail address [reference](https://help.github.com/articles/setting-your-commit-email-address-in-git/)

Once Git and your GitHub account have been properly configured, you can add the -S flag to the git commits:

```console
$ git commit -S -m "your commit message"
# Creates a signed commit
```

### Commit message format

This repository uses the guidelines for commit messages outlined in [How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/)

## Pull Requests

If this is your first time contributing to an open-source project on GitHub, make sure you read about [Creating a pull request](https://help.github.com/en/articles/creating-a-pull-request).

A pull request must always link to at least one GitHub issue. If that is not the case, create a GitHub issue and link it.

To increase the chance of having your pull request accepted, make sure your pull request follows these guidelines:

* Title and description matches the implementation.
* Commits within the pull request follow the formatting guidelines.
* The pull request closes one related issue.
* The pull request contains necessary tests that verify the intended behavior.
* If your pull request has conflicts, rebase your branch onto the main branch.

If the pull request fixes a bug:

* The pull request description must include `Fixes #<issue number>`.
* To avoid regressions, the pull request should include tests that replicate the fixed bug.

The team _squashes_ all commits into one when we accept a pull request. The title of the pull request becomes the subject line of the squashed commit message. We still encourage contributors to write informative commit messages, as they becomes a part of the Git commit body.

We use the pull request title when we generate change logs for releases. As such, we strive to make the title as informative as possible.

Make sure that the title for your pull request uses the same format as the subject line in the commit message.

### Quality Gates for pull requests

GitHub Actions are used to enforce quality gates when a pull request is created or when any commit is made to the pull request. These GitHub Actions enforce our minimum code quality requirement for any code that get checked into the Go code repository. If any of the quality gates fail, it is expected that the contributor will look into the check log, understand the problem and resolve the issue. If help is needed, please feel free to reach out the maintainers of the project for [support](SUPPORT.md).

#### Security scans

* [Golang Security Checker](https://github.com/securego/gosec) inspects source code for security vulnerabilities by scanning the Go AST. It is integrated via [golangci-lint](https://golangci-lint.run/).
* [Malware Scanner](https://github.com/dell/common-github-actions/tree/main/malware-scanner) inspects source code for malware.

#### Code linting, scanning and vetting

[GitHub action](https://github.com/golangci/golangci-lint-action) that analyzes source code to flag programming errors, stylistics errors, and suspicious constructs. Please refer to [golangci-lint](https://golangci-lint.run/) for more information.

#### Code build/test/coverage

[GitHub action](https://github.com/dell/common-github-actions/tree/main/go-code-tester) that runs Go unit tests and checks that the code coverage of each package meets a configured threshold (currently 90%). An error is flagged if a given pull request does not meet the test coverage threshold and blocks the pull request from being merged.

## Code Reviews

All submissions, including submissions by project members, require review. We use GitHub pull requests for this purpose. Consult [GitHub Help](https://help.github.com/articles/about-pull-requests/) for more information on using pull requests.

A pull request must satisfy following for it to be merged:

* A pull request will require at least 2 maintainer approvals, one of which must come from a code owner.
* Maintainers must perform a review to ensure the changes adhere to guidelines laid out in this document.
* If any commits are made after the PR has been approved, the PR approval will automatically be removed and the above process must happen again.

## Code Style

For the Go code in the repository, we expect the code styling outlined in [Effective Go](https://golang.org/doc/effective_go.html). In addition to this, we have the following supplements:

### Handle Errors

See [Effective Go](https://golang.org/doc/effective_go.html#errors) for details on handling errors.

Do not discard errors using _ variables. If a function returns an error, check it to make sure the function succeeded.  Handle the error, return it, or, in truly exceptional situations, panic.  This can be checked using the errcheck tool if you have it installed locally.

Do not log the error if it will also be logged by a caller higher up the call stack;  doing so causes the logs to become repetitive.  Instead, consider wrapping the error in order to provide more detail.  To see practical examples of this, see this bad practice and this preferred practice:

#### Bad

```go
package main

import (
    "errors"
    "log"
)

func main() {
    err := foo()
    if err != nil {
        log.Printf("error: %+v", err)
        return
    }
}

func foo() error {
    err := bar()
    if err != nil {
        log.Printf("error: %+v", err)
        return err
    }
    return nil
}

func bar() error {
    return errors.New("something bad happened")
}
```

#### Preferred

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

var (
    ErrSomethingBad = errors.New("something bad happened")
)

func main() {
    err := foo()
    if err != nil {
        log.Printf("error: %+v", err)
        return
    }
}

func foo() error {
    err := bar()
    if err != nil {
        return fmt.Errorf("calling bar: %w", err)
    }
    return nil
}

func bar() error {
    return ErrSomethingBad
}
```

Do not use the github.com/pkg/errors package as it is now in maintenance mode since Go 1.13+ added official support for error wrapping. See [go1.13-errors](https://blog.golang.org/go1.13-errors) and [errwrap](https://github.com/fatih/errwrap) for more information.

I see that there's a section titled "golangci-lint" in the README file, but it doesn't currently contain any description. Typically, this section would provide information about what golangci-lint is, why it's used in the project, and how contributors should use it.

Here's a sample description you can add to the "golangci-lint" section:

### golangci-lint

[golangci-lint](https://golangci-lint.run/) is a popular static analysis tool for Go code. It helps maintain code quality and consistency by identifying various issues, including code style violations, potential bugs, and other code quality concerns.

In our project, we use golangci-lint to ensure that contributed code meets our coding standards and is free from common issues. Before submitting a pull request, it's important to run golangci-lint on your code to catch any potential problems and ensure that your code adheres to our coding guidelines.

To run golangci-lint on your code, follow these steps:

1. Install golangci-lint by referring to the [official installation guide](https://golangci-lint.run/usage/install/).

2. Navigate to your project directory.

3. Run golangci-lint using the following command:

    ```shell
    make lint
    ```

   This command will analyze your code and provide feedback on any issues it detects.

4. Address any reported issues before submitting your pull request. This helps maintain code quality and ensures that your contribution aligns with our project's standards.

By using golangci-lint, we can collectively improve the quality of our codebase and create a more maintainable project. If you have any questions or encounter issues related to golangci-lint, feel free to reach out to the project maintainers for assistance.

### TODOs in the code

We don't like TODOs in the code or documentation. It is really best if you sort out all issues you can see with the changes before we check the changes in.

<!-- URLs here -->

[pulls]: https://github.com/dell/goobjectscale/pulls
[issues]: https://github.com/dell/goobjectscale/issues
[beginner-friendly]: https://github.com/dell/goobjectscale/issues?q=is%3Aopen+is%3Aissue+label%3A%22beginner+friendly%22
[help-wanted]: https://github.com/dell/goobjectscale/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22
[new-bug-report]: https://github.com/dell/goobjectscale/issues/new?labels=type%2Fbug%2C+needs-triage&template=bug_report.yml&title=%5BBUG%5D%3A
[new-feature-request]: https://github.com/dell/goobjectscale/issues/new?labels=type%2Ffeature-request%2C+needs-triage&template=feature_request.yml&title=%5BFEATURE%5D%3A
[new-question]: https://github.com/dell/goobjectscale/issues/new?labels=type%2Fquestion&template=ask_a_question.yml&title=%5BQUESTION%5D%3A
