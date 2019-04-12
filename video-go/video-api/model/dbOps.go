package model

import (
	videoGoConfig "config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	dbConf := videoGoConfig.VideoConf.DBConfig
	var err error
	DB, err = sql.Open("mysql", dbConf.DbUser+":"+dbConf.DbPwd+"@("+dbConf.DbHost+")/"+dbConf.DataBaseName)
	// DB, err = sql.Open("mysql", "root:root@(127.0.0.1:3306)" + dbConf.DataBaseName)
	if err != nil {
		log.Printf("sql.Open mysql have error: %v", err)
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Printf("DB.Ping have error: %v", err)
		panic(err)
	}
}

// TruncateTable 用于测试清理数据表的数据
func TruncateTable(tableName string) {
	DB.Exec("truncate table " + tableName)
}
