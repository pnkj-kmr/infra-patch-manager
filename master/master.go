package master

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
	"github.com/pnkj-kmr/infra-patch-manager/rpc"
	"github.com/pnkj-kmr/infra-patch-manager/rpc/pb"
)

var maxSessionTimeout time.Duration = 60 * time.Second
var defaultFileExt string = ".gz"

type _master struct {
	remote  remote.Remote
	connect pb.PatchClient
	logger  bool
}

// NewPatchMaster - pointer of a remote with extra binder
func NewPatchMaster(name string, logger bool) (PatchMaster, error) {
	remote, err := remote.NewRemote(name)
	if err != nil {
		return nil, err
	}
	conn, err := rpc.Connection(remote.AgentAddress())
	if err != nil {
		return nil, err
	}
	return &_master{remote, pb.NewPatchClient(conn), logger}, nil
}

func (m *_master) log(v ...interface{}) {
	if m.logger {
		log.Println(v...)
	}
}

func (m *_master) Ping() (ok bool, err error) {
	m.log(m.remote.Name(), "PING: request with data request - PING")
	req := &pb.PingReq{Msg: "PING"}
	res, err := m.connect.Ping(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "Ping request failed:", err)
		return false, err
	}
	if strings.EqualFold(res.GetMsg(), "PONG") && err == nil {
		ok = true
	}
	m.log(m.remote.Name(), "PING: receieved response with data - ", res.GetMsg())
	return
}

func (m *_master) UploadFileToRemote(in entity.File) (out entity.Entity, ok bool, err error) {
	m.log(m.remote.Name(), "Patch file upload request received", in.Path())

	ctx, cancel := context.WithTimeout(context.Background(), maxSessionTimeout)
	defer cancel()
	stream, err := m.connect.Upload(ctx)
	if err != nil {
		m.log(m.remote.Name(), "Unable to connect remote during upload", err)
		return
	}

	req := &pb.UploadReq{
		Data: &pb.UploadReq_Info{
			Info: &pb.FileInfo{
				FileName: in.Name(),
				FileExt:  defaultFileExt, // filepath.Ext(path) --> gives .gz if x.tar.gz file
				FileInfo: rpc.EntityToFILE(in),
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		m.log(m.remote.Name(), "Cannot send file info to server", err, stream.RecvMsg(nil))
		return
	}
	file, err := os.Open(in.Path())
	if err != nil {
		m.log(m.remote.Name(), "Unable to open the given path file", in.Path(), in.Name(), err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			m.log(m.remote.Name(), "cannot read chunk to buffer", err)
			return nil, ok, err
		}

		req := &pb.UploadReq{
			Data: &pb.UploadReq_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			m.log(m.remote.Name(), "cannot send chunk to server", err, stream.RecvMsg(nil))
			return nil, ok, err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		m.log(m.remote.Name(), "Upload file receieve response failed", err)
		return
	}

	out = rpc.FILEToEntity(res.GetInfo().GetFileInfo())

	if out.Size() == in.Size() {
		ok = true
	}

	m.log(m.remote.Name(), "File uploaded :", out.Name(), ", size: ", out.Size())
	return
}

func (m *_master) ExtractFileToRemote(path, name string) (files []entity.Entity, ok bool, err error) {
	m.log(m.remote.Name(), "EXTRACT: request with data request  - ", path, name)

	req := &pb.ExtractReq{
		Path: path,
		Name: name,
	}
	res, err := m.connect.Extract(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "ExtractFile request failed:", err)
		return
	}

	for _, f := range res.GetData() {
		files = append(files, rpc.FILEToEntity(f))
	}
	ok = res.GetVerified()
	m.log(m.remote.Name(), "EXTRACT: receieved response with status - ", ok)
	return
}

func (m *_master) RightsCheckFor(reqApps []remote.App) (apps []remote.App, err error) {
	m.log(m.remote.Name(), "RIGHTS: request with data request  - ", len(reqApps))
	var rr []*pb.APP

	for _, a := range reqApps {
		rr = append(rr, m.RemoteAppToAPP(a))
	}

	req := &pb.RightsReq{
		Applications: rr,
	}
	res, err := m.connect.Checks(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "Rights request failed:", err)
		return
	}

	resApps := res.GetApplications()
	for _, r := range resApps {
		apps = append(apps, m.APPToRemoteApp(r.GetApp(), r.GetOk(), m.remote.Name(), nil))
	}
	m.log(m.remote.Name(), "Rights receieved response with status")
	return
}

func (m *_master) PatchTo(reqApps []remote.App) (apps []remote.App, err error) {
	m.log(m.remote.Name(), "PATCH: request with data request  - ", len(reqApps))
	var rr []*pb.APP

	for _, a := range reqApps {
		rr = append(rr, m.RemoteAppToAPP(a))
	}

	req := &pb.ApplyReq{
		Applications: rr,
	}
	stream, err := m.connect.Apply(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "patch request failed:", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			m.log(m.remote.Name(), "PATCH: Cannot receieve stream data from server", err)
			return apps, err
		}

		appInfo := res.GetApplications()
		for _, x := range appInfo {
			apps = append(apps, m.APPToRemoteApp(x.GetApp(), x.GetVerified(), m.remote.Name(), x.GetData()))
		}
	}
	return
}

