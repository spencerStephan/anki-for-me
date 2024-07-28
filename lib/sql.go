package lib

import (
	"io/fs"
	"strings"
)

type Database interface {
	Exists([]fs.DirEntry) bool
	Create() (*Sqlite, error)
	//Delete() error
	//Insert() error
	//Remove() error
}

type Sqlite struct{}

func (s Sqlite) Exists(userConfigDir []fs.DirEntry) bool {
	exists := false
	for i := 0; i < len(userConfigDir); i++ {
		if strings.Contains(userConfigDir[i].Name(), "sqlite") {
			exists = true
		}
	}
	return exists
}

func (s Sqlite) Create() (*Sqlite, error) {
	return &Sqlite{}, nil
}
