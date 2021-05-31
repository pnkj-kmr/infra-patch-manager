package main

import (
	"log"
	"path/filepath"

	"github.com/pnkj-kmr/patch/task"
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

	task, err := task.NewPatchTask()
	if err != nil {
		log.Fatal("task object errors:", err)
	}

	msg := "ping"
	remoteStat := task.PingToAll(msg)
	for k, v := range remoteStat {
		log.Println(">>>>> HOST host:", k, "req:", msg, "res:", v)
	}

	path := filepath.Join("tmp3/test.tar.gz")
	result := task.PatchFileUploadToAll(path)
	for _, r := range result {
		log.Println("FILE UPLOAD : ", r.Remote, r.File, r.Size, r.Ok, r.Err)
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
