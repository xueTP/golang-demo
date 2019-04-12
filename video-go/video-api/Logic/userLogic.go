package Logic

import (
	"github.com/Sirupsen/logrus"
	"golang-demo/video-go/video-api/config"
	"golang-demo/video-go/video-api/model"
	"time"
	"util"
)

type UserLogic struct {
	userModel model.UserModeler
}

func NewUserLogic() UserLogic {
	return UserLogic{
		userModel: model.UserModel{},
	}
}

// RegisterUser 注册用户，主要只是添加用户信息
func (this UserLogic) RegisterUser(user model.UserTable) (int32, string) {
	info, err := this.userModel.GetUserInfo(model.UserTable{UName: user.UName})
	if err != nil || info.UID > 0 {
		logrus.Warnf("UserLogic.RegisterUser userModel.GetUserInfo error: %v, uname: %v", err, user.UName)
		return config.UNameIsExitErr, "注册账户已经存在"
	}
	user.CreateTime = time.Now()
	user.PWD = util.GetMD5(user.PWD)
	user.UID, err = this.userModel.AddUser(user)
	if err != nil {
		logrus.Errorf("UserLogic RegisterUser userModel.AddUser have error: %v, param: %+v", err, user)
		return config.AddUserErr, err.Error()
	}
	return config.Success, "用户注册成功"
}

// Login 登录
func (this UserLogic) Login(user model.UserTable) (int32, string, string) {
	var sessionId string
	info, err := this.userModel.GetUserInfo(model.UserTable{UName: user.UName})
	if err != nil || info.UID <= 0 {
		logrus.Errorf("UserLogic.Login userModel.GetUserInfo error: %v, uname: %v", err, user.UName)
		return config.UserInfoNotFound, "用户不存在", sessionId
	}
	if info.PWD != util.GetMD5(user.PWD) {
		logrus.Warnf("UserLogic.Login pwd is not true, uname: %v", user.UName)
		return config.UserPWDErr, "密码错误", sessionId
	}
	sessionId = util.GetUUid(user.UName)
	NewSessionLogic().SetSession(sessionId, info)
	return config.Success, "", sessionId
}

// Logout 用户登出
func (this UserLogic) Logout(sessionId string) (int32, string) {
	NewSessionLogic().DelSession(sessionId)
	return config.Success, ""
}
