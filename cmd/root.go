package cmd

import (
	"github.com/spf13/cobra"

	config "github.com/tts2k/anitrack/cmd/config"
	"github.com/tts2k/anitrack/cmd/logger"
	site "github.com/tts2k/anitrack/cmd/site"
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
		logger.Println(VERSION)
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	conf := config.GetConfig()
	// Initialize logger
	logger.InitLogger(conf.Verbose)

	// Persistent Flags
	rootCmd.PersistentFlags().BoolVarP(&conf.Verbose, "verbose", "v", false, "verbose output")

	// Commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(config.Command)
	rootCmd.AddCommand(site.Command)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Errorln(err)
	}
}
