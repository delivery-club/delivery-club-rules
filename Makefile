test:
	go test --count=1 -race .

lint:
	@golangci-lint cache clean
	@echo "Running golangci-lint..."
	@golangci-lint run --skip-dirs testdata --disable deadcode,unused
	@echo "Running go-critic"
	@gocritic check -enable='#experimental' ./...
