package task

import (
	"github.com/pnkj-kmr/infra-patch-manager/service"
)

// // Task defines all task action/operation functions
// type Task interface {
// 	Ping(string) string
// 	SaveFile(string, string, bytes.Buffer) error
// 	GetFile(string) (module.I, error)
// }

// PatchTask declares all patch action related tasks
type PatchTask struct {
	r service.Remote
}

// NewPatchTask return new task struct
func NewPatchTask() (t *PatchTask, err error) {
	r, err := service.NewRemoteConfig()
	if err != nil {
		return &PatchTask{nil}, err
	}
	return &PatchTask{r}, nil
}
