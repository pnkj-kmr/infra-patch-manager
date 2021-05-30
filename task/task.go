package task

import "strings"

// Task defines all task action/operation functions
type Task interface {
	Ping(string) string
}

// PatchTask declares all patch action related tasks
type PatchTask struct{}

// NewPatchTask return new task struct
func NewPatchTask() *PatchTask {
	return &PatchTask{}
}

// Ping is poking funcation input "PING" returns "PONG"
func (t *PatchTask) Ping(in string) (out string) {
	if ok := strings.EqualFold(in, "PING"); ok {
		out = "PONG"
	}
	return
}
