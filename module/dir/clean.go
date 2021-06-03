package dir

import (
	"log"
	"os"
	"path/filepath"
)

// Clean removes evenything from current directory
// Returns error if any
func (d *D) Clean() (err error) {
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
