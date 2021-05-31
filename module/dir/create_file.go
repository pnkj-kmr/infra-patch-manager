package dir

import (
	"bytes"
	"os"
	"path/filepath"
)

// CreateFile creates a new file in current directory
// Returns error if any
func (d *D) CreateFile(file string, data []byte) (err error) {
	err = os.WriteFile(filepath.Join(d.Path(), file), data, d.Mode())
	return
}

// CreateAndWriteFile creates a new file in current directory
// Returns error if any
func (d *D) CreateAndWriteFile(file string, data bytes.Buffer) (size int64, err error) {
	f, err := os.Create(filepath.Join(d.Path(), file))
	if err != nil {
		return
	}
	defer f.Close()
	return data.WriteTo(f)
}
