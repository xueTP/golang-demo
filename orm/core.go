package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func core() {
	db, err := gorm.Open("mysql", "root:12345678@/gormDemo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = db

}
