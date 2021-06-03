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
	err = os.RemoveAll(fullPath)
	if err != nil {
		log.Println("Cannot remove the dir", fullPath, err)
		return err
	}
	log.Println("REMOVE: dir -", fullPath)
	return
}
