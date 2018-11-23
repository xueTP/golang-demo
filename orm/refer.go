package orm

import "github.com/jinzhu/gorm"

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
	if !DB.HasTable(&UserDetail{}) {
		DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&UserDetail{})
	} else {
		DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&UserDetail{})
	}
	if !DB.HasTable(&Email{}) {
		DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Email{})
	} else {
		DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&Email{})
	}
	// 设置相应外键
	//DB.Model(&User{}).AddForeignKey("class_id", "class(`class_id`)", "RESTRICT", "RESTRICT")
}

func DropTable() {
	DB.DropTableIfExists(&User{})
	DB.DropTableIfExists(&Class{})
	DB.DropTableIfExists(&UserDetail{})
	DB.DropTableIfExists(&Email{})
}

type User struct {
	UserId     int32      `gorm:"primary_key"`
	Name       string     `gorm:"column:userName;type: varchar(20);not null"`
	Age        int8       `gorm:"default:0"`
	Class      Class      `gorm:"ForeignKey:ClassRefer;AssociationForeignKey:ClassId"` // ClassRefer 为外键
	ClassRefer int32      // user 属于 class classRefer 是外键
	UserDetail UserDetail // 包含
	Emails     []Email    // 包含多个
}

type UserDetail struct {
	UDId      int32  `gorm:primary_key`
	UserRefer int32  // userDetail 属于 user userRefer 是外键
	IdCard    string `gorm:"column:idCard;type:varchar(11)"`
}

type Class struct {
	gorm.Model
	ClassId int32
	Name    string `gorm:"type: varchar(20);unique"`
}

func (Class) TableName() string {
	return "class"
}

type Email struct {
	Id     int32 `gorm:"primary_key"`
	UserId int32
	Email  string `gorm:"type:varchar(20)"`
}