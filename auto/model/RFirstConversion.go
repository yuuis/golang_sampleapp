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
func GetByPKRFirstConversion(sid string, r *http.Request) (autodbdto.DBRFirstConversionDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from r_first_conversion where sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBRFirstConversionDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBRFirstConversionDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBRFirstConversionDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.MSettingSid, 
			&columns.Url, 
			&columns.UserKey, 
			&columns.FirstTBeaconSid, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBRFirstConversionDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBRFirstConversionDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKRFirstConversionForUpdate(db *sql.Tx, sid string, r *http.Request) (autodbdto.DBRFirstConversionDTO, bool, bool) { 
 
	sql := "select * from r_first_conversion where sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBRFirstConversionDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBRFirstConversionDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBRFirstConversionDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.MSettingSid, 
			&columns.Url, 
			&columns.UserKey, 
			&columns.FirstTBeaconSid, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBRFirstConversionDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBRFirstConversionDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInRFirstConversion(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBRFirstConversionDTO, bool) { 
	var rdto []autodbdto.DBRFirstConversionDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from r_first_conversion where " + targetColumn + " in (" 
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
		var columns autodbdto.DBRFirstConversionDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.MSettingSid, 
			&columns.Url, 
			&columns.UserKey, 
			&columns.FirstTBeaconSid, 
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
func SelectByInRFirstConversionForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBRFirstConversionDTO, bool) { 
	var rdto []autodbdto.DBRFirstConversionDTO 
 
 
	var arr []string 
 
	sql := "select * from r_first_conversion where " + targetColumn + " in (" 
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
		var columns autodbdto.DBRFirstConversionDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.MSettingSid, 
			&columns.Url, 
			&columns.UserKey, 
			&columns.FirstTBeaconSid, 
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
func SelectRFirstConversion(search *autodbdto.DBRFirstConversionForm, r *http.Request) ([]autodbdto.DBRFirstConversionDTO, bool) { 
	var rdto []autodbdto.DBRFirstConversionDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from r_first_conversion" 
	where := "" 
 
	where, arr = AppendWhere(search.Sid, "sid", where, arr) 
	where, arr = AppendWhere(search.MSettingSid, "m_setting_sid", where, arr) 
	where, arr = AppendWhere(search.Url, "url", where, arr) 
	where, arr = AppendWhere(search.UserKey, "user_key", where, arr) 
	where, arr = AppendWhere(search.FirstTBeaconSid, "first_t_beacon_sid", where, arr) 
 
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
		var columns autodbdto.DBRFirstConversionDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.MSettingSid, 
			&columns.Url, 
			&columns.UserKey, 
			&columns.FirstTBeaconSid, 
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
func InsertRFirstConversion(db *sql.Tx, search *autodbdto.DBRFirstConversionForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"m_setting_sid", 
		"url", 
		"user_key", 
		"first_t_beacon_sid", 
	} 
	values := []string{ 
		search.Sid, 
		search.MSettingSid, 
		search.Url, 
		search.UserKey, 
		search.FirstTBeaconSid, 
	} 
	sql := "INSERT INTO r_first_conversion (" 
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
func UpdateRFirstConversion(db *sql.Tx, search *autodbdto.DBRFirstConversionForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"m_setting_sid", 
		"url", 
		"user_key", 
		"first_t_beacon_sid", 
	} 
	values := []string{ 
		search.Sid, 
		search.MSettingSid, 
		search.Url, 
		search.UserKey, 
		search.FirstTBeaconSid, 
	} 
	sql := "UPDATE r_first_conversion SET " 
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
func BulkInsertRFirstConversion(db *sql.Tx, search []*autodbdto.DBRFirstConversionForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"m_setting_sid", 
		"url", 
		"user_key", 
		"first_t_beacon_sid", 
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
					arr = append(arr, search[dataInd].MSettingSid) 
				case 2: 
					arr = append(arr, search[dataInd].Url) 
				case 3: 
					arr = append(arr, search[dataInd].UserKey) 
				case 4: 
					arr = append(arr, search[dataInd].FirstTBeaconSid) 
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
func BulkUpdateRFirstConversion(db *sql.Tx, search []*autodbdto.DBRFirstConversionForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"sid", 
		"m_setting_sid", 
		"url", 
		"user_key", 
		"first_t_beacon_sid", 
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
					arr = append(arr, search[dataInd].MSettingSid) 
				case 2: 
					arr = append(arr, search[dataInd].Url) 
				case 3: 
					arr = append(arr, search[dataInd].UserKey) 
				case 4: 
					arr = append(arr, search[dataInd].FirstTBeaconSid) 
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
 
