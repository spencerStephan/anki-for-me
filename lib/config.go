package lib

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Conf struct {
	Dir  *ConfigDir
	File *ConfigFile
}

type ConfigDir struct {
	Path   string
	Dir    []fs.DirEntry
	Exists bool
}

type ConfigFile struct {
	File   fs.FileInfo
	Path   string
	Exists bool
}

var (
	homeDir               = getUserHomeDir()
	defaultConfigDirPath  = filepath.Join(homeDir, configDirName, folderName)
	defaultConfigFilePath = filepath.Join(homeDir, configDirName, folderName, "config.yaml")
)

const (
	configDirName = ".config"
	folderName    = "anki-for-me"
)

func getUserHomeDir() string {
	operatingSystem := runtime.GOOS
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home Dir: %v", err)
	}

	if operatingSystem == "windows" {
		log.Fatalln("Windows is not supported, please use either Linux of MacOS")
	}
	return homeDir
}

func (d *ConfigDir) DirExists() bool {
	return d.Exists
}

func (f *ConfigFile) FileExists() bool {
	return f.Exists
}

func (d *ConfigDir) GetDir(isInit bool) {
	configDirPath := filepath.Join(homeDir, configDirName, folderName)
	dir, err := os.ReadDir(configDirPath)
	if err != nil {
		if isInit {
			d.Exists = false
			return
		}
		log.Fatal("Config directory not found, please run anki-for-me init to create one.")
	}

	d.Exists = true
	d.Path = configDirPath
	d.Dir = dir
}

func CreateConfigDir(newDir *ConfigDir) *ConfigDir {
	newDir.GetDir(true)
	configDirPath := filepath.Join(homeDir, configDirName, folderName)

	err := os.Mkdir(configDirPath, 0755)
	if err != nil {
		log.Fatalf("failed to create config dir: %v", err)
	}
	dir, err := os.ReadDir(configDirPath)
	if err != nil {
		log.Fatalf("failed to read newly created config Dir: %v \n", err)
	}
	newDir.Dir = dir
	newDir.Exists = true
	newDir.Path = configDirPath
	return newDir
}

func (f *ConfigFile) GetFile(isInit bool, Dir *ConfigDir) {
	configFilePath := filepath.Join(Dir.Path, "config.yaml")
	file, err := os.Stat(configFilePath)
	if err != nil {
		if isInit {
			f.Exists = false
			return
		}
		log.Fatal("Config file not found, please run anki-for-me init to create one.")
	}
	f.Path = configFilePath
	f.Exists = true
	f.File = file
}

func CreateConfigFile(Dir *ConfigDir) *ConfigFile {
	fp := filepath.Join(Dir.Path, "config.yaml")
	file, err := os.Create(fp)
	if err != nil {
		log.Fatalf("failed to create config file: %v", err)
	}
	if err != nil {
		log.Fatalf("failed to create config file: %v", err)
	}
	newFile, _ := os.Stat(fp)

	return &ConfigFile{Path: file.Name(), Exists: true, File: newFile}
}

func OverrideConfigFile(overridePath string) {
	currentConf := InitConfig()
	_, err := os.Stat(overridePath)
	if err != nil {
		log.Fatalln("Invalid file, file must exist and be in YAML format.")
	}
	currentConfExists := CheckIfConfigFileExists(defaultConfigDirPath)

	if overridePath == defaultConfigFilePath {
		fmt.Println("Your selected configuration file is the default one, skipping...")
		return
	}

	if !currentConfExists {
		createFile(defaultConfigDirPath, overridePath)
		return
	}
	ArchivePreviousConfig(currentConf.Dir)
	ReplaceConfig(overridePath)
	fmt.Printf("Success! Set new default config file. Located at %s\n", defaultConfigFilePath)
}

func CheckIfConfigFileExists(overridePath string) bool {
	_, err := os.Stat(defaultConfigDirPath)
	return err == nil
}

func createFile(parentDir string, overridePath string) {
	destFile, err := os.Create(parentDir)
	if err != nil {
		log.Fatalln("failed to create config file")
	}
	sourceFile, err := os.Open(overridePath)
	if err != nil {
		log.Fatalln("failed to read override file")
	}
	defer sourceFile.Close()
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		log.Fatalln("failed to copy file")
	}
}

func ArchivePreviousConfig(dir *ConfigDir) {
	backupFilePath := filepath.Join(dir.Path, "config.bak.yaml")
	_, err := os.Stat(backupFilePath)
	if err == nil {
		err = os.Remove(backupFilePath)
		if err != nil {
			log.Fatal("there was an unexpected error removing the backup file")
		}
	}
	os.Rename(defaultConfigFilePath, filepath.Join(backupFilePath))
}

func ReplaceConfig(userInputtedPath string) {
	err := os.Rename(userInputtedPath, defaultConfigFilePath)
	if err != nil {
		log.Fatal("there was an unexpected problem replacing your config file.")
	}
}
