package lib

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type InitParams struct {
	ConfigDirPath string
	ConfigExists  bool
}

func (p *InitParams) SetConfigDirPath() error {
	operatingSystem := runtime.GOOS
	homeDir, err := os.UserHomeDir()
	const FOLDER_NAME = "anki-for-me"

	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	switch operatingSystem {
	case "linux":
		configDirName := ".config"
		configDirPath := filepath.Join(homeDir, configDirName)
		p.ConfigDirPath = filepath.Join(configDirPath, FOLDER_NAME)
		return nil
	case "darwin":
		p.ConfigDirPath = filepath.Join(homeDir, FOLDER_NAME)
		return nil
	default:
		return errors.New("please use either Linux or MacOS")
	}
}

func (p *InitParams) SetConfigExists(s string) error {
	_, err := os.Stat(s)
	if err == nil {
		p.ConfigExists = true
		return nil // File exists
	} else if errors.Is(err, os.ErrNotExist) {
		return nil // File does not exist
	} else {
		return fmt.Errorf("error checking config file existence: %w", err) // Unexpected error
	}
}

func StartProgram() *InitParams {
	params := &InitParams{
		ConfigExists:  false,
		ConfigDirPath: "",
	}

	if err := params.SetConfigDirPath(); err != nil {
		log.Fatal(err)
	}

	if err := params.SetConfigExists(params.ConfigDirPath); err != nil {
		log.Fatal(err)
	}
	return params
}
