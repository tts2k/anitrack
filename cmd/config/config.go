package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AuthData struct {
	AccessToken  string `mapstructure:"access_token"`
	RefreshToken string `mapstructure:"refresh_Token"`
}

type Config struct {
	ActiveSite string
	ConfigPath string
	Verbose    bool
	Mal        AuthData `mapstructure:"mal"`
	Kitsu      AuthData `mapstructure:"kitsu"`
	AniList    AuthData `mapstructure:"aniList"`
}

var config Config

func getConfigDir() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Println(err)
		log.Println("Cannot get user config dir")
		return "", err
	}

	return filepath.Join(userConfigDir, "anitrack"), nil
}

// Check config directory. Return true if exists, false if not, together with config path
// If return path is an empty string, it means that there is another problem which made
// config dir inaccessible
func CheckConfigDir() (string, bool) {
	anitrackConfigDir, err := getConfigDir()
	if err != nil {
		return anitrackConfigDir, false
	}

	_, err = os.Stat(anitrackConfigDir)
	if errors.Is(err, os.ErrNotExist) {
		return anitrackConfigDir, false
	}
	if err != nil {
		log.Println(err)
		return "", false
	}

	return anitrackConfigDir, true
}

func InitConfigDir() error {
	anitrackConfigDir, _ := CheckConfigDir()
	if anitrackConfigDir == "" {
		return errors.New("init config dir failed")
	}
	err := os.MkdirAll(anitrackConfigDir, 0700)
	if err != nil {
		return err
	}
	fmt.Printf("A new directory has been created at %s\n", anitrackConfigDir)

	return nil
}

func initViper() error {
	configDir, _ := CheckConfigDir()

	// set config path
	viper.SetConfigFile(filepath.Join(configDir, "anitrack.toml"))
	viper.SetConfigType("toml")

	return viper.ReadInConfig()
}

func InitConfig() {
	if err := initViper(); err != nil && !os.IsNotExist(err) {
		log.Println(err)
		log.Println("Error reading config file: " + viper.ConfigFileUsed())
	}

	// Default active site
	viper.SetDefault("default_site", "Kitsu")

	config.ActiveSite = viper.GetString("default_site")
	config.ConfigPath = viper.ConfigFileUsed()

	// Unmarshal config file
	if err := viper.Unmarshal(&config); err != nil {
		log.Println("Error parsing config file")
		log.Println("Using default config")
	}
}

func GetConfig() *Config {
	return &config
}
