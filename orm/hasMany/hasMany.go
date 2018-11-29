package hasMany

import (
	"golang-demo/orm"
	"fmt"
)

// `gorm:"ForeignKey:UserID"`

type User struct {
	ID         int32  `gorm:"primary_key"`
	Name       string `gorm:"column:userName;type: varchar(20);not null"`
	Age        int8   `gorm:"default:0"`
	Refer      int32
	Email      []Email  `gorm:"ForeignKey:UserID"` // 包含多个
}

type Email struct {
	ID      int32 `gorm:"primary_key"`
	UserID 	int32
	Email   string `gorm:"type:varchar(20)"`
}

func HasManyDemo() {
	orm.DB.DropTableIfExists(&User{})
	orm.DB.DropTableIfExists(&Email{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Email{})

	// user
	user1 := User{Name: "sum"}
	user2 := User{Name: "bod"}
	orm.DB.Create(&user1)
	orm.DB.Create(&user2)

	// Email
	orm.DB.Create(&Email{UserID: user1.ID, Email: "t1@email.com"})
	orm.DB.Create(&Email{UserID: user1.ID, Email: "t2@email.com"})
	orm.DB.Create(&Email{UserID: user2.ID, Email: "t3@email.com"})
	orm.DB.Create(&Email{UserID: user2.ID, Email: "t4@email.com"})

	// 包含多个
	userInfo := User{}
	orm.DB.Model(&User{}).Where("userName = ?", "bod").Find(&userInfo)
	fmt.Printf("%+v", userInfo)
	orm.DB.Model(&userInfo).Related(&userInfo.Email)
	fmt.Printf("%+v", userInfo.Email)

	// orm.DB.Model(&user_3).Association("Emails").Find(&email_3s)
	//
}