package actions

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/client"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
)

// ApplyResult defines the returns result of remotes
type ApplyResult struct {
	Remote    string         `json:"remote"`
	RemoteApp string         `json:"remote_app"`
	Verified  bool           `json:"is_verified"`
	Data      []*pb.FILE     `json:"data"`
	Status    service.Status `json:"status"`
}

// ApplyPatchTo defines the grpc Ping method call
func (a *Action) ApplyPatchTo(remote string, apptype string) (out []*ApplyResult) {
	c := a.r.Get(remote)
	return getApplyResult(c, apptype)
}

// ApplyPatchToAll defines the grpc Ping method call to all remotes
func (a *Action) ApplyPatchToAll(apptype string) (out []*ApplyResult) {
	start := time.Now()
	// for _, c := range a.r.GetAll() {
	// 	out = append(out, getApplyResult(c, apptype)...)
	// }
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	data := a.r.GetAll()
	wg.Add(len(data))
	for _, cc := range data {
		// concurrency with mutli host environment
		go func(r *[]*ApplyResult, c *client.Client, apptype string) {
			defer wg.Done()
			res := getApplyResult(c, apptype)
			mutex.Lock()
			*r = append(*r, res...)
			mutex.Unlock()
		}(&out, cc, apptype)
	}
	wg.Wait()
	log.Println("PATCH: receieved response with data -", len(out), "T:", time.Since(start))
	return
}

func getApplyResult(c *client.Client, apptype string) (out []*ApplyResult) {
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
			err := fmt.Errorf("Remote is not found: %s", c.Remote.Name)
			out = []*ApplyResult{{
				Remote:    c.Remote.Name,
				RemoteApp: "",
				Verified:  false,
				Data:      nil,
				Status:    service.Status{Ok: false, Err: err.Error()},
			}}
			return
		}
		log.Println(c.Remote.Name, "APPLY: sending request for applying patch for apps", targets)
		res, err := c.ApplyPatch(targets)
		if err != nil {
			out = []*ApplyResult{{
				Remote:    c.Remote.Name,
				RemoteApp: "",
				Verified:  false,
				Data:      nil,
				Status:    service.Status{Ok: false, Err: err.Error()},
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
				Status:    service.Status{Ok: true},
			}
		}
	} else {
		err := fmt.Errorf("Remote is not connected: %s", c.Remote.Name)
		out = []*ApplyResult{{
			Remote:    c.Remote.Name,
			RemoteApp: "",
			Verified:  false,
			Data:      nil,
			Status:    service.Status{Ok: false, Err: err.Error()},
		}}
	}
	return
}
