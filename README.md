# delivery-club-rules

Delivery club go rules linter

[![Tests](https://github.com/delivery-club/delivery-club-rules/workflows/Tests/badge.svg)](https://github.com/delivery-club/delivery-club-rules/blob/main/.github/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/delivery-club/delivery-club-rules)](https://goreportcard.com/report/github.com/delivery-club/delivery-club-rules)
[![Go Reference](https://pkg.go.dev/badge/github.com/delivery-club/delivery-club-rules.svg)](https://pkg.go.dev/github.com/delivery-club/delivery-club-rules)

### How to use:

Full installation example: https://github.com/peakle/dc-rules-example

1. Install rules:
      ```shell
      go get -v github.com/delivery-club/delivery-club-rules
      ```
2. Create rules.go file in your project like
   in [example](https://github.com/delivery-club/delivery-club-rules/tree/main/example/rules.go)
3. Add linter to your pipeline:
    1. Like another one check in golangci-lint (will work for golangci-lint >v1.27.0):

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
                rules: "YourDir/rules.go"
        ```
    2. Like explicit check WITHOUT golangci-lint:
        1. install binary by `curl` or `go install`:
       ```shell
       curl -sSfL https://raw.githubusercontent.com/delivery-club/delivery-club-rules/master/install.sh | sh -s -- -d -b ${YOUR_DIR} latest
       ```
       OR
       ``` shell
       go install -v github.com/delivery-club/delivery-club-rules/cmd/dcRules
       ```
        2. Run lint:
       ``` shell
       dcRules ./...
       ```
    3. Like file watcher in Goland IDE (will work for golangci-lint >v1.27.0):
        1. add golangci-lint as `File Watcher` in IDE (Preferences -> Tools -> File Watchers -> Add)
        2. set `Arguments` field where `.golangci.yml` file will be like example above:

       ```
       run $FileDir$ --config=$ProjectFileDir$/.golangci.yml
       ```

### How to update to new rules version:

1. update rules version in your go.mod file
2. download new rules version:
   ```shell
   go get github.com/delivery-club/delivery-club-rules@newVersion
   ```
3. if you using golangci-lint update cache:
   ```shell
   golangci-lint cache clean
   ```

### How to add new checks:

1. Ruleguard tour for newbees: https://go-ruleguard.github.io/by-example
2. Fork repo && open PR :D
