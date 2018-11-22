package main

import (
	"fmt"
	"golang-demo/orm"
)

func main() {
	//closeBag.FeibonaciNumsList()
	//errorDefer.HttpErrDemo()
	//maze.MazeBaseBreadthFirstDemo()
	//reptile.GetHtml()

	defer orm.DB.Close()
	//class
	class1 := orm.Class{Name: "t1 class", ClassId: 3}
	class2 := orm.Class{Name: "t2 class", ClassId: 4}
	orm.DB.Create(&class1)
	orm.DB.Create(&class2)
	// user
	user1 := orm.User{Name: "sum", ClassRefer: class2.ClassId}
	user2 := orm.User{Name: "bod", ClassRefer: class1.ClassId}
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
	fmt.Printf("%+v, -- %d", user_1, user_1.ClassRefer)
	//err := orm.DB.Model(&user_1).Related(&class_1, "ClassRefer").Error
	err := orm.DB.Model(&user_1).Association("Class").Find(&class_1).Error
	fmt.Printf("%+v \n %+v \n %v", user_1, class_1, err)
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
