package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config interface {
	GetDir() []fs.DirEntry
	Create() []fs.DirEntry
}

type UserConfig struct{}

func (c UserConfig) GetDir() []fs.DirEntry {
	operatingSystem := runtime.GOOS
	var configDir []fs.DirEntry
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home dir: %v", err)
	}
	const folderName = "anki-for-me"

	switch operatingSystem {
	case "windows":
		log.Fatalf("Windows is not support, please use either Linux of MacOS")
	default:
		configDirName := ".config"
		configDirPath := filepath.Join(homeDir, configDirName, folderName)
		configDir, err = os.ReadDir(configDirPath)
		if err != nil {
			fmt.Println("No config found, creating...")
		}
	}
	return configDir
}

func (c UserConfig) Create() []fs.DirEntry {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("failed to get user home dir: %v", err)
	}
	parentDir := filepath.Join(homeDir, ".config", "anki-for-me")
	err = os.Mkdir(parentDir, 0755)
	if err != nil {
		log.Fatalf("failed to create config dir: %v", err)
	}
	return c.GetDir()
}
