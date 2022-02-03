.PHONY:
	clean run gen test tidy master agent clean_grpc clean_build release


APP_NAME = master
BUILD_DIR = $(PWD)/build

clean_rpc:
	rm -rf rpc/pb/*.go

gen: clean_rpc
	protoc --proto_path=rpc/proto --go_out=. --go-grpc_out=. rpc/proto/*.proto

tidy:
	go mod tidy

run:
	go run exec/cli/master/main.go

swag:
	swag init -g exec/api/master/main.go

clean_build:
	rm -rf ./build

# build: swag clean_build
# 	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) exec/cli/master/main.go

agent:
	go run exec/cli/agent/main.go -port 8081

master:
	go run -race exec/cli/master/main.go

test:
	go test -cover -race ./...

# folders:
# 	mkdir -p resources/{patch,rollback,assets}

# Adding new agents
agent1:
	go run exec/cli/agent/main.go -port 8081

agent2:
	go run exec/cli/agent/main.go -port 8082

agent3:
	go run exec/cli/agent/main.go -port 8083

agent4:
	go run exec/cli/agent/main.go -port 8084

release:
	goreleaser release --snapshot --rm-dist

build:
	goreleaser build --single-target --snapshot --rm-dist
