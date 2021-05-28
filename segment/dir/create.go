package dir

import (
	"os"
	"path/filepath"
)

// Create creates a new folder in current directory
// Returns error if any
func (d *D) Create(path string) (err error) {
	err = os.MkdirAll(filepath.Join(d.Path(), path), d.Mode())
	return
}
