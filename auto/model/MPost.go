package automodel 
 
import ( 
	"../../config" 
	"../../service/common" 
	"../../service/db" 
	"../dto" 
	"net/http" 
	"strings" 
	"database/sql" 
) 
 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInMPost(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMPostDTO, bool) { 
	var rdto []autodbdto.DBMPostDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_post where " + targetColumn + " in (" 
	where := "" 
	for ind, _ := range in { 
		if where != "" { 
			where = where + "," 
		} 
		where = where + "?" 
		arr = append(arr, in[ind]) 
	} 
	sql = sql + where + ")" 
	if where == "" { 
		return rdto, false 
	} 
 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return rdto, false 
	} 
	defer stmt.Close() 
 
	dest := make([]interface{}, len(arr)) 
	for i, _ := range arr { 
		dest[i] = &arr[i] 
	} 
	rows, err := stmt.Query(dest...) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, strings.Join(arr, ","), r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return rdto, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMPostDTO 
 
		err := rows.Scan( 
			&columns.PostId, 
			&columns.OldPostCode, 
			&columns.PostCode, 
			&columns.PrefectureNameKana, 
			&columns.CityNameKana, 
			&columns.AddressNameKana, 
			&columns.PrefectureName, 
			&columns.CityName, 
			&columns.AddressName, 
			&columns.Flag1, 
			&columns.Flag2, 
			&columns.Flag3, 
			&columns.Flag4, 
			&columns.Flag5, 
			&columns.Flag6, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateDate, 
			&columns.UpdateDate, 
			&columns.DeleteFlag, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			return rdto, false 
		} 
		rdto = append(rdto, columns) 
	} 
 
	return rdto, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行しつつロックをかける 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInMPostForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMPostDTO, bool) { 
	var rdto []autodbdto.DBMPostDTO 
 
 
	var arr []string 
 
	sql := "select * from m_post where " + targetColumn + " in (" 
	where := "" 
	for ind, _ := range in { 
		if where != "" { 
			where = where + "," 
		} 
		where = where + "?" 
		arr = append(arr, in[ind]) 
	} 
	sql = sql + where + ")" 
	if where == "" { 
		return rdto, false 
	} 
	sql = sql + " for update " 
 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return rdto, false 
	} 
	defer stmt.Close() 
 
	dest := make([]interface{}, len(arr)) 
	for i, _ := range arr { 
		dest[i] = &arr[i] 
	} 
	rows, err := stmt.Query(dest...) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, strings.Join(arr, ","), r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return rdto, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMPostDTO 
 
		err := rows.Scan( 
			&columns.PostId, 
			&columns.OldPostCode, 
			&columns.PostCode, 
			&columns.PrefectureNameKana, 
			&columns.CityNameKana, 
			&columns.AddressNameKana, 
			&columns.PrefectureName, 
			&columns.CityName, 
			&columns.AddressName, 
			&columns.Flag1, 
			&columns.Flag2, 
			&columns.Flag3, 
			&columns.Flag4, 
			&columns.Flag5, 
			&columns.Flag6, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateDate, 
			&columns.UpdateDate, 
			&columns.DeleteFlag, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			return rdto, false 
		} 
		rdto = append(rdto, columns) 
	} 
 
	return rdto, true 
} 
 
/* 
 * Formデータに含まれる情報を元にSQLを発行する 
 * 空白の場合は検索対象としない 
 * 取得できなかった場合は、空の配列を返す 
 */ 
