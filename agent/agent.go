package agent

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/rpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type _agent struct {
	assets   entity.Dir
	patch    entity.Dir
	target   entity.Dir
	isBackup bool
	rollback entity.Dir
}

// NewPatchAgent - action pointer
func NewPatchAgent(path string, isBackup bool) (PatchAgent, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	assets, err := entity.NewDir(entity.C.AssetPath())
	if err != nil {
		log.Println("Assets folder does not exist or any", err)
		return nil, err
	}
	patch, err := entity.NewDir(entity.C.PatchPath())
	if err != nil {
		log.Println("Patch folder does not exist or any", err)
		return nil, err
	}
	var rollback entity.Dir
	if isBackup {
		rollback, err = entity.NewDir(entity.C.RollbackPath())
		if err != nil {
			log.Println("Rollback folder does not exist or any", err)
			return nil, err
		}
	}
	target, err := entity.NewDir(path)
	if err != nil {
		log.Println("Target folder does not exist or any", err)
		return nil, err
	}
	return &_agent{assets, patch, target, isBackup, rollback}, nil
}

func (a *_agent) RightsCheck() (ok bool, err error) {
	start := time.Now()
	randomStr := entity.RandomString(10)
	err = a.target.Create(randomStr)
	// fmt.Println("Folder Create..", err)
	if err == nil {
		err = a.target.Remove(randomStr)
		// fmt.Println("Folder Delete..", err)
		if err == nil {
			err = a.target.CreateFile(randomStr, []byte("temp data into file"))
			// fmt.Println("File Create..", err)
			if err == nil {
				err = a.target.RemoveFile(randomStr)
				// fmt.Println("File Delete..", err)
				if err == nil {
					ok = true
				}
			}
		}
	}
	log.Println("R/W CHECKS: ", a.target.Path(), "OK:", ok, "T:", time.Since(start))
	return
}

func (a *_agent) WriteUploaded(in entity.Entity, data bytes.Buffer) (out entity.Entity, err error) {
	// Writeing the file into directory
	_, err = a.assets.CreateAndWriteFile(in.Name(), data)
	if err != nil {
		return nil, rpc.LogError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}
	out, err = entity.NewFile(filepath.Join(a.assets.Path(), in.Name()), a.assets.Path())
	return
}

func (a *_agent) PatchNow() (err error) {
	start := time.Now()
	// R/W check for target folder
	rw, err := a.RightsCheck()
	if !rw || err != nil {
		return rpc.LogError(status.Errorf(codes.Internal, "target remote is not having read/write permission: %v | %v", a.target.Path(), err))
	}
	if a.isBackup {
		// taking a rollbackup backup
		err = a.PatchRollback()
		if err != nil {
			return rpc.LogError(status.Errorf(codes.Internal, "rollback backup failed: %v | %v", a.target.Path(), err))
		}
		// verifying the rollback patch
		_, err := a.VerifyRollbacked()
		if err != nil {
			return rpc.LogError(status.Errorf(codes.Internal, "rollback files are improper: %v | %v", a.target.Path(), err))
		}
	}

	// applying patch to given folder
	err = a.patch.Copy(a.target.Path())
	if err != nil {
		return rpc.LogError(status.Errorf(codes.Internal, "Apply patch failed: %v | %v", a.target.Path(), err))
	}

	log.Println("PATCHED T:", time.Since(start), "target: ", a.target.Path())
	return
}

func (a *_agent) VerifyPatched() (files []entity.File, ok bool, err error) {
	start := time.Now()

	patchFiles, err := a.patch.Scan()
	if err != nil {
		log.Println("Unable to scan patch folder", err)
		return
	}
	for _, file := range patchFiles {
		dstInfo, e := entity.NewFile(filepath.Join(a.target.Path(), file.RPath()), a.target.Path())
		if e != nil {
			log.Println("Cannot find the file path", filepath.Join(a.target.Path(), file.RPath()), e)
			continue
		}
		match, _ := file.IsSameFileAt(dstInfo, false)
		ok = match
		if !match {
			break
		}
		if dstInfo.IsDir() {
			files = append(files, file)
		} else {
			files = append(files, dstInfo)
		}
	}
	if ok {
		// Crosssing verifying the file length applied and patched
		ok = len(files) == len(patchFiles)
	}
	log.Println("PATCH VERIFED: ", a.target.Path(), "OK:", ok, "T:", time.Since(start))
	return
}

