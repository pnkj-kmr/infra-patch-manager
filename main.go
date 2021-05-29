package main

import (
	"fmt"

	"github.com/pnkj-kmr/patch/module/tar"
	"github.com/pnkj-kmr/patch/task"
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

	t := tar.New("ROLLBACK__1622290827", "tar.gz", "uploads")
	err = task.ExtractIntoPatchDir(t)
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

	// err = task.CleanRollbackDir()
	// if err != nil {
	// 	fmt.Println("CleanRollbackDir>>>>\n", err)
	// }

	// t := tar.T{Name: utility.RandomString(6), Ext: "tar.gz"}
	// t := tar.New(utility.RandomString(6), "tar.gz", "")
	// t := tar.New("HFfiAW", "tar.gz", "")
	// err := t.Tar([]string{filepath.Join("portfolio", "redo")})
	// err := t.Untar(filepath.Join("asset", "patch"))
	// fmt.Println("TAR : >>>", err)

	// s, _ := file.F{P: "aaaa/bbb/c.txt", R: "bbb/c.txt", S: "", I: nil}
	// s, _ := file.New("asset/patch/redo/rent_jan_mar_2021.pdf", "asset/patch")
	// fmt.Println(s.RPath())

}
