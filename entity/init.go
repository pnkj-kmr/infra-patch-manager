package entity

import (
	"math/rand"
	"time"
)

// letters - taking alphanumeric variables for randoms
const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// setting up the required default folders
const resourceDir string = "resources"
const assetsDir string = "resources/assets"
const patchDir string = "resources/patch"
const rollbackDir string = "resources/rollback"

// ConfPath - default configuration loaction
const ConfPath string = "conf"

// C contains all system level configurations
var C Conf

func init() {
	// setting random seed
	rand.Seed(time.Now().UnixNano())
	// setting up the system level variable
	C = &_conf{patchDir, rollbackDir, assetsDir}
}
