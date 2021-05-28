package dir

import (
	"io/fs"
	"time"
)

// Path presents the current directory/folder
// Returns string
func (d *D) Path() string {
	return d.P
}

// Name fs.FileInfo
func (d *D) Name() string {
	return d.I.Name()
}

// Size fs.FileInfo
func (d *D) Size() int64 {
	return d.I.Size()
}

// Mode fs.FileInfo
func (d *D) Mode() fs.FileMode {
	return d.I.Mode()
}

// ModTime fs.FileInfo
func (d *D) ModTime() time.Time {
	return d.I.ModTime()
}

// IsDir fs.FileInfo
func (d *D) IsDir() bool {
	return d.I.IsDir()
}

// Sys fs.FileInfo
func (d *D) Sys() interface{} {
	return d.I.Sys()
}
