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
func GetByPKMSettingConnectionMessage(sid string, r *http.Request) (autodbdto.DBMSettingConnectionMessageDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from m_setting_connection_message where sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingConnectionMessageDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingConnectionMessageDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMSettingConnectionMessageDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.LastComparedTimeFlag, 
			&columns.LastXdShopTimeFlag, 
			&columns.LastXdCartTimeFlag, 
			&columns.LastConvertionTimeFlag, 
			&columns.StatusMessage, 
			&columns.DetailMessage, 
			&columns.RemarkMessage, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMSettingConnectionMessageDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMSettingConnectionMessageDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKMSettingConnectionMessageForUpdate(db *sql.Tx, sid string, r *http.Request) (autodbdto.DBMSettingConnectionMessageDTO, bool, bool) { 
 
	sql := "select * from m_setting_connection_message where sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingConnectionMessageDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingConnectionMessageDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMSettingConnectionMessageDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.LastComparedTimeFlag, 
			&columns.LastXdShopTimeFlag, 
			&columns.LastXdCartTimeFlag, 
			&columns.LastConvertionTimeFlag, 
			&columns.StatusMessage, 
			&columns.DetailMessage, 
			&columns.RemarkMessage, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMSettingConnectionMessageDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMSettingConnectionMessageDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInMSettingConnectionMessage(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMSettingConnectionMessageDTO, bool) { 
	var rdto []autodbdto.DBMSettingConnectionMessageDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_setting_connection_message where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMSettingConnectionMessageDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.LastComparedTimeFlag, 
			&columns.LastXdShopTimeFlag, 
			&columns.LastXdCartTimeFlag, 
			&columns.LastConvertionTimeFlag, 
			&columns.StatusMessage, 
			&columns.DetailMessage, 
			&columns.RemarkMessage, 
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
func SelectByInMSettingConnectionMessageForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMSettingConnectionMessageDTO, bool) { 
	var rdto []autodbdto.DBMSettingConnectionMessageDTO 
 
 
	var arr []string 
 
	sql := "select * from m_setting_connection_message where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMSettingConnectionMessageDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.LastComparedTimeFlag, 
			&columns.LastXdShopTimeFlag, 
			&columns.LastXdCartTimeFlag, 
			&columns.LastConvertionTimeFlag, 
			&columns.StatusMessage, 
			&columns.DetailMessage, 
			&columns.RemarkMessage, 
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
func SelectMSettingConnectionMessage(search *autodbdto.DBMSettingConnectionMessageForm, r *http.Request) ([]autodbdto.DBMSettingConnectionMessageDTO, bool) { 
	var rdto []autodbdto.DBMSettingConnectionMessageDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_setting_connection_message" 
	where := "" 
 
	where, arr = AppendWhere(search.Sid, "sid", where, arr) 
	where, arr = AppendWhere(search.LastComparedTimeFlag, "last_compared_time_flag", where, arr) 
	where, arr = AppendWhere(search.LastXdShopTimeFlag, "last_xd_shop_time_flag", where, arr) 
	where, arr = AppendWhere(search.LastXdCartTimeFlag, "last_xd_cart_time_flag", where, arr) 
	where, arr = AppendWhere(search.LastConvertionTimeFlag, "last_convertion_time_flag", where, arr) 
	where, arr = AppendWhere(search.StatusMessage, "status_message", where, arr) 
	where, arr = AppendWhere(search.DetailMessage, "detail_message", where, arr) 
	where, arr = AppendWhere(search.RemarkMessage, "remark_message", where, arr) 
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
		var columns autodbdto.DBMSettingConnectionMessageDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.LastComparedTimeFlag, 
			&columns.LastXdShopTimeFlag, 
			&columns.LastXdCartTimeFlag, 
			&columns.LastConvertionTimeFlag, 
			&columns.StatusMessage, 
			&columns.DetailMessage, 
			&columns.RemarkMessage, 
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
func InsertMSettingConnectionMessage(db *sql.Tx, search *autodbdto.DBMSettingConnectionMessageForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"last_compared_time_flag", 
		"last_xd_shop_time_flag", 
		"last_xd_cart_time_flag", 
		"last_convertion_time_flag", 
		"status_message", 
		"detail_message", 
		"remark_message", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Sid, 
		search.LastComparedTimeFlag, 
		search.LastXdShopTimeFlag, 
		search.LastXdCartTimeFlag, 
		search.LastConvertionTimeFlag, 
		search.StatusMessage, 
		search.DetailMessage, 
		search.RemarkMessage, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "INSERT INTO m_setting_connection_message (" 
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
func UpdateMSettingConnectionMessage(db *sql.Tx, search *autodbdto.DBMSettingConnectionMessageForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"last_compared_time_flag", 
		"last_xd_shop_time_flag", 
		"last_xd_cart_time_flag", 
		"last_convertion_time_flag", 
		"status_message", 
		"detail_message", 
		"remark_message", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Sid, 
		search.LastComparedTimeFlag, 
		search.LastXdShopTimeFlag, 
		search.LastXdCartTimeFlag, 
		search.LastConvertionTimeFlag, 
		search.StatusMessage, 
		search.DetailMessage, 
		search.RemarkMessage, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "UPDATE m_setting_connection_message SET " 
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
	sql = sql + " where  sid = ? " 
	arr = append(arr, search.Sid) 
 
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
func BulkInsertMSettingConnectionMessage(db *sql.Tx, search []*autodbdto.DBMSettingConnectionMessageForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"last_compared_time_flag", 
		"last_xd_shop_time_flag", 
		"last_xd_cart_time_flag", 
		"last_convertion_time_flag", 
		"status_message", 
		"detail_message", 
		"remark_message", 
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
					arr = append(arr, search[dataInd].Sid) 
				case 1: 
					arr = append(arr, search[dataInd].LastComparedTimeFlag) 
				case 2: 
					arr = append(arr, search[dataInd].LastXdShopTimeFlag) 
				case 3: 
					arr = append(arr, search[dataInd].LastXdCartTimeFlag) 
				case 4: 
					arr = append(arr, search[dataInd].LastConvertionTimeFlag) 
				case 5: 
					arr = append(arr, search[dataInd].StatusMessage) 
				case 6: 
					arr = append(arr, search[dataInd].DetailMessage) 
				case 7: 
					arr = append(arr, search[dataInd].RemarkMessage) 
				case 8: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 9: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 10: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 11: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 12: 
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
func BulkUpdateMSettingConnectionMessage(db *sql.Tx, search []*autodbdto.DBMSettingConnectionMessageForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"sid", 
		"last_compared_time_flag", 
		"last_xd_shop_time_flag", 
		"last_xd_cart_time_flag", 
		"last_convertion_time_flag", 
		"status_message", 
		"detail_message", 
		"remark_message", 
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
					arr = append(arr, search[dataInd].Sid) 
				case 1: 
					arr = append(arr, search[dataInd].LastComparedTimeFlag) 
				case 2: 
					arr = append(arr, search[dataInd].LastXdShopTimeFlag) 
				case 3: 
					arr = append(arr, search[dataInd].LastXdCartTimeFlag) 
				case 4: 
					arr = append(arr, search[dataInd].LastConvertionTimeFlag) 
				case 5: 
					arr = append(arr, search[dataInd].StatusMessage) 
				case 6: 
					arr = append(arr, search[dataInd].DetailMessage) 
				case 7: 
					arr = append(arr, search[dataInd].RemarkMessage) 
				case 8: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 9: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 10: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 11: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 12: 
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
 
