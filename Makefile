# cannot be migrated to go tool
# reason: https://golangci-lint.run/welcome/install/#install-from-sources
GOLANGCI_LINT_IMAGE=golangci/golangci-lint:v2.1-alpine

.PHONY: deps
deps:
	docker pull $(GOLANGCI_LINT_IMAGE)

.PHONY: lint
lint:
	$(MAKE) deps

	docker run -t --rm -v $(PWD):/src -w /src $(GOLANGCI_LINT_IMAGE) golangci-lint run

.PHONY: test
test:
	./scripts/test.sh
