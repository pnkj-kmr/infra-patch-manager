package task

import (
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/module/jsn"
)

// RightsCheckFor defines the grpc
func (t *PatchTask) RightsCheckFor(remote string, apps []string) jsn.Remote {
	c := t.r.Get(remote)
	log.Println(c.Remote.Name, "RIGHTS: sending request with data -", remote, apps)
	var targetApps []jsn.RemoteApp
	var notExistsApps []jsn.RemoteApp
	for _, app := range apps {
		rApp, err := isAppExistsInRemote(c.Remote, app)
		if err != nil {
			log.Println(c.Remote.Name, "RIGHTS: error -", err)
			notExistsApps = append(notExistsApps, rApp)
			continue
		}
		targetApps = append(targetApps, rApp)
	}
	res, err := c.RightsCheck(targetApps)
	if err != nil {
		log.Println(c.Remote.Name, "RIGHTS: error -", err)
		return jsn.Remote{
			Name:    c.Remote.Name,
			Address: c.Remote.Address,
			Status:  jsn.RemoteStatus{Ok: false, Err: err},
			Apps:    append(targetApps, notExistsApps...),
		}
	}
	log.Println(c.Remote.Name, "RIGHTS: success -", err)
	return jsn.Remote{
		Name:    c.Remote.Name,
		Address: c.Remote.Address,
		Status:  jsn.RemoteStatus{Ok: len(notExistsApps) == 0, Err: err},
		Apps:    append(res, notExistsApps...),
	}
}

// RightsCheckForAll defines the grpc
func (t *PatchTask) RightsCheckForAll() (out []jsn.Remote) {
	for _, c := range t.r.GetAll() {
		log.Println(c.Remote.Name, "RIGHTS: sending request with data -", c.Remote.Apps)
		res, err := c.RightsCheck(c.Remote.Apps)
		if err != nil {
			log.Println(c.Remote.Name, "RIGHTS: error -", err)
			out = append(out, jsn.Remote{
				Name:    c.Remote.Name,
				Address: c.Remote.Address,
				Apps:    c.Remote.Apps,
				Status:  jsn.RemoteStatus{Ok: false, Err: err},
			})
			continue
		}
		log.Println(c.Remote.Name, "RIGHTS: success -", err)
		out = append(out, jsn.Remote{
			Name:    c.Remote.Name,
			Address: c.Remote.Address,
			Apps:    res,
			Status:  jsn.RemoteStatus{Ok: true, Err: err},
		})
	}
	return
}
