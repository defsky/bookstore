
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
	protoc -I../proto --go_out=plugins=grpc:../proto ../proto/user/user.proto

.PHONY: build
build: proto

	go build -o user-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest
