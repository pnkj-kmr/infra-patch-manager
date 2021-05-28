package main

import (
	"fmt"

	"github.com/pnkj-kmr/patch/request"
	"github.com/pnkj-kmr/patch/segment/tar"
)

func main() {
	dsts := []string{"tmp3"}

	err := request.RollbackFrom(dsts[0])
	if err != nil {
		fmt.Println("Rollback>>>>\n", err)
	}

	dmap, err := request.VerifyRollback(dsts[0])
	if err != nil {
		fmt.Println("Watch  >>>", dmap, err)
	}

	t := tar.New("nLZe4C__1622227634638736000", "tar.gz", "uploads")
	err = request.ExtractIntoPatchDir(t)
	if err != nil {
		fmt.Printf("Extract--- %q\n", err)
	}

	err = request.ApplyPatchTo(dsts)
	if err != nil {
		fmt.Printf("Copy failed %q\n", err)
	}

	dmap, err = request.VerifyPatch(dsts)
	if err != nil {
		fmt.Println("Watch  >>>", dmap, err)
	}

	dmap, err = request.VerifyRollback(dsts[0])
	if err != nil {
		fmt.Println("Watch  >>>", dmap, err)
	}

	dmap, err = request.CheckRights(dsts)
	if err != nil {
		fmt.Println("Check Rights>>>>\n", dmap, err)
	}

	// err = request.CleanRollbackDir()
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
