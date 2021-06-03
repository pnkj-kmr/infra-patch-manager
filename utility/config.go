package utility

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// Config holds the all configuration variables of application
// Variables loaded from env file and os env by viber
type Config struct {
	RemedyDir    string `mapstructure:"DIR_PATCH"`
	RevokeDir    string `mapstructure:"DIR_PATCH_ROLLBACK"`
	AssetsDir    string `mapstructure:"DIR_ASSETS"`
	Port         string `mapstructure:"PORT"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
}

// LoadConfig helps to setup configuration from file or env variable
func LoadConfig() (config Config, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	// Looks multiple folder to match these files
	viper.AddConfigPath(filepath.Join(filepath.Dir(wd), ConfDirectory))
	viper.AddConfigPath(filepath.Join(wd, ConfDirectory))
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

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig(c Config) fiber.Config {

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(c.ReadTimeout),
	}
}
