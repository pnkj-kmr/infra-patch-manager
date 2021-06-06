package actions

import (
	"fmt"

	"github.com/pnkj-kmr/infra-patch-manager/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func isAppExistsInRemote(r service.Remote, apppath string) (out service.RemoteApp, err error) {
	for _, app := range r.Apps {
		if app.Source == apppath {
			return app, nil
		}
	}
	err = status.Error(codes.InvalidArgument, "App does not exist in remote")
	return service.RemoteApp{Source: apppath, Status: service.RemoteStatus{Ok: false, Err: err}}, err
}

func getAppsByItsType(r service.Remote, apptype string) (out []service.RemoteApp, err error) {
	for _, app := range r.Apps {
		if app.Type == apptype {
			out = append(out, app)
		}
	}
	if len(out) == 0 {
		err = status.Error(codes.InvalidArgument, fmt.Sprintf("No app exist in remotes with app type - %s", apptype))
	}
	return
}
