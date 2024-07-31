package cmd

import (
	"fmt"
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
		config := lib.InitConfig()
		dir := config.Dir
		dir.GetDir(false)
		db, err := lib.NewSqlite(dir)
		if err != nil {
			log.Fatal("There was an error", err)
		}
		exists := db.SqlExists(dir)
		fmt.Println(exists)

		lib.InitServices(config, db)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var viperReadType string
		if CfgFile != "" {
			viperReadType = "flag"
			viper.SetConfigFile(CfgFile)
			if ConfigIsSaved {
				lib.OverrideConfigFile(CfgFile)
				viper.SetConfigType("yaml")
				viper.SetConfigName("config")
				viper.AddConfigPath("$HOME/.config/anki-for-me")
			}
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
				log.Fatal("Invalid file, file must exist and be in YAML format.")
			case "automatic":
				if !strings.Contains(cmd.Name(), "init") {
					log.Fatal("Please run anki-for-me init to initialize your config or pass in a config file using --config.")
				}
			}
		}
	},
}

var (
	CfgFile       string
	sqlFile       string
	ConfigIsSaved bool
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&CfgFile, "config", "c", "", "config file (default is $HOME/.config/anki-for-me/config.yaml")
	rootCmd.PersistentFlags().BoolVarP(&ConfigIsSaved, "save", "s", false, "Using this flag will override your config file with the one you entered using --config")
}
