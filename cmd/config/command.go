package cmd

import (
	"github.com/spf13/cobra"

	logger "github.com/tts2k/anitrack/cmd/logger"
	utils "github.com/tts2k/anitrack/cmd/utils"
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
			logger.Error(err)
			return
		}
	} else if !ok {
		return
	}

	// Open external editor
	ok = utils.OpenExternalEditor(config.ConfigPath)
	if ok {
		logger.Info("Config changed")
	} else {
		logger.Info("Config unchanged")
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
