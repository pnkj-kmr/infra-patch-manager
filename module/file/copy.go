package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pnkj-kmr/patch/module"
)

// Copy helps mkdir
func (f *F) Copy(dst module.I) (err error) {
	if !f.Mode().IsRegular() {
		return fmt.Errorf("Copy: non-regular source file %s (%q)", f.Name(), f.Mode().String())
	}
	if !(dst.IsDir()) {
		return fmt.Errorf("Copy: non-regular destination folder %s (%q)", dst.Name(), dst.Mode().String())
	}
	err = copy(f.Path(), filepath.Join(dst.Path(), f.Name()))
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
	log.Println("Copy: SRC:", src, " | DST:", dst)
	return
}
