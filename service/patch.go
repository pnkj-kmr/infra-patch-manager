package service

import (
	"context"
	"log"

	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/task"
)

// PatchServer struct to handle the ping service request
type PatchServer struct {
	task task.Task
	pb.UnimplementedPatchServer
}

// NewPatchServer returns a new ping server
func NewPatchServer() *PatchServer {
	return &PatchServer{&task.PatchTask{}, struct{}{}}
}

func (p *PatchServer) mustEmbedUnimplementedPatchServer() {}

// Ping defines the PING-PONG
func (p *PatchServer) Ping(
	ctx context.Context,
	req *pb.PingRequest,
) (res *pb.PingResponse, err error) {
	ping := req.GetMsg()
	log.Println("PING: data - ", ping)
	res = &pb.PingResponse{
		Msg: p.task.Ping(ping),
	}
	return
}

func (p *PatchServer) mustEmbedUnimplementedPingServiceServer() {}
