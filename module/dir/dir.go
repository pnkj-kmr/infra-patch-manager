package dir

import (
	"io/fs"
	"os"
	"time"
)

// D defines the Directory Information
type D struct {
	p string
	i fs.FileInfo
}

// New creates the directory struct
// Returns directory and error
func New(path string) (*D, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &D{path, info}, nil
}

// Path presents the current directory/folder
// Returns string
func (d *D) Path() string {
	return d.p
}

// Name fs.FileInfo
func (d *D) Name() string {
	return d.i.Name()
}

// Size fs.FileInfo
func (d *D) Size() int64 {
	return d.i.Size()
}

// Mode fs.FileInfo
func (d *D) Mode() fs.FileMode {
	return d.i.Mode()
}

// ModTime fs.FileInfo
func (d *D) ModTime() time.Time {
	return d.i.ModTime()
}

// IsDir fs.FileInfo
func (d *D) IsDir() bool {
	return d.i.IsDir()
}

// Sys fs.FileInfo
func (d *D) Sys() interface{} {
	return d.i.Sys()
}
