package dir

import (
	"log"
	"os"
	"path/filepath"
)

// Remove removes the given folder from current directory
// Returns error if any
func (d *D) Remove(path string) (err error) {
	fullPath := filepath.Join(d.Path(), path)
	log.Println("Remove: DIR: ", fullPath)
	err = os.RemoveAll(fullPath)
	if err != nil {
		return err
	}
	return
}
