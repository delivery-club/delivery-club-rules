test:
	go test --count=1 -race .

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --skip-dirs testdata --disable deadcode,unused
	@echo "Running go-critic"
	@gocritic check -enable='#experimental' ./...
