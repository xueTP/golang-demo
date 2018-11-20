package orm

import (
	"database/sql"
	"time"
)

//select * from INFORMATION_SCHEMA.KEY_COLUMN_USAGE  where REFERENCED_TABLE_NAME='t_stu'

func CreateTable() {
	if !DB.HasTable(&User{}) {
		DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	} else {
		DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&User{})
	}
	if !DB.HasTable(&Class{}) {
		DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Class{})
	} else {
		DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&Class{})
	}
	// 设置相应外键
	DB.Model(&User{}).AddForeignKey("class_id", "class(`classId`)", "RESTRICT", "RESTRICT")
}

func DropTable() {
	DB.DropTableIfExists(&User{})
	DB.DropTableIfExists(&Class{})
}

type User struct {
	UserId  int32  `gorm:"primary_key"`
	Name    string `gorm:"column:userName;type: varchar(20);not null;"`
	Age     int8   `gorm:"default 0"`
	Class   Class  `gorm:"ForeignKey:ClassId;AssociationForeignKey:ClassId"` // classId 为外键
	ClassId int32
}

type Class struct {
	ClassId int32  `gorm:"column:classId;primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"type: varchar(20);unique"`
	Branch Branch
}

func (Class) TableName() string {
	return "class"
}

type Branch struct {
	BranchId   int32          `gorm:"column:branchId;primary_key;AUTO_INCREMENT;"`
	BranchName sql.NullString `gorm:"column:branchName;type:varchar(40)"`
	CreateTime time.Time      `gorm:"column:createTime"`
}
