package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	config "github.com/tts2k/anitrack/cmd/config"
	logger "github.com/tts2k/anitrack/cmd/logger"
	site "github.com/tts2k/anitrack/cmd/site"
	anime "github.com/tts2k/anitrack/cmd/anime"
)

const VERSION = "0.1"

var rootCmd = &cobra.Command{
	Use:   "anitrack",
	Short: "A CLI client for various anime progress tracking sites",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Anitrack version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	// Persistent Flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	// Viper bind
	_ = viper.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))

	// Commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(config.Command)
	rootCmd.AddCommand(site.Command)
	rootCmd.AddCommand(anime.Command)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err)
	}
}
