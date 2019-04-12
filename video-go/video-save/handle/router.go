package handle

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterRouter() *httprouter.Router {
	r := httprouter.New()
	videoServer := VideoServer{}
	r.POST("/upLoad", videoServer.Upload)
	r.GET("/downLoad/:path/:fileName", videoServer.DownLoad)
	r.GET("/uploadFile", videoServer.UploadFile)
	return r
}

type MiddleHandle struct {
	connLimit ConnLimit
	r         *httprouter.Router
}

func NewMiddleHandle(r *httprouter.Router) MiddleHandle {
	return MiddleHandle{
		r:         r,
		connLimit: NewConnLimit(),
	}
}

func (this MiddleHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取conn
	this.connLimit.GetConn()
	this.r.ServeHTTP(w, r)
	defer this.connLimit.GcConn()
}
