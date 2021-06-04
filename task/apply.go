package task

import (
	"fmt"
	"log"

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

func getApplyResult(remote, app string, verified bool, data []*pb.FILE, err error) *ApplyResult {
	return &ApplyResult{
		Remote:    remote,
		RemoteApp: app,
		Verified:  verified,
		Data:      data,
		Err:       err,
	}
}

// ApplyPatchTo defines the grpc Ping method call
func (t *PatchTask) ApplyPatchTo(remote string) (out []*ApplyResult) {
	c := t.r.Get(remote)
	if c.Ok {
		var targets []string
		for _, app := range c.Remote.Apps {
			targets = append(targets, app.Source)
		}
		log.Println(c.Remote.Name, "APPLY: sending request for applying patch for apps", targets)
		res, err := c.ApplyPatch(targets)
		if err != nil {
			out = []*ApplyResult{getApplyResult(remote, "", false, nil, fmt.Errorf("Remote is not found: %s", remote))}
			return
		}
		out = make([]*ApplyResult, len(res))
		for i, r := range res {
			log.Println(c.Remote.Name, "APPLY: received response for app", r.GetRemoteApp(), r.GetVerified())
			out[i] = getApplyResult(remote, r.GetRemoteApp(), r.GetVerified(), r.GetData(), err)
		}
	} else {
		out = []*ApplyResult{getApplyResult(remote, "", false, nil, fmt.Errorf("Remote is not found: %s", remote))}
	}
	return
}

// ApplyPatchToAll defines the grpc Ping method call to all remotes
func (t *PatchTask) ApplyPatchToAll() (out map[string][]*ApplyResult) {
	out = make(map[string][]*ApplyResult)
	var targets []string
	var remote string
	for _, c := range t.r.GetAll() {
		remote = c.Remote.Name
		if c.Ok {
			targets = []string{}
			for _, app := range c.Remote.Apps {
				targets = append(targets, app.Source)
			}
			log.Println(remote, "APPLY: sending request for applying patch for apps", targets)
			res, err := c.ApplyPatch(targets)
			if err != nil {
				out[remote] = []*ApplyResult{getApplyResult(remote, "", false, nil, fmt.Errorf("Remote is not found: %s", remote))}
				continue
			}
			result := make([]*ApplyResult, len(res))
			for i, r := range res {
				log.Println(c.Remote.Name, "APPLY: received response for app", r.GetRemoteApp(), r.GetVerified())
				result[i] = getApplyResult(remote, r.GetRemoteApp(), r.GetVerified(), r.GetData(), err)
			}
			out[remote] = result
		} else {
			out[remote] = []*ApplyResult{getApplyResult(remote, "", false, nil, fmt.Errorf("Remote is not found: %s", remote))}
		}
	}
	return
}
