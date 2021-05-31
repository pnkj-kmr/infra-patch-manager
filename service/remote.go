package service

import (
	"log"

	"github.com/pnkj-kmr/patch/module/jsn"
	"github.com/pnkj-kmr/patch/service/pb"
)

// Remote defines the remote config interface
// Actions method with related to Remote Config
type Remote interface {
	Get(string) *ClientInfo
	GetAll() []*ClientInfo
}

// RemoteConfig holds all remotes details
type RemoteConfig struct {
	RemoteMap map[string]jsn.Remote
}

// NewRemoteConfig returns the new remote clients objects
func NewRemoteConfig() (r *RemoteConfig, err error) {
	remotes, err := jsn.GetRemotes()
	if err != nil {
		return nil, err
	}
	rmap := make(map[string]jsn.Remote)
	for _, r := range remotes {
		rmap[r.Name] = r
	}
	return &RemoteConfig{RemoteMap: rmap}, nil
}

// Get returns the remote client by name
func (r *RemoteConfig) Get(remote string) *ClientInfo {
	site, ok := r.RemoteMap[remote]
	if !ok {
		log.Println("Remote not exists:", remote, site)
		return &ClientInfo{
			Ok:   false,
			Name: remote,
			pc:   pb.NewPatchClient(nil),
		}
	}
	return NewClientInfo(site)
}

// GetAll returns all remote clients
func (r *RemoteConfig) GetAll() (c []*ClientInfo) {
	for _, site := range r.RemoteMap {
		c = append(c, NewClientInfo(site))
	}
	return
}
