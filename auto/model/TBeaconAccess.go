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
func GetByPKTBeaconAccess(sid string, r *http.Request) (autodbdto.DBTBeaconAccessDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from t_beacon_access where sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTBeaconAccessDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTBeaconAccessDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBTBeaconAccessDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.Key, 
			&columns.Name, 
			&columns.Url, 
			&columns.Percent, 
			&columns.Encode, 
			&columns.Forward, 
			&columns.RequestUrl, 
			&columns.RefererUrl, 
			&columns.AccessTime, 
			&columns.SessionId, 
			&columns.CookiesKey, 
			&columns.Price, 
			&columns.UserKey, 
			&columns.IpAddress, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.OrderNumber, 
			&columns.UserAgent, 
			&columns.RegularlyFlag, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBTBeaconAccessDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBTBeaconAccessDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKTBeaconAccessForUpdate(db *sql.Tx, sid string, r *http.Request) (autodbdto.DBTBeaconAccessDTO, bool, bool) { 
 
	sql := "select * from t_beacon_access where sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTBeaconAccessDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTBeaconAccessDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBTBeaconAccessDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.Key, 
			&columns.Name, 
			&columns.Url, 
			&columns.Percent, 
			&columns.Encode, 
			&columns.Forward, 
			&columns.RequestUrl, 
			&columns.RefererUrl, 
			&columns.AccessTime, 
			&columns.SessionId, 
			&columns.CookiesKey, 
			&columns.Price, 
			&columns.UserKey, 
			&columns.IpAddress, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.OrderNumber, 
			&columns.UserAgent, 
			&columns.RegularlyFlag, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBTBeaconAccessDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBTBeaconAccessDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInTBeaconAccess(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBTBeaconAccessDTO, bool) { 
	var rdto []autodbdto.DBTBeaconAccessDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from t_beacon_access where " + targetColumn + " in (" 
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
		var columns autodbdto.DBTBeaconAccessDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.Key, 
			&columns.Name, 
			&columns.Url, 
			&columns.Percent, 
			&columns.Encode, 
			&columns.Forward, 
			&columns.RequestUrl, 
			&columns.RefererUrl, 
			&columns.AccessTime, 
			&columns.SessionId, 
			&columns.CookiesKey, 
			&columns.Price, 
			&columns.UserKey, 
			&columns.IpAddress, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.OrderNumber, 
			&columns.UserAgent, 
			&columns.RegularlyFlag, 
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
func SelectByInTBeaconAccessForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBTBeaconAccessDTO, bool) { 
	var rdto []autodbdto.DBTBeaconAccessDTO 
 
 
	var arr []string 
 
	sql := "select * from t_beacon_access where " + targetColumn + " in (" 
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
		var columns autodbdto.DBTBeaconAccessDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.Key, 
			&columns.Name, 
			&columns.Url, 
			&columns.Percent, 
			&columns.Encode, 
			&columns.Forward, 
			&columns.RequestUrl, 
			&columns.RefererUrl, 
			&columns.AccessTime, 
			&columns.SessionId, 
			&columns.CookiesKey, 
			&columns.Price, 
			&columns.UserKey, 
			&columns.IpAddress, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.OrderNumber, 
			&columns.UserAgent, 
			&columns.RegularlyFlag, 
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
func SelectTBeaconAccess(search *autodbdto.DBTBeaconAccessForm, r *http.Request) ([]autodbdto.DBTBeaconAccessDTO, bool) { 
	var rdto []autodbdto.DBTBeaconAccessDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from t_beacon_access" 
	where := "" 
 
	where, arr = AppendWhere(search.Sid, "sid", where, arr) 
	where, arr = AppendWhere(search.Key, "key", where, arr) 
	where, arr = AppendWhere(search.Name, "name", where, arr) 
	where, arr = AppendWhere(search.Url, "url", where, arr) 
	where, arr = AppendWhere(search.Percent, "percent", where, arr) 
	where, arr = AppendWhere(search.Encode, "encode", where, arr) 
	where, arr = AppendWhere(search.Forward, "forward", where, arr) 
	where, arr = AppendWhere(search.RequestUrl, "request_url", where, arr) 
	where, arr = AppendWhere(search.RefererUrl, "referer_url", where, arr) 
	where, arr = AppendWhere(search.AccessTime, "access_time", where, arr) 
	where, arr = AppendWhere(search.SessionId, "session_id", where, arr) 
	where, arr = AppendWhere(search.CookiesKey, "cookies_key", where, arr) 
	where, arr = AppendWhere(search.Price, "price", where, arr) 
	where, arr = AppendWhere(search.UserKey, "user_key", where, arr) 
	where, arr = AppendWhere(search.IpAddress, "ip_address", where, arr) 
	where, arr = AppendWhere(search.DeleteFlag, "delete_flag", where, arr) 
	where, arr = AppendWhere(search.CreateUserId, "create_user_id", where, arr) 
	where, arr = AppendWhere(search.UpdateUserId, "update_user_id", where, arr) 
	where, arr = AppendWhere(search.CreateTime, "create_time", where, arr) 
	where, arr = AppendWhere(search.UpdateTime, "update_time", where, arr) 
	where, arr = AppendWhere(search.OrderNumber, "order_number", where, arr) 
	where, arr = AppendWhere(search.UserAgent, "user_agent", where, arr) 
	where, arr = AppendWhere(search.RegularlyFlag, "regularly_flag", where, arr) 
 
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
		var columns autodbdto.DBTBeaconAccessDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.Key, 
			&columns.Name, 
			&columns.Url, 
			&columns.Percent, 
			&columns.Encode, 
			&columns.Forward, 
			&columns.RequestUrl, 
			&columns.RefererUrl, 
			&columns.AccessTime, 
			&columns.SessionId, 
			&columns.CookiesKey, 
			&columns.Price, 
			&columns.UserKey, 
			&columns.IpAddress, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.OrderNumber, 
			&columns.UserAgent, 
			&columns.RegularlyFlag, 
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
func InsertTBeaconAccess(db *sql.Tx, search *autodbdto.DBTBeaconAccessForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"key", 
		"name", 
		"url", 
		"percent", 
		"encode", 
		"forward", 
		"request_url", 
		"referer_url", 
		"access_time", 
		"session_id", 
		"cookies_key", 
		"price", 
		"user_key", 
		"ip_address", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"order_number", 
		"user_agent", 
		"regularly_flag", 
	} 
	values := []string{ 
		search.Sid, 
		search.Key, 
		search.Name, 
		search.Url, 
		search.Percent, 
		search.Encode, 
		search.Forward, 
		search.RequestUrl, 
		search.RefererUrl, 
		search.AccessTime, 
		search.SessionId, 
		search.CookiesKey, 
		search.Price, 
		search.UserKey, 
		search.IpAddress, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
		search.OrderNumber, 
		search.UserAgent, 
		search.RegularlyFlag, 
	} 
	sql := "INSERT INTO t_beacon_access (" 
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
func UpdateTBeaconAccess(db *sql.Tx, search *autodbdto.DBTBeaconAccessForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"key", 
		"name", 
		"url", 
		"percent", 
		"encode", 
		"forward", 
		"request_url", 
		"referer_url", 
		"access_time", 
		"session_id", 
		"cookies_key", 
		"price", 
		"user_key", 
		"ip_address", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"order_number", 
		"user_agent", 
		"regularly_flag", 
	} 
	values := []string{ 
		search.Sid, 
		search.Key, 
		search.Name, 
		search.Url, 
		search.Percent, 
		search.Encode, 
		search.Forward, 
		search.RequestUrl, 
		search.RefererUrl, 
		search.AccessTime, 
		search.SessionId, 
		search.CookiesKey, 
		search.Price, 
		search.UserKey, 
		search.IpAddress, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
		search.OrderNumber, 
		search.UserAgent, 
		search.RegularlyFlag, 
	} 
	sql := "UPDATE t_beacon_access SET " 
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
func BulkInsertTBeaconAccess(db *sql.Tx, search []*autodbdto.DBTBeaconAccessForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"key", 
		"name", 
		"url", 
		"percent", 
		"encode", 
		"forward", 
		"request_url", 
		"referer_url", 
		"access_time", 
		"session_id", 
		"cookies_key", 
		"price", 
		"user_key", 
		"ip_address", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"order_number", 
		"user_agent", 
		"regularly_flag", 
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
					arr = append(arr, search[dataInd].Key) 
				case 2: 
					arr = append(arr, search[dataInd].Name) 
				case 3: 
					arr = append(arr, search[dataInd].Url) 
				case 4: 
					arr = append(arr, search[dataInd].Percent) 
				case 5: 
					arr = append(arr, search[dataInd].Encode) 
				case 6: 
					arr = append(arr, search[dataInd].Forward) 
				case 7: 
					arr = append(arr, search[dataInd].RequestUrl) 
				case 8: 
					arr = append(arr, search[dataInd].RefererUrl) 
				case 9: 
					arr = append(arr, search[dataInd].AccessTime) 
				case 10: 
					arr = append(arr, search[dataInd].SessionId) 
				case 11: 
					arr = append(arr, search[dataInd].CookiesKey) 
				case 12: 
					arr = append(arr, search[dataInd].Price) 
				case 13: 
					arr = append(arr, search[dataInd].UserKey) 
				case 14: 
					arr = append(arr, search[dataInd].IpAddress) 
				case 15: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 16: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 17: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 18: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 19: 
					arr = append(arr, search[dataInd].UpdateTime) 
				case 20: 
					arr = append(arr, search[dataInd].OrderNumber) 
				case 21: 
					arr = append(arr, search[dataInd].UserAgent) 
				case 22: 
					arr = append(arr, search[dataInd].RegularlyFlag) 
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
func BulkUpdateTBeaconAccess(db *sql.Tx, search []*autodbdto.DBTBeaconAccessForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"sid", 
		"key", 
		"name", 
		"url", 
		"percent", 
		"encode", 
		"forward", 
		"request_url", 
		"referer_url", 
		"access_time", 
		"session_id", 
		"cookies_key", 
		"price", 
		"user_key", 
		"ip_address", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"order_number", 
		"user_agent", 
		"regularly_flag", 
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
					arr = append(arr, search[dataInd].Key) 
				case 2: 
					arr = append(arr, search[dataInd].Name) 
				case 3: 
					arr = append(arr, search[dataInd].Url) 
				case 4: 
					arr = append(arr, search[dataInd].Percent) 
				case 5: 
					arr = append(arr, search[dataInd].Encode) 
				case 6: 
					arr = append(arr, search[dataInd].Forward) 
				case 7: 
					arr = append(arr, search[dataInd].RequestUrl) 
				case 8: 
					arr = append(arr, search[dataInd].RefererUrl) 
				case 9: 
					arr = append(arr, search[dataInd].AccessTime) 
				case 10: 
					arr = append(arr, search[dataInd].SessionId) 
				case 11: 
					arr = append(arr, search[dataInd].CookiesKey) 
				case 12: 
					arr = append(arr, search[dataInd].Price) 
				case 13: 
					arr = append(arr, search[dataInd].UserKey) 
				case 14: 
					arr = append(arr, search[dataInd].IpAddress) 
				case 15: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 16: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 17: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 18: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 19: 
					arr = append(arr, search[dataInd].UpdateTime) 
				case 20: 
					arr = append(arr, search[dataInd].OrderNumber) 
				case 21: 
					arr = append(arr, search[dataInd].UserAgent) 
				case 22: 
					arr = append(arr, search[dataInd].RegularlyFlag) 
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
 
