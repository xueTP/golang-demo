package hasOne

import (
	"fmt"
	"golang-demo/orm"
)

type User struct {
	ID         int32  `gorm:"primary_key"`
	Name       string `gorm:"column:userName;type: varchar(20);not null"`
	Age        int8   `gorm:"default:0"`
	Refer      int32
	UserDetail UserDetail `gorm:"ForeignKey:UserRefer;Association_ForeignKey:Refer"` // 包含
}

type UserDetail struct {
	UDId      int32  `gorm:"primary_key"`
	UserRefer int32  // userDetail 属于 user UserID 是外键
	IdCard    string `gorm:"column:idCard;type:varchar(11)"`
}

func HasOneDemo() {
	orm.DB.DropTableIfExists(&User{})
	orm.DB.DropTableIfExists(&UserDetail{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&UserDetail{})

	// user1
	user1 := User{Name: "sum"}
	user2 := User{Name: "bod"}
	orm.DB.Create(&user1)
	orm.DB.Create(&user2)

	// userDetail
	orm.DB.Create(&UserDetail{UserRefer: user2.ID, IdCard: "12345678"})
	orm.DB.Create(&UserDetail{UserRefer: user1.ID, IdCard: "87654321"})

	// 包含
	userInfo := User{}
	orm.DB.Where("userName = ?", "bod").Find(&userInfo)
	fmt.Printf("User is : %+v", userInfo)
	// orm.DB.Model(&userInfo).Related(&userInfo.UserDetail, "UserRefer")
	orm.DB.Model(&userInfo).Association("UserDetail").Find(&userInfo.UserDetail)
	fmt.Printf("%+v", userInfo.UserDetail)
}
