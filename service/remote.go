package service

import (
	"log"

	"github.com/pnkj-kmr/patch/module/jsn"
	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/task"
)

// RemoteClient holds all remotes details
type RemoteClient struct {
	task      task.Task
	RemoteMap map[string]jsn.Remote
}

// NewRemoteClient returns the new remote clients objects
func NewRemoteClient() (r *RemoteClient, err error) {
	remotes, err := jsn.GetRemotes()
	if err != nil {
		return nil, err
	}
	rmap := make(map[string]jsn.Remote)
	for _, r := range remotes {
		rmap[r.Name] = r
	}
	return &RemoteClient{task: task.NewPatchTask(), RemoteMap: rmap}, nil
}

// Get returns the remote client by name
func (r *RemoteClient) Get(remote string) *ClientInfo {
	site, ok := r.RemoteMap[remote]
	if !ok {
		log.Println("Remote not exists:", remote, site)
		return &ClientInfo{
			Ok:     false,
			Name:   remote,
			Client: pb.NewPatchClient(nil),
		}
	}
	return NewClientInfo(site)
}

// GetAll returns all remote clients
func (r *RemoteClient) GetAll() (c []*ClientInfo) {
	for _, site := range r.RemoteMap {
		c = append(c, NewClientInfo(site))
	}
	return
}

// PingTo defines the grpc Ping method call
func (r *RemoteClient) PingTo(remote, msg string) (out string) {
	c := r.Get(remote)
	return c.PingTo(msg)
}

// PingToAll defines the grpc Ping method call to all remotes
func (r *RemoteClient) PingToAll(msg string) (out map[string]string) {
	out = make(map[string]string)
	for _, c := range r.GetAll() {
		out[c.Name] = c.PingTo(msg)
	}
	return
}

// UploadFileTo defines the grpc upload file method call
func (r *RemoteClient) UploadFileTo(remote, path string) (fileName string, fileSize uint64, err error) {
	c := r.Get(remote)
	f, err := r.task.GetFile(path)
	if err != nil {
		return "", 0, err
	}
	return c.FileUploadTo(f)
}

// UploadFileToAll defines the grpc upload file method call to all remotes
func (r *RemoteClient) UploadFileToAll(msg string) (out map[string]string) {
	out = make(map[string]string)
	for _, c := range r.GetAll() {
		out[c.Name] = c.PingTo(msg)
	}
	return
}
