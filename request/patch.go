package request

import (
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/pnkj-kmr/patch/segment/dir"
	"github.com/pnkj-kmr/patch/segment/tar"
)

// CleanPatchDir cleans the patch folder
func CleanPatchDir() (err error) {
	patchDir, err := dir.New(LocationPatch)
	if err != nil {
		return err
	}
	return patchDir.Clean()
}

// ExtractIntoPatchDir helps to extract tar file into patch folder
func ExtractIntoPatchDir(t *tar.T) (err error) {
	err = CleanPatchDir()
	if err != nil {
		return
	}
	return t.Untar(LocationPatch)
}

// ApplyPatchTo helps to apply patch to dst
func ApplyPatchTo(dst []string) (err error) {
	start := time.Now()
	wg := sync.WaitGroup{}
	src, err := dir.New(LocationPatch)
	if err != nil {
		return err
	}
	wg.Add(len(dst))
	for _, d := range dst {
		go func(dst string) {
			err := src.Copy(dst)
			if err != nil {
				log.Println("Apply: ERROR  ", dst, err)
			} else {
				log.Println("Apply: SUCCESS", dst)
			}
			wg.Done()
		}(d)
	}
	wg.Wait()
	log.Println("Apply: TIME   ", time.Since(start))
	return
}

// VerifyPatch helps to verify the applied patch
func VerifyPatch(dst []string) (dmap map[string]bool, err error) {
	start := time.Now()
	dmap = make(map[string]bool)
	src, err := dir.New(LocationPatch)
	if err != nil {
		return nil, err
	}
	files, err := src.Scan()
	if err != nil {
		return
	}
	var match bool
	for _, d := range dst {
		match = false
		for _, file := range files {
			dstInfo, e := dir.New(filepath.Join(d, file.RPath()))
			if e != nil {
				break
			}
			ok, _ := file.IsSameFileAt(dstInfo, false)
			match = ok
			if !ok {
				break
			}
		}
		dmap[d] = match
	}
	log.Println("Value:", dmap)
	log.Println("Check: TIME   ", time.Since(start))
	return
}
