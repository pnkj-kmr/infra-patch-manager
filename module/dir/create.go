package dir

import (
	"log"
	"os"
	"path/filepath"
)

// Create creates a new folder in current directory
// Returns error if any
func (d *D) Create(path string) (err error) {
	newPath := filepath.Join(d.Path(), path)
	_, err = New(newPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(newPath, d.Mode())
			log.Println("Create: DIR: ", newPath)
		}
	}
	return
}
