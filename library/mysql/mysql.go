package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// New mysql
func New(user, password, dbname, addr string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+addr+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	// defer db.Close()
	return db, err
}
