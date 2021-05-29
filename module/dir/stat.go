package dir

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Stat helps to get the file/folder details if exists
// Return FileInfo and err
func (d *D) Stat(file string) (fs.FileInfo, error) {
	return os.Stat(filepath.Join(d.Path(), file))
}
