package lib

import (
	"log"
)

type Services struct {
	DB     Database
	Config *Conf
}

func InitServices(config *Conf, db *Sqlite) *Services {
	return &Services{
		Config: config,
		DB:     db,
	}
}

func CreateConfig() *Conf {
	dir := &ConfigDir{}
	dir.GetDir(true)
	if !dir.DirExists() {
		dir = CreateConfigDir(dir)
	}
	file := &ConfigFile{}
	file.GetFile(true, dir)
	if !file.FileExists() {
		file = CreateConfigFile(dir)
	}
	sqliteFile, err := NewSqlite(dir)
	if err != nil {
		log.Fatal("There was an error creating the database file.")
	}
	if !sqliteFile.SqlExists(dir) {
		sqliteFile.Create(dir)
	}

	return &Conf{Dir: dir, File: file}
}

func InitConfig() *Conf {
	dir := &ConfigDir{}
	dir.GetDir(false)
	file := &ConfigFile{}
	file.GetFile(false, dir)
	return &Conf{Dir: dir, File: file}
}
