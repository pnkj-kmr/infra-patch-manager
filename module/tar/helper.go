package tar

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func scanAndCopy(source string, tw *tar.Writer) (err error) {
	info, err := os.Stat(source)
	if err != nil {
		return
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.ToSlash(filepath.Join(baseDir, strings.TrimPrefix(path, source)))
			}

			if err := tw.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tw, file)
			return err
		})
}
