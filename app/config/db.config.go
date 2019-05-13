package config

import (
	"fmt"

	"github.com/iphuket/pkt/library/mysql"
	"github.com/iphuket/pkt/library/sqlite3"
	"github.com/jinzhu/gorm"
)

const (
	sqlite3PATH = "./iuu.db"
)

// DB 配置函数
func DB(drive string) (db *gorm.DB, err error) {
	switch drive {
	case "sqlite3":
		db, err := sqlite3.New(sqlite3PATH)
		return db, err
	case "mysql":
		db, err := mysql.New("iuu", "iuu", "iuu")
		return db, err
	}
	return nil, fmt.Errorf("not this drive %s", drive)
}
