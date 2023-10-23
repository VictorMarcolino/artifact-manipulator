# Variables
GINKGO_FLAGS_COMMON = -r -randomize-all -randomize-suites
ACK_GINKGO_DEPRECATIONS = 2.12.1

deps:
	@go install github.com/onsi/ginkgo/v2/ginkgo

.PHONY: tests
tests: deps
	@ACK_GINKGO_DEPRECATIONS=$(ACK_GINKGO_DEPRECATIONS) ginkgo run -r --succinct

.PHONY: inject-githooks
inject-githook:
	@echo "Injecting Git hooks..."
	@echo '#!/bin/sh \nif ! make pre-commit; then echo "pre-commit failed. Please fix before committing."; exit 1; fi; exit 0' > .git/hooks/pre-commit
	@chmod 777 .git/hooks/pre-commit
	@echo '#!/bin/sh \nif ! make pre-push; then echo "pre-push failed. Please fix before committing."; exit 1; fi; exit 0' > .git/hooks/pre-push
	@chmod 777 .git/hooks/pre-push
	@echo "Done"

.PHONY: remove-githook
remove-githook:
	@rm -rf .git/hooks/pre-commit
	@rm -rf .git/hooks/pre-push

.PHONY: govet
govet:
	@go vet ./...

.PHONY: gofmt
gofmt:
	@go fmt ./...

.PHONY: gotidy
gotidy:
	@go mod tidy && git diff --exit-code go.mod go.sum

.PHONY: pre-commit
pre-commit: gofmt govet tests

.PHONY: pre-push
pre-push: gotidy pre-commit