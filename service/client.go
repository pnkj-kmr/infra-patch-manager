package service

import (
	"log"

	"github.com/pnkj-kmr/patch/module/jsn"
	"github.com/pnkj-kmr/patch/service/pb"
	"google.golang.org/grpc"
)

// Clients defines available remotes in the system
type Clients interface {
	Get(string) ClientInfo
	GetAll() []ClientInfo
}

// ClientInfo defines the grpc client with availability status
type ClientInfo struct {
	Ok     bool
	Client pb.PatchClient
}

// RemoteClient holds all remotes details
type RemoteClient struct {
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
	return &RemoteClient{RemoteMap: rmap}, nil
}

// Get returns the remote client by name
func (r *RemoteClient) Get(remote string) ClientInfo {
	site, ok := r.RemoteMap[remote]
	if !ok {
		log.Println("Remote not exists:", remote, site)
		return ClientInfo{
			Ok:     false,
			Client: pb.NewPatchClient(nil),
		}
	}
	return getClientInfo(site)
}

// GetAll returns all remote clients
func (r *RemoteClient) GetAll() (c []ClientInfo) {
	for _, site := range r.RemoteMap {
		c = append(c, getClientInfo(site))
	}
	return
}

func getClientInfo(remote jsn.Remote) ClientInfo {
	conn, err := grpc.Dial(remote.Address, grpc.WithInsecure())
	if err != nil {
		log.Println("Connection dial check:", remote.Address, err)
		return ClientInfo{
			Ok:     false,
			Client: pb.NewPatchClient(nil),
		}
	}
	return ClientInfo{
		Ok:     true,
		Client: pb.NewPatchClient(conn),
	}
}
