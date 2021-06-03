package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

// F defines the file information
type F struct {
	P string      `json:"path"`
	R string      `json:"relative_path"`
	S string      `json:"sub_directory"`
	I fs.FileInfo `json:"info"`
}

// New creates the file struct
// Returns file and error
func New(path, basepath string) (*F, error) {
	rpath, err := filepath.Rel(basepath, path)
	if err != nil {
		return nil, err
	}
	spath, _ := filepath.Split(rpath)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &F{path, rpath, spath, info}, nil
}
