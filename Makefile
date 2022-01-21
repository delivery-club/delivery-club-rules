test:
	go test --count=1 -race .

lint:
	ruleguard -rules=rules.go ./...
