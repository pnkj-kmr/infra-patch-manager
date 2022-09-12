

APP_NAME = master
BUILD_DIR = $(PWD)/build

clean_rpc:
	rm -rf rpc/pb/*.go

gen: clean_rpc
	protoc --proto_path=rpc/proto --go_out=. --go-grpc_out=. rpc/proto/*.proto

tidy:
	go mod tidy
	
test:
	go test -cover -race ./...

# swag:
# 	swag init -g endpoint/api/master/main.go

# clean_build:
# 	rm -rf ./build

# build: swag clean_build
# 	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) endpoint/cli/master/main.go

# folders:
# 	mkdir -p resources/{patch,rollback,assets}

run:
	go run endpoint/cli/master/main.go

# Adding new agents
agent1:
	go run endpoint/cli/agent/main.go -port 8081

agent2:
	go run endpoint/cli/agent/main.go -port 8082

agent3:
	go run endpoint/cli/agent/main.go -port 8083

agent4:
	go run endpoint/cli/agent/main.go -port 8084

release:
	goreleaser release --snapshot --rm-dist

build:
	goreleaser build --single-target --snapshot --rm-dist

.PHONY:
	clean run gen test tidy master agent clean_grpc clean_build release

# git tag -a v0.2.2 -m "mag"
# git push