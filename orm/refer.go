package orm

import (
	"github.com/jinzhu/gorm"
)

//select * from INFORMATION_SCHEMA.KEY_COLUMN_USAGE  where REFERENCED_TABLE_NAME='t_stu'

func CreateTable() {
	if !DB.HasTable(&User{}) {
		DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	}else {
		DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&User{})
	}
	if !DB.HasTable(&Class{}) {
		DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Class{})
	}else {
		DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&Class{})
	}
}

type User struct {
	UserId int32 `gorm:"primary_key"`
	Name string `gorm:"column:userName;type: varchar(20);not null;"`
	Age int8 `gorm:"default 0"`
	Class Class // classId 为外键
	ClassId int32
}

type Class struct {
	gorm.Model
	Name string `gorm:"type: varchar(20);unique"`
}

func (Class) TableName() string {
	return "class"
}