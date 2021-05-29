package module

import (
	"io/fs"
)

// I helps to share os.Stat into with extra info
type I interface {
	Create(string) error
	Path() string
	fs.FileInfo
}
