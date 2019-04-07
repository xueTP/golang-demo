package handle

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"golang-demo/video-go/video-api/Logic"
	"golang-demo/video-go/video-api/config"
	"golang-demo/video-go/video-api/model"
	"io/ioutil"
	"net/http"
)

type UserHandle struct{}

// RegisterUser 注册用户
func (this UserHandle) RegisterUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user model.UserTable
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if user.UName == "" || user.PWD == "" || err != nil {
		logrus.Errorf("UserHandle.CreateUser params is empty")
		SendResponse(w, PARAM_ERR)
		return
	}
	code, msg := Logic.NewUserLogic().RegisterUser(user)
	if code != config.Success {
		SendResponse(w, ResponseRes{ResponseCode: http.StatusInternalServerError, Err: Err{ErrCode: code, ErrMsg: msg}})
		return
	}
	SendResponse(w, ResponseRes{ResponseCode: http.StatusOK, Err: Err{ErrCode: code, ErrMsg: msg}})
}

// Logic 登录
func (this UserHandle) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user model.UserTable
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if user.UName == "" || user.PWD == "" || err != nil {
		logrus.Errorf("UserHandle.CreateUser params is empty")
		SendResponse(w, PARAM_ERR)
		return
	}
	code, msg, sessionId := Logic.NewUserLogic().Login(user)
	w.Header().Set(config.VideoConf.SessionIdHeadKey, sessionId)
	SendResponse(w, ResponseRes{ResponseCode: 200, Err: Err{ErrCode: code, ErrMsg: msg}})
	return
}

// Logout 登出
func (this UserHandle) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sessionId := r.Header.Get(config.VideoConf.SessionIdHeadKey)
	Logic.NewSessionLogic().DelSession(sessionId)
	SendResponse(w, ResponseRes{ResponseCode: 200, Err: Err{ErrCode: config.Success}})
}
