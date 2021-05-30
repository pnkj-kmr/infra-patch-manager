package utility

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds the all configuration variables of application
// Variables loaded from env file and os env by viber
type Config struct {
	RemedyDir string `mapstructure:"DIR_PATCH"`
	RevokeDir string `mapstructure:"DIR_PATCH_ROLLBACK"`
	AssetsDir string `mapstructure:"DIR_ASSETS"`
	Port      string `mapstructure:"PORT"`
}

// LoadConfig helps to setup configuration from file or env variable
func LoadConfig() (config Config, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	// Looks multiple folder to match these files
	viper.AddConfigPath(filepath.Join(filepath.Dir(wd), "conf"))
	viper.AddConfigPath(filepath.Join(wd, "conf"))
	// filename with extensions
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	// override the variable from env
	viper.AutomaticEnv()
	// reading the found config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	f, _ := filepath.Rel(wd, viper.ConfigFileUsed())
	log.Println("Config Name:", f)
	return

}
