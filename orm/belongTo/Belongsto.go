package belongTo

import (
	"fmt"
	"golang-demo/orm"
)

// `gorm:

type User struct {
	ID         int32  `gorm:"primary_key;AUTO_INCREMENT"`
	Name       string `gorm:"column:userName;type: varchar(20);not null"`
	Age        int8   `gorm:"default:0"`
	Class      Class
	ClassRefer int32
}

type Class struct {
	ClassId int32 `gorm:"primary_key;AUTO_INCREMENT"`
	Refer   int32
	Name    string `gorm:"type: varchar(20);unique"`
}

func BelongsToDemo() {
	orm.DB.DropTableIfExists(&User{})
	orm.DB.DropTableIfExists(&Class{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Class{})

	//class
	class1 := Class{Name: "t1 class", Refer: 1}
	class2 := Class{Name: "t2 class", Refer: 2}
	orm.DB.Create(&class1)
	orm.DB.Create(&class2)
	// user
	user1 := User{Name: "sum", ClassRefer: class2.ClassId}
	user2 := User{Name: "bod", ClassRefer: class1.ClassId}
	orm.DB.Create(&user1)
	orm.DB.Create(&user2)

	// belongs to
	userInfo := User{Name: "sum"}
	orm.DB.Model(&userInfo).Where(userInfo).Find(&userInfo)
	fmt.Printf("User is : %+v", userInfo)
	orm.DB.Model(&userInfo).Association("Class").Find(&userInfo.Class)
	// orm.DB.Model(&userInfo).Related(&userInfo.Class, "ClassRefer")
	fmt.Printf("Class is : %+v", userInfo.Class)

	// orm.DB.Model(&userInfo).Association("Class").Find(&userInfo.Class)
}
