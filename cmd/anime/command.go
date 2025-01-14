package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/briandowns/spinner"

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

	s := spinner.New(spinner.CharSets[11], 100 * time.Millisecond)
	s.Suffix = " Loading..."

	s.Start()

	resp, err := site.Trending()
	if err != nil {
		logger.Error(err)
		return
	}
	
	s.Stop()

	fmt.Println()

	// TODO: Pretty print the result
	fmt.Println(resp)
}

func listShowsCommandRun(cmd *cobra.Command, _ []string) {
	site, err := utils.InitSite()
	if err != nil {
		logger.Error(err)
		return
	}

	flags := cmd.Flags()
	page, _ := flags.GetUint("page")
	limit, _ := flags.GetUint("limit")

	s := spinner.New(spinner.CharSets[11], 100 * time.Millisecond)
	s.Suffix = " Loading..."

	s.Start()

	site.UserAnime(page, limit)

	s.Stop()
}

func init() {
	Command = &cobra.Command{
		Use:   "anime",
		Short: "Search/Query anime. Running without subcommand with print a list of trending anime",
		Run: animeCommandRun,
	}

	listShowsCommand := &cobra.Command{
		Use: "list-shows",
		Short: "List all shows from your library",
		Run: listShowsCommandRun,
	}

	Command.AddCommand(listShowsCommand)
	Command.PersistentFlags().UintP("page", "p", 0, "Which page to get")
	Command.PersistentFlags().UintP("limit", "l", 10, "Which number of item per page")
}
