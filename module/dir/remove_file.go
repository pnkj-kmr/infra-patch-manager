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
	log.Println("Remove: FILE: ", fullPath)
	err = os.Remove(fullPath)
	if err != nil {
		return err
	}
	return
}
