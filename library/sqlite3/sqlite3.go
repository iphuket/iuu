package sqlite3

import (
	"github.com/jinzhu/gorm"

	// sqlite 3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

// New sql
func New(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	defer db.Close()
	return db, err
}
