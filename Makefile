GOPATH_DIR=`go env GOPATH`

test:
	go test --count=1 -race .

lint:
	ruleguard -rules=rules.go ./...

ci-lint: install-linter lint

install-linter:
	@go install  github.com/quasilyte/go-ruleguard/cmd/ruleguard@cb19258d2ade88dbf466420bb4585dc747bcec57
