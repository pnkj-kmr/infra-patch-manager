package agent

import (
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
)

func backupExistingRollback() (err error) {
	d, err := entity.NewDir(entity.C.RollbackPath())
	if err != nil {
		log.Println("Unable to load rollback folder", entity.C.RollbackPath(), err)
		return err
	}
	assetDir, err := entity.NewDir(entity.C.AssetPath())
	if err != nil {
		log.Println("Unable to load assets folder", entity.C.AssetPath(), err)
		return err
	}
	t := entity.NewTar(entity.RandomStringWithTime(0, "ROLLBACK"), ".tar.gz", assetDir.Path())
	return t.Tar([]string{d.Path()})
}
