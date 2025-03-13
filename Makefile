restaking:
	rm -rf ./generated/restaking
	go build 
	./solana-anchor-go -src=./example/restaking_idl.json -pkg=restaking -dst=./generated/restaking

test: restaking
	go test ./generated/restaking
	go test ./example/restaking_test.go