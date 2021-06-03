package dir

import (
	"log"
	"os"
	"path/filepath"
)

// RemoveFile removes the given file from current directory
// Returns error if any
func (d *D) RemoveFile(file string) (err error) {
	fullPath := filepath.Join(d.Path(), file)
	err = os.Remove(fullPath)
	if err != nil {
		log.Println("Cannot remove the file", fullPath, err)
		return err
	}
	log.Println("REMOVE: file -", fullPath)
	return
}
