package Logic

import (
	videoGoConfig "config"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang-demo/video-go/video-api/config"
	"golang-demo/video-go/video-api/model"
	"net/http"
	"strings"
	"sync"
	"time"
	"util"
)

type SessionLogic struct {
	AliveTime int64
}

func NewSessionLogic() SessionLogic {
	return SessionLogic{
		AliveTime: videoGoConfig.VideoConf.SessionLiveTime,
	}
}

// 使用线程安全的map将session数据存储在内存里
var session sync.Map

// sessionData session 存储的主要数据结构有创建时间与用户数据组成
type sessionData struct {
	CreateTime int64           `json:"createTime"`
	User       model.UserTable `json:"user"`
}

// SetSession 设置session数据到 session 中
func (this SessionLogic) SetSession(key string, user model.UserTable) {
	data := sessionData{User: user, CreateTime: time.Now().Unix()}
	jsonData, _ := json.Marshal(data)
	session.Store(key, jsonData)
}

// GetSession 获取 session 里对于key的值
func (this SessionLogic) GetSession(key string) (model.UserTable, error) {
	var data sessionData
	jsonData, ok := session.Load(key)
	if !ok {
		return data.User, errors.New("SessionLogic.GetSession this key is empty on session")
	}
	d, ok := jsonData.([]byte)
	if !ok {
		return data.User, errors.New("SessionLogic.GetSession type string is not ok")
	}
	json.Unmarshal(d, &data)
	if this.isOutTimeSession(data, key) {
		return model.UserTable{}, errors.New("SessionLogic.GetSession this key is out tiome")
	}
	return data.User, nil
}

// DelSession 删除session
func (this SessionLogic) DelSession(key string) {
	session.Delete(key)
}

// isOutTimeSession session是否已经过期
func (this SessionLogic) isOutTimeSession(data sessionData, key string) bool {
	if time.Now().Unix()-data.CreateTime > this.AliveTime {
		session.Load(key)
		return true
	}
	return false
}

// AuthCheckSession session 权限验证
func (this SessionLogic) AuthCheckSession(r *http.Request) (int32, string) {
	// 例外判断
	path := strings.Split(r.RequestURI, "/")
	sessionId := r.Header.Get(videoGoConfig.VideoConf.SessionIdHeadKey)
	if len(path) <= 1 || util.InArrayString(path[1], videoGoConfig.VideoConf.NotAuthCheck) == -1 {
		// 是否有效判断
		_, err := this.GetSession(sessionId)
		if err != nil {
			logrus.Errorf("UserLogic.Logout sessionLogic.GetSession error: %v, sessionId: %v", err, sessionId)
			return config.UserLogoutSessionErr, "session 获取失败或者已经失效"
		}
	}
	return config.Success, ""
}