func SelectMPost(search *autodbdto.DBMPostForm, r *http.Request) ([]autodbdto.DBMPostDTO, bool) { 
	var rdto []autodbdto.DBMPostDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_post" 
	where := "" 
 
	where, arr = AppendWhere(search.PostId, "post_id", where, arr) 
	where, arr = AppendWhere(search.OldPostCode, "old_post_code", where, arr) 
	where, arr = AppendWhere(search.PostCode, "post_code", where, arr) 
	where, arr = AppendWhere(search.PrefectureNameKana, "prefecture_name_kana", where, arr) 
	where, arr = AppendWhere(search.CityNameKana, "city_name_kana", where, arr) 
	where, arr = AppendWhere(search.AddressNameKana, "address_name_kana", where, arr) 
	where, arr = AppendWhere(search.PrefectureName, "prefecture_name", where, arr) 
	where, arr = AppendWhere(search.CityName, "city_name", where, arr) 
	where, arr = AppendWhere(search.AddressName, "address_name", where, arr) 
	where, arr = AppendWhere(search.Flag1, "flag_1", where, arr) 
	where, arr = AppendWhere(search.Flag2, "flag_2", where, arr) 
	where, arr = AppendWhere(search.Flag3, "flag_3", where, arr) 
	where, arr = AppendWhere(search.Flag4, "flag_4", where, arr) 
	where, arr = AppendWhere(search.Flag5, "flag_5", where, arr) 
	where, arr = AppendWhere(search.Flag6, "flag_6", where, arr) 
	where, arr = AppendWhere(search.CreateUserId, "create_user_id", where, arr) 
	where, arr = AppendWhere(search.UpdateUserId, "update_user_id", where, arr) 
	where, arr = AppendWhere(search.CreateDate, "create_date", where, arr) 
	where, arr = AppendWhere(search.UpdateDate, "update_date", where, arr) 
	where, arr = AppendWhere(search.DeleteFlag, "delete_flag", where, arr) 
 
	if where != "" { 
		sql = sql + " where " + where 
 
	} 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return rdto, false 
	} 
	defer stmt.Close() 
 
	dest := make([]interface{}, len(arr)) 
	for i, _ := range arr { 
		dest[i] = &arr[i] 
	} 
	rows, err := stmt.Query(dest...) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, strings.Join(arr, ","), r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return rdto, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMPostDTO 
 
		err := rows.Scan( 
			&columns.PostId, 
			&columns.OldPostCode, 
			&columns.PostCode, 
			&columns.PrefectureNameKana, 
			&columns.CityNameKana, 
			&columns.AddressNameKana, 
			&columns.PrefectureName, 
			&columns.CityName, 
			&columns.AddressName, 
			&columns.Flag1, 
			&columns.Flag2, 
			&columns.Flag3, 
			&columns.Flag4, 
			&columns.Flag5, 
			&columns.Flag6, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateDate, 
			&columns.UpdateDate, 
			&columns.DeleteFlag, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			return rdto, false 
		} 
		rdto = append(rdto, columns) 
	} 
 
	return rdto, true 
} 
 
/* 
 * Formデータに含まれる情報を追加する 
 * 除外したいカラムはexcludesへ配列で格納する 
 * 成功時には追加したIDを返す 
 */ 
func InsertMPost(db *sql.Tx, search *autodbdto.DBMPostForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"post_id", 
		"old_post_code", 
		"post_code", 
		"prefecture_name_kana", 
		"city_name_kana", 
		"address_name_kana", 
		"prefecture_name", 
		"city_name", 
		"address_name", 
		"flag_1", 
		"flag_2", 
		"flag_3", 
		"flag_4", 
		"flag_5", 
		"flag_6", 
		"create_user_id", 
		"update_user_id", 
		"create_date", 
		"update_date", 
		"delete_flag", 
	} 
	values := []string{ 
		search.PostId, 
		search.OldPostCode, 
		search.PostCode, 
		search.PrefectureNameKana, 
		search.CityNameKana, 
		search.AddressNameKana, 
		search.PrefectureName, 
		search.CityName, 
		search.AddressName, 
		search.Flag1, 
		search.Flag2, 
		search.Flag3, 
		search.Flag4, 
		search.Flag5, 
		search.Flag6, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateDate, 
		search.UpdateDate, 
		search.DeleteFlag, 
	} 
	sql := "INSERT INTO m_post (" 
	for ind, _ := range columns { 
		if IsColumnExcludes(columns[ind], excludes) == false { 
			if columnSql != "" { 
				columnSql = columnSql + "," 
			} 
			columnSql = columnSql + "`" + columns[ind] + "`" 
		} 
	} 
	sql = sql + columnSql + ") VALUES (" 
	for ind, _ := range columns { 
		if IsColumnExcludes(columns[ind], excludes) == false { 
			if valueSql != "" { 
				valueSql = valueSql + "," 
			} 
			valueSql = valueSql + "?" 
			arr = append(arr, values[ind]) 
		} 
	} 
	sql = sql + valueSql + ")" 
 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, strings.Join(arr, ","), r) 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return 0, false 
	} 
	defer stmt.Close() 
 
	dest := make([]interface{}, len(arr)) 
	for i, _ := range arr { 
		dest[i] = &arr[i] 
	} 
	res, err := stmt.Exec(dest...) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return 0, false 
	} 
 
	id, _ := res.LastInsertId() 
	return id, true 
} 
 
