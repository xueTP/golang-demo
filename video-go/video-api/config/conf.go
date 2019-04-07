package config

var VideoConf videoConf

func init() {
	//初始化videoConf
	VideoConf = videoConf{
		ServerAddress:    ":9009",
		SessionLiveTime:  30 * 60,
		SessionIdHeadKey: "X-Video-Go",
		NotAuthCheck: []string{
			"register",
			"login",
		},
		DBConfig: DBConf{
			DbHost:       "127.0.0.1:3306",
			DbUser:       "root",
			DbPwd:        "root",
			DataBaseName: "videogo",
		},
		MD5Salt: "~!-video*go$xueheo@163.com$",
	}
}

type videoConf struct {
	ServerAddress    string
	SessionLiveTime  int64
	SessionIdHeadKey string
	NotAuthCheck     []string
	DBConfig         DBConf
	MD5Salt          string
}

type DBConf struct {
	DbHost       string
	DbUser       string
	DbPwd        string
	DataBaseName string
}
