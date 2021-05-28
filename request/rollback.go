package request

import (
	"log"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/patch/segment/dir"
	"github.com/pnkj-kmr/patch/segment/file"
	"github.com/pnkj-kmr/patch/segment/tar"
	"github.com/pnkj-kmr/patch/utility"
)

// CleanRollbackDir cleans the rollback folder
func CleanRollbackDir() (err error) {
	rollback, err := dir.New(LocationRollback)
	if err != nil {
		return err
	}
	return rollback.Clean()
}

// CompressRollbackDir tar a compress backup to folder
func CompressRollbackDir() (err error) {
	rollback, err := dir.New(LocationRollback)
	if err != nil {
		return err
	}
	upload, err := dir.New(LocationTarFiles)
	if err != nil {
		return err
	}
	t := tar.New(utility.RandomStringWithTime(6), "tar.gz", upload.Path())
	return t.Tar([]string{rollback.Path()})
}

// RollbackFrom helps to apply patch to dst
func RollbackFrom(dst string) (err error) {
	start := time.Now()
	err = CompressRollbackDir() // TODO - rollback backup
	err = CleanRollbackDir()    // cleaning the dir
	if err != nil {
		return err
	}
	rollbackPath, err := dir.New(LocationRollback)
	if err != nil {
		return err
	}
	fromDir, err := dir.New(dst)
	if err != nil {
		return err
	}
	patchPath, err := dir.New(LocationPatch)
	if err != nil {
		return
	}
	files, err := patchPath.Scan()
	if err != nil {
		return
	}
	for _, f := range files {
		file, e := file.New(filepath.Join(fromDir.Path(), f.RPath()), fromDir.Path())
		if e != nil {
			log.Println("Rollback: ", e)
			continue
		}
		if len(file.SPath()) > 0 {
			err = rollbackPath.Create(file.SPath())
			if err != nil {
				log.Println("Rollback: ", err)
				continue
			}
			dstInfo, e := dir.New(filepath.Join(rollbackPath.Path(), file.SPath()))
			if e != nil {
				log.Println("Rollback: ", err)
				continue
			}
			err = file.Copy(dstInfo)
			if err != nil {
				log.Println("Rollback: ", err)
				continue
			}
		} else {
			err = file.Copy(rollbackPath)
			if err != nil {
				log.Println("Rollback: ", err)
				continue
			}
		}
	}
	log.Println("Rollback: TIME   ", time.Since(start))
	return
}

// VerifyRollback helps to verify the applied rollback
func VerifyRollback(dst string) (dmap map[string]bool, err error) {
	start := time.Now()
	dmap = make(map[string]bool)
	src, err := dir.New(LocationRollback)
	if err != nil {
		return nil, err
	}
	files, err := src.Scan()
	if err != nil {
		return
	}
	var match bool
	for _, file := range files {
		dstInfo, e := dir.New(filepath.Join(dst, file.RPath()))
		if e != nil {
			break
		}
		ok, _ := file.IsSameFileAt(dstInfo, true)
		match = ok
		if !ok {
			break
		}
	}
	dmap[dst] = match

	log.Println("Value:", dmap)
	log.Println("Check: TIME   ", time.Since(start))
	return
}
