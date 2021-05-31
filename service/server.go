package service

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/module/tar"
	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
func (p *PatchServer) Ping(
	ctx context.Context,
	req *pb.PingRequest,
) (res *pb.PingResponse, err error) {
	msg := req.GetMsg()
	res = &pb.PingResponse{
		Msg: utility.Ping(msg),
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
		// log.Println("Receiving file data...")
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

	// Assets directory - Default patch hold directory
	assetDir, err := dir.New(utility.AssetsDirectory)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}
	// Writeing the file into directory
	fileSizeWritten, err := assetDir.CreateAndWriteFile(fileName+fileType, fileData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}
	// Patch (remedy) directory
	remedyDir, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot patch directory.. update config.env file: %v", err))
	}
	err = remedyDir.Clean()
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Cannot clean patch directory: %v", err))
	}
	// untaring the uploaded file
	t := tar.New(fileName, fileType, utility.AssetsDirectory)
	err = t.Untar(utility.RemedyDirectory)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Unable to extract file into patch directory: %v", err))
	}
	// Scan the remedy dir for all files
	files, err := remedyDir.Scan()
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Unable to scan patch directory: %v", err))
	}
	var fileList []*pb.FILE
	for _, f := range files {
		fileList = append(fileList, &pb.FILE{
			Isdir: f.IsDir(),
			File:  f.Name(),
			Path:  f.Path(),
			Size:  f.Size(),
			Time:  timestamppb.New(f.ModTime()),
		})
	}

	res := &pb.UploadFileResponse{
		Name: fileName + fileType,
		Size: uint64(fileSize),
		Data: fileList,
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "response error: %v", err))
	}
	log.Println("File saved into assets successfully |", fileName, fileSize, fileSizeWritten, (int64(fileSize) == fileSizeWritten))
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