func (a *_agent) PatchRollback() (err error) {
	start := time.Now()
	// backing up the existing rollback if any
	err = backupExistingRollback()
	if err != nil {
		return err
	}

	// cleaning the rollback directory
	a.rollback.Clean()

	// scanning the patchable files
	files, err := a.patch.Scan()
	if err != nil {
		log.Println("Unable to scan - ", a.patch.Path(), err)
		return
	}

	for _, f := range files {
		file, e := entity.NewFile(filepath.Join(a.target.Path(), f.RPath()), a.target.Path())
		if e != nil {
			log.Println("Rollback: ", e)
			continue
		}
		if len(file.SPath()) > 0 {
			err = a.rollback.Create(file.SPath())
			if err != nil {
				log.Println("Rollback: ", err)
				continue
			}
			dstInfo, e := entity.NewDir(filepath.Join(a.rollback.Path(), file.SPath()))
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
			err = file.Copy(a.rollback)
			if err != nil {
				log.Println("Rollback: ", err)
				continue
			}
		}
	}
	log.Println("ROLLBACK FROM: ", a.target.Path(), "T:", time.Since(start))
	return
}

func (a *_agent) VerifyRollbacked() (ok bool, err error) {
	start := time.Now()
	files, err := a.rollback.Scan()
	if err != nil {
		log.Println("Unable to scan - ", a.rollback.Path(), err)
		return
	}
	for _, file := range files {
		dstInfo, e := entity.NewFile(filepath.Join(a.target.Path(), file.RPath()), a.target.Path())
		if e != nil {
			ok = false
			break
		}
		match, _ := file.IsSameFileAt(dstInfo, true)
		ok = match
		if !match {
			break
		}
	}
	log.Println("ROLLBACK VERIFIED: ", a.target.Path(), "OK:", ok, "T:", time.Since(start))
	return
}

func (a *_agent) PatchExtract(name, ext string) (err error) {
	start := time.Now()
	// cleaning the patch directory
	err = a.patch.Clean()
	if err != nil {
		return rpc.LogError(status.Errorf(codes.Internal, "Cannot clean default patch folder: %v | %v", a.patch.Path(), err))
	}
	if strings.HasSuffix(ext, ".gz") {
		t := entity.NewTar(name, ext, a.assets.Path())
		err = t.Untar(a.patch.Path())
	} else {
		f, err := entity.NewFile(filepath.Join(a.assets.Path(), name+ext), a.assets.Path())
		if err != nil {
			return err
		}
		err = f.Copy(a.patch)
	}
	log.Println("PATCH EXTRACT: ", a.patch.Path(), "T:", time.Since(start))
	return
}

func (a *_agent) VerifyExtracted() (files []entity.File, ok bool, err error) {
	start := time.Now()
	files, err = a.patch.Scan()
	if err != nil {
		log.Println("Unable to scan patch folder", err)
		return
	}
	ok = true // todo - cross check file count if needed
	log.Println("PATCH VERIFED: ", a.target.Path(), "OK:", ok, "T:", time.Since(start))
	return
}

func (a *_agent) ListAssets() (out []entity.Entity, err error) {
	start := time.Now()
	files, err := a.assets.Scan()
	if err != nil {
		log.Println("Unable to scan assets folder", err)
		return
	}
	var s entity.MultiEntity
	for _, f := range files {
		s = append(s, f)
	}
	sort.Sort(s)
	c := 0
	if len(s) > 10 {
		c = len(s) - 10
	}
	out = s[c:]
	log.Println("LIST: found", a.assets.Path(), len(s), len(out), "T:", time.Since(start))
	return out, nil
}
