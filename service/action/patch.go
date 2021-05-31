package action

import (
	"log"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ApplyPatchTo helps to apply patch to target folder
func ApplyPatchTo(target string, backup bool) (err error) {
	// R/W check for target folder
	rw, err := RemoteRWRights(target)
	if !rw || err != nil {
		return logError(status.Errorf(codes.Internal, "target remote is not having read/write permission: %v | %v", target, err))
	}
	if backup {
		// taking a rollbackup backup
		err = RollbackFrom(target)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "rollback backup failed: %v", err))
		}
		// verifying the rollback patch
		ok, err := VerifyRollback(target)
		if !ok || err != nil {
			return logError(status.Errorf(codes.Internal, "rollback files is improper: %v | %v", target, err))
		}
	}
	// applying patch to given folder
	err = PatchTo(target)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "patch failed: %v", err))
	}
	return
}

// VerifyPatch helps to verify the applied patch
func VerifyPatch(target string) (f []*pb.FILE, match bool, err error) {
	start := time.Now()
	src, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return
	}
	files, err := src.Scan()
	if err != nil {
		return
	}
	f = make([]*pb.FILE, len(files))
	for i, file := range files {
		dstInfo, e := dir.New(filepath.Join(target, file.RPath()))
		if e != nil {
			break
		}
		ok, _ := file.IsSameFileAt(dstInfo, false)
		match = ok
		if !ok {
			break
		}
		f[i] = convertToFILE(file)
	}
	log.Println("Value:", match)
	log.Println("Check: TIME   ", time.Since(start))
	return
}

// PatchTo helps to apply patch to target folder
func PatchTo(target string) (err error) {
	start := time.Now()
	src, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return err
	}
	err = src.Copy(target)
	if err != nil {
		log.Println("Apply: ERROR  ", target, err)
	} else {
		log.Println("Apply: SUCCESS", target)
	}
	log.Println("Apply: TIME   ", time.Since(start))
	return
}
