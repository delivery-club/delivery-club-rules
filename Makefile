GOPATH_DIR=`go env GOPATH`

test:
	go test --count=1 -race .

lint:
	ruleguard -rules=rules.go ./...

ci-lint: install-linter lint

install-linter:
	@go install  github.com/quasilyte/go-ruleguard/cmd/ruleguard@20831c4d6142bf041244dea4b8e749e0ea323581
