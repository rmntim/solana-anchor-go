.PHONY: build
build:
	go build

.PHONY: restaking
restaking: build
	rm -rf ./generated/restaking
	./solana-anchor-go -src=./example/restaking_idl.json -pkg=restaking -dst=./generated/restaking

.PHONY: test
test: restaking
	go test ./generated/restaking
	go test ./example/restaking_test.go

.PHONY: clean
clean:
	rm -rf ./generated/restaking
	rm -rf ./solana-anchor-go