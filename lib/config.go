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

func SaveConfig(userInputtedPath string) {
	dir := GetDir(true)
	userPathAlreadyExists := checkIfUserInputtedConfigPathExists(userInputtedPath)
	if userPathAlreadyExists && strings.Contains(userInputtedPath, "/config.yaml") {
		return
	}

	targetConfigPath := filepath.Join(dir.Path, "config.yaml")
	configFileExists := checkIfConfigFileExists(targetConfigPath)

	if !configFileExists {
		moveAndRenameConfigFile(userInputtedPath, targetConfigPath)
		return
	}

	archivePreviousConfig(targetConfigPath, dir)
	replaceConfig(userInputtedPath, targetConfigPath)
}

func checkIfConfigFileExists(defaultConfigPath string) bool {
	_, err := os.Stat(defaultConfigPath)
	return err == nil
}

func checkIfUserInputtedConfigPathExists(userInputtedPath string) bool {
	_, err := os.Stat(userInputtedPath)
	return err == nil
}

func archivePreviousConfig(targetConfigPath string, dir Directory) {
	// check if backup file already exists
	backupFilePath := filepath.Join(dir.Path, "config.bak.yaml")
	_, err := os.Stat(backupFilePath)
	if err == nil {
		err = os.Remove(backupFilePath)
		if err != nil {
			log.Fatal("there was an unexpected error removing the backup file")
		}
	}
	os.Rename(targetConfigPath, filepath.Join(backupFilePath))
}

func replaceConfig(userInputtedPath string, targetConfigPath string) {
	err := os.Rename(userInputtedPath, targetConfigPath)
	if err != nil {
		log.Fatal("there was an unexpected probleming replacing your config file.")
	}
}

func moveAndRenameConfigFile(userInputtedPath string, targetConfigPath string) {
	os.Rename(userInputtedPath, targetConfigPath)
}

func NewConfig() Config {
	GetDir(false)
	return UserConfig
}
