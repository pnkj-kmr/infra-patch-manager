package entity

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// letters - taking alphanumeric variables for randoms
const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// setting up the required default folders
var resourceDir string = "resources"
var assetsDir string = "resources/assets"
var patchDir string = "resources/patch"
var rollbackDir string = "resources/rollback"

// ConfPath - default configuration loaction
var ConfPath string = "conf"

// Conf holds the all configuration variables of application
// Variables loaded from env file and os env by viber
type Conf struct {
	PatchPath    string `mapstructure:"DIR_PATCH"`
	RollbackPath string `mapstructure:"DIR_PATCH_ROLLBACK"`
	AssetPath    string `mapstructure:"DIR_ASSETS"`
}

// C contains all system level configurations
var C Conf

// TODO - rollbackPath resource directory
var rollbackPath string = "conf"

func init() {
	// setting random seed
	rand.Seed(time.Now().UnixNano())
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	// Creating the default folders for applications
	for _, d := range []string{resourceDir, assetsDir, patchDir, rollbackDir} {
		_, err = CreateDirectoryIfNotExists(d)
		if err != nil {
			log.Fatal(filepath.Join(wd, d), "ERROR:", err)
		}
	}

	C = Conf{
		PatchPath:    patchDir,
		RollbackPath: rollbackDir,
		AssetPath:    assetsDir,
	}

	// // init the system default configuration
	// err := initConf(&C)
	// if err != nil {
	// 	log.Fatal("Cannot load the configuration files: ", err)
	// }
	// cwd, _ := os.Getwd()
	// precheck(cwd, C.PatchPath)
	// precheck(cwd, C.RollbackPath)
	// precheck(cwd, C.AssetPath)
}

// func precheck(cwd, d string) {
// 	_, err := NewDir(d)
// 	if err != nil {
// 		log.Fatal("CREATE A FOLDER AS:- ", filepath.Join(cwd, d))
// 	}
// }

// // initConf helps to setup configuration from file or env variable
// func initConf(c *Conf) (err error) {
// 	// Viper Configuraton Setup -- conf/config.env or os env
// 	// Looks multiple folder to match these files
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		return
// 	}
// 	viper.AddConfigPath(filepath.Join(filepath.Dir(wd), ConfPath))
// 	viper.AddConfigPath(filepath.Join(wd, ConfPath))
// 	// filename with extensions
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("env")
// 	// override the variable from env
// 	viper.AutomaticEnv()
// 	// reading the found config file
// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}
// 	err = viper.Unmarshal(c)
// 	filepath.Rel(wd, viper.ConfigFileUsed())
// 	// f, _ := filepath.Rel(wd, viper.ConfigFileUsed())
// 	// fmt.Println("Conf Path	:", f)
// 	return
// }
