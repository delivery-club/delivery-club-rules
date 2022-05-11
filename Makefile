GOPATH_DIR=$(shell go env GOPATH)
VERSION=$(shell git describe --tags)

test:
	go test --count=1 -race .

lint:
	ruleguard -rules=rules.go ./...

ci-lint: install-linter lint

install-linter:
	@go install  github.com/quasilyte/go-ruleguard/cmd/ruleguard@cb19258d2ade88dbf466420bb4585dc747bcec57

generate:
	go generate ./...

ci-generate: generate
	git diff --exit-code --quiet || (echo "Please run 'go generate ./...' to update precompiled rules."; false)

install: generate
	go install -ldflags "-s -w -X ./cmd/dcRules.VERSION=${VERSION}" ./cmd/dcRules

build:
	go build -o bin/dcRules -ldflags "-s -w -X ./cmd/dcRules.VERSION=${VERSION}" ./cmd/dcRules

draft-release:
	go run releaser/release.go -version ${VERSION}
