package task

import (
	"fmt"
	"log"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/service/pb"
)

// UploadResult defines the returns result of remotes
type UploadResult struct {
	Remote string
	File   string
	Size   uint64
	Data   []*pb.FILE
	Ok     bool
	Err    error
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
	res, err := c.UploadFile(file.Path())
	uFile = res.GetName()
	if res.GetSize() != uint64(file.Size()) {
		return uFile, fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), res.GetSize())
	}
	return
}

// PatchFileUploadToAll uploads the file to all remotes
func (t *PatchTask) PatchFileUploadToAll(path string) (out map[string]*UploadResult) {
	out = make(map[string]*UploadResult)
	for _, c := range t.r.GetAll() {
		if !c.Ok {
			log.Println("Connection is not stable with remote", c.Remote.Name)
			out[c.Remote.Name] = &UploadResult{
				Remote: c.Remote.Name,
				File:   "",
				Size:   0,
				Ok:     c.Ok,
				Err:    fmt.Errorf("connection to remote failed: %s", c.Remote.Name),
			}
			continue
		}
		file, err := dir.New(path)
		if err != nil {
			out[c.Remote.Name] = &UploadResult{
				Remote: c.Remote.Name,
				File:   "",
				Size:   0,
				Ok:     c.Ok,
				Err:    err,
			}
			continue
		}
		res, err := c.UploadFile(file.Path())
		if err != nil {
			out[c.Remote.Name] = &UploadResult{
				Remote: c.Remote.Name,
				File:   res.GetName(),
				Size:   res.GetSize(),
				Data:   res.GetData(),
				Ok:     c.Ok,
				Err:    err,
			}
			continue
		}
		if res.GetSize() != uint64(file.Size()) {
			out[c.Remote.Name] = &UploadResult{
				Remote: c.Remote.Name,
				File:   res.GetName(),
				Size:   res.GetSize(),
				Data:   res.GetData(),
				Ok:     c.Ok,
				Err:    fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), res.GetSize()),
			}
			continue
		}
		out[c.Remote.Name] = &UploadResult{
			Remote: c.Remote.Name,
			File:   res.GetName(),
			Size:   res.GetSize(),
			Data:   res.GetData(),
			Ok:     c.Ok,
			Err:    err,
		}
	}
	return
}
