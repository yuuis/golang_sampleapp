package model

import (
	"../../config"
	"../common"
	"../db"
	"../dto/db"
	"net/http"
)

func TestSelectSQL(r *http.Request) {
	db := db.DbConn()
	sql := "select * from t_test where sid = ? and delete_flag = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(0, 0)

	common.WriteLog(config.DEBUG, sql, r)
	common.WriteLog(config.DEBUG, "0, 0", r)

	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var columns dtodb.DbTtest //取得したいカラム名と揃える必要がある

		err := rows.Scan(
			&columns.Sid,
			&columns.Number,
			&columns.Input_date,
			&columns.Float,
			&columns.Delete_flag,
			&columns.Create_user_id,
			&columns.Update_user_id,
			&columns.Create_time,
			&columns.Update_time)

		if err != nil {
			common.WriteErrorLog(config.FATAL, err, r)
			panic(err)
		}
		//		common.WriteLog( fmt.Printf("id:%s\\tname:%s", columns.Sid, columns.Input_date)
		common.WriteLog(config.DEBUG, columns.Input_date, r)
	}
}
