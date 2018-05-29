package db

import (
	"../../config"
	"../common"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
 * DB接続
 * golang は自前でコネクションプールを管理しているため、プール設定は不要
 */
func DbInit() (*sql.DB, error) {
	var err error
	db, err = sql.Open(
		config.DBTYPE,
		""+config.USER+":"+config.PASSWORD+"@tcp("+config.HOSTNAME+":"+config.HOSTPORT+")/"+config.DBNAME+"?parseTime=true")
	db.SetMaxIdleConns(config.DBMAXIDOLECONNECTION)
	db.SetMaxOpenConns(config.DBMAXCONNECTION)
	return db, err
}

// データベースから切断する
func DbClose() {
	if db != nil {
		db.Close()
	}
}

// データベースハンドラを取得する
func DbConn() *sql.DB {
	if db == nil {
		_, err := DbInit()
		if err != nil {
			common.WriteErrorLog(config.DEBUG, err, nil)
			panic(err)
		}
	}
	return db
}
