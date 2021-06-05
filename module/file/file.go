package file

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// F defines the file information
type F struct {
	p string
	r string
	s string
	i fs.FileInfo
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

// SPath gives the sub directory of file
// Returns string
func (f *F) SPath() string {
	return f.s
}

// RPath gives the relative path
// Returns string
func (f *F) RPath() string {
	return f.r
}

// Path presents the current file
// Returns string
func (f *F) Path() string {
	return f.p
}

// Name fs.FileInfo
func (f *F) Name() string {
	return f.i.Name()
}

// Size fs.FileInfo
func (f *F) Size() int64 {
	return f.i.Size()
}

// Mode fs.FileInfo
func (f *F) Mode() fs.FileMode {
	return f.i.Mode()
}

// ModTime fs.FileInfo
func (f *F) ModTime() time.Time {
	return f.i.ModTime()
}

// IsDir fs.FileInfo
func (f *F) IsDir() bool {
	return f.i.IsDir()
}

// Sys fs.FileInfo
func (f *F) Sys() interface{} {
	return f.i.Sys()
}
