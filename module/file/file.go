package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

// F defines the file information
type F struct {
	P string
	R string
	S string
	I fs.FileInfo
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
