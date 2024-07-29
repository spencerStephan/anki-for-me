package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Directory struct {
	Path  string
	Files []fs.DirEntry
}

type Config struct {
	Directory Directory
}

var UserConfig Config

func GetDir(savingConfig bool) Directory {
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
			CreateDir(configDirPath)
		}

		if !savingConfig {
			_, err = os.Stat(filepath.Join(configDirPath, "config.yaml"))
			if err != nil {

				fmt.Println("No config file found, creating...")
				CreateFile(configDirPath)
			}
		}

		UserConfig.Directory = Directory{Files: configDir, Path: configDirPath}
	}
	return UserConfig.Directory
}

func CreateDir(parentDir string) {
	err := os.Mkdir(parentDir, 0755)
	if err != nil {
		log.Fatalf("failed to create config dir: %v", err)
	}
}

func CreateFile(parentDir string) {
	_, err := os.Create(filepath.Join(parentDir, "config.yaml"))
	if err != nil {
		log.Fatal("failed to create config file")
	}
}

func SaveConfig(path string) {
	dir := GetDir(true)
	for i := 0; i < len(dir.Files); i++ {
		if strings.Contains(dir.Files[i].Name(), "config.yaml") {
			oldConfig := filepath.Join(dir.Path, "old.config.yaml")
			_, err := os.Stat(oldConfig)
			if err == nil {
				os.Remove(oldConfig)
			}
			old := filepath.Join(dir.Path, "config.yaml")
			new := filepath.Join(dir.Path, "old.config.yaml")
			os.Rename(old, new)
			file, err := os.Stat(path)
			if err == nil && !file.IsDir() {
				os.Rename(path, old)
			} else {
				log.Fatal("cannot open config file, no file exists with this name")
			}
		}
	}
}

func NewConfig() Config {
	GetDir(false)
	return UserConfig
}
