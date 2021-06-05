package actions

import (
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/service"
)

// RightsCheckFor defines the grpc
func (a *Action) RightsCheckFor(remote string, apps []string) service.Remote {
	c := a.r.Get(remote)
	log.Println(c.Remote.Name, "RIGHTS: sending request with data -", remote, apps)
	var targetApps []service.RemoteApp
	var notExistsApps []service.RemoteApp
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
		return service.Remote{
			Name:    c.Remote.Name,
			Address: c.Remote.Address,
			Status:  service.RemoteStatus{Ok: false, Err: err},
			Apps:    append(targetApps, notExistsApps...),
		}
	}
	log.Println(c.Remote.Name, "RIGHTS: success -", err)
	return service.Remote{
		Name:    c.Remote.Name,
		Address: c.Remote.Address,
		Status:  service.RemoteStatus{Ok: len(notExistsApps) == 0, Err: err},
		Apps:    append(res, notExistsApps...),
	}
}

// RightsCheckForAll defines the grpc
func (a *Action) RightsCheckForAll() (out []service.Remote) {
	for _, c := range a.r.GetAll() {
		log.Println(c.Remote.Name, "RIGHTS: sending request with data -", c.Remote.Apps)
		res, err := c.RightsCheck(c.Remote.Apps)
		if err != nil {
			log.Println(c.Remote.Name, "RIGHTS: error -", err)
			out = append(out, service.Remote{
				Name:    c.Remote.Name,
				Address: c.Remote.Address,
				Apps:    c.Remote.Apps,
				Status:  service.RemoteStatus{Ok: false, Err: err},
			})
			continue
		}
		log.Println(c.Remote.Name, "RIGHTS: success -", err)
		out = append(out, service.Remote{
			Name:    c.Remote.Name,
			Address: c.Remote.Address,
			Apps:    res,
			Status:  service.RemoteStatus{Ok: true, Err: err},
		})
	}
	return
}
