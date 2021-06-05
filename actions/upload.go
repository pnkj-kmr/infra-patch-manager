package actions

import (
	"fmt"
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
	"github.com/pnkj-kmr/infra-patch-manager/service/client"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
)

// UploadResult defines the returns result of remotes
type UploadResult struct {
	Remote string
	File   string
	Size   uint64
	Data   []*pb.FILE
	Err    error
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
		out = &UploadResult{
			Remote: c.Remote.Name,
			File:   "",
			Size:   0,
			Data:   nil,
			Err:    fmt.Errorf("connection to remote failed: %s", c.Remote.Name),
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
			Err:    err,
		}
		return
	}
	// Calling the uplaod file client function to upload the file
	res, err := c.UploadFile(file.Path())
	out = &UploadResult{
		Remote: c.Remote.Name,
		File:   res.GetName(),
		Size:   res.GetSize(),
		Data:   res.GetData(),
		Err:    err,
	}
	if err != nil {
		return
	}
	if res.GetSize() != uint64(file.Size()) {
		out = &UploadResult{
			Remote: c.Remote.Name,
			File:   res.GetName(),
			Size:   res.GetSize(),
			Data:   res.GetData(),
			Err:    fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), res.GetSize()),
		}
		return
	}
	log.Println(out.Remote, "UPLOAD: received response data -", out.File, out.Size)
	return
}
