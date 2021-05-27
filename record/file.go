package record

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// File returns a directory interface
type File interface {
	Copy(string) error
}

// F defines the file
type F struct {
	Path    string
	RPath   string
	SubPath string
	Info    fs.FileInfo
}

// NewFile returns the file interface obj
func NewFile(basepath, path string) (*F, error) {
	rPath, err := filepath.Rel(basepath, path)
	if err != nil {
		return nil, err
	}
	subPath, _ := filepath.Split(rPath)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &F{path, rPath, subPath, info}, nil
}

// Copy helps mkdir
func (f *F) Copy(dst *D) (err error) {
	srcInfo := f.Info
	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("Copy: non-regular source file %s (%q)", srcInfo.Name(), srcInfo.Mode().String())
	}
	if !(dst.Info.IsDir()) {
		return fmt.Errorf("Copy: non-regular destination folder %s (%q)", dst.Info.Name(), dst.Info.Mode().String())
	}
	dstInfo, err := NewDir(filepath.Join(dst.Path, f.SubPath))
	if err != nil {
		if os.IsNotExist(err) {
			if err = dst.Create(f.SubPath); err != nil {
				return
			}
			dstInfo, err = NewDir(filepath.Join(dst.Path, f.SubPath))
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	// // checking same files are same or not
	// if os.SameFile(srcInfo, dstInfo) {
	// 	fmt.Println("Files found to same")
	// 	return
	// }
	err = copy(f.Path, filepath.Join(dstInfo.Path, srcInfo.Name()))
	return
}

func copy(src, dst string) (err error) {
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
	return
}
