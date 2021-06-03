package dir

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
)

// CreateFile creates a new file in current directory
// Returns error if any
func (d *D) CreateFile(file string, data []byte) (err error) {
	fullPath := filepath.Join(d.Path(), file)
	err = os.WriteFile(fullPath, data, d.Mode())
	log.Println("FILE_CREATE: file -", fullPath, err)
	return
}

// CreateAndWriteFile creates a new file in current directory
// Returns error if any
func (d *D) CreateAndWriteFile(file string, data bytes.Buffer) (size int64, err error) {
	fullPath := filepath.Join(d.Path(), file)
	f, err := os.Create(fullPath)
	if err != nil {
		log.Println("Cannot open the path", fullPath, err)
		return
	}
	defer f.Close()
	log.Println("FILE_CREATE_WRITE: file -", fullPath)
	return data.WriteTo(f)
}
