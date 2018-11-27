package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	core()
}

var DB *gorm.DB

func core() {
	// root/12345678
	db, err := gorm.Open("mysql", "root:12345678@/gormdemo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = db
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gorm_" + defaultTableName
	}
	// 开启调试
	DB.LogMode(true)
	// 删除表
	//DropTable()
	//创建及迁移表
	//CreateTable()
}
