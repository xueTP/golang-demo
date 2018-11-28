package polymorphism

import (
	"fmt"
	"golang-demo/orm"
)

type User struct {
	ID   int32  `gorm:"primary_key"`
	Name string `gorm:"column:userName;type: varchar(20);not null"`
	Age  int8   `gorm:"default:0"`
	Role Role   `gorm:"polymorphic:Owner;"`
}

type Admin struct {
	ID   int32  `gorm:"primary_key"`
	Name string `gorm:"column:adminName;type: varchar(20);not null"`
	Age  int8   `gorm:"default:0"`
	Role []Role `gorm:"polymorphic:Owner;"`
}

type Role struct {
	ID        int32  `gorm:"primary_key"`
	RoleName  string `gorn:"column:roleName;size:20"`
	OwnerId   int32
	OwnerType string
}

func PolymorphismDemo() {
	orm.DB.DropTableIfExists(&User{})
	orm.DB.DropTableIfExists(&Admin{})
	orm.DB.DropTableIfExists(&Role{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Admin{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Role{})

	// user
	user1 := User{Name: "sum"}
	orm.DB.Create(&user1)

	// admin
	admin1 := Admin{Name: "tom"}
	orm.DB.Create(&admin1)

	// role
	role1 := Role{RoleName: "查看"}
	role2 := Role{RoleName: "管理修改"}
	role3 := Role{RoleName: "管理查看"}
	orm.DB.Model(&user1).Association("Role").Append(role1)
	orm.DB.Model(&admin1).Association("Role").Append(role2)
	orm.DB.Model(&admin1).Association("Role").Append(role3)

	// 包含多态
	userInfo := User{}
	orm.DB.Where("userName = ?", "sum").Find(&userInfo)
	fmt.Printf("User is : %+v", userInfo)
	orm.DB.Model(&userInfo).Association("Role").Find(&userInfo.Role)
	fmt.Printf("Role is : %+v", userInfo.Role)

	adminInfo := Admin{}
	orm.DB.Where("adminName = ?", "tom").Find(&adminInfo)
	fmt.Printf("Admin is : %+v", userInfo)
	orm.DB.Model(&adminInfo).Association("Role").Find(&adminInfo.Role)
	fmt.Printf("Role is : %+v", adminInfo.Role)
}
