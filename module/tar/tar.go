package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// T declares tar declaration
type T struct {
	Name string
	Ext  string
	Path string
}

// New creates a tar object with name and extension
func New(name string, ext string, path string) *T {
	if path != "" {
		return &T{Name: name, Ext: ext, Path: path}
	}
	return &T{Name: name, Ext: ext, Path: "asset"}
}

// TarFilePath represent the full path of tar/tar.gz file
func (t *T) TarFilePath() string {
	return filepath.Join(t.Path, t.Name+t.Ext)
}

// Tar helps to compress the given folder
func (t *T) Tar(paths []string) (err error) {
	file, err := os.Create(t.TarFilePath())
	if err != nil {
		return
	}
	defer file.Close()

	var fw io.WriteCloser = file
	if strings.HasSuffix(t.TarFilePath(), ".gz") {
		fw = gzip.NewWriter(file)
		defer fw.Close()
	}

	tw := tar.NewWriter(fw)
	defer tw.Close()

	for _, i := range paths {
		if err := scanAndCopy(i, tw); err != nil {
			return err
		}
	}

	log.Println("TAR: ", t.TarFilePath(), err)
	return
}
