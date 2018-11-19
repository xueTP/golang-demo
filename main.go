package main

import (
	"golang-demo/orm"
	"github.com/Sirupsen/logrus"
)

func main() {
	// closeBag.FeibonaciNumsList()
	// errorDefer.HttpErrDemo()
	//maze.MazeBaseBreadthFirstDemo()
	// reptile.GetHtml()

	class := orm.Class{Name: "t1 class"}
	err := orm.DB.Create(&class).Error
	user := orm.User{Name: "sum", ClassId: int32(class.ID)}
	err = orm.DB.Create(&user).Error
	orm.DB.Model(&user).Related(&class)
	if err != nil {
		logrus.Errorf("err : %v", err)
	}
}
