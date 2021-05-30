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

// CreateDirectoryIfNotExists helps to create a dir if not exists and returns the dir path
func CreateDirectoryIfNotExists(path string) (newpath string, err error) {
	// creating the directories if not exists
	d, err := New(path)
	newpath = d.Path()
	if err != nil {
		if os.IsNotExist(err) {
			d, err := New(filepath.Dir(path))
			if err != nil {
				log.Println("Cannot load directory information: ", filepath.Dir(path), err)
			}
			log.Println("Directory: ", d.Path(), path)
			err = d.Create(path)
			if err != nil {
				log.Println("Cannot create directory: ", d.Path(), path, err)
			}
			return filepath.Join(d.Path(), path), err
		}
	}
	return
}
