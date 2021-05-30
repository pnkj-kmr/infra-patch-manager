package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pnkj-kmr/patch/utility"
)

// Untar helps to extract the given tar/tar.gz file
func (t *T) Untar(extractPath string) (err error) {
	file, err := os.Open(t.TarFilePath())
	if err != nil {
		return
	}
	defer file.Close()

	var fr io.ReadCloser = file
	if strings.HasSuffix(t.TarFilePath(), ".gz") {
		if fr, err = gzip.NewReader(file); err != nil {
			return
		}
		defer fr.Close()
	}

	tarBallReader := tar.NewReader(fr)

	var skipBaseDir string
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if skipBaseDir == "" {
			if strings.Split(header.Name, string(os.PathSeparator))[0] == utility.RevokeDirectory {
				skipBaseDir = utility.RevokeDirectory
			}
		}

		rpath, err := filepath.Rel(skipBaseDir, header.Name)
		if err != nil {
			return err
		}

		filename := filepath.Join(extractPath, filepath.FromSlash(rpath))

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
	return nil
}
