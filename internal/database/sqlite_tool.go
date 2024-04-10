package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open("./db/ts.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
}

func GetDB() *gorm.DB {
	return db
}
