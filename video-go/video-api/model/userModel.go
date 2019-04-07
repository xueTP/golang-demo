package model

import (
	"strings"
	"time"
)

type UserModel struct{}

const UserTableName = "vg_user"

type UserModeler interface {
	AddUser(user UserTable) (int32, error)
	GetUserInfo(params UserTable) (UserTable, error)
}

func (this UserModel) AddUser(user UserTable) (int32, error) {
	stmt, err := DB.Prepare("INSERT INTO " + UserTableName + "(`uname`,`pwd`,`createTime`) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.UName, user.PWD, time.Now())
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

func (this UserModel) getCondition(params UserTable) (string, []interface{}) {
	var whereArr []string
	var value []interface{}
	if params.UName != "" {
		whereArr = append(whereArr, "uname = ?")
		value = append(value, params.UName)
	}
	if params.UID > 0 {
		whereArr = append(whereArr, "uid = ?")
		value = append(value, params.UID)
	}
	return strings.Join(whereArr, " AND "), value
}

func (this UserModel) GetUserInfo(params UserTable) (UserTable, error) {
	info := UserTable{}
	where, value := this.getCondition(params)
	stmt, err := DB.Prepare("select * from " + UserTableName + " where " + where)
	if err != nil {
		return info, err
	}
	defer stmt.Close()
	infoRes := stmt.QueryRow(value...)
	var createTime []byte
	infoRes.Scan(&info.UID, &info.UName, &info.PWD, &createTime)
	info.CreateTime, err = time.Parse("2006-01-02 15:04:05", string(createTime))
	if err != nil {
		return info, err
	}
	return info, err
}
