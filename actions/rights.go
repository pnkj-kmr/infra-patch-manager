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
			Status:  service.Status{Ok: false, Err: err.Error()},
			Apps:    append(targetApps, notExistsApps...),
		}
	}
	log.Println(c.Remote.Name, "RIGHTS: success -", err)
	return service.Remote{
		Name:    c.Remote.Name,
		Address: c.Remote.Address,
		Status:  service.Status{Ok: len(notExistsApps) == 0, Err: err.Error()},
		Apps:    append(res, notExistsApps...),
	}
}

// RightsCheckForAll defines the grpc
func (a *Action) RightsCheckForAll(apptype string) (out []service.Remote) {
	for _, c := range a.r.GetAll() {
		log.Println(c.Remote.Name, "RIGHTS: sending request with data -", apptype, c.Remote.Apps)
		var targetApps []service.RemoteApp
		var err error
		if apptype != "" {
			targetApps, err = getAppsByItsType(c.Remote, apptype)
		} else {
			targetApps = c.Remote.Apps
		}
		if err != nil {
			log.Println(c.Remote.Name, "RIGHTS: error -", err)
			out = append(out, service.Remote{
				Name:    c.Remote.Name,
				Address: c.Remote.Address,
				Apps:    c.Remote.Apps,
				Status:  service.Status{Ok: false, Err: err.Error()},
			})
			continue
		}
		// calling into rights check with target apps
		res, err := c.RightsCheck(targetApps)
		if err != nil {
			log.Println(c.Remote.Name, "RIGHTS: error -", err)
			out = append(out, service.Remote{
				Name:    c.Remote.Name,
				Address: c.Remote.Address,
				Apps:    c.Remote.Apps,
				Status:  service.Status{Ok: false, Err: err.Error()},
			})
			continue
		}
		log.Println(c.Remote.Name, "RIGHTS: success -", err)
		out = append(out, service.Remote{
			Name:    c.Remote.Name,
			Address: c.Remote.Address,
			Apps:    res,
			Status:  service.Status{Ok: true},
		})
	}
	return
}
