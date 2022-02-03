package entity

import (
	"bytes"
	"io/fs"
)

// Entity helps to share os.Stat into with extra info
type Entity interface {
	Create(string) error
	Path() string
	fs.FileInfo
}

// Dir declares the directory
type Dir interface {
	Entity
	Clean() error
	CreateFile(string, []byte) error
	CreateAndWriteFile(string, bytes.Buffer) (int64, error)
	Remove(string) error
	RemoveFile(string) error
	Scan() ([]File, error)
	Copy(to string) error
}

// File declares the directory
type File interface {
	Entity
	SPath() string
	RPath() string
	IsSameFileAt(Entity, bool) (bool, error)
	Copy(Entity) error
}

// Tar helps to deal with tar type files
type Tar interface {
	Tar([]string) error
	Untar(string) error
}

// Conf help to deal with configuration default
type Conf interface {
	AssetPath() string
	PatchPath() string
	RollbackPath() string
}
