package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "afm",
	Version: "0.0.1",
	Short:   "A CLI tool for memorizing things you've learned.",
	Long:    "An interactive TUI that helps you learn by using techniques such as spaced repetition.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var cfgFile string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/anki-for-me/config.yaml")
	rootCmd.PersistentFlags()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath("$HOME/.config/anki-for-me")

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {
			fmt.Println(err)
		}
	}
	viper.AutomaticEnv()
}
