package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	logger "github.com/tts2k/anitrack/cmd/logger"
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
		logger.Error(err)
		logger.Error("Cannot get user config dir")
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
		logger.Error(err)
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
	logger.Warn("A new directory has been created at " + anitrackConfigDir)

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
		logger.Error(err)
		logger.Error("Error reading config file: " + viper.ConfigFileUsed())
	}

	// Default active site
	viper.SetDefault("active_site", "Kitsu")

	config.ActiveSite = viper.GetString("active_site")
	config.ConfigPath = viper.ConfigFileUsed()

	// Unmarshal config file
	if err := viper.Unmarshal(&config); err != nil {
		logger.Error("Error parsing config file")
		logger.Warn("Using default config")
	}
}

func GetConfig() *Config {
	return &config
}

func SetTokens(accessToken string, refreshToken string) {
	switch strings.ToLower(config.ActiveSite) {
	case "kitsu":
		config.Kitsu.AccessToken = accessToken
		config.Kitsu.RefreshToken = refreshToken
		viper.Set("Kitsu", config.Kitsu)
	case "mal":
	case "anilist":
	}

	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error writing config to disk")
	}
}

// https://github.com/spf13/viper/issues/632
func viperUnset(key string) error {
    configMap := viper.AllSettings()
    delete(configMap, key)
    encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
    err := viper.ReadConfig(bytes.NewReader(encodedConfig))
    if err != nil {
        return err
    }
	return nil
}

func RemoveTokens() {
	switch strings.ToLower(config.ActiveSite) {
	case "kitsu":
		config.Kitsu.AccessToken = ""
		config.Kitsu.RefreshToken = ""
		viperUnset("Kitsu")
	case "mal":
		config.Mal.AccessToken = ""
		config.Mal.RefreshToken = ""
		viperUnset("Mal")
	case "anilist":
		config.AniList.AccessToken = ""
		config.AniList.RefreshToken = ""
		viperUnset("Anilist")
	}

	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error writing config to disk")
	}
}
