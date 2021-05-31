gen:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. proto/*.proto

clean:
	rm -rf service/pb/*.go

tidy:
	go mod tidy

run:
	go run -race main.go

agent:
	go run execute/agent/main.go --port 8080

master:
	go run execute/master/main.go --address 0.0.0.0:8080

test:
	go test -cover -race ./...

.PHONY:
	clean run gen test tidy master agent 