package task

import (
	"github.com/pnkj-kmr/patch/module/jsn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func isAppExistsInRemote(r jsn.Remote, apppath string) (out jsn.RemoteApp, err error) {
	for _, app := range r.Apps {
		if app.Path == apppath {
			return app, nil
		}
	}
	err = status.Error(codes.InvalidArgument, "App does not exist in remote")
	return jsn.RemoteApp{Path: apppath, Status: jsn.RemoteStatus{Ok: false, Err: err}}, err
}
