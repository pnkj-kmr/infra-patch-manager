package entity

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// _tar declares tar declaration
type _tar struct {
	Name string
	Ext  string
	Path string
}

// NewTar creates a tar object with name and extension
func NewTar(name string, ext string, path string) Tar {
	if path != "" {
		return &_tar{Name: name, Ext: ext, Path: path}
	}
	return &_tar{Name: name, Ext: ext, Path: ""}
}

// tarFilePath represent the full path of tar/tar.gz file
func (t *_tar) tarFilePath() string {
	return filepath.Join(t.Path, t.Name+t.Ext)
}

// Tar helps to compress the given folder
func (t *_tar) Tar(paths []string) (err error) {
	file, err := os.Create(t.tarFilePath())
	if err != nil {
		return
	}
	defer file.Close()

	var fw io.WriteCloser = file
	if strings.HasSuffix(t.tarFilePath(), ".gz") {
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

	log.Println("TAR: ", t.tarFilePath(), err)
	return
}

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

// Untar helps to extract the given tar/tar.gz file
func (t *_tar) Untar(extractPath string) (err error) {
	file, err := os.Open(t.tarFilePath())
	if err != nil {
		return
	}
	defer file.Close()

	var fr io.ReadCloser = file
	if strings.HasSuffix(t.tarFilePath(), ".gz") {
		if fr, err = gzip.NewReader(file); err != nil {
			return
		}
		defer fr.Close()
	}

	tarBallReader := tar.NewReader(fr)
	// var skipBaseDir string
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// This code specific to source dir - as
		// if skipBaseDir == "" {
		// 	if strings.Split(header.Name, string(os.PathSeparator))[0] == "source" {
		// 		skipBaseDir = "source"
		// 	}
		// }
		// rpath, err := filepath.Rel(skipBaseDir, header.Name)
		// if err != nil {
		// 	return err
		// }
		// filename := filepath.Join(extractPath, filepath.FromSlash(rpath))

		filename := filepath.Join(extractPath, filepath.FromSlash(header.Name))

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer
			if err != nil {
				return err
			}

		case tar.TypeReg:
			writer, err := os.Create(filename)
			if err != nil {
				return err
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			writer.Close()
		default:
			log.Println("ERROR : ", header.Typeflag, filename)
		}
	}
	log.Println("UNTAR: ", t.tarFilePath(), extractPath)
	return nil
}
