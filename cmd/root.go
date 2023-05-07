package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/tts2k/anitrack/cmd/config"
)

const VERSION = "0.1"

var rootCmd = &cobra.Command{
	Use:   "anitrack",
	Short: "A CLI client for various anime progress tracking sites",
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print Anitrack version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	conf := config.GetConfig()
	// Persistent Flags
	rootCmd.PersistentFlags().BoolVarP(&conf.Verbose, "verbose", "v", false, "verbose output")

	// Commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(config.Command)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
