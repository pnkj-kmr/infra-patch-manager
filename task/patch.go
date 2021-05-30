package task

import (
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/module/tar"
	"github.com/pnkj-kmr/patch/utility"
)

// CleanRemedyDir cleans the patch folder
func CleanRemedyDir() (err error) {
	d, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return err
	}
	return d.Clean()
}

// ExtractIntoRemedyDir helps to extract tar file into patch folder
func ExtractIntoRemedyDir(t *tar.T) (err error) {
	err = CleanRemedyDir()
	if err != nil {
		return
	}
	return t.Untar(utility.RemedyDirectory)
}

// ApplyPatchTo helps to apply patch to target folder
func ApplyPatchTo(targets []string) (err error) {
	start := time.Now()
	wg := sync.WaitGroup{}
	src, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return err
	}
	wg.Add(len(targets))
	for _, d := range targets {
		go func(target string) {
			err := src.Copy(target)
			if err != nil {
				log.Println("Apply: ERROR  ", target, err)
			} else {
				log.Println("Apply: SUCCESS", target)
			}
			wg.Done()
		}(d)
	}
	wg.Wait()
	log.Println("Apply: TIME   ", time.Since(start))
	return
}

// VerifyPatch helps to verify the applied patch
func VerifyPatch(targets []string) (dmap map[string]bool, err error) {
	start := time.Now()
	dmap = make(map[string]bool)
	src, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return nil, err
	}
	files, err := src.Scan()
	if err != nil {
		return
	}
	var match bool
	for _, d := range targets {
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
