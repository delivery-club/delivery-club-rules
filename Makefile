GOPATH_DIR=`go env GOPATH`

test:
	go test --count=1 -race .

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --skip-dirs testdata --config=.golangci.yml

ci-lint: install-linter lint

install-linter:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_DIR)/bin v1.43.0
	@$(GOPATH_DIR)/bin/golangci-lint run -v
