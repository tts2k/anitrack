package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"

	logger "github.com/tts2k/anitrack/cmd/logger"
	utils "github.com/tts2k/anitrack/cmd/utils"
)

var Command *cobra.Command

func userCommandRun(_ *cobra.Command, _ []string) {
	site, err := utils.InitSite()
	if err != nil {
		logger.Error(err)
		return
	}

	s := spinner.New(spinner.CharSets[11], 100 * time.Millisecond)
	s.Suffix = " Loading..."

	s.Start()

	user, err := site.User()
	if err != nil {
		logger.Error(err)
	}

	s.Stop()

	// TODO: pretty print
	fmt.Println()
	fmt.Println(user)

}

func init() {
	Command = &cobra.Command{
		Use:   "user",
		Short: "Get user information",
		Run: userCommandRun,
	}
}
