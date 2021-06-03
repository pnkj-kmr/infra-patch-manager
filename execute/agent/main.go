package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/pnkj-kmr/patch/service"
	"github.com/pnkj-kmr/patch/service/pb"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("server start on port : %d", *port)

	pingServer := service.NewPatchServer()
	grpcServer := grpc.NewServer()
	pb.RegisterPatchServer(grpcServer, pingServer)
	// // TO DEBUG THE gRPC SERVICE with help to
	// // EVANS Client --- https://github.com/ktr0731/evans
	// reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start the server agent ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start the grpc server agent ", err)
	}

	// d, err := jsn.GetRemotes()
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(">>>>remotes", d)
}
