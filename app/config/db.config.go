package config

import (
	"github.com/iphuket/iuu/library/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/iphuket/iuu/library/sqlite3"
)

const(
	sqlite3PATH = "./iuu.db"
) 

// DB 配置函数
func DB(drive string)(db *gorm.DB, err error){
	switch drive {
	case "sqlite3":
		db, err := sqlite3.New(sqlite3PATH)
		return db, err
	case "mysql":
		db, err := mysql.New("iuu","iuu","iuu")
		return db, err
	}
	return nil, fmt.Errorf("not this drive %s", drive)
}