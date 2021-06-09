package actions

import (
	"fmt"
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/client"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
)

// VerifyResult defines the returns result of remotes
type VerifyResult struct {
	Remote    string         `json:"remote"`
	RemoteApp string         `json:"remote_app"`
	Verified  bool           `json:"is_verified"`
	Data      []*pb.FILE     `json:"data"`
	Status    service.Status `json:"status"`
}

// VerifyPatchTo defines the grpc Ping method call
func (a *Action) VerifyPatchTo(remote string, apptype string) (out []*VerifyResult) {
	c := a.r.Get(remote)
	return getVerifyResult(c, apptype)
}

// VerifyPatchToAll defines the grpc Ping method call to all remotes
func (a *Action) VerifyPatchToAll(apptype string) (out map[string][]*VerifyResult) {
	out = make(map[string][]*VerifyResult)
	for _, c := range a.r.GetAll() {
		out[c.Remote.Name] = getVerifyResult(c, apptype)
	}
	return
}

func getVerifyResult(c *client.Client, apptype string) (out []*VerifyResult) {
	if c.Ok {
		var targets []string
		for _, app := range c.Remote.Apps {
			if apptype == "" {
				targets = append(targets, app.Source)
			} else if apptype == app.Type {
				targets = append(targets, app.Source)
			}
		}
		if len(targets) == 0 {
			err := fmt.Errorf("VERIFY: Remote is not found: %s", c.Remote.Name)
			out = []*VerifyResult{{
				Remote:    c.Remote.Name,
				RemoteApp: "",
				Verified:  false,
				Data:      nil,
				Status:    service.Status{Ok: false, Err: err.Error()},
			}}
			return
		}
		log.Println(c.Remote.Name, "VERIFY: sending request for applying patch for apps", targets)
		res, err := c.VerifyPatch(targets)
		if err != nil {
			out = []*VerifyResult{{
				Remote:    c.Remote.Name,
				RemoteApp: "",
				Verified:  false,
				Data:      nil,
				Status:    service.Status{Ok: false, Err: err.Error()},
			}}
			return
		}
		out = make([]*VerifyResult, len(res))
		for i, r := range res {
			log.Println(c.Remote.Name, "VERIFY: received response for app", r.GetRemoteApp(), r.GetVerified())
			out[i] = &VerifyResult{
				Remote:    c.Remote.Name,
				RemoteApp: r.GetRemoteApp(),
				Verified:  r.GetVerified(),
				Data:      r.GetData(),
				Status:    service.Status{Ok: true},
			}
		}
	} else {
		err := fmt.Errorf("VERIFY: Remote is not connected: %s", c.Remote.Name)
		out = []*VerifyResult{{
			Remote:    c.Remote.Name,
			RemoteApp: "",
			Verified:  false,
			Data:      nil,
			Status:    service.Status{Ok: false, Err: err.Error()},
		}}
	}
	return
}
