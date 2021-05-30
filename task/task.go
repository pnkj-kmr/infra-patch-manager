package task

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/pnkj-kmr/patch/module"
	"github.com/pnkj-kmr/patch/module/file"
	"github.com/pnkj-kmr/patch/utility"
)

// Task defines all task action/operation functions
type Task interface {
	Ping(string) string
	SaveFile(string, string, bytes.Buffer) error
	GetFile(string) (module.I, error)
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

// SaveFile helps to save the tar into target
func (t *PatchTask) SaveFile(file string, ext string, data bytes.Buffer) (err error) {
	// TODO - NEED TO SAVE THE FILE
	return
}

// GetFile helps to get the tar info
func (t *PatchTask) GetFile(path string) (f module.I, err error) {
	f, err = file.New(filepath.Join(utility.AssetsDirectory, path), utility.AssetsDirectory)
	return
}
