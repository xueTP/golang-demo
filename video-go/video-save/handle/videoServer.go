package handle

import (
	"config"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"util"
)

type VideoServer struct{}

func (this VideoServer) UploadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("video-go/video/upload.html")
	t.Execute(w, nil)
}

func (this VideoServer) Upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, fileHandle, err := r.FormFile("upload")
	if err != nil {
		logrus.Errorf("http.Request.FormFile error: %v", err)
		SendResponse(w, PARAM_ERR)
		return
	}

	if fileHandle.Size >= config.VideoConf.MaxUpLoadSize {
		logrus.Errorf("upload file is to big, file size: %v", fileHandle.Size)
		SendResponse(w, FILESIZE_ERR)
		return
	}

	tld := strings.Split(fileHandle.Filename, ".")
	if util.InArrayString(strings.ToLower(tld[len(tld)-1]), config.VideoConf.AllowUploatFuffix) == -1 {
		logrus.Errorf("suffix is not allow to upload")
		SendResponse(w, FILESIZE_ERR)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Errorf("ioutil.ReadAll error: %v", err)
		SendResponse(w, FILEREAD_ERR)
		return
	}

	path := config.VideoConf.VideoTempDir + time.Now().Format("2006-01-02") + "/"
	err = os.Mkdir(path, 0666)
	if err != nil {
		logrus.Errorf("video-go.Upload os.MakeDir error: %v, path: %v", err, path)
		SendResponse(w, FILEWRITE_ERR)
		return
	}

	newFileName := util.GetMD5(util.GetUUid(fileHandle.Filename)) + "." + tld[len(tld)-1]
	err = ioutil.WriteFile(path+newFileName, data, 0666)
	if err != nil {
		logrus.Errorf("videoServer.Upload ioutil.WriteFile error: %v", err)
		SendResponse(w, FILEWRITE_ERR)
		return
	}
	SendResponse(w, ResponseRes{ResponseCode: http.StatusOK, Err: Err{ErrCode: 1000, ErrMsg: path}})
	return
}

func (this VideoServer) DownLoad(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := config.VideoConf.VideoRealDir + p.ByName("path") + "/" + p.ByName("fileName")
	logrus.Infof("return path :%v", path)
	fileHand, err := os.Open(path)
	if err != nil {
		logrus.Errorf("os.Open error: %v, path: %s", err, path)
		SendResponse(w, PARAM_ERR)
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), fileHand)
	defer fileHand.Close()
	//content, err := ioutil.ReadFile(path)
	//if err != nil {
	//	logrus.Errorf("ioutil.ReadFile error: %v, path: %v", content, path)
	//	SendResponse(w, PARAM_ERR)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "video/mp4")
	//w.Write(content)
	//return
}
