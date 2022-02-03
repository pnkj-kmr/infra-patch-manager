package agent

import (
	"bytes"
	"context"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/rpc"
	"github.com/pnkj-kmr/infra-patch-manager/rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxFileSize = 50 * 1 << 20 // 50 MB file - max file

type _ps struct {
	pb.UnimplementedPatchServer
	agentDefault PatchAgent
}

// NewPatchServer returns a new ping server
func NewPatchServer() pb.PatchServer {
	a, err := NewPatchAgent(entity.C.AssetPath(), false)
	if err != nil {
		log.Fatal("Unable to locate conf path")
		return nil
	}
	return &_ps{struct{}{}, a}
}

func (p *_ps) mustEmbedUnimplementedPatchServer() {}

func (p *_ps) Ping(ctx context.Context, req *pb.PingReq) (res *pb.PingResp, err error) {
	res = &pb.PingResp{Msg: entity.Ping(req.GetMsg())}
	log.Println("PING: request -", req.GetMsg(), "| response -", res.GetMsg())
	return
}

func (p *_ps) Checks(ctx context.Context, req *pb.RightsReq) (res *pb.RightsResp, err error) {
	rapps := req.GetApplications()
	log.Println("RIGHTS: check request receieved - ", len(rapps))
	var appInfo []*pb.RightsInfo
	for i, a := range rapps {
		agent, err := NewPatchAgent(a.GetSource(), i == 0) // default backup enlabed for first app
		if err != nil {
			appInfo = append(appInfo, &pb.RightsInfo{App: a, Ok: false})
			continue
		}
		Ok, _ := agent.RightsCheck()
		appInfo = append(appInfo, &pb.RightsInfo{App: a, Ok: Ok})
	}
	res = &pb.RightsResp{Applications: appInfo}
	log.Println("RIGHTS: completed ", len(rapps))
	return
}

func (p *_ps) Upload(stream pb.Patch_UploadServer) (err error) {
	start := time.Now()
	req, err := stream.Recv()
	if err != nil {
		return rpc.LogError(status.Errorf(codes.Unknown, "cannot receive file info"))
	}
	// getting file detail
	fileName := req.GetInfo().GetFileName()
	fileExt := req.GetInfo().GetFileExt()
	fileInfo := req.GetInfo().GetFileInfo()
	log.Println("UPLOAD: files info received - ", fileName, fileExt, fileInfo.GetSize())

	fileData := bytes.Buffer{}
	fileSize := 0
	// loop - getting all file chunk into buffer
	for {
		// checking upload is cancel by send
		err := rpc.ContextError(stream.Context())
		if err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("UPLOAD: No more data")
			break
		}
		if err != nil {
			return rpc.LogError(status.Errorf(codes.Unknown, "Cannot receieve chunk data: %v", err))
		}
		// reading the file chunk
		chunk := req.GetChunkData()
		fileSize += len(chunk)
		if fileSize > maxFileSize {
			return rpc.LogError(status.Errorf(codes.InvalidArgument, "File is too large: %d > %d", fileSize, maxFileSize))
		}
		// slow writing data into buffer
		// time.Sleep(time.Second)
		// writing into buffer
		_, err = fileData.Write(chunk)
		if err != nil {
			return rpc.LogError(status.Errorf(codes.Internal, "Cannot write chunk data: %v", err))
		}
	}

	fileWritten, err := p.agentDefault.WriteUploaded(rpc.FILEToEntity(fileInfo), fileData)
	if fileWritten.Size() != fileInfo.GetSize() {
		return rpc.LogError(status.Errorf(codes.Internal, "Written file is not same as uploaded"))
	}
	log.Println("UPLOAD: files info written - T:", time.Since(start))

	res := &pb.UploadResp{
		Info: &pb.FileInfo{
			FileName: fileWritten.Name(),
			FileExt:  "",
			FileInfo: rpc.EntityToFILE(fileWritten),
		},
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return rpc.LogError(status.Errorf(codes.Unknown, "cannot sent file upload response: %v", err))
	}
	log.Println("UPLOAD : File uploaded successfully |", fileName, "written-", fileWritten.Size(), "received-", fileInfo.GetSize())
	return
}

