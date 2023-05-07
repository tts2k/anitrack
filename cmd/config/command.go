package config

import (
	"github.com/spf13/cobra"

	"github.com/tts2k/anitrack/cmd/logger"
	"github.com/tts2k/anitrack/cmd/utils"
)

var Command *cobra.Command

func editCommandRun(_ *cobra.Command, _ []string) {
	config := GetConfig()

	// Config dir access check
	configDir, ok := CheckConfigDir()
	if !ok && configDir != "" {
		proceed := utils.Prompt(
			"Config dir does not exist. Do you want to create a new one?",
			false,
		)

		var err error
		if proceed {
			err = InitConfigDir()
		} else {
			return
		}
		if err != nil {
			logger.Debugln(err)
			return
		}
	} else if !ok {
		return
	}

	// Open external editor
	ok = utils.OpenExternalEditor(config.ConfigPath)
	if ok {
		logger.Println("Config changed")
	} else {
		logger.Println("Config unchanged")
	}
}

func init() {
	Command = &cobra.Command{
		Use:   "config",
		Short: "Print or change anitrack config",
	}

	editComand := &cobra.Command{
		Use:   "edit",
		Short: "Edit config file with an external editor",
		Long:  "Call an external text editor to edit config file. Default to $EDTIOR environment variable but fallback to nano if not set",
		Run:   editCommandRun,
	}

	Command.AddCommand(editComand)
}
