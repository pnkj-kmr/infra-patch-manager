package service

import (
	"bytes"
	"context"
	"io"
	"log"
	"strconv"

	"github.com/pnkj-kmr/infra-patch-manager/service/action"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxFileSize = 5 * 1 << 20 // 5 MB file - max file

// PatchServer struct to handle the ping service request
type PatchServer struct {
	pb.UnimplementedPatchServer
}

// NewPatchServer returns a new ping server
func NewPatchServer() *PatchServer {
	return &PatchServer{struct{}{}}
}

func (p *PatchServer) mustEmbedUnimplementedPatchServer() {}

// Ping defines the PING-PONG
func (p *PatchServer) Ping(ctx context.Context, req *pb.PingRequest) (res *pb.PingResponse, err error) {
	msg := req.GetMsg()
	res = &pb.PingResponse{
		Msg: utility.Ping(msg),
	}
	log.Println("PING: request -", msg, "| response -", res.GetMsg())
	return
}

// RightsCheck helps to verify the target folder read/write rights
func (p *PatchServer) RightsCheck(ctx context.Context, req *pb.RightsCheckRequest) (res *pb.RightsCheckResponse, err error) {
	apps := req.GetRemoteApps()
	log.Println("Rights check request receieved for apps", apps)

	checkApps := make([]*pb.AppRightsInfo, len(apps))
	for i, app := range apps {
		match, err := action.RemoteRWRights(app.GetSource())
		if err != nil {
			log.Println("Rights check error", err)
		}
		checkApps[i] = &pb.AppRightsInfo{
			RemoteApp: app,
			HasRights: match,
		}
	}
	res = &pb.RightsCheckResponse{Apps: checkApps}
	return
}

// UploadFile uploads the file data to server(remote) with client streaming rpc
func (p *PatchServer) UploadFile(stream pb.Patch_UploadFileServer) (err error) {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive file info"))
	}
	fileName := req.GetInfo().GetFileName()
	fileType := req.GetInfo().GetFileType()
	log.Println("UPLOAD: files info received", fileName, fileType)

	fileData := bytes.Buffer{}
	fileSize := 0
	for {
		// checking upload is cancel by send
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		// log.Println("Receiving file data...")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("UPLOAD: No more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "Cannot receieve chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		fileSize += len(chunk)
		if fileSize > maxFileSize {
			return logError(status.Errorf(codes.InvalidArgument, "File is too large: %d > %d", fileSize, maxFileSize))
		}

		// slow writing data into buffer
		// time.Sleep(time.Second)

		_, err = fileData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "Cannot write chunk data: %v", err))
		}
	}

	// post actions - file save, untar, dir scan
	fileList, fileSizeWritten, err := action.PostActionAfterUploadFile(fileName, fileType, fileData)
	if err != nil {
		return
	}

	res := &pb.UploadFileResponse{
		Name: fileName + fileType,
		Size: uint64(fileSize),
		Data: fileList,
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot sent file upload response: %v", err))
	}
	log.Println("File uploaded |", fileName, fileSize, fileSizeWritten, (int64(fileSize) == fileSizeWritten))
	return
}

// ApplyPatch helps to apply patch at given remote applications with server-streaming
func (p *PatchServer) ApplyPatch(req *pb.ApplyPatchRequest, stream pb.Patch_ApplyPatchServer) (err error) {
	apps := req.GetRemoteApps()
	log.Println("Apply patch request receieved for apps", apps)

	found := func(r string, v bool, d []*pb.FILE) error {
		res := &pb.ApplyPatchResponse{
			RemoteApp: r, Verified: v, Data: d,
		}
		err := stream.Send(res)
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot sent data: %v", err))
		}
		log.Println("PATCH applied for remote app:", res.GetRemoteApp(), ", Verified patch:", strconv.FormatBool(res.GetVerified()))
		return nil
	}

	for i, path := range apps {
		// checking upload is cancel by send
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		err = action.ApplyPatchTo(path, i == 0)
		if err != nil {
			return err
		}
		files, verified, err := action.VerifyPatch(path)
		if err != nil {
			return err
		}
		err = found(path, verified, files)
		if err != nil {
			return err
		}
	}
	return
}
