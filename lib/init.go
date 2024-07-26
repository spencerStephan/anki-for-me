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
	ConfigExists bool
}

func (p *InitParams) CheckIfConfigExists() error {
	operatingSystem := runtime.GOOS
	homeDir, err := os.UserHomeDir()
	const FOLDER_NAME = "wofi"
	var configFilePath string

	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	switch operatingSystem {
	case "linux":
		configDirName := ".config"
		configDirPath := filepath.Join(homeDir, configDirName)
		configFilePath = filepath.Join(configDirPath, FOLDER_NAME)
	case "darwin":
		configFilePath = filepath.Join(homeDir, FOLDER_NAME)
	default:
		return errors.New("please use either Linux or MacOS")
	}

	_, err = os.Stat(configFilePath)
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
		ConfigExists: false,
	}
	if err := params.CheckIfConfigExists(); err != nil {
		log.Fatal(err)
	}
	return params
}
