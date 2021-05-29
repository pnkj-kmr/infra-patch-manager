package dir

import (
	"io/fs"
	"os"
)

// D defines the Directory Information
type D struct {
	P string
	I fs.FileInfo
}

// New creates the directory struct
// Returns directory and error
func New(path string) (*D, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &D{path, info}, nil
}
