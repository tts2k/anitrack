package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	config "github.com/tts2k/anitrack/cmd/config"
	site "github.com/tts2k/anitrack/lib"
)

var Command *cobra.Command

func loginCommandRun(_ *cobra.Command, _ []string) {
	conf := config.GetConfig()

	s := site.New(conf.ActiveSite)
	accessToken, refreshToken, err := s.Login()
	if err != nil {
		return
	}

	config.SetToken(accessToken, refreshToken)
	fmt.Println("Login success!")
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
		Run:  loginCommandRun,
	}

	Command.AddCommand(setCommand)
	Command.AddCommand(loginCommand)
}
