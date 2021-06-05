package server

import (
	"log"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func applyPatchTo(target string, backup bool) (err error) {
	start := time.Now()
	// R/W check for target folder
	rw, err := remoteRWRights(target)
	if !rw || err != nil {
		return service.LogError(status.Errorf(codes.Internal, "target remote is not having read/write permission: %v | %v", target, err))
	}
	if backup {
		// taking a rollbackup backup
		err = rollbackFrom(target)
		if err != nil {
			return service.LogError(status.Errorf(codes.Internal, "rollback backup failed: %v | %v", target, err))
		}
		// verifying the rollback patch
		_, err := verifyRollback(target)
		if err != nil {
			return service.LogError(status.Errorf(codes.Internal, "rollback files are improper: %v | %v", target, err))
		}
	}
	// applying patch to given folder
	err = patchTo(target)
	if err != nil {
		return service.LogError(status.Errorf(codes.Internal, "Apply patch failed: %v | %v", target, err))
	}
	log.Println("PATCH APPLY TO", target, "T:", time.Since(start))
	return
}

func verifyPatch(target string) (f []*pb.FILE, match bool, err error) {
	start := time.Now()
	src, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		log.Println("Patch folder does not exist or any", err)
		return
	}
	files, err := src.Scan()
	if err != nil {
		log.Println("Unable to scan patch folder", err)
		return
	}
	f = make([]*pb.FILE, len(files))
	for i, file := range files {
		dstInfo, e := dir.New(filepath.Join(target, file.RPath()))
		if e != nil {
			log.Println("Cannot find the file path", filepath.Join(target, file.RPath()), e)
			continue
		}
		ok, _ := file.IsSameFileAt(dstInfo, false)
		match = ok
		if !ok {
			break
		}
		f[i] = service.ConvertToFILE(dstInfo)
	}
	log.Println("PATCH VERIFED FOR", target, "OK:", match, "T:", time.Since(start))
	return
}

func patchTo(target string) (err error) {
	start := time.Now()
	src, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		log.Println("Patch folder does not exist or any", err)
		return err
	}
	err = src.Copy(target)
	log.Println("PATCHING T:", time.Since(start))
	return
}
