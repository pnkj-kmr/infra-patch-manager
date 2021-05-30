package utility

import (
	"log"
)

// RemedyDirectory - default patch location
var RemedyDirectory string

// RevokeDirectory - default last rollback location
var RevokeDirectory string

// AssetsDirectory - default tar files location
var AssetsDirectory string

func init() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Cannot load the configuration files: ", err)
	}

	RemedyDirectory = config.RemedyDir
	RevokeDirectory = config.RevokeDir
	AssetsDirectory = config.AssetsDir
}
