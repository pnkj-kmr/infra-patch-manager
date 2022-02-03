package entity

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

// _file defines the file information
type _file struct {
	p string
	r string
	s string
	i fs.FileInfo
}

// NewFile creates the file struct
// Returns file and error
func NewFile(path, basepath string) (File, error) {
	rpath, err := filepath.Rel(basepath, path)
	if err != nil {
		return nil, err
	}
	spath, _ := filepath.Split(rpath)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &_file{path, rpath, spath, info}, nil
}

// SPath gives the sub directory of file
// Returns string
func (f *_file) SPath() string {
	return f.s
}

// RPath gives the relative path
// Returns string
func (f *_file) RPath() string {
	return f.r
}

// Path presents the current file
// Returns string
func (f *_file) Path() string {
	return f.p
}

// Name fs.FileInfo
func (f *_file) Name() string {
	return f.i.Name()
}

// Size fs.FileInfo
func (f *_file) Size() int64 {
	return f.i.Size()
}

// Mode fs.FileInfo
func (f *_file) Mode() fs.FileMode {
	return f.i.Mode()
}

// ModTime fs.FileInfo
func (f *_file) ModTime() time.Time {
	return f.i.ModTime()
}

// IsDir fs.FileInfo
func (f *_file) IsDir() bool {
	return f.i.IsDir()
}

// Sys fs.FileInfo
func (f *_file) Sys() interface{} {
	return f.i.Sys()
}

// IsSameFileAt helps to verify src and dst file - same or not
func (f *_file) IsSameFileAt(d Entity, old bool) (match bool, err error) {
	if old {
		match = (f.Name() == d.Name()) && (f.Size() == d.Size()) && (f.ModTime().After(d.ModTime()))
	} else {
		match = (f.Name() == d.Name()) && (f.Size() == d.Size()) && (f.ModTime().Before(d.ModTime()))
	}
	log.Println("FILE_COMPARE: matched: ", match, "flag_old:", old, "from-", f.Path(), "to-", d.Path())
	return
}

// Create help to create file
func (f *_file) Create(string) error {
	// TODO - create a file with data
	return nil
}

// Copy helps cp -rf
func (f *_file) Copy(d Entity) (err error) {
	if !f.Mode().IsRegular() {
		return fmt.Errorf("Copy: non-regular source file %s (%q)", f.Name(), f.Mode().String())
	}
	if !d.IsDir() {
		return fmt.Errorf("Copy: non-regular destination folder %s (%q)", d.Name(), d.Mode().String())
	}
	// source file path
	src := f.Path()
	// dest file path
	dst := filepath.Join(d.Path(), f.Name())
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	log.Println("COPY: src-", src, "| dst-", dst)
	return
}
