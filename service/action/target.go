package action

import (
	"log"
	"time"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/utility"
)

// RemoteRWRights helps to check read/write at given target
func RemoteRWRights(target string) (match bool, err error) {
	start := time.Now()
	randomStr := utility.RandomString(10)
	dst, err := dir.New(target)
	if err == nil {
		err = dst.Create(randomStr)
		// fmt.Println("Folder Create..", err)
		if err == nil {
			err = dst.Remove(randomStr)
			// fmt.Println("Folder Delete..", err)
			if err == nil {
				err = dst.CreateFile(randomStr, []byte("temp data into file"))
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
	log.Println("R/W CHECKS FOR", target, "OK:", match, "T:", time.Since(start))
	return
}
