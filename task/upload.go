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
	Err    error
}

func getUploadResult(remote, file string, size uint64, data []*pb.FILE, err error) *UploadResult {
	return &UploadResult{
		Remote: remote,
		File:   file,
		Size:   size,
		Data:   data,
		Err:    err,
	}
}

// PatchFileUploadTo uploads the file to given remote
func (t *PatchTask) PatchFileUploadTo(remote, path string) (out *UploadResult) {
	log.Println(remote, "UPLOAD: sending request data -", path)
	c := t.r.Get(remote)
	if !c.Ok {
		log.Println("Connection is not stable with remote", remote)
		out = getUploadResult(remote, "", 0, nil, fmt.Errorf("connection to remote failed: %s", remote))
		return
	}
	file, err := dir.New(path)
	if err != nil {
		out = getUploadResult(remote, "", 0, nil, err)
		return
	}
	res, err := c.UploadFile(file.Path())
	out = getUploadResult(remote, res.GetName(), res.GetSize(), res.GetData(), err)
	if res.GetSize() != uint64(file.Size()) {
		out = getUploadResult(remote, res.GetName(), res.GetSize(), res.GetData(), fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), res.GetSize()))
		return
	}
	log.Println(out.Remote, "UPLOAD: received response data -", out.File, out.Size)
	return
}

// PatchFileUploadToAll uploads the file to all remotes
func (t *PatchTask) PatchFileUploadToAll(path string) (out map[string]*UploadResult) {
	out = make(map[string]*UploadResult)
	for _, c := range t.r.GetAll() {
		log.Println(c.Remote.Name, "UPLOAD: sending request data -", path)
		if !c.Ok {
			log.Println("Connection is not stable with remote", c.Remote.Name)
			out[c.Remote.Name] = getUploadResult(c.Remote.Name, "", 0, nil, fmt.Errorf("connection to remote failed: %s", c.Remote.Name))
			continue
		}
		file, err := dir.New(path)
		if err != nil {
			out[c.Remote.Name] = getUploadResult(c.Remote.Name, "", 0, nil, err)
			continue
		}
		res, err := c.UploadFile(file.Path())
		if err != nil {
			out[c.Remote.Name] = getUploadResult(c.Remote.Name, res.GetName(), res.GetSize(), res.GetData(), err)
			continue
		}
		if res.GetSize() != uint64(file.Size()) {
			out[c.Remote.Name] = getUploadResult(
				c.Remote.Name, res.GetName(), res.GetSize(), res.GetData(),
				fmt.Errorf("Uploaded file size does not match: uploaded :%d | given: %d", file.Size(), res.GetSize()),
			)
			continue
		}
		out[c.Remote.Name] = getUploadResult(c.Remote.Name, res.GetName(), res.GetSize(), res.GetData(), err)
		log.Println(c.Remote.Name, "UPLOAD: received response data -", out[c.Remote.Name].File, out[c.Remote.Name].Size)
	}
	return
}
