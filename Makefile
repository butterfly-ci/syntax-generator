gen:
	protoc  --go_out=generated proto/*.proto

test: gen
	go test ./... -v