package config

import (
	"fmt"

	"github.com/iphuket/pkt/library/mysql"
	"github.com/iphuket/pkt/library/sqlite3"
	"github.com/jinzhu/gorm"
)

// SQLDrive ...
var drive = "mysql"

// DB 配置函数
func DB() (db *gorm.DB, err error) {
	switch drive {
	case "sqlite3":
		db, err := sqlite3.New("./sqlite3.db")
		return db, err
	case "mysql":
		db, err := mysql.New("pkt", "pkt", "pkt", "127.0.0.1:3306")
		return db, err
	}
	return nil, fmt.Errorf("not this drive %s", drive)
}
