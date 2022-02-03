package remote

import "github.com/pnkj-kmr/infra-patch-manager/entity"

// Remote defines the remote config interface
// Actions method with related to Remote
type Remote interface {
	Name() string
	Status() bool
	Type() string
	Apps() ([]App, error)
	App(string) (App, error)
	AppByType(string) ([]App, error)
	AgentAddress() string
	UpdateStatus(bool)
}

// App defines the remote application
// Actions method with related to target App
type App interface {
	Name() string
	Status() bool
	SourcePath() string
	ServiceName() string
	Type() string
	GetFiles() []entity.Entity
	UpdateStatus(bool)
	UpdateFiles([]entity.Entity)
}
