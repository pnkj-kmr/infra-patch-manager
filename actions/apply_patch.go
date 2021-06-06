package actions

import (
	"fmt"
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/service/client"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
)

// ApplyResult defines the returns result of remotes
type ApplyResult struct {
	Remote    string
	RemoteApp string
	Verified  bool
	Data      []*pb.FILE
	Err       error
}

// ApplyPatchTo defines the grpc Ping method call
func (a *Action) ApplyPatchTo(remote string, apps []string) (out []*ApplyResult) {
	c := a.r.Get(remote)
	return getApplyResult(c)
}

// ApplyPatchToAll defines the grpc Ping method call to all remotes
func (a *Action) ApplyPatchToAll(apptype string) (out map[string][]*ApplyResult) {
	out = make(map[string][]*ApplyResult)
	for _, c := range a.r.GetAll() {
		out[c.Remote.Name] = getApplyResult(c)
	}
	return
}

func getApplyResult(c *client.Client) (out []*ApplyResult) {
	if c.Ok {
		var targets []string
		for _, app := range c.Remote.Apps {
			targets = append(targets, app.Source)
		}
		log.Println(c.Remote.Name, "APPLY: sending request for applying patch for apps", targets)
		res, err := c.ApplyPatch(targets)
		if err != nil {
			out = []*ApplyResult{{
				Remote:    c.Remote.Name,
				RemoteApp: "",
				Verified:  false,
				Data:      nil,
				Err:       fmt.Errorf("Remote is not found: %s", c.Remote.Name),
			}}
			return
		}
		out = make([]*ApplyResult, len(res))
		for i, r := range res {
			log.Println(c.Remote.Name, "APPLY: received response for app", r.GetRemoteApp(), r.GetVerified())
			out[i] = &ApplyResult{
				Remote:    c.Remote.Name,
				RemoteApp: r.GetRemoteApp(),
				Verified:  r.GetVerified(),
				Data:      r.GetData(),
				Err:       nil,
			}
		}
	} else {
		out = []*ApplyResult{{
			Remote:    c.Remote.Name,
			RemoteApp: "",
			Verified:  false,
			Data:      nil,
			Err:       fmt.Errorf("Remote is not found: %s", c.Remote.Name),
		}}
	}
	return
}
