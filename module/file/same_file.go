package file

import (
	"log"

	"github.com/pnkj-kmr/patch/module"
)

// IsSameFileAt helps to verify src and dst file - same or not
func (f *F) IsSameFileAt(d module.I, old bool) (match bool, err error) {
	if old {
		match = isSameAndOld(f, d)
	} else {
		match = isSameAndNew(f, d)
	}
	log.Println("FILE_COMPARE: matched-", match, "old flag-", false, "from-", f.Path(), "to-", d.Path())
	return
}

func isSameAndNew(f1, f2 module.I) (match bool) {
	// fmt.Println(">>>", match, f1, f2)
	// fmt.Println("-----f1", f1.Name(), f1.Size(), f1.ModTime())
	// fmt.Println("-----f2", f2.Name(), f2.Size(), f2.ModTime())
	match = (f1.Name() == f2.Name()) && (f1.Size() == f2.Size()) && (f1.ModTime().Before(f2.ModTime()))
	return
}

func isSameAndOld(f1, f2 module.I) (match bool) {
	match = (f1.Name() == f2.Name()) && (f1.Size() == f2.Size()) && (f1.ModTime().After(f2.ModTime()))
	return
}
