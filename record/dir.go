package record

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Dir returns a directory interface
type Dir interface {
	Create(string) error
	Remove() error
	Clean() error
	Scan() (files []F, err error)
	Copy(string) error
}

// D defines the Directory
type D struct {
	Path string
	Info fs.FileInfo
}

// NewDir returns the Dir interface object
func NewDir(path string) (*D, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &D{path, info}, nil
}

// Create helps mkdir
func (d *D) Create(path string) (err error) {
	err = os.MkdirAll(filepath.Join(d.Path, path), d.Info.Mode())
	return
}

// Remove helps rmdir
func (d *D) Remove() (err error) {
	err = os.RemoveAll(d.Path)
	if err != nil {
		return err
	}
	return
}

// Clean helps rm -rf *
func (d *D) Clean() (err error) {
	dir, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(d.Path, name))
		if err != nil {
			return err
		}
	}
	return
}

func closure(files *[]*F, basepath string) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rpath, err := filepath.Rel(basepath, path)
		if err != nil {
			return err
		}
		subpath, _ := filepath.Split(rpath)
		*files = append(*files, &F{path, rpath, subpath, info})
		return nil
	}
}

// Scan helps to scan the directory
func (d *D) Scan() (files []*F, err error) {
	err = filepath.Walk(d.Path, closure(&files, d.Path))
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Copy helps cp -rf
func (d *D) Copy(to string) (err error) {
	dst, err := NewDir(to)
	if err != nil {
		return
	}
	files, err := d.Scan()
	if err != nil {
		return
	}
	for _, file := range files {
		err = file.Copy(dst)
		if err != nil {
			return
		}
	}
	return
}
