package handle

import (
	"github.com/julienschmidt/httprouter"
	"golang-demo/video-go/video-api/Logic"
	"golang-demo/video-go/video-api/config"
	"net/http"
)

func RegisterRouter() *httprouter.Router {
	r := httprouter.New()
	// user
	useHandle := UserHandle{}
	r.POST("/register", useHandle.RegisterUser)
	r.POST("/login", useHandle.Login)
	r.POST("/logout", useHandle.Logout)
	return r
}

type MiddleHandle struct {
	r *httprouter.Router
}

func NewMiddleHandle(r *httprouter.Router) MiddleHandle {
	return MiddleHandle{
		r: r,
	}
}

func (this MiddleHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 验证session
	code, msg := Logic.NewSessionLogic().AuthCheckSession(r)
	if code != config.Success {
		SendResponse(w, ResponseRes{ResponseCode: 201, Err: Err{ErrCode: code, ErrMsg: msg}})
		return
	}
	this.r.ServeHTTP(w, r)
}
