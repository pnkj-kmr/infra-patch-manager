package task

import (
	"fmt"
	"log"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/service"
	"github.com/pnkj-kmr/patch/service/pb"
)

// // Task defines all task action/operation functions
// type Task interface {
// 	Ping(string) string
// 	SaveFile(string, string, bytes.Buffer) error
// 	GetFile(string) (module.I, error)
// }

// PatchTask declares all patch action related tasks
type PatchTask struct {
	r service.Remote
}

// UploadResult defines the returns result of remotes
type UploadResult struct {
	Remote string
	File   string
	Size   uint64
	Data   []*pb.FILE
	Ok     bool
	Err    error
}

// NewPatchTask return new task struct
func NewPatchTask() (t *PatchTask, err error) {
	r, err := service.NewRemoteConfig()
	if err != nil {
		return &PatchTask{nil}, err
	}
	return &PatchTask{r}, nil
}

// PingTo defines the grpc Ping method call
func (t *PatchTask) PingTo(remote, msg string) (out string) {
	c := t.r.Get(remote)
	return c.PingTo(msg)
}

// PingToAll defines the grpc Ping method call to all remotes
func (t *PatchTask) PingToAll(msg string) (out map[string]string) {
	out = make(map[string]string)
	for _, c := range t.r.GetAll() {
		out[c.Name] = c.PingTo(msg)
	}
	return
}

// PatchFileUploadTo uploads the file to given remote
func (t *PatchTask) PatchFileUploadTo(remote, path string) (uFile string, err error) {
	c := t.r.Get(remote)
	if !c.Ok {
		log.Println("Connection is not stable with remote", remote)
		return uFile, fmt.Errorf("connection to remote failed: %s", remote)
	}
	file, err := dir.New(path)
	if err != nil {
		return "", err
	}
	uFile, uSize, _, err := c.FileUploadTo(file.Path())
	if uSize != uint64(file.Size()) {
		return uFile, fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), uSize)
	}
	return
}

// PatchFileUploadToAll uploads the file to all remotes
func (t *PatchTask) PatchFileUploadToAll(path string) (out map[string]*UploadResult) {
	out = make(map[string]*UploadResult)
	for _, c := range t.r.GetAll() {
		if !c.Ok {
			log.Println("Connection is not stable with remote", c.Name)
			out[c.Name] = &UploadResult{
				Remote: c.Name,
				File:   "",
				Size:   0,
				Ok:     c.Ok,
				Err:    fmt.Errorf("connection to remote failed: %s", c.Name),
			}
			continue
		}
		file, err := dir.New(path)
		if err != nil {
			out[c.Name] = &UploadResult{
				Remote: c.Name,
				File:   "",
				Size:   0,
				Ok:     c.Ok,
				Err:    err,
			}
			continue
		}
		uFile, uSize, uList, err := c.FileUploadTo(file.Path())
		if err != nil {
			uFile = ""
			out[c.Name] = &UploadResult{
				Remote: c.Name,
				File:   uFile,
				Size:   uSize,
				Ok:     c.Ok,
				Err:    err,
			}
			continue
		}
		if uSize != uint64(file.Size()) {
			out[c.Name] = &UploadResult{
				Remote: c.Name,
				File:   uFile,
				Size:   uSize,
				Ok:     c.Ok,
				Err:    fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), uSize),
			}
			continue
		}
		out[c.Name] = &UploadResult{
			Remote: c.Name,
			File:   uFile,
			Size:   uSize,
			Data:   uList,
			Ok:     c.Ok,
			Err:    err,
		}
	}
	return
}
