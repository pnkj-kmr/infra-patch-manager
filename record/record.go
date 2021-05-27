package record

import (
	"log"
	"path/filepath"
	"sync"
	"time"
)

// LocationPatch - default patch location
var LocationPatch string

// LocationRollback - default last rollback location
var LocationRollback string

func init() {
	LocationPatch = filepath.Join("portfolio", "redo")
	LocationRollback = filepath.Join("portfolio", "undo")
}

// ApplyPatch helps to apply patch to dst
func ApplyPatch(dst []string) (err error) {
	start := time.Now()
	wg := sync.WaitGroup{}
	src, err := NewDir(LocationPatch)
	if err != nil {
		return err
	}
	wg.Add(len(dst))
	for _, d := range dst {
		go func(dst string) {
			err := src.Copy(dst)
			if err != nil {
				log.Println("Copy: ERROR  ", dst, err)
			} else {
				log.Println("Copy: SUCCESS", dst)
			}
			wg.Done()
		}(d)
	}
	wg.Wait()
	duration := time.Since(start)
	log.Println("Copy: TIME   ", duration)
	return
}
