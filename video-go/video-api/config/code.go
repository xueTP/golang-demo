package config

const (
	Success              = 1000 // 操作成功
	AddUserErr           = 2001 // 用户注册失败
	UserInfoNotFound     = 2002 // 用户信息不存在
	UserPWDErr           = 2003 // 用户密码错误
	UNameIsExitErr       = 2004 // 账号已经存在
	UserLogoutSessionErr = 2005 // 用户退出session错误
)
