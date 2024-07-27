package cmd

import (
	"fmt"
	"github.com/spencerStephan/anki-for-me/lib"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(initializeCmd)
}

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up Anki-for-me to work on your system",
	Long:  "Enable your database and set up your user defaults for Anki-for-me to work properly",
	Run: func(cmd *cobra.Command, args []string) {
		params := lib.StartProgram()
		if params.ConfigExists {
			fmt.Println("Anki-for-me is already setup, use anki-for-me or afm to get started")
			os.Exit(1)
		}
	},
}
