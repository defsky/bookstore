
GOPATH:=$(shell go env GOPATH)
# MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

# .PHONY: proto
# proto:
# 	protoc -I../proto --go_out=plugins=grpc:../proto ../proto/user/user.proto

.PHONY: user
user:
	cd user && make build && cd ..

.PHONY: user-cli
user-cli: user
	cd user-cli && make build && cd ..

.PHONY: user-api
user-api: user
	cd user-api && make build && cd ..

.PHONY: docker
docker: user user-cli user-api
	docker-compose build

.PHONY: run
run: docker
	docker-compose up -d

.PHONY: stop
stop: 
	docker-compose stop

.PHONY: test
test:
	go test -v ./... -cover

