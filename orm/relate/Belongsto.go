package relate

import (
	"fmt"
	"golang-demo/orm"
)

// `gorm:"ForeignKey:ClassRefer;AssociationForeignKey:ClassId"` // ClassRefer 为外键
// user 属于 class classRefer 是外键

type User struct {
	ID         int32  `gorm:"primary_key;AUTO_INCREMENT"`
	Name       string `gorm:"column:userName;type: varchar(20);not null"`
	Age        int8   `gorm:"default:0"`
	Class      Class  `gorm:"ForeignKey:ClassRefer"`
	ClassRefer int32
}

type Class struct {
	ID   int32  `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type: varchar(20);unique"`
}

func BelongsToDemo() {
	orm.DB.DropTableIfExists(&User{})
	orm.DB.DropTableIfExists(&Class{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Class{})

	//class
	class1 := Class{Name: "t1 class"}
	class2 := Class{Name: "t2 class"}
	orm.DB.Create(&class1)
	orm.DB.Create(&class2)
	// user
	user1 := User{Name: "sum", ClassRefer: class2.ID}
	user2 := User{Name: "bod", ClassRefer: class1.ID}
	orm.DB.Create(&user1)
	orm.DB.Create(&user2)

	// belongs to
	userInfo := User{Name: "sum"}
	orm.DB.Model(&userInfo).Where(userInfo).Find(&userInfo)
	fmt.Printf("User is : %+v", userInfo)
	//orm.DB.Model(&userInfo).Related(&userInfo.Class, "ClassRefer")
	orm.DB.Model(&userInfo).Association("Class").Find(&userInfo.Class)
	fmt.Printf("Class is : %+v", userInfo.Class)

}
