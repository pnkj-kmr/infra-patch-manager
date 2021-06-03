package action

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
		log.Println("Unable to load rollback folder", err)
		return err
	}
	return d.Clean()
}

// BackupRevokeDir takes a tar backup from rollback folder
func BackupRevokeDir() (err error) {
	d, err := dir.New(utility.RevokeDirectory)
	if err != nil {
		log.Println("Unable to load rollback folder", utility.RevokeDirectory, err)
		return err
	}
	assetDir, err := dir.New(utility.AssetsDirectory)
	if err != nil {
		log.Println("Unable to load assets folder", utility.AssetsDirectory, err)
		return err
	}
	t := tar.New(utility.RandomStringWithTime(0, "ROLLBACK"), "tar.gz", assetDir.Path())
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
		log.Println("Unable to load rollback folder", utility.RevokeDirectory, err)
		return err
	}
	fromDir, err := dir.New(target)
	if err != nil {
		log.Println("Unable to load given target folder", target, err)
		return err
	}
	remedyDir, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		log.Println("Unable to load default patch folder", utility.RemedyDirectory, err)
		return
	}
	files, err := remedyDir.Scan()
	if err != nil {
		log.Println("Unable to scan", utility.RemedyDirectory, err)
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
	log.Println("ROLLBACK FROM", target, "T:", time.Since(start))
	return
}

// VerifyRollback helps to verify the applied rollback
func VerifyRollback(target string) (match bool, err error) {
	start := time.Now()
	src, err := dir.New(utility.RevokeDirectory)
	if err != nil {
		log.Println("Unable to load rollback folder", utility.RevokeDirectory, err)
		return
	}
	files, err := src.Scan()
	if err != nil {
		log.Println("Unable to scan", utility.RevokeDirectory, err)
		return
	}
	for _, file := range files {
		dstInfo, e := dir.New(filepath.Join(target, file.RPath()))
		if e != nil {
			match = false
			break
		}
		ok, _ := file.IsSameFileAt(dstInfo, true)
		match = ok
		if !ok {
			break
		}
	}
	log.Println("ROLLBACK VERIFIED FOR", target, "OK:", match, "T:", time.Since(start))
	return
}
