package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Directory struct {
	Path  string
	Files []fs.DirEntry
}

type UserConfig struct {
	Directory *Directory
}

func (c *UserConfig) GetDir() {
	operatingSystem := runtime.GOOS
	var configDir []fs.DirEntry
	var configDirPath string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home dir: %v", err)
	}
	const folderName = "anki-for-me"

	switch operatingSystem {
	case "windows":
		log.Fatalf("Windows is not supported, please use either Linux of MacOS")
	default:
		configDirName := ".config"
		configDirPath = filepath.Join(homeDir, configDirName, folderName)
		configDir, err = os.ReadDir(configDirPath)
		if err != nil {
			fmt.Println("No config directory found, creating...")
			c.CreateDir(configDirPath)
		}
		_, err = os.Stat(filepath.Join(configDirPath, "config.yaml"))
		if err != nil {

			fmt.Println("No config file found, creating...")
			c.CreateFile(configDirPath)
		}
	}

	c.Directory = &Directory{Files: configDir, Path: configDirPath}
}

func (c *UserConfig) CreateDir(parentDir string) {
	err := os.Mkdir(parentDir, 0755)
	if err != nil {
		log.Fatalf("failed to create config dir: %v", err)
	}
}

func (c *UserConfig) CreateFile(parentDir string) {
	_, err := os.Create(filepath.Join(parentDir, "config.yaml"))
	if err != nil {
		log.Fatal("failed to create config file")
	}
}

func NewConfig() *UserConfig {
	config := &UserConfig{
		Directory: nil,
	}
	config.GetDir()
	return config
}
