package dir

import (
	"path/filepath"
)

// Copy helps copy everyting from current dir to given path
// Returns err if any
func (d *D) Copy(to string) (err error) {
	dst, err := New(to)
	if err != nil {
		return
	}
	files, err := d.Scan()
	if err != nil {
		return
	}
	for _, file := range files {
		if len(file.SPath()) > 0 {
			err = dst.Create(file.SPath())
			if err != nil {
				return
			}
			dstInfo, e := New(filepath.Join(dst.Path(), file.SPath()))
			if e != nil {
				return e
			}
			err = file.Copy(dstInfo)
			if err != nil {
				return

			}
		} else {
			err = file.Copy(dst)
			if err != nil {
				return
			}
		}
	}
	return
}
