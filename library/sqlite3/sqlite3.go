package sqlite3

import (
	"github.com/iphuket/iuu/app/config"
	"github.com/jinzhu/gorm"

	// sqlite 3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

// New sql
func New() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.Sqlite3DBFilePath)
	defer db.Close()
	return db, err
}
