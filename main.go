package main

import (
	"golang-demo/orm"
	"fmt"
)

func main() {
	//closeBag.FeibonaciNumsList()
	//errorDefer.HttpErrDemo()
	//maze.MazeBaseBreadthFirstDemo()
	//reptile.GetHtml()

	defer orm.DB.Close()
	// class
	class1 := orm.Class{Name: "t1 class"}
	class2 := orm.Class{Name: "t2 class"}
	orm.DB.Create(&class1)
	orm.DB.Create(&class2)
	// user
	user1 := orm.User{Name: "sum", ClassRefer: class1.ClassId}
	user2 := orm.User{Name: "bod", ClassRefer: class2.ClassId}
	orm.DB.Create(&user1)
	orm.DB.Create(&user2)
	// userDetail
	orm.DB.Create(&orm.UserDetail{UserId: 2, IdCard: "12345678"})
	orm.DB.Create(&orm.UserDetail{UserId: user1.UserId, IdCard: "32342342"})
	//email
	orm.DB.Create(&orm.Email{UserId: user1.UserId, Email: "t1@email.com"})
	orm.DB.Create(&orm.Email{UserId: user1.UserId, Email: "t2@email.com"})
	orm.DB.Create(&orm.Email{UserId: user2.UserId, Email: "t3@email.com"})
	orm.DB.Create(&orm.Email{UserId: user2.UserId, Email: "t4@email.com"})

	// 属于
	user_1 := orm.User{}
	class_1 := orm.Class{}
	orm.DB.Where("userName = ?", "bod").Find(&user_1)
	fmt.Printf("%+v, -- %d", user_1, user1.ClassRefer)
	orm.DB.Model(&user_1).Related(&class_1)
	fmt.Printf("%+v \n %+v", user_1, class_1)
	// 包含
	// user_2 := orm.User{}
	// orm.DB.Where("userName = ?", "bod").Find(&user_2)
	// orm.DB.Model(&user_2).Related(&user_2.UserDetail)
	// fmt.Printf("%+v", user_2)
	// 包含多个
	// user_3 := orm.User{}
	// orm.DB.Where("userName = ?", "bod").Find(&user_3)
	// orm.DB.Model(&user_3).Related(&user_3.Emails, "Emails")
	// fmt.Printf("%+v", user_3.Emails)
}
