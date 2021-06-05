.PHONY:
	clean run gen test tidy master agent clean_grpc clean_build


APP = master
BUILD_DIR = $(PWD)/build

clean_grpc:
	rm -rf service/pb/*.go

gen: clean_grpc
	protoc --proto_path=service/proto --go_out=. --go-grpc_out=. service/proto/*.proto

tidy:
	go mod tidy

run:
	go run -race main.go

swag:
	swag init -g execute/master/main.go

clean_build:
	rm -rf ./build

build: swag clean_build
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) execute/master/main.go

agent:
	go run execute/agent/main.go -port 8081

master: swag
	go run execute/master/main.go

test:
	go test -cover -race ./...
