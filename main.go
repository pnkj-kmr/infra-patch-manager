package main

import (
	"fmt"

	"github.com/pnkj-kmr/patch/module/tar"
	"github.com/pnkj-kmr/patch/task"
	"github.com/pnkj-kmr/patch/utility"
)

func main() {
	dsts := []string{"tmp3"}

	err := task.RollbackFrom(dsts[0])
	if err != nil {
		fmt.Println("Rollback>>>>\n", err)
	}

	dmap, err := task.VerifyRollback(dsts[0])
	if err != nil {
		fmt.Println("Watch  >>>", dmap, err)
	}

	t := tar.New("test", "tar.gz", utility.AssetsDirectory)
	err = task.ExtractIntoRemedyDir(t)
	if err != nil {
		fmt.Printf("Extract--- %q\n", err)
	}

	err = task.ApplyPatchTo(dsts)
	if err != nil {
		fmt.Printf("Copy failed %q\n", err)
	}

	dmap, err = task.VerifyPatch(dsts)
	if err != nil {
		fmt.Println("Watch  >>>", dmap, err)
	}

	dmap, err = task.VerifyRollback(dsts[0])
	if err != nil {
		fmt.Println("Watch  >>>", dmap, err)
	}

	dmap, err = task.CheckRights(dsts)
	if err != nil {
		fmt.Println("Check Rights>>>>\n", dmap, err)
	}

}
