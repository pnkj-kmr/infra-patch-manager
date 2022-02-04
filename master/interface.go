package master

import (
	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/master/remote"
)

// PatchMaster declare all master actions
type PatchMaster interface {
	Ping() (bool, error)
	UploadFileToRemote(entity.File) (entity.Entity, bool, error)
	ExtractFileToRemote(string, string) ([]entity.Entity, bool, error)
	RightsCheckFor([]remote.App) ([]remote.App, error)
	PatchTo([]remote.App) ([]remote.App, error)
	VerifyFrom([]remote.App) ([]remote.App, error)
	ExecuteCmdOnRemote(string) ([]byte, error)
	ListAvailablePatches() ([]string, error)
}
