package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	logger "github.com/tts2k/anitrack/cmd/logger"
	utils "github.com/tts2k/anitrack/cmd/utils"
)

var Command *cobra.Command

func animeCommandRun(_ *cobra.Command, _ []string) {
	site, err := utils.InitSite()
	if err != nil {
		logger.Error(err)
		return
	}

	resp, err := site.Trending()
	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(resp)
}

func init() {
	Command = &cobra.Command{
		Use:   "anime",
		Short: "Search/Query anime. Running without subcommand with print a list of trending anime",
		Run: animeCommandRun,
	}
}