func (m *_master) VerifyFrom(reqApps []remote.App) (apps []remote.App, err error) {
	m.log(m.remote.Name(), "VERIFY: request with data request  - ", len(reqApps))
	var rr []*pb.APP

	for _, a := range reqApps {
		rr = append(rr, m.RemoteAppToAPP(a))
	}

	req := &pb.VerifyReq{
		Applications: rr,
	}
	stream, err := m.connect.Verify(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "verify request failed:", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			m.log(m.remote.Name(), "VERIFY: Cannot receieve stream data from server", err)
			return apps, err
		}

		appInfo := res.GetApplications()
		for _, x := range appInfo {
			apps = append(apps, m.APPToRemoteApp(x.GetApp(), x.GetVerified(), m.remote.Name(), x.GetData()))
		}
	}
	return
}

func (m *_master) ExecuteCmdOnRemote(in, pass string) (out []byte, err error) {
	m.log(m.remote.Name(), "EXECUTE: request with data request: ", in)
	req := &pb.CmdReq{Cmd: in, Pass: pass}
	res, err := m.connect.Execute(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "EXECUTE: request failed:", err)
		return
	}
	out = res.GetOut()
	if res.GetErr() != "" {
		err = errors.New(res.GetErr())
	}
	m.log(m.remote.Name(), "EXECUTE: receieved response data:", string(out), err)
	return
}

func (m *_master) ListAvailablePatches() (out []string, err error) {
	m.log(m.remote.Name(), "LIST: request with data request: ")
	req := &pb.ListUploadedReq{}
	res, err := m.connect.ListUploaded(context.Background(), req)
	if err != nil {
		m.log(m.remote.Name(), "EXECUTE: request failed:", err)
		return
	}
	out = res.GetItems()
	m.log(m.remote.Name(), "LIST: receieved response data:", len(out), err)
	return
}

// RemoteAppToAPP - helps to convert
func (m *_master) RemoteAppToAPP(r remote.App) *pb.APP {
	return &pb.APP{Name: r.Name(), Source: r.SourcePath(), Service: r.ServiceName()}
}

// APPToRemoteApp - helps to convert
func (m *_master) APPToRemoteApp(a *pb.APP, ok bool, r string, f []*pb.FILE) remote.App {
	app, err := remote.NewRemoteApp(a.Name, r)
	if err != nil {
		log.Print(err)
		return nil
	}
	app.UpdateStatus(ok)
	if len(f) > 0 {
		var ef []entity.Entity
		for _, x := range f {
			ef = append(ef, rpc.FILEToEntity(x))
		}
		app.UpdateFiles(ef)
	}
	return app
}
