package utility

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
)

// RemedyDirectory - default patch location
var RemedyDirectory string

// RevokeDirectory - default last rollback location
var RevokeDirectory string

// AssetsDirectory - default tar files location
var AssetsDirectory string

// ConfDirectory - default configuration loaction
var ConfDirectory string = "conf"

func init() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Cannot load the configuration files: ", err)
	}

	cwd, _ := os.Getwd()
	precheck(cwd, config.RemedyDir)
	precheck(cwd, config.RevokeDir)
	precheck(cwd, config.AssetsDir)

	RemedyDirectory = config.RemedyDir
	RevokeDirectory = config.RevokeDir
	AssetsDirectory = config.AssetsDir
}

func precheck(cwd, d string) {
	_, err := dir.New(d)
	if err != nil {
		log.Fatal("Create new folder: ", filepath.Join(cwd, d))
	}
}
