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
 * PKからデータを取得する 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKTSettingConnection(m_setting_sid string, r *http.Request) (autodbdto.DBTSettingConnectionDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from t_setting_connection where m_setting_sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTSettingConnectionDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(m_setting_sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, m_setting_sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTSettingConnectionDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBTSettingConnectionDTO 
 
		err := rows.Scan( 
			&columns.MSettingSid, 
			&columns.LastComparedTime, 
			&columns.LastXdShopTime, 
			&columns.LastXdCartTime, 
			&columns.LastConvertionTime, 
			&columns.LastComparedTimeSp, 
			&columns.LastXdShopTimeSp, 
			&columns.LastXdCartTimeSp, 
			&columns.LastConvertionTimeSp, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBTSettingConnectionDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBTSettingConnectionDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKTSettingConnectionForUpdate(db *sql.Tx, m_setting_sid string, r *http.Request) (autodbdto.DBTSettingConnectionDTO, bool, bool) { 
 
	sql := "select * from t_setting_connection where m_setting_sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTSettingConnectionDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(m_setting_sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, m_setting_sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTSettingConnectionDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBTSettingConnectionDTO 
 
		err := rows.Scan( 
			&columns.MSettingSid, 
			&columns.LastComparedTime, 
			&columns.LastXdShopTime, 
			&columns.LastXdCartTime, 
			&columns.LastConvertionTime, 
			&columns.LastComparedTimeSp, 
			&columns.LastXdShopTimeSp, 
			&columns.LastXdCartTimeSp, 
			&columns.LastConvertionTimeSp, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBTSettingConnectionDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBTSettingConnectionDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInTSettingConnection(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBTSettingConnectionDTO, bool) { 
	var rdto []autodbdto.DBTSettingConnectionDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from t_setting_connection where " + targetColumn + " in (" 
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
		var columns autodbdto.DBTSettingConnectionDTO 
 
		err := rows.Scan( 
			&columns.MSettingSid, 
			&columns.LastComparedTime, 
			&columns.LastXdShopTime, 
			&columns.LastXdCartTime, 
			&columns.LastConvertionTime, 
			&columns.LastComparedTimeSp, 
			&columns.LastXdShopTimeSp, 
			&columns.LastXdCartTimeSp, 
			&columns.LastConvertionTimeSp, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
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
func SelectByInTSettingConnectionForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBTSettingConnectionDTO, bool) { 
	var rdto []autodbdto.DBTSettingConnectionDTO 
 
 
	var arr []string 
 
	sql := "select * from t_setting_connection where " + targetColumn + " in (" 
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
		var columns autodbdto.DBTSettingConnectionDTO 
 
		err := rows.Scan( 
			&columns.MSettingSid, 
			&columns.LastComparedTime, 
			&columns.LastXdShopTime, 
			&columns.LastXdCartTime, 
			&columns.LastConvertionTime, 
			&columns.LastComparedTimeSp, 
			&columns.LastXdShopTimeSp, 
			&columns.LastXdCartTimeSp, 
			&columns.LastConvertionTimeSp, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
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
func SelectTSettingConnection(search *autodbdto.DBTSettingConnectionForm, r *http.Request) ([]autodbdto.DBTSettingConnectionDTO, bool) { 
	var rdto []autodbdto.DBTSettingConnectionDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from t_setting_connection" 
	where := "" 
 
	where, arr = AppendWhere(search.MSettingSid, "m_setting_sid", where, arr) 
	where, arr = AppendWhere(search.LastComparedTime, "last_compared_time", where, arr) 
	where, arr = AppendWhere(search.LastXdShopTime, "last_xd_shop_time", where, arr) 
	where, arr = AppendWhere(search.LastXdCartTime, "last_xd_cart_time", where, arr) 
	where, arr = AppendWhere(search.LastConvertionTime, "last_convertion_time", where, arr) 
	where, arr = AppendWhere(search.LastComparedTimeSp, "last_compared_time_sp", where, arr) 
	where, arr = AppendWhere(search.LastXdShopTimeSp, "last_xd_shop_time_sp", where, arr) 
	where, arr = AppendWhere(search.LastXdCartTimeSp, "last_xd_cart_time_sp", where, arr) 
	where, arr = AppendWhere(search.LastConvertionTimeSp, "last_convertion_time_sp", where, arr) 
	where, arr = AppendWhere(search.DeleteFlag, "delete_flag", where, arr) 
	where, arr = AppendWhere(search.CreateUserId, "create_user_id", where, arr) 
	where, arr = AppendWhere(search.UpdateUserId, "update_user_id", where, arr) 
	where, arr = AppendWhere(search.CreateTime, "create_time", where, arr) 
	where, arr = AppendWhere(search.UpdateTime, "update_time", where, arr) 
 
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
		var columns autodbdto.DBTSettingConnectionDTO 
 
		err := rows.Scan( 
			&columns.MSettingSid, 
			&columns.LastComparedTime, 
			&columns.LastXdShopTime, 
			&columns.LastXdCartTime, 
			&columns.LastConvertionTime, 
			&columns.LastComparedTimeSp, 
			&columns.LastXdShopTimeSp, 
			&columns.LastXdCartTimeSp, 
			&columns.LastConvertionTimeSp, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
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
func InsertTSettingConnection(db *sql.Tx, search *autodbdto.DBTSettingConnectionForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"m_setting_sid", 
		"last_compared_time", 
		"last_xd_shop_time", 
		"last_xd_cart_time", 
		"last_convertion_time", 
		"last_compared_time_sp", 
		"last_xd_shop_time_sp", 
		"last_xd_cart_time_sp", 
		"last_convertion_time_sp", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.MSettingSid, 
		search.LastComparedTime, 
		search.LastXdShopTime, 
		search.LastXdCartTime, 
		search.LastConvertionTime, 
		search.LastComparedTimeSp, 
		search.LastXdShopTimeSp, 
		search.LastXdCartTimeSp, 
		search.LastConvertionTimeSp, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "INSERT INTO t_setting_connection (" 
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
 * Formデータに含まれる情報を更新する 
 * 除外したいカラムはexcludesへ配列で格納する 
 */ 
func UpdateTSettingConnection(db *sql.Tx, search *autodbdto.DBTSettingConnectionForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"m_setting_sid", 
		"last_compared_time", 
		"last_xd_shop_time", 
		"last_xd_cart_time", 
		"last_convertion_time", 
		"last_compared_time_sp", 
		"last_xd_shop_time_sp", 
		"last_xd_cart_time_sp", 
		"last_convertion_time_sp", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.MSettingSid, 
		search.LastComparedTime, 
		search.LastXdShopTime, 
		search.LastXdCartTime, 
		search.LastConvertionTime, 
		search.LastComparedTimeSp, 
		search.LastXdShopTimeSp, 
		search.LastXdCartTimeSp, 
		search.LastConvertionTimeSp, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "UPDATE t_setting_connection SET " 
	for ind, _ := range columns { 
		if IsColumnExcludes(columns[ind], excludes) == false { 
			if columnSql != "" { 
				columnSql = columnSql + "," 
			} 
			columnSql = columnSql + "`" + columns[ind] + "` = ?" 
			arr = append(arr, values[ind]) 
		} 
	} 
	sql = sql + columnSql 
	sql = sql + " where  m_setting_sid = ? " 
	arr = append(arr, search.MSettingSid) 
 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return false 
	} 
	defer stmt.Close() 
 
	dest := make([]interface{}, len(arr)) 
	for i, _ := range arr { 
		dest[i] = &arr[i] 
	} 
	stmt.Exec(dest...) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, strings.Join(arr, ","), r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		return false 
	} 
	return true 
} 
/* 
 * バルクインサート 
 */ 
func BulkInsertTSettingConnection(db *sql.Tx, search []*autodbdto.DBTSettingConnectionForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"m_setting_sid", 
		"last_compared_time", 
		"last_xd_shop_time", 
		"last_xd_cart_time", 
		"last_convertion_time", 
		"last_compared_time_sp", 
		"last_xd_shop_time_sp", 
		"last_xd_cart_time_sp", 
		"last_convertion_time_sp", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
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
					arr = append(arr, search[dataInd].MSettingSid) 
				case 1: 
					arr = append(arr, search[dataInd].LastComparedTime) 
				case 2: 
					arr = append(arr, search[dataInd].LastXdShopTime) 
				case 3: 
					arr = append(arr, search[dataInd].LastXdCartTime) 
				case 4: 
					arr = append(arr, search[dataInd].LastConvertionTime) 
				case 5: 
					arr = append(arr, search[dataInd].LastComparedTimeSp) 
				case 6: 
					arr = append(arr, search[dataInd].LastXdShopTimeSp) 
				case 7: 
					arr = append(arr, search[dataInd].LastXdCartTimeSp) 
				case 8: 
					arr = append(arr, search[dataInd].LastConvertionTimeSp) 
				case 9: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 10: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 11: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 12: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 13: 
					arr = append(arr, search[dataInd].UpdateTime) 
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
 
/* 
 * バルクアップデート 
 */ 
func BulkUpdateTSettingConnection(db *sql.Tx, search []*autodbdto.DBTSettingConnectionForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"m_setting_sid", 
		"last_compared_time", 
		"last_xd_shop_time", 
		"last_xd_cart_time", 
		"last_convertion_time", 
		"last_compared_time_sp", 
		"last_xd_shop_time_sp", 
		"last_xd_cart_time_sp", 
		"last_convertion_time_sp", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
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
					if dataInd == 0 { 
						duplicateKeys = duplicateKeys + "," 
					} 
				} 
				valueSqlSub = valueSqlSub + "?" 
				if dataInd == 0 { 
					duplicateKeys = duplicateKeys + "" + columns[ind] + " = VALUES(" + columns[ind] + ")" 
				} 
				switch ind { 
				case 0: 
					arr = append(arr, search[dataInd].MSettingSid) 
				case 1: 
					arr = append(arr, search[dataInd].LastComparedTime) 
				case 2: 
					arr = append(arr, search[dataInd].LastXdShopTime) 
				case 3: 
					arr = append(arr, search[dataInd].LastXdCartTime) 
				case 4: 
					arr = append(arr, search[dataInd].LastConvertionTime) 
				case 5: 
					arr = append(arr, search[dataInd].LastComparedTimeSp) 
				case 6: 
					arr = append(arr, search[dataInd].LastXdShopTimeSp) 
				case 7: 
					arr = append(arr, search[dataInd].LastXdCartTimeSp) 
				case 8: 
					arr = append(arr, search[dataInd].LastConvertionTimeSp) 
				case 9: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 10: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 11: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 12: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 13: 
					arr = append(arr, search[dataInd].UpdateTime) 
				} 
			} 
		} 
		valueSql = valueSql + valueSqlSub + ")" 
	} 
	sql = sql + valueSql + "" 
	sql = sql + " ON DUPLICATE KEY UPDATE " + duplicateKeys 
 
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
 
