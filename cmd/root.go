package cmd

import (
	"log"
	"strings"

	"github.com/spencerStephan/anki-for-me/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "afm",
	Version: "0.0.1",
	Short:   "A CLI tool for memorizing things you've learned.",
	Long:    "An interactive TUI that helps you learn by using techniques such as spaced repetition.",
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.NewConfig()
		db, err := lib.NewSqlite(config.Directory)
		if err != nil {
			log.Fatal("There was an error", err)
		}
		lib.InitServices(config, db)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var viperReadType string
		if cfgFile != "" {
			viperReadType = "flag"
			viper.SetConfigFile(cfgFile)
		} else {
			viperReadType = "automatic"
			viper.SetConfigType("yaml")
			viper.SetConfigName("config")
			viper.AddConfigPath("$HOME/.config/anki-for-me")
		}

		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			switch viperReadType {
			case "flag":
				log.Fatal("Invalid file type, must be YAML format.")
			case "automatic":
				if !strings.Contains(cmd.Name(), "init") {
					log.Fatal("Please run anki-for-me init to initialize your config or pass in a config file using --config.")
				}
			}
		}
	},
}

var (
	cfgFile string
	sqlFile string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/anki-for-me/config.yaml")
}