/* 
 * バルクインサート 
 */ 
func BulkInsertMPost(db *sql.Tx, search []*autodbdto.DBMPostForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"post_id", 
		"old_post_code", 
		"post_code", 
		"prefecture_name_kana", 
		"city_name_kana", 
		"address_name_kana", 
		"prefecture_name", 
		"city_name", 
		"address_name", 
		"flag_1", 
		"flag_2", 
		"flag_3", 
		"flag_4", 
		"flag_5", 
		"flag_6", 
		"create_user_id", 
		"update_user_id", 
		"create_date", 
		"update_date", 
		"delete_flag", 
	} 
	sql := "INSERT INTO t_inquiry (" 
	for ind, _ := range columns { 
		if IsColumnExcludes(columns[ind], excludes) == false { 
			if columnSql != "" { 
				columnSql = columnSql + "," 
			} 
			columnSql = columnSql + "`" + columns[ind] + "`" 
		} 
	} 
	sql = sql + columnSql + ") VALUES " 
	for dataInd, _ := range search { 
		if dataInd != 0 { 
			valueSql = valueSql + "," 
		} 
		valueSql = valueSql + "(" 
		valueSqlSub := "" 
		for ind, _ := range columns { 
			if IsColumnExcludes(columns[ind], excludes) == false { 
				if valueSqlSub != "" { 
					valueSqlSub = valueSqlSub + "," 
				} 
				valueSqlSub = valueSqlSub + "?" 
				switch ind { 
				case 0: 
					arr = append(arr, search[dataInd].PostId) 
				case 1: 
					arr = append(arr, search[dataInd].OldPostCode) 
				case 2: 
					arr = append(arr, search[dataInd].PostCode) 
				case 3: 
					arr = append(arr, search[dataInd].PrefectureNameKana) 
				case 4: 
					arr = append(arr, search[dataInd].CityNameKana) 
				case 5: 
					arr = append(arr, search[dataInd].AddressNameKana) 
				case 6: 
					arr = append(arr, search[dataInd].PrefectureName) 
				case 7: 
					arr = append(arr, search[dataInd].CityName) 
				case 8: 
					arr = append(arr, search[dataInd].AddressName) 
				case 9: 
					arr = append(arr, search[dataInd].Flag1) 
				case 10: 
					arr = append(arr, search[dataInd].Flag2) 
				case 11: 
					arr = append(arr, search[dataInd].Flag3) 
				case 12: 
					arr = append(arr, search[dataInd].Flag4) 
				case 13: 
					arr = append(arr, search[dataInd].Flag5) 
				case 14: 
					arr = append(arr, search[dataInd].Flag6) 
				case 15: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 16: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 17: 
					arr = append(arr, search[dataInd].CreateDate) 
				case 18: 
					arr = append(arr, search[dataInd].UpdateDate) 
				case 19: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				} 
			} 
		} 
		valueSql = valueSql + valueSqlSub + ")" 
	} 
	sql = sql + valueSql + "" 
 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, strings.Join(arr, ","), r) 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return 0, false 
	} 
	defer stmt.Close() 
 
	dest := make([]interface{}, len(arr)) 
	for i, _ := range arr { 
		dest[i] = &arr[i] 
	} 
	res, err := stmt.Exec(dest...) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return 0, false 
	} 
 
	id, _ := res.LastInsertId() 
	return id, true 
} 
 
