package entity

import (
	"bytes"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

// _dir defines the Directory Information
type _dir struct {
	p string
	i fs.FileInfo
}

// NewDir creates the directory struct
// Returns directory and error
func NewDir(path string) (Dir, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &_dir{path, info}, nil
}

// Path presents the current directory/folder
// Returns string
func (d *_dir) Path() string {
	return d.p
}

// Name fs.FileInfo
func (d *_dir) Name() string {
	return d.i.Name()
}

// Size fs.FileInfo
func (d *_dir) Size() int64 {
	return d.i.Size()
}

// Mode fs.FileInfo
func (d *_dir) Mode() fs.FileMode {
	return d.i.Mode()
}

// ModTime fs.FileInfo
func (d *_dir) ModTime() time.Time {
	return d.i.ModTime()
}

// IsDir fs.FileInfo
func (d *_dir) IsDir() bool {
	return d.i.IsDir()
}

// Sys fs.FileInfo
func (d *_dir) Sys() interface{} {
	return d.i.Sys()
}

// Stat helps to get the file/folder details if exists
// Return FileInfo and err
func (d *_dir) Stat(file string) (fs.FileInfo, error) {
	return os.Stat(filepath.Join(d.Path(), file))
}

// Clean removes evenything from current directory
// Returns error if any
func (d *_dir) Clean() (err error) {
	dstInfo, err := os.Open(d.Path())
	if err != nil {
		return err
	}
	defer dstInfo.Close()
	names, err := dstInfo.Readdirnames(-1)
	if err != nil {
		return err
	}
	var path string
	for _, name := range names {
		path = filepath.Join(d.Path(), name)
		log.Println("CLEAN: content -", path)
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return
}

// Create creates a new folder in current directory
// Returns error if any
func (d *_dir) Create(path string) (err error) {
	newPath := filepath.Join(d.Path(), path)
	_, err = NewDir(newPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(newPath, d.Mode())
			log.Println("CREATE: dir -", newPath)
		}
	}
	return
}

// CreateFile creates a new file in current directory
// Returns error if any
func (d *_dir) CreateFile(file string, data []byte) (err error) {
	fullPath := filepath.Join(d.Path(), file)
	err = os.WriteFile(fullPath, data, d.Mode())
	log.Println("FILE_CREATE: file -", fullPath, err)
	return
}

// CreateAndWriteFile creates a new file in current directory
// Returns error if any
func (d *_dir) CreateAndWriteFile(file string, data bytes.Buffer) (size int64, err error) {
	fullPath := filepath.Join(d.Path(), file)
	f, err := os.Create(fullPath)
	if err != nil {
		log.Println("Cannot open the path", fullPath, err)
		return
	}
	defer f.Close()
	log.Println("FILE_CREATE_WRITE: file -", fullPath)
	return data.WriteTo(f)
}

// Remove removes the given folder from current directory
// Returns error if any
func (d *_dir) Remove(path string) (err error) {
	fullPath := filepath.Join(d.Path(), path)
	err = os.RemoveAll(fullPath)
	if err != nil {
		log.Println("Cannot remove the dir", fullPath, err)
		return err
	}
	log.Println("REMOVE: dir -", fullPath)
	return
}

// RemoveFile removes the given file from current directory
// Returns error if any
func (d *_dir) RemoveFile(file string) (err error) {
	fullPath := filepath.Join(d.Path(), file)
	err = os.Remove(fullPath)
	if err != nil {
		log.Println("Cannot remove the file", fullPath, err)
		return err
	}
	log.Println("REMOVE: file -", fullPath)
	return
}

// Scan helps to scan the directory
func (d *_dir) Scan() (files []File, err error) {
	err = filepath.Walk(d.Path(), func(fileList *[]File, basepath string) filepath.WalkFunc {
		return func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			f, err := NewFile(path, basepath)
			if err != nil {
				return err
			}
			*fileList = append(*fileList, f)
			return nil
		}
	}(&files, d.Path()))
	if err != nil {
		log.Println("Cannot scan the path ", d.Path(), err)
		return nil, err
	}
	log.Println("SCAN: files -", len(files))
	return files, nil
}

// Copy helps copy everyting from current dir to given path
// Returns err if any
func (d *_dir) Copy(to string) (err error) {
	dst, err := NewDir(to)
	if err != nil {
		return
	}
	results, err := d.Scan()
	if err != nil {
		return
	}
	for _, f := range results {
		if len(f.SPath()) > 0 {
			err = dst.Create(f.SPath())
			if err != nil {
				return
			}
			dstInfo, e := NewDir(filepath.Join(dst.Path(), f.SPath()))
			if e != nil {
				return e
			}
			err = f.Copy(dstInfo)
			if err != nil {
				return

			}
		} else {
			err = f.Copy(dst)
			if err != nil {
				return
			}
		}
	}
	return
}
