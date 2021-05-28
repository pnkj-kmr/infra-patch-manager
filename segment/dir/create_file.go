package dir

import (
	"os"
	"path/filepath"
)

// CreateFile creates a new file in current directory
// Returns error if any
func (d *D) CreateFile(file string) (err error) {
	_, err = os.Create(filepath.Join(d.Path(), file))
	return
}
