package entity

import (
	"math/rand"
	"time"
)

const (
	// letters - taking alphanumeric variables for randoms
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// setting up the required default folders
	resourceDir string = "resources"
	assetsDir   string = "resources/assets"
	patchDir    string = "resources/patch"
	rollbackDir string = "resources/rollback"

	// ConfPath - default configuration loaction
	ConfPath string = "conf"
)

// C contains all system level configurations
var C Conf

var passcode string

func init() {
	// setting random seed
	rand.Seed(time.Now().UnixNano())
	// setting up the system level variable
	C = &_conf{patchDir, rollbackDir, assetsDir}
}
