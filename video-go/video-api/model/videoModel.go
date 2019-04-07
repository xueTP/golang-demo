package model

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type VideoModel struct{}

const VideoTableName = "vg_video"

type VideoModeler interface {
	AddVideo(video VideoTable) (int32, error)
	GetVideoList(params GetVideoParam) ([]VideoTable, error)
	GetVideoInfo(params GetVideoParam) (VideoTable, error)
	DeleteVideo(params GetVideoParam) (int32, error)
}

func (this VideoModel) AddVideo(video VideoTable) (int32, error) {
	stmtIn, err := DB.Prepare("INSERT into " + VideoTableName + "(vname, path, vdesc, createTime, updateTime, uid, isDelete) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmtIn.Close()
	res, err := stmtIn.Exec(video.VName, video.Path, video.VDesc, time.Now(), time.Now(), video.UID, video.IsDelete)
	if err != nil {
		return 0, err
	}
	vid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int32(vid), nil
}

type GetVideoParam struct {
	VideoTable
	Field  string
	Limit  int32
	Offset int32
	Order  string
}

func (this VideoModel) getCondition(params *GetVideoParam) (string, []interface{}) {
	var whereArr []string
	var value []interface{}
	if params.Field == "" {
		params.Field = "*"
	}
	if params.VID > 0 {
		whereArr = append(whereArr, "vid = ?")
		value = append(value, params.VID)
	}
	if params.UID > 0 {
		whereArr = append(whereArr, "uid = ?")
		value = append(value, params.UID)
	}
	if params.IsDelete == 1 {
		whereArr = append(whereArr, "isDelete = 1")
	}
	whereArr = append(whereArr, "1 = 1")
	where := strings.Join(whereArr, " AND ")
	return where, value
}

// getOtherSql 后续的limit、offset、order处理
func (this VideoModel) getOtherSql(params GetVideoParam) string {
	var otherWhere string
	if params.Order == "" {
		otherWhere += " ORDER BY vid desc"
	} else {
		otherWhere += " ORDER BY " + params.Order
	}
	if params.Limit > 0 {
		otherWhere += " LIMIT " + strconv.Itoa(int(params.Limit))
		if params.Offset >= 0 {
			otherWhere += " OFFSET " + strconv.Itoa(int(params.Offset))
		}
	}
	return otherWhere
}

func (this VideoModel) GetVideoList(params GetVideoParam) ([]VideoTable, error) {
	list := []VideoTable{}
	where, value := this.getCondition(&params)
	otherWhere := this.getOtherSql(params)
	stmtOut, err := DB.Prepare("SELECT " + params.Field + " FROM " + VideoTableName + " WHERE " + where + otherWhere)
	if err != nil {
		log.Printf("DB.Prepare error: %v", err)
		return list, err
	}
	rows, err := stmtOut.Query(value...)
	if err != nil {
		log.Printf("stmtOut.Query error: %v", err)
		return list, err
	}
	for rows.Next() {
		tempInfo := VideoTable{}
		var createTime, updTime []byte
		rows.Scan(&tempInfo.VID, &tempInfo.VName, &tempInfo.VDesc, &tempInfo.Path, &createTime, &updTime, &tempInfo.UID, &tempInfo.IsDelete)
		tempInfo.CreateTime, err = time.Parse("2006-01-02 15:04:05", string(createTime))
		if err != nil {
			log.Printf("time.Parse error: %vtime: %v", err, string(createTime))
			return list, err
		}
		tempInfo.UpdateTime, err = time.Parse("2006-01-02 15:04:05", string(updTime))
		if err != nil {
			log.Printf("time.Parse error: %v, time: %v", err, string(updTime))
			return list, err
		}
		list = append(list, tempInfo)
	}
	return list, nil
}

func (this VideoModel) GetVideoInfo(params GetVideoParam) (VideoTable, error) {
	info := VideoTable{}
	where, value := this.getCondition(&params)
	stemtOut, err := DB.Prepare("SELECT * FROM " + VideoTableName + " WHERE " + where + " LIMIT 1")
	if err != nil {
		return info, err
	}
	row := stemtOut.QueryRow(value...)
	var createTime, updTime []byte
	row.Scan(&info.VID, &info.VName, &info.VDesc, &info.Path, &createTime, &updTime, &info.UID, &info.IsDelete)
	info.CreateTime, err = time.Parse("2006-01-02 15:04:05", string(createTime))
	if err != nil {
		return info, err
	}
	info.UpdateTime, err = time.Parse("2006-01-02 15:04:05", string(updTime))
	if err != nil {
		return info, err
	}
	return info, nil
}

func (this VideoModel) DeleteVideo(params GetVideoParam) (int32, error) {
	where, value := this.getCondition(&params)
	if strings.Trim(where, " ") == "" {
		return 0, errors.New("where is Empty to Delete table " + VideoTableName)
	}
	stmtIns, err := DB.Prepare("DELETE FROM " + VideoTableName + " WHERE " + where)
	if err != nil {
		return 0, err
	}
	res, err := stmtIns.Exec(value...)
	if err != nil {
		return 0, err
	}
	rowsAffect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int32(rowsAffect), nil
}
