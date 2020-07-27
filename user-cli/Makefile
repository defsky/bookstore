
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
    
	protoc --proto_path=../proto --micro_out=${MODIFY}:../proto --go_out=${MODIFY}:../proto ../proto/user/user.proto
    

.PHONY: build
build: 

	go build -o user-cli *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-cli:latest