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
	//class := orm.Class{Name: "t1 class"}
	//err := orm.DB.Create(&class).Error
	//user := orm.User{Name: "sum", ClassId: int32(class.ClassId)}
	//err = orm.DB.Create(&user).Error
	//user2 := orm.User{Name: "bod", ClassId: class.ClassId}
	//orm.DB.Create(&user2)
	//orm.DB.Create(&orm.UserDetail{UserId: 2, IdCard: "12345678"})
	//orm.DB.Create(&orm.UserDetail{UserId: user.UserId, IdCard: "32342342"})
	////email
	//orm.DB.Create(&orm.Email{UserId: user.UserId, Email: "t1@email.com"})
	//orm.DB.Create(&orm.Email{UserId: user.UserId, Email: "t2@email.com"})
	//orm.DB.Create(&orm.Email{UserId: user2.UserId, Email: "t3@email.com"})
	//orm.DB.Create(&orm.Email{UserId: user2.UserId, Email: "t4@email.com"})
	//if err != nil {
	//	logrus.Errorf("err : %v", err)
	//}
	// 属于
	//user_1 := orm.User{}
	//orm.DB.Where("userName = ?", "bod").Find(&user_1)
	//fmt.Printf("%+v", user_1)
	//orm.DB.Model(&user_1).Related(&user_1.Class)
	//fmt.Printf("%+v", user_1)
	// 包含
	//user_2 := orm.User{}
	//orm.DB.Where("userName = ?", "bod").Find(&user_2)
	//orm.DB.Model(&user_2).Related(&user_2.UserDetail)
	//fmt.Printf("%+v", user_2)
	// 包含多个
	user_3 := orm.User{}
	orm.DB.Where("userName = ?", "bod").Find(&user_3)
	orm.DB.Model(&user_3).Related(&user_3.Emails, "Emails")
	fmt.Printf("%+v", user_3.Emails)
}
