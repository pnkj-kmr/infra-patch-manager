package file

import (
	"io/fs"
	"time"
)

// SPath gives the sub directory of file
// Returns string
func (f *F) SPath() string {
	return f.S
}

// RPath gives the relative path
// Returns string
func (f *F) RPath() string {
	return f.R
}

// Path presents the current file
// Returns string
func (f *F) Path() string {
	return f.P
}

// Name fs.FileInfo
func (f *F) Name() string {
	return f.I.Name()
}

// Size fs.FileInfo
func (f *F) Size() int64 {
	return f.I.Size()
}

// Mode fs.FileInfo
func (f *F) Mode() fs.FileMode {
	return f.I.Mode()
}

// ModTime fs.FileInfo
func (f *F) ModTime() time.Time {
	return f.I.ModTime()
}

// IsDir fs.FileInfo
func (f *F) IsDir() bool {
	return f.I.IsDir()
}

// Sys fs.FileInfo
func (f *F) Sys() interface{} {
	return f.I.Sys()
}
