package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	config "github.com/tts2k/anitrack/cmd/config"
	logger "github.com/tts2k/anitrack/cmd/logger"
	utils "github.com/tts2k/anitrack/cmd/utils"
)

var Command *cobra.Command

func loginCommandRun(_ *cobra.Command, _ []string) {
	s, err := utils.InitSite()
	if err != nil {
		logger.Error(err)
		return
	}

	accessToken, refreshToken, err := s.Login()
	if err != nil {
		logger.Error(err)
		return
	}

	config.SetTokens(accessToken, refreshToken)
	fmt.Println("Login success!")
}

func logoutCommandRun(_ *cobra.Command, _ []string) {
	config.RemoveTokens()
	fmt.Println("Logout success!")
}

func setCommandRun(_ *cobra.Command, _ []string) {
	fmt.Println("Not implemented")
}

func init() {
	Command = &cobra.Command{
		Use:   "site",
		Short: "Switch site or login",
	}

	setCommand := &cobra.Command{
		Use:   "set",
		Short: "Change site",
		Run:   setCommandRun,
	}

	loginCommand := &cobra.Command{
		Use:   "login",
		Short: "Login to current site",
		Run:   loginCommandRun,
	}

	logoutCommandRun := &cobra.Command{
		Use:   "logout",
		Short: "Logout to current site",
		Run:   logoutCommandRun,
	}

	Command.AddCommand(setCommand)
	Command.AddCommand(loginCommand)
	Command.AddCommand(logoutCommandRun)
}
