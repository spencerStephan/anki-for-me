package cmd

import (
	"fmt"
	"github.com/spencerStephan/anki-for-me/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initializeCmd)
}

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Anki-for-me to work on your system",
	Long:  "Enable your database and set up your user defaults for Anki-for-me to work properly",
	Run: func(cmd *cobra.Command, args []string) {
		init := lib.InitService{
			Config: lib.UserConfig{},
			DB:     lib.Sqlite{},
		}
		fmt.Println(init.DB.Exists(init.Config.GetDir()))
	},
}
