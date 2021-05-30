package service

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/patch/module"
	"github.com/pnkj-kmr/patch/module/jsn"
	"github.com/pnkj-kmr/patch/service/pb"
	"google.golang.org/grpc"
)

// // Client defines available remotes in the system
// type Client interface {
// 	Get(string) ClientInfo
// 	GetAll() []ClientInfo
// }

// ClientInfo defines the grpc client with availability status
type ClientInfo struct {
	Ok     bool
	Name   string
	Client pb.PatchClient
}

// NewClientInfo return client object
func NewClientInfo(remote jsn.Remote) *ClientInfo {
	conn, err := grpc.Dial(remote.Address, grpc.WithInsecure())
	if err != nil {
		log.Println("Connection dial check:", remote.Address, err)
		return &ClientInfo{
			Ok:     false,
			Name:   remote.Name,
			Client: pb.NewPatchClient(nil),
		}
	}
	return &ClientInfo{
		Ok:     true,
		Name:   remote.Name,
		Client: pb.NewPatchClient(conn),
	}
}

// PingTo calls the gRPC client
func (c *ClientInfo) PingTo(in string) (out string) {
	if c.Ok {
		req := &pb.PingRequest{Msg: in}
		res, err := c.Client.Ping(context.Background(), req)
		if err != nil {
			log.Println("cannot start the server agent ", err)
		}
		out = res.GetMsg()
		return
	}
	return
}

// FileUploadTo calls upload file gRPC client
func (c *ClientInfo) FileUploadTo(f module.I) (fileName string, fileSize uint64, err error) {
	file, err := os.Open(f.Path())
	if err != nil {
		log.Println("cannot open tar file: ", err)
		return
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := c.Client.UploadFile(ctx)
	if err != nil {
		log.Println("cannot upload file: ", err)
		return
	}

	req := &pb.UploadFileRequest{
		Data: &pb.UploadFileRequest_Info{
			Info: &pb.FileInfo{
				FileName: f.Name(),
				FileType: filepath.Ext(f.Name()),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Println("cannot send file info to server: ", err, stream.RecvMsg(nil))
		return
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("cannot read chunk to buffer: ", err)
			return "", 0, err
		}

		req := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return "", 0, err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("cannot receive response: ", err)
		return "", 0, err
	}

	fileName = res.GetFileName()
	fileSize = res.GetFileSize()
	log.Printf("file uploaded with name: %s, size: %d", fileName, fileSize)
	return
}
