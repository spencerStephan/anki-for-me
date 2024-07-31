package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	SqlExists(*ConfigDir) bool
	Create(*ConfigDir) error
	Connect(string) (*Sqlite, error)
	Close()
	// Delete() error
	// Insert() error
	// Remove() error
}

type Sqlite struct {
	Connection *sql.DB
}

func (s *Sqlite) SqlExists(d *ConfigDir) bool {
	filename := filepath.Join(d.Path, "collections.db")
	_, err := os.Stat(filename)
	return err == nil
}

func (s *Sqlite) Create(d *ConfigDir) error {
	_, err := os.Create(path.Join(d.Path, "collections.db"))
	return err
}

func (s *Sqlite) Connect(path string) (*Sqlite, error) {
	var err error

	s.Connection, err = sql.Open("sqlite3", filepath.Join("file:", path, "collections.db"))

	fmt.Println(filepath.Join("file:", path, "collections.db"))
	if err != nil {
		log.Fatal(err)
	}
	return s, nil
}

func (s *Sqlite) Close() {
	err := s.Connection.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func NewSqlite(d *ConfigDir) (*Sqlite, error) {
	sqlite := &Sqlite{
		nil,
	}
	exists := sqlite.SqlExists(d)

	if !exists {
		err := sqlite.Create(d)
		if err != nil {
			return nil, err
		}
	}
	return sqlite, nil
}
