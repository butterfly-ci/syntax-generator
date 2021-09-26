gen:
	protoc -I proto/ --go_out=generated proto/*.proto 

test: gen
	go test ./... -v