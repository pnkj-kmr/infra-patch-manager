package entity

import (
	"math/rand"
	"path/filepath"
	"time"
)

const (
	// letters - taking alphanumeric variables for randoms
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// ConfPath - default configuration loaction
	ConfPath string = "conf"
)

var passcode string

// setting up the required default folders
var resourceDir string = "resources"
var assetsDir string = filepath.Join(resourceDir, "assets")
var patchDir string = filepath.Join(resourceDir, "patch")
var rollbackDir string = filepath.Join(resourceDir, "rollback")

// C contains all system level configurations
var C Conf

func init() {
	// setting random seed
	rand.Seed(time.Now().UnixNano())
	// setting up the system level variable
	C = &_conf{patchDir, rollbackDir, assetsDir}
}
