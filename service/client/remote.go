package client

import (
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
)

// // Remote defines the remote config interface
// // Actions method with related to Remote Config
// type Remote interface {
// 	Get(string) *Client
// 	GetAll() []*Client
// }

// RemoteConfig holds all remotes details
type RemoteConfig struct {
	RemoteMap map[string]service.Remote
}

// NewRemoteConfig returns the new remote clients objects
func NewRemoteConfig() (r *RemoteConfig, err error) {
	remotes, err := getRemotes()
	if err != nil {
		return nil, err
	}
	rmap := make(map[string]service.Remote)
	for _, r := range remotes {
		rmap[r.Name] = r
	}
	return &RemoteConfig{RemoteMap: rmap}, nil
}

// Get returns the remote client by name
func (r *RemoteConfig) Get(remote string) *Client {
	site, ok := r.RemoteMap[remote]
	if !ok {
		log.Println("Remote does not exist:", remote, site)
		return &Client{
			Ok:     false,
			Remote: service.Remote{Name: remote},
			pc:     pb.NewPatchClient(nil),
		}
	}
	return NewClient(site)
}

// GetAll returns all remote clients
func (r *RemoteConfig) GetAll() (c []*Client) {
	for _, site := range r.RemoteMap {
		c = append(c, NewClient(site))
	}
	return
}
