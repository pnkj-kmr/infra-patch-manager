package task

import (
	"log"
	"time"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/utility"
)

// CheckRights helps to check read/write at given targets
func CheckRights(targets []string) (dmap map[string]bool, err error) {
	start := time.Now()
	dmap = make(map[string]bool)
	randomStr := utility.RandomString(10)
	for _, d := range targets {
		var match bool
		dst, err := dir.New(d)
		if err == nil {
			err = dst.Create(randomStr)
			// fmt.Println("Folder Create..", err)
			if err == nil {
				err = dst.Remove(randomStr)
				// fmt.Println("Folder Delete..", err)
				if err == nil {
					err = dst.CreateFile(randomStr)
					// fmt.Println("File Create..", err)
					if err == nil {
						err = dst.RemoveFile(randomStr)
						// fmt.Println("File Delete..", err)
						if err == nil {
							match = true
						}
					}
				}
			}
		}
		dmap[d] = match
	}
	log.Println("Returns    :", dmap)
	log.Println("CheckRights: TIME   ", time.Since(start))
	return
}