func (p *_ps) Extract(ctx context.Context, req *pb.ExtractReq) (res *pb.ExtractResp, err error) {
	start := time.Now()
	name := req.GetName()
	path := req.GetPath()
	log.Println("EXTRACT: request receieved -", name, path)
	err = p.agentDefault.PatchExtract("", name)
	if err != nil {
		return
	}
	files, Ok, err := p.agentDefault.VerifyExtracted()
	var ff []*pb.FILE
	for _, f := range files {
		ff = append(ff, rpc.EntityToFILE(f))
	}
	res = &pb.ExtractResp{
		Path:     path,
		Verified: Ok,
		Data:     ff,
	}
	log.Println("EXTRACT: file -", path, name, "T:", time.Since(start))
	return
}

func (p *_ps) Apply(req *pb.ApplyReq, stream pb.Patch_ApplyServer) (err error) {
	start := time.Now()
	apps := req.GetApplications()
	log.Println("Apply patch request receieved for apps", apps)

	found := func(r *pb.APP, v bool, d []*pb.FILE) error {
		var apps []*pb.ApplyInfo
		apps = append(apps, &pb.ApplyInfo{
			App:      r,
			Verified: v,
			Data:     d,
		})
		res := &pb.ApplyResp{
			Applications: apps,
		}
		err := stream.Send(res)
		if err != nil {
			return rpc.LogError(status.Errorf(codes.Unknown, "cannot sent data: %v", err))
		}
		log.Println("PATCH applied for remote app:", r.GetName(), ", Verified patch:", strconv.FormatBool(v))
		return nil
	}

	for i, app := range apps {
		// checking upload is cancel by send
		err := rpc.ContextError(stream.Context())
		if err != nil {
			return err
		}
		agent, err := NewPatchAgent(app.GetSource(), i == 0) // taking first app backup
		if err != nil {
			return err
		}
		// applying patch
		err = agent.PatchNow()
		if err != nil {
			return err
		}
		// verifying the applied patch
		files, Ok, err := agent.VerifyPatched()
		var ff []*pb.FILE
		for _, f := range files {
			ff = append(ff, rpc.EntityToFILE(f))
		}
		err = found(app, Ok, ff)
		if err != nil {
			return err
		}
	}
	log.Println("PATCHED: T:", time.Since(start))
	return
}

func (p *_ps) Verify(req *pb.VerifyReq, stream pb.Patch_VerifyServer) (err error) {
	start := time.Now()
	apps := req.GetApplications()
	log.Println("Verifying patch request receieved for apps", apps)

	found := func(r *pb.APP, v bool, d []*pb.FILE) error {
		var apps []*pb.VerifyInfo
		apps = append(apps, &pb.VerifyInfo{
			App:      r,
			Verified: v,
			Data:     d,
		})
		res := &pb.VerifyResp{
			Applications: apps,
		}
		err := stream.Send(res)
		if err != nil {
			return rpc.LogError(status.Errorf(codes.Unknown, "cannot sent data: %v", err))
		}
		log.Println("Verified for remote app:", r.GetName(), ", Verified patch:", strconv.FormatBool(v))
		return nil
	}

	for _, app := range apps {
		// checking upload is cancel by send
		err := rpc.ContextError(stream.Context())
		if err != nil {
			return err
		}
		agent, err := NewPatchAgent(app.GetSource(), false)
		if err != nil {
			return err
		}
		// verifying the applied patch
		files, Ok, err := agent.VerifyPatched()
		var ff []*pb.FILE
		for _, f := range files {
			ff = append(ff, rpc.EntityToFILE(f))
		}
		err = found(app, Ok, ff)
		if err != nil {
			return err
		}
	}
	log.Println("VERIFIED: T:", time.Since(start))
	return
}

func (p *_ps) Execute(ctx context.Context, req *pb.CmdReq) (res *pb.CmdResp, err error) {
	cmd := req.GetCmd()
	log.Println("EXECUTE: request receieved - ", cmd)
	out, err := entity.ExecuteCmd(cmd)
	log.Println("EXECUTE: completed -", string(out), "\nerr:", err)
	var e string
	if err != nil {
		e = err.Error()
	}
	res = &pb.CmdResp{
		Out: out,
		Err: e,
	}
	return
}

func (p *_ps) ListUploaded(ctx context.Context, req *pb.ListUploadedReq) (res *pb.ListUploadedResp, err error) {
	log.Println("LIST: request receieved")
	out, err := p.agentDefault.ListAssets()
	log.Println("LIST: completed -", len(out), "\nerr:", err)
	var items []string
	for _, i := range out {
		items = append(items, i.Name())
	}
	res = &pb.ListUploadedResp{
		Items: items,
	}
	return
}
