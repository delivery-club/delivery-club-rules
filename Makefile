test:
	go test --count=1 -race .

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --skip-dirs testdata --disable deadcode,unused --config=.golangci.yml
	@echo "Running go-critic"
	@gocritic check -enable='#experimental' ./...
