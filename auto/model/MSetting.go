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
func GetByPKMSetting(sid string, r *http.Request) (autodbdto.DBMSettingDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from m_setting where sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMSettingDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.BusinessId, 
			&columns.MUserSid, 
			&columns.MShopSid, 
			&columns.Key, 
			&columns.Context, 
			&columns.ActiveFlag, 
			&columns.UrlMatchingPattern, 
			&columns.Status, 
			&columns.ScreenshotGetTime, 
			&columns.RegularlyItemCode, 
			&columns.SortOrder, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMSettingDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMSettingDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKMSettingForUpdate(db *sql.Tx, sid string, r *http.Request) (autodbdto.DBMSettingDTO, bool, bool) { 
 
	sql := "select * from m_setting where sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMSettingDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMSettingDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.BusinessId, 
			&columns.MUserSid, 
			&columns.MShopSid, 
			&columns.Key, 
			&columns.Context, 
			&columns.ActiveFlag, 
			&columns.UrlMatchingPattern, 
			&columns.Status, 
			&columns.ScreenshotGetTime, 
			&columns.RegularlyItemCode, 
			&columns.SortOrder, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMSettingDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMSettingDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInMSetting(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMSettingDTO, bool) { 
	var rdto []autodbdto.DBMSettingDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_setting where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMSettingDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.BusinessId, 
			&columns.MUserSid, 
			&columns.MShopSid, 
			&columns.Key, 
			&columns.Context, 
			&columns.ActiveFlag, 
			&columns.UrlMatchingPattern, 
			&columns.Status, 
			&columns.ScreenshotGetTime, 
			&columns.RegularlyItemCode, 
			&columns.SortOrder, 
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
func SelectByInMSettingForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMSettingDTO, bool) { 
	var rdto []autodbdto.DBMSettingDTO 
 
 
	var arr []string 
 
	sql := "select * from m_setting where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMSettingDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.BusinessId, 
			&columns.MUserSid, 
			&columns.MShopSid, 
			&columns.Key, 
			&columns.Context, 
			&columns.ActiveFlag, 
			&columns.UrlMatchingPattern, 
			&columns.Status, 
			&columns.ScreenshotGetTime, 
			&columns.RegularlyItemCode, 
			&columns.SortOrder, 
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
func SelectMSetting(search *autodbdto.DBMSettingForm, r *http.Request) ([]autodbdto.DBMSettingDTO, bool) { 
	var rdto []autodbdto.DBMSettingDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_setting" 
	where := "" 
 
	where, arr = AppendWhere(search.Sid, "sid", where, arr) 
	where, arr = AppendWhere(search.BusinessId, "business_id", where, arr) 
	where, arr = AppendWhere(search.MUserSid, "m_user_sid", where, arr) 
	where, arr = AppendWhere(search.MShopSid, "m_shop_sid", where, arr) 
	where, arr = AppendWhere(search.Key, "key", where, arr) 
	where, arr = AppendWhere(search.Context, "context", where, arr) 
	where, arr = AppendWhere(search.ActiveFlag, "active_flag", where, arr) 
	where, arr = AppendWhere(search.UrlMatchingPattern, "url_matching_pattern", where, arr) 
	where, arr = AppendWhere(search.Status, "status", where, arr) 
	where, arr = AppendWhere(search.ScreenshotGetTime, "screenshot_get_time", where, arr) 
	where, arr = AppendWhere(search.RegularlyItemCode, "regularly_item_code", where, arr) 
	where, arr = AppendWhere(search.SortOrder, "sort_order", where, arr) 
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
		var columns autodbdto.DBMSettingDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.BusinessId, 
			&columns.MUserSid, 
			&columns.MShopSid, 
			&columns.Key, 
			&columns.Context, 
			&columns.ActiveFlag, 
			&columns.UrlMatchingPattern, 
			&columns.Status, 
			&columns.ScreenshotGetTime, 
			&columns.RegularlyItemCode, 
			&columns.SortOrder, 
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
func InsertMSetting(db *sql.Tx, search *autodbdto.DBMSettingForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"business_id", 
		"m_user_sid", 
		"m_shop_sid", 
		"key", 
		"context", 
		"active_flag", 
		"url_matching_pattern", 
		"status", 
		"screenshot_get_time", 
		"regularly_item_code", 
		"sort_order", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Sid, 
		search.BusinessId, 
		search.MUserSid, 
		search.MShopSid, 
		search.Key, 
		search.Context, 
		search.ActiveFlag, 
		search.UrlMatchingPattern, 
		search.Status, 
		search.ScreenshotGetTime, 
		search.RegularlyItemCode, 
		search.SortOrder, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "INSERT INTO m_setting (" 
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
func UpdateMSetting(db *sql.Tx, search *autodbdto.DBMSettingForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"business_id", 
		"m_user_sid", 
		"m_shop_sid", 
		"key", 
		"context", 
		"active_flag", 
		"url_matching_pattern", 
		"status", 
		"screenshot_get_time", 
		"regularly_item_code", 
		"sort_order", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Sid, 
		search.BusinessId, 
		search.MUserSid, 
		search.MShopSid, 
		search.Key, 
		search.Context, 
		search.ActiveFlag, 
		search.UrlMatchingPattern, 
		search.Status, 
		search.ScreenshotGetTime, 
		search.RegularlyItemCode, 
		search.SortOrder, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "UPDATE m_setting SET " 
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
func BulkInsertMSetting(db *sql.Tx, search []*autodbdto.DBMSettingForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"business_id", 
		"m_user_sid", 
		"m_shop_sid", 
		"key", 
		"context", 
		"active_flag", 
		"url_matching_pattern", 
		"status", 
		"screenshot_get_time", 
		"regularly_item_code", 
		"sort_order", 
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
					arr = append(arr, search[dataInd].BusinessId) 
				case 2: 
					arr = append(arr, search[dataInd].MUserSid) 
				case 3: 
					arr = append(arr, search[dataInd].MShopSid) 
				case 4: 
					arr = append(arr, search[dataInd].Key) 
				case 5: 
					arr = append(arr, search[dataInd].Context) 
				case 6: 
					arr = append(arr, search[dataInd].ActiveFlag) 
				case 7: 
					arr = append(arr, search[dataInd].UrlMatchingPattern) 
				case 8: 
					arr = append(arr, search[dataInd].Status) 
				case 9: 
					arr = append(arr, search[dataInd].ScreenshotGetTime) 
				case 10: 
					arr = append(arr, search[dataInd].RegularlyItemCode) 
				case 11: 
					arr = append(arr, search[dataInd].SortOrder) 
				case 12: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 13: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 14: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 15: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 16: 
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
func BulkUpdateMSetting(db *sql.Tx, search []*autodbdto.DBMSettingForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"sid", 
		"business_id", 
		"m_user_sid", 
		"m_shop_sid", 
		"key", 
		"context", 
		"active_flag", 
		"url_matching_pattern", 
		"status", 
		"screenshot_get_time", 
		"regularly_item_code", 
		"sort_order", 
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
					arr = append(arr, search[dataInd].BusinessId) 
				case 2: 
					arr = append(arr, search[dataInd].MUserSid) 
				case 3: 
					arr = append(arr, search[dataInd].MShopSid) 
				case 4: 
					arr = append(arr, search[dataInd].Key) 
				case 5: 
					arr = append(arr, search[dataInd].Context) 
				case 6: 
					arr = append(arr, search[dataInd].ActiveFlag) 
				case 7: 
					arr = append(arr, search[dataInd].UrlMatchingPattern) 
				case 8: 
					arr = append(arr, search[dataInd].Status) 
				case 9: 
					arr = append(arr, search[dataInd].ScreenshotGetTime) 
				case 10: 
					arr = append(arr, search[dataInd].RegularlyItemCode) 
				case 11: 
					arr = append(arr, search[dataInd].SortOrder) 
				case 12: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 13: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 14: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 15: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 16: 
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
 
