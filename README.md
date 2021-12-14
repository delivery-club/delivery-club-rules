# delivery-club-rules

Delivery club go rules linter

[![Go Report Card](https://goreportcard.com/badge/github.com/delivery-club/delivery-club-rules)](https://goreportcard.com/report/github.com/delivery-club/delivery-club-rules)
[![Go Reference](https://pkg.go.dev/badge/github.com/delivery-club/delivery-club-rules.svg)](https://pkg.go.dev/github.com/delivery-club/delivery-club-rules)

### How to use:

1. Copy rules.go file to your project
2. Add check to your pipeline:
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

### How to add new checks:

Ruleguard tour for newbees: https://go-ruleguard.github.io/by-example
