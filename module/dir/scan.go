package dir

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/pnkj-kmr/patch/module/file"
)

// Scan helps to scan the directory
func (d *D) Scan() (files []*file.F, err error) {
	err = filepath.Walk(d.Path(), closure(&files, d.Path()))
	if err != nil {
		log.Println("Cannot scan the path ", d.Path(), err)
		return nil, err
	}
	log.Println("SCAN: files -", len(files))
	return files, nil
}

func closure(files *[]*file.F, basepath string) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rpath, err := filepath.Rel(basepath, path)
		if err != nil {
			return err
		}
		spath, _ := filepath.Split(rpath)
		*files = append(*files, &file.F{P: path, R: rpath, S: spath, I: info})
		return nil
	}
}
