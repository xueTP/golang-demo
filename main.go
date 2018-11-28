package main

import (
	"golang-demo/orm"
	"fmt"
	"time"
)

func main() {
	//closeBag.FeibonaciNumsList()
	//errorDefer.HttpErrDemo()
	//maze.MazeBaseBreadthFirstDemo()
	//reptile.GetHtml()

	defer orm.DB.Close()
	//class
	//class1 := orm.Class{Name: "t1 class", ClassId: 3}
	//class2 := orm.Class{Name: "t2 class", ClassId: 4}
	//orm.DB.Create(&class1)
	//orm.DB.Create(&class2)
	//// user
	//user1 := orm.User{Name: "sum", ClassRefer: class2.ClassId}
	//user2 := orm.User{Name: "bod", ClassRefer: class1.ClassId}
	//orm.DB.Create(&user1)
	//orm.DB.Create(&user2)
	//// userDetail
	//orm.DB.Create(&orm.UserDetail{UserRefer: user2.ID, IdCard: "12345678"})
	//orm.DB.Create(&orm.UserDetail{UserRefer: user1.ID, IdCard: "32342342"})
	////email
	//orm.DB.Create(&orm.Email{UserTid: user1.ID, Email: "t1@email.com"})
	//orm.DB.Create(&orm.Email{UserTid: user1.ID, Email: "t2@email.com"})
	//orm.DB.Create(&orm.Email{UserTid: user2.ID, Email: "t3@email.com"})
	//orm.DB.Create(&orm.Email{UserTid: user2.ID, Email: "t4@email.com"})
	////language
	//language1 := orm.Language{Language: "english", Level: 1}
	//language2 := orm.Language{Language: "中文", Level: 1}
	//orm.DB.Create(&language1)
	//orm.DB.Create(&language2)

	//属于
	//user_1 := orm.User{}
	//class_1 := orm.Class{}
	//orm.DB.Where("userName = ?", "bod").Find(&user_1)
	//fmt.Printf("%+v, -- %d", user_1, user_1.ClassRefer)
	//err := orm.DB.Model(&user_1).Association("Class").Find(&class_1).Error
	//fmt.Printf("%+v \n %+v \n %v", user_1, class_1, err)
	//// 包含
	//user_2 := orm.User{}
	//userDetial_2 := orm.UserDetail{}
	//orm.DB.Where("userName = ?", "bod").Find(&user_2)
	//fmt.Printf("%+v", user_2)
	//orm.DB.Model(&user_2).Related(&userDetial_2, "UserRefer")
	//fmt.Printf("%+v", userDetial_2)
	//// 包含多个
	//user_3 := orm.User{}
	//email_3s := []orm.Email{}
	//orm.DB.Where("userName = ?", "bod").Find(&user_3)
	//fmt.Printf("%+v", user_3)
	//orm.DB.Model(&user_3).Related(&email_3s, "UserTid")
	//// orm.DB.Model(&user_3).Association("Emails").Find(&email_3s)
	//fmt.Printf("%+v", email_3s)
	//// 多对多
	//user_4 := orm.User{}
	//language_4 := []orm.Language{}
	//orm.DB.Model(&orm.User{}).Where("userName = ?", "sum").Find(&user_4)
	//fmt.Printf("%+v", user_4)
	//orm.DB.Model(&user_4).Association("Languages").Append(&language1)
	//orm.DB.Model(&user_4).Association("Languages").Append(&language2)
	//
	//orm.DB.Model(&user_4).Related(&language_4, "Languages")
	//// orm.DB.Model(&user_4).Association("Languages").Find(&language_4)
	//fmt.Printf("%+v", language_4)

	orm.DB.LogMode(true)
	// preload
	user_5 := orm.User{}
	orm.DB.Preload("Languages").Preload("UserDetail").Where("userName = ?", "sum").Find(&user_5)
	//orm.DB.Preload("Class").Preload("Email").Where("userName = ?", "sum").Find(&user_5) // Preload
	// DefaultTableNameHandler 与 model.TableName 同时使用导致class 不能使用
	// User 包含多个 Email 的外键设置没有使用默认类名 + ID 导致email 不能使用
	fmt.Printf("%+v", &user_5)

	time.Sleep(5 * time.Second)
	orm.DB.Model(&orm.User{}).Where("userName = ?", "sum").Update("age", 12)


	//manyToMany.ManyToManyDemo()
}
