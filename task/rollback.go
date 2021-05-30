package task

import (
	"log"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/module/file"
	"github.com/pnkj-kmr/patch/module/tar"
	"github.com/pnkj-kmr/patch/utility"
)

// CleanRevokeDir cleans the rollback folder
func CleanRevokeDir() (err error) {
	d, err := dir.New(utility.RevokeDirectory)
	if err != nil {
		return err
	}
	return d.Clean()
}

// BackupRevokeDir takes a tar backup from rollback folder
func BackupRevokeDir() (err error) {
	d, err := dir.New(utility.RevokeDirectory)
	if err != nil {
		return err
	}
	assetDir, err := dir.New(utility.AssetsDirectory)
	if err != nil {
		return err
	}
	t := tar.New("ROLLBACK"+utility.RandomStringWithTime(0), "tar.gz", assetDir.Path())
	return t.Tar([]string{d.Path()})
}

// RollbackFrom helps to take a rollback patch from target folder
func RollbackFrom(target string) (err error) {
	start := time.Now()
	err = BackupRevokeDir() // backup
	err = CleanRevokeDir()  // cleaning the dir
	if err != nil {
		return err
	}
	revokePath, err := dir.New(utility.RevokeDirectory)
	if err != nil {
		return err
	}
	fromDir, err := dir.New(target)
	if err != nil {
		return err
	}
	remedyDir, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return
	}
	files, err := remedyDir.Scan()
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
			err = revokePath.Create(file.SPath())
			if err != nil {
				log.Println("Rollback: ", err)
				continue
			}
			dstInfo, e := dir.New(filepath.Join(revokePath.Path(), file.SPath()))
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
			err = file.Copy(revokePath)
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
func VerifyRollback(target string) (dmap map[string]bool, err error) {
	start := time.Now()
	dmap = make(map[string]bool)
	src, err := dir.New(utility.RevokeDirectory)
	if err != nil {
		return nil, err
	}
	files, err := src.Scan()
	if err != nil {
		return
	}
	var match bool
	for _, file := range files {
		dstInfo, e := dir.New(filepath.Join(target, file.RPath()))
		if e != nil {
			break
		}
		ok, _ := file.IsSameFileAt(dstInfo, true)
		match = ok
		if !ok {
			break
		}
	}
	dmap[target] = match

	log.Println("Value:", dmap)
	log.Println("Check: TIME   ", time.Since(start))
	return
}
