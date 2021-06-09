package actions

import (
	"fmt"
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/client"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
)

// UploadResult defines the returns result of remotes
type UploadResult struct {
	Remote string         `json:"remote"`
	File   string         `json:"filename"`
	Size   uint64         `json:"size"`
	Data   []*pb.FILE     `json:"data"`
	Status service.Status `json:"status"`
}

// PatchFileUploadTo uploads the file to given remote
func (a *Action) PatchFileUploadTo(remote, path string) (out *UploadResult) {
	log.Println(remote, "UPLOAD: sending request data -", path)
	c := a.r.Get(remote)
	return getUploadResult(c, path)
}

// PatchFileUploadToAll uploads the file to all remotes
func (a *Action) PatchFileUploadToAll(path string) (out map[string]*UploadResult) {
	log.Println("UPLOAD: sending request data -", path)
	out = make(map[string]*UploadResult)
	for _, c := range a.r.GetAll() {
		res := getUploadResult(c, path)
		out[c.Remote.Name] = res
		log.Println(c.Remote.Name, "UPLOAD: received response data -", res.File, res.Size)
	}
	return
}

func getUploadResult(c *client.Client, path string) (out *UploadResult) {
	if !c.Ok {
		log.Println("Connection is not stable with remote", c.Remote.Name)
		err := fmt.Errorf("connection to remote failed: %s", c.Remote.Name)
		out = &UploadResult{
			Remote: c.Remote.Name,
			File:   "",
			Size:   0,
			Data:   nil,
			Status: service.Status{Ok: false, Err: err.Error()},
		}
		return
	}
	file, err := dir.New(path)
	if err != nil {
		out = &UploadResult{
			Remote: c.Remote.Name,
			File:   "",
			Size:   0,
			Data:   nil,
			Status: service.Status{Ok: false, Err: err.Error()},
		}
		return
	}
	// Calling the uplaod file client function to upload the file
	res, err := c.UploadFile(file.Path())
	if err != nil {
		out = &UploadResult{
			Remote: c.Remote.Name,
			File:   res.GetName(),
			Size:   res.GetSize(),
			Data:   res.GetData(),
			Status: service.Status{Ok: err == nil, Err: err.Error()},
		}
		return
	}
	if res.GetSize() != uint64(file.Size()) {
		err = fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), res.GetSize())
		out = &UploadResult{
			Remote: c.Remote.Name,
			File:   res.GetName(),
			Size:   res.GetSize(),
			Data:   res.GetData(),
			Status: service.Status{Ok: false, Err: err.Error()},
		}
		return
	}
	out = &UploadResult{
		Remote: c.Remote.Name,
		File:   res.GetName(),
		Size:   res.GetSize(),
		Data:   res.GetData(),
		Status: service.Status{Ok: true},
	}
	log.Println(out.Remote, "UPLOAD: received response data -", out.File, out.Size)
	return
}
