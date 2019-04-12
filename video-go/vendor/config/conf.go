package config

var VideoConf videoConf

func init() {
	//初始化videoConf
	VideoConf = videoConf{
		ApiServerAddress:  ":9009",
		SaveServerAddress: ":9010",
		MaxUpLoadSize:     50 * 1024 * 1024, // 50M
		VideoSaveMaxConn:  2,
		AllowUploatFuffix: []string{
			"mp4", "avi",
		},
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
	ApiServerAddress  string
	SaveServerAddress string
	MaxUpLoadSize     int64
	AllowUploatFuffix []string
	VideoSaveMaxConn  int
	SessionLiveTime   int64
	SessionIdHeadKey  string
	NotAuthCheck      []string
	DBConfig          DBConf
	MD5Salt           string
}

type DBConf struct {
	DbHost       string
	DbUser       string
	DbPwd        string
	DataBaseName string
}
