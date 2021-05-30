package main

import (
	"context"
	"log"

	"github.com/pnkj-kmr/patch/service"
	"github.com/pnkj-kmr/patch/service/pb"
)

func main() {
	// address := flag.String("address", "", "the server port")
	// flag.Parse()
	// log.Printf("dail address : %s", *address)

	// conn, err := grpc.Dial(*address, grpc.WithInsecure())
	// addrs := []string{"0.0.0.0:8080", "0.0.0.0:8081"}
	// for _, addr := range addrs {
	// 	rpc(addr)
	// }

	clients, err := service.NewRemoteClient()
	if err != nil {
		log.Fatal("Remote errors:", err)
	}

	for _, c := range clients.GetAll() {
		log.Println("--------c", c)
		if c.Ok {
			req := &pb.PingRequest{Msg: "ping"}
			res, err := c.Client.Ping(context.Background(), req)
			if err != nil {
				log.Println("cannot start the server agent ", err)
			}
			log.Println(">>>>>> RES <<<<<<", res.GetMsg())
		} else {
			log.Println(">>>>> CONNECTION <<<< ", c.Ok, c.Client)
		}
	}
}

// func rpc(addr string) {
// 	conn, err := grpc.Dial(addr, grpc.WithInsecure())
// 	if err != nil {
// 		log.Println("cannot start the server agent ", err)
// 	}

// 	pingClient := pb.NewPatchClient(conn)

// 	req := &pb.PingRequest{Msg: "ping"}
// 	res, err := pingClient.Ping(context.Background(), req)
// 	if err != nil {
// 		log.Println("cannot start the server agent ", err)
// 	}
// 	log.Println(">>>>>> RES <<<<<<", res.GetMsg())
// }
