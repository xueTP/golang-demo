package main

import (
	"golang-demo/orm"
	"fmt"
	"github.com/Sirupsen/logrus"
)

func main() {
	// closeBag.FeibonaciNumsList()
	// errorDefer.HttpErrDemo()
	//maze.MazeBaseBreadthFirstDemo()
	// reptile.GetHtml()

	class := orm.Class{Name: "t1 class"}
	err := orm.DB.Create(&class).Error
	user := orm.User{Name: "sum", ClassId: int32(class.ClassId)}
	err = orm.DB.Create(&user).Error
	if err != nil {
		logrus.Errorf("err : %v", err)
	}
	otherUser := orm.User{}
	orm.DB.Where("userName = ?", "sum").Find(&otherUser)
	fmt.Printf("%+v", otherUser)
	orm.DB.Model(&otherUser).Related(&otherUser.Class)
	fmt.Printf("%+v", otherUser)
}
