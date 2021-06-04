package action

import (
	"log"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ApplyPatchTo helps to apply patch to target folder
func ApplyPatchTo(target string, backup bool) (err error) {
	start := time.Now()
	// R/W check for target folder
	rw, err := RemoteRWRights(target)
	if !rw || err != nil {
		return logError(status.Errorf(codes.Internal, "target remote is not having read/write permission: %v | %v", target, err))
	}
	if backup {
		// taking a rollbackup backup
		err = RollbackFrom(target)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "rollback backup failed: %v | %v", target, err))
		}
		// verifying the rollback patch
		_, err := VerifyRollback(target)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "rollback files are improper: %v | %v", target, err))
		}
	}
	// applying patch to given folder
	err = PatchTo(target)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Apply patch failed: %v | %v", target, err))
	}
	log.Println("PATCH APPLY TO", target, "T:", time.Since(start))
	return
}

// VerifyPatch helps to verify the applied patch
func VerifyPatch(target string) (f []*pb.FILE, match bool, err error) {
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
		f[i] = ConvertDToFILE(dstInfo)
	}
	log.Println("PATCH VERIFED FOR", target, "OK:", match, "T:", time.Since(start))
	return
}

// PatchTo helps to apply patch to target folder
func PatchTo(target string) (err error) {
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
