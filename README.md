# delivery-club-rules

Delivery club go rules linter

[![Go Report Card](https://goreportcard.com/badge/github.com/delivery-club/delivery-club-rules)](https://goreportcard.com/report/github.com/delivery-club/delivery-club-rules)
[![Go Reference](https://pkg.go.dev/badge/github.com/delivery-club/delivery-club-rules.svg)](https://pkg.go.dev/github.com/delivery-club/delivery-club-rules)

### How to use:

1. Install [ruleguard](https://github.com/quasilyte/go-ruleguard) and DSL package:
      ```shell
      go get -v -u github.com/quasilyte/go-ruleguard/dsl
      go install -v github.com/quasilyte/go-ruleguard/cmd/ruleguard
      ```
2. Copy rules.go file to your project
3. Add check to your pipeline:
   1. Like explicit check:
      ``` shell
      ruleguard -rules rules.go ./...
      ```

   2. Or add like another one check in golangci-lint:

       ``` yaml
       linters:
         enable:
           - gocritic
       linters-settings:
         gocritic:
           enabled-checks:
             - ruleguard
           settings:
             ruleguard:
               rules: "rules.go"
       ```
   3. Like file watcher in Goland IDE (will work for golangci-lint >v1.27.0):
      1. add golangci-lint as `File Watcher` in IDE (Preferences -> Tools -> File Watchers -> Add)
      2. set `Arguments` field where `.golangci.yml` file will be like example above:

      ```
      run $FileDir$ --config=$ProjectFileDir$/.golangci.yml
      ```

### How to add new checks:

Ruleguard tour for newbees: https://go-ruleguard.github.io/by-example
