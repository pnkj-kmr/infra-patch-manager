package service

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/pnkj-kmr/patch/module/jsn"
	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/utility"
	"google.golang.org/grpc"
)

const (
	maxSessionTimeout = 5 * time.Second // default timeout for file upload
	defaultFileExt    = ".tar.gz"       // default file extension
)

// ClientInfo defines the grpc client with availability status
type ClientInfo struct {
	Ok     bool
	Remote jsn.Remote
	pc     pb.PatchClient
}

// NewClientInfo return client object
func NewClientInfo(remote jsn.Remote) *ClientInfo {
	conn, err := grpc.Dial(remote.Address, grpc.WithInsecure())
	if err != nil {
		log.Println("Connection dial check for remote:", remote.Address, err)
		return &ClientInfo{
			Ok:     false,
			Remote: remote,
			pc:     pb.NewPatchClient(nil),
		}
	}
	return &ClientInfo{
		Ok:     true,
		Remote: remote,
		pc:     pb.NewPatchClient(conn),
	}
}

// Ping calls the gRPC client
func (c *ClientInfo) Ping(in string) (out string) {
	if c.Ok {
		req := &pb.PingRequest{Msg: in}
		res, err := c.pc.Ping(context.Background(), req)
		if err != nil {
			log.Println(c.Remote.Name, "Ping request failed:", err)
		}
		out = res.GetMsg()
		return
	}
	return
}

// RightsCheck calls the gRPC client for read/write check
func (c *ClientInfo) RightsCheck(apps []jsn.RemoteApp) (out []jsn.RemoteApp, err error) {
	log.Println(c.Remote.Name, "Rights check receieved for apps", apps)
	var rApps []*pb.APP
	for _, app := range apps {
		rApps = append(rApps, &pb.APP{
			Name:    app.Name,
			Path:    app.Path,
			Service: app.Service,
		})
	}
	req := &pb.RightsCheckRequest{RemoteApps: rApps}
	res, err := c.pc.RightsCheck(context.Background(), req)
	if err != nil {
		log.Println(c.Remote.Name, "Rights check response has error as:", err)
		return apps, err
	}
	resApps := res.GetApps()
	var app *pb.APP
	for _, appinfo := range resApps {
		app = appinfo.GetRemoteApp()
		out = append(out, jsn.RemoteApp{
			Name:    app.GetName(),
			Path:    app.GetPath(),
			Service: app.GetService(),
			Status:  jsn.RemoteStatus{Ok: appinfo.GetHasRights()},
		})
	}
	return
}

// UploadFile calls upload file gRPC client
func (c *ClientInfo) UploadFile(path string) (res *pb.UploadFileResponse, err error) {
	log.Println(c.Remote.Name, "Patch file upload request received", path)
	fileName := utility.RandomStringWithTime(0, "PATCH")
	file, err := os.Open(path)
	if err != nil {
		log.Println("cannot open tar file: ", err)
		log.Println(c.Remote.Name, "Unable to open the given path file", path, err)
		return
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), maxSessionTimeout)
	defer cancel()
	stream, err := c.pc.UploadFile(ctx)
	if err != nil {
		log.Println(c.Remote.Name, "Unable to open the path file", err)
		return
	}

	req := &pb.UploadFileRequest{
		Data: &pb.UploadFileRequest_Info{
			Info: &pb.FileInfo{
				FileName: fileName,
				FileType: defaultFileExt, // filepath.Ext(path) --> gives .gz if x.tar.gz file
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Println(c.Remote.Name, "Cannot send file info to server", err, stream.RecvMsg(nil))
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
			log.Println(c.Remote.Name, "cannot read chunk to buffer", err)
			return nil, err
		}

		req := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Println(c.Remote.Name, "cannot send chunk to server", err, stream.RecvMsg(nil))
			return nil, err
		}
	}

	res, err = stream.CloseAndRecv()
	if err != nil {
		log.Println(c.Remote.Name, "Upload file receieve response failed", err)
		return
	}
	log.Println(c.Remote.Name, "File uploaded :", res.GetName(), ", size: ", res.GetSize())
	return
}

// ApplyPatch sending a patch request to remote server
func (c *ClientInfo) ApplyPatch(apps []string) (out []*pb.ApplyPatchResponse, err error) {
	log.Println(c.Remote.Name, "Apply patch receieved for apps", apps)
	ctx, cancel := context.WithTimeout(context.Background(), maxSessionTimeout)
	defer cancel()

	req := &pb.ApplyPatchRequest{RemoteApps: apps}
	stream, err := c.pc.ApplyPatch(ctx, req)
	if err != nil {
		log.Println(c.Remote.Name, "Cannot send apply patch request to server", err)
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(c.Remote.Name, "Cannot receieve stream data from server", err)
			return out, err
		}
		out = append(out, res)
	}
	return
}
