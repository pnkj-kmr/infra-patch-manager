package actions

import (
	"github.com/pnkj-kmr/infra-patch-manager/service/client"
)

// // Task defines all task action/operation functions
// type Task interface {
// 	Ping(string) string
// 	SaveFile(string, string, bytes.Buffer) error
// 	GetFile(string) (module.I, error)
// }

// Action declares all patch action related tasks
type Action struct {
	r *client.RemoteConfig
}

// NewAction return new task struct
func NewAction() (t *Action, err error) {
	r, err := client.NewRemoteConfig()
	if err != nil {
		return &Action{nil}, err
	}
	return &Action{r}, nil
}
