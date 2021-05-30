package main

import (
	"log"

	"github.com/pnkj-kmr/patch/service"
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

	client, err := service.NewRemoteClient()
	if err != nil {
		log.Fatal("Remote errors:", err)
	}

	msg := "ping"
	remoteStat := client.PingToAll(msg)
	for k, v := range remoteStat {
		log.Println(">>>>> HOST host:", k, "req:", msg, "res:", v)
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
