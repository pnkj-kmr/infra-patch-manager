package main

import (
	"fmt"

	"github.com/pnkj-kmr/patch/record"
)

func main() {
	// READ CHECK
	// d, err := record.NewDir(record.LocationPatch)
	// fmt.Println(d)
	// fmt.Println(d.Info.Mode())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// files, err := d.Scan()
	// // fmt.Println(files)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {
	// 	fmt.Println(file.Path, file.RPath, file.SubPath, file.Info.Name())
	// 	// fmt.Println(file)
	// }

	// FILE WRITE CHECK
	// src := "portfolio/redo/rent_jan_mar_2021.pdf"
	// src := "portfolio/redo/sub1/sub1-2/rent_receipt.pdf"
	// f, err := record.NewFile(record.LocationPatch, src)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dst, err := record.NewDir("tmp")
	// if err != nil && os.IsNotExist(err) {
	// 	fmt.Println("Copy: ", err)
	// 	return
	// }

	// err = f.Copy(dst)
	// if err != nil {
	// 	fmt.Printf("Copy failed %q\n", err)
	// } else {
	// 	fmt.Printf("Copy succeeded\n")
	// }

	dsts := []string{"tmp", "tmp2", "tmp3"}
	err := record.ApplyPatch(dsts)
	if err != nil {
		fmt.Printf("Copy failed %q\n", err)
	}
	// // FOLDER WRITE CHECK
	// src, err := record.NewDir(record.LocationPatch)
	// if err != nil && os.IsNotExist(err) {
	// 	fmt.Println("Copy: ", err)
	// 	return
	// }
	// dst := "tmp"
	// err = src.Copy(dst)
	// if err != nil {
	// 	fmt.Printf("Folder Copy failed %q\n", err)
	// } else {
	// 	fmt.Printf("Folder Copy succeeded\n")
	// }

	// // FOLDER WRITE CHECK
	// src, err = record.NewDir("tmp")
	// if err != nil && os.IsNotExist(err) {
	// 	fmt.Println("Copy: ", err)
	// 	return
	// }
	// dst = record.LocationRollback
	// err = src.Copy(dst)
	// if err != nil {
	// 	fmt.Printf("Folder Copy failed %q\n", err)
	// } else {
	// 	fmt.Printf("Folder Copy succeeded\n")
	// }
}
