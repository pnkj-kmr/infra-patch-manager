package service

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxFileSize = 5 * 1 << 20 // 5 MB file - max file

// PatchServer struct to handle the ping service request
type PatchServer struct {
	task task.Task
	pb.UnimplementedPatchServer
}

// NewPatchServer returns a new ping server
func NewPatchServer() *PatchServer {
	return &PatchServer{task.NewPatchTask(), struct{}{}}
}

func (p *PatchServer) mustEmbedUnimplementedPatchServer() {}

// Ping defines the PING-PONG
func (p *PatchServer) Ping(
	ctx context.Context,
	req *pb.PingRequest,
) (res *pb.PingResponse, err error) {
	msg := req.GetMsg()
	res = &pb.PingResponse{
		Msg: p.task.Ping(msg),
	}
	log.Println("PING: request -", msg, "| response -", res.GetMsg())
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
	log.Println("UPLOAD: files info ", fileName, fileType)

	fileData := bytes.Buffer{}
	fileSize := 0
	for {
		// checking upload is cancel by send
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		log.Println("Receiving file data...")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("No more data")
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

	err = p.task.SaveFile(fileName, fileType, fileData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}

	res := &pb.UploadFileResponse{
		FileName: fileName,
		FileSize: uint64(fileSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "response error: %v", err))
	}
	log.Println("File saved into assets successfully |", fileName, fileSize)
	return
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "canceled by sender"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
