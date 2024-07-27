package main

import (
	"fmt"
	"os"

	"github.com/spencerStephan/anki-for-me/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		fmt.Println(err)
		os.Exit(1)
	}
}
