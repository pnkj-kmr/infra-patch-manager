package remote

import (
	"errors"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
)

type _app struct {
	AppName string `json:"name"`
	Source  string `json:"source"`
	Port    string `json:"port"`
	Service string `json:"service"`
	Apptype string `json:"type"`
	_status bool
	_files  []entity.Entity
}

// NewRemoteApp - new remote app pointer
func NewRemoteApp(name, remote string) (App, error) {
	for _, r := range _remotes {
		if r.RemoteName == remote {
			for _, a := range r.Applications {
				if a.AppName == name {
					return &a, nil
				}
			}
		}
	}
	return nil, errors.New("REMOTE APP: does not exists or alter conf/remotes.json")
}

func (a *_app) Name() string              { return a.AppName }
func (a *_app) SourcePath() string        { return a.Source }
func (a *_app) AppPort() string           { return a.Port }
func (a *_app) ServiceName() string       { return a.Service }
func (a *_app) Type() string              { return a.Apptype }
func (a *_app) Status() bool              { return a._status }
func (a *_app) GetFiles() []entity.Entity { return a._files }

// update function
func (a *_app) UpdateStatus(ok bool)          { a._status = ok }
func (a *_app) UpdateFiles(f []entity.Entity) { a._files = f }
