package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Sx180107@tcp(192.168.80.128:3306)/simpletest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	DBConn = d
}

func GetDB() *gorm.DB {
	return DBConn
}
