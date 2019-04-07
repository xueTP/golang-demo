package model

import "time"

// DataBase: videogo; table: vg_user
type UserTable struct {
	UID        int32     `json:"uid"`
	UName      string    `json:"uname"`
	PWD        string    `json:"pwd"`
	CreateTime time.Time `json:"createTime"`
}

// DataBase: videogo; table: vg_video
type VideoTable struct {
	VID        int32     `json:"vid"`
	VName      string    `json:"vname"`
	Path       string    `json:"path"`
	VDesc      string    `json:"vdesc"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	UID        int32     `json:"uid"`
	IsDelete   int8      `json:"isDelete"`
}
