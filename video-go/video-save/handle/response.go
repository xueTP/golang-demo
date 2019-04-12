package handle

import (
	"encoding/json"
	"io"
	"net/http"
)

type Err struct {
	ErrCode int32
	ErrMsg  string
}

type ResponseRes struct {
	Err
	ResponseCode int
}

var (
	PARAM_ERR      = ResponseRes{ResponseCode: http.StatusBadRequest, Err: Err{ErrCode: 10001, ErrMsg: "参数错误"}}
	FILEREAD_ERR   = ResponseRes{ResponseCode: http.StatusBadRequest, Err: Err{ErrCode: 10002, ErrMsg: "上传文件读错误"}}
	FILEWRITE_ERR  = ResponseRes{ResponseCode: http.StatusBadRequest, Err: Err{ErrCode: 10003, ErrMsg: "写文件错误"}}
	FILESIZE_ERR   = ResponseRes{ResponseCode: http.StatusBadRequest, Err: Err{ErrCode: 10004, ErrMsg: "文件太大"}}
	FILESUFFIX_ERR = ResponseRes{ResponseCode: http.StatusBadRequest, Err: Err{ErrCode: 10005, ErrMsg: "文件后缀不支持"}}
)

func SendResponse(w http.ResponseWriter, param ResponseRes) {
	w.WriteHeader(param.ResponseCode)
	resStr, _ := json.Marshal(param.Err)
	io.WriteString(w, string(resStr))
}
