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
	PARAM_ERR = ResponseRes{ResponseCode: http.StatusBadRequest, Err: Err{ErrCode: 10001, ErrMsg: "参数错误"}}
)

func SendResponse(w http.ResponseWriter, param ResponseRes) {
	w.WriteHeader(param.ResponseCode)
	resStr, _ := json.Marshal(param.Err)
	io.WriteString(w, string(resStr))
}
