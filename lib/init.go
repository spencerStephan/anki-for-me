package lib

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type InitParams struct {
	ConfigDir    []fs.DirEntry
	SqliteFile   string
	ConfigExists bool
}

func (p *InitParams) SetConfigDirPath() error {
	operatingSystem := runtime.GOOS
	homeDir, err := os.UserHomeDir()
	//const folderName = "anki-for-me"
	const folderName = "anki-for-me"

	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	switch operatingSystem {
	case "linux":
		configDirName := ".config"
		configDirPath := filepath.Join(homeDir, configDirName, folderName)
		p.ConfigDir, err = os.ReadDir(configDirPath)
		if err != nil {
			p.ConfigExists = false
		} else {
			p.ConfigExists = true
		}

	case "darwin":
		configDirName := ".config"
		configDirPath := filepath.Join(homeDir, configDirName, folderName)
		p.ConfigDir, err = os.ReadDir(configDirPath)
		if err != nil {
			p.ConfigExists = false
		} else {
			p.ConfigExists = true
		}
	default:
		return errors.New("please use either Linux or MacOS")
	}
	return nil
}

func StartProgram() *InitParams {
	params := &InitParams{
		ConfigExists: false,
	}
	if err := params.SetConfigDirPath(); err != nil {
		log.Fatal(err)
	}
	return params
}
