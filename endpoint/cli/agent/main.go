package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pnkj-kmr/infra-patch-manager/agent"
	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	// enabling the agent mode
	entity.EnableAgentMode()

	port := flag.Int("port", 8008, "the server port")
	flag.Parse()
	log.Printf("server start on port : %d", *port)

	patch := agent.NewPatchServer()
	grpcServer := grpc.NewServer()
	pb.RegisterPatchServer(grpcServer, patch)
	// // TO DEBUG THE gRPC SERVICE with help to
	// // EVANS Client --- https://github.com/ktr0731/evans
	// reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start the server agent ", err)
	}

	// err = grpcServer.Serve(listener)
	// if err != nil {
	// 	log.Fatal("cannot start the grpc server agent ", err)
	// }

	// Graceful shutdwon of server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatalf("listen: %s\n", err)
			done <- syscall.SIGTERM
		}
	}()
	log.Print("server started")

	<-done
	grpcServer.GracefulStop()
	log.Print("server exited properly")

}
