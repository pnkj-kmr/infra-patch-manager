package task

import (
	"log"

	"github.com/pnkj-kmr/patch/module/jsn"
)

// RightsCheckFor defines the grpc
func (t *PatchTask) RightsCheckFor(remote, app string) (out jsn.Remote) {
	c := t.r.Get(remote)
	log.Println(c.Remote.Name, "RIGHTS: sending request with data -", remote, app)
	rApp, err := isAppExistsInRemote(c.Remote, app)
	if err != nil {
		log.Println(c.Remote.Name, "RIGHTS: error -", err)
		c.Remote.Apps = []jsn.RemoteApp{rApp}
		return c.Remote
	}
	res, err := c.RightsCheck([]jsn.RemoteApp{rApp})
	if err != nil {
		log.Println(c.Remote.Name, "RIGHTS: error -", err)
		c.Remote.Apps = []jsn.RemoteApp{rApp}
		return c.Remote
	}
	log.Println(c.Remote.Name, "RIGHTS: success -", err)
	c.Remote.Apps = res
	return c.Remote
}

// RightsCheckForApps defines the grpc
func (t *PatchTask) RightsCheckForApps(remote string, apps []string) (out jsn.Remote) {
	c := t.r.Get(remote)
	log.Println(c.Remote.Name, "RIGHTS: sending request with data -", remote, apps)
	var targetApps []jsn.RemoteApp
	for _, app := range apps {
		rApp, err := isAppExistsInRemote(c.Remote, app)
		if err != nil {
			log.Println(c.Remote.Name, "RIGHTS: error -", err)
			c.Remote.Apps = []jsn.RemoteApp{rApp}
			return c.Remote
		}
		targetApps = append(targetApps, rApp)
	}
	res, err := c.RightsCheck(targetApps)
	if err != nil {
		log.Println(c.Remote.Name, "RIGHTS: error -", err)
		c.Remote.Apps = targetApps
		return c.Remote
	}
	log.Println(c.Remote.Name, "RIGHTS: success -", err)
	c.Remote.Apps = res
	return c.Remote
}

// RightsCheckForAllApps defines the grpc
func (t *PatchTask) RightsCheckForAllApps(remote string) (out jsn.Remote) {
	c := t.r.Get(remote)
	log.Println(c.Remote.Name, "RIGHTS: sending request with data -", remote)
	res, err := c.RightsCheck(c.Remote.Apps)
	if err != nil {
		log.Println(c.Remote.Name, "RIGHTS: error -", err)
		return c.Remote
	}
	log.Println(c.Remote.Name, "RIGHTS: success -", err)
	c.Remote.Apps = res
	return c.Remote
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
				Status:  false,
			})
			continue
		}
		log.Println(c.Remote.Name, "RIGHTS: success -", err)
		out = append(out, jsn.Remote{
			Name:    c.Remote.Name,
			Address: c.Remote.Address,
			Apps:    res,
			Status:  true,
		})
	}
	return
}
