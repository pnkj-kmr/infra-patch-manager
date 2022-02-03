package remote

import (
	"errors"
)

// Remote defines the server basic details
type _remote struct {
	RemoteName   string `json:"name"`
	RemoteType   string `json:"type"`
	Agent        string `json:"agent_address"`
	Applications []_app `json:"apps"`
	_status      bool
}

// NewRemote - declare a remote host
func NewRemote(name string) (Remote, error) {
	for _, r := range _remotes {
		if r.RemoteName == name {
			return &r, nil
		}
	}
	return nil, errors.New("REMOTE: does not exists or alter conf/remotes.json")
}

func (r *_remote) Name() (name string) {
	return r.RemoteName
}

func (r *_remote) AgentAddress() (address string) {
	return r.Agent
}

func (r *_remote) Status() (ok bool) {
	return r._status
}

func (r *_remote) Type() string {
	return r.RemoteType
}

func (r *_remote) UpdateStatus(ok bool) {
	r._status = ok
}

func (r *_remote) Apps() (apps []App, err error) {
	for _, a := range r.Applications {
		app, err := NewRemoteApp(a.AppName, r.RemoteName)
		if err != nil {
			continue
		}
		apps = append(apps, app)
	}
	return
}

func (r *_remote) App(name string) (app App, err error) {
	for _, a := range r.Applications {
		if a.AppName == name {
			app, err := NewRemoteApp(a.AppName, r.RemoteName)
			if err == nil {
				return app, nil
			}
		}
	}
	return nil, errors.New("REMOTE APP: does not exists or alter conf/remotes.json")
}

func (r *_remote) AppByType(apptype string) (apps []App, err error) {
	for _, a := range r.Applications {
		if a.Apptype == apptype {
			app, err := NewRemoteApp(a.AppName, r.RemoteName)
			if err == nil {
				apps = append(apps, app)
			}
		}
	}
	return
}
