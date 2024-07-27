package lib

import (
	"io/fs"
	"strings"
)

type DB struct{}

func CreateSqliteDB() *DB {
	return &DB{}
}

func CheckIfSqliteFileExists(userConfigDir []fs.DirEntry) bool {
	exists := false

	for i := 0; i < len(userConfigDir); i++ {
		if strings.Contains(userConfigDir[i].Name(), "sqlite") {
			exists = true
		}
	}
	return exists
}
