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
func GetByPKMShop(sid string, r *http.Request) (autodbdto.DBMShopDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from m_shop where sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMShopDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMShopDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMShopDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.ShopId, 
			&columns.BusinessId, 
			&columns.ShopName, 
			&columns.ManagerAuthKey, 
			&columns.LastAccessTime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMShopDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMShopDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKMShopForUpdate(db *sql.Tx, sid string, r *http.Request) (autodbdto.DBMShopDTO, bool, bool) { 
 
	sql := "select * from m_shop where sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMShopDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMShopDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMShopDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.ShopId, 
			&columns.BusinessId, 
			&columns.ShopName, 
			&columns.ManagerAuthKey, 
			&columns.LastAccessTime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMShopDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMShopDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInMShop(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMShopDTO, bool) { 
	var rdto []autodbdto.DBMShopDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_shop where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMShopDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.ShopId, 
			&columns.BusinessId, 
			&columns.ShopName, 
			&columns.ManagerAuthKey, 
			&columns.LastAccessTime, 
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
func SelectByInMShopForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMShopDTO, bool) { 
	var rdto []autodbdto.DBMShopDTO 
 
 
	var arr []string 
 
	sql := "select * from m_shop where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMShopDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.ShopId, 
			&columns.BusinessId, 
			&columns.ShopName, 
			&columns.ManagerAuthKey, 
			&columns.LastAccessTime, 
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
func SelectMShop(search *autodbdto.DBMShopForm, r *http.Request) ([]autodbdto.DBMShopDTO, bool) { 
	var rdto []autodbdto.DBMShopDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_shop" 
	where := "" 
 
	where, arr = AppendWhere(search.Sid, "sid", where, arr) 
	where, arr = AppendWhere(search.ShopId, "shop_id", where, arr) 
	where, arr = AppendWhere(search.BusinessId, "business_id", where, arr) 
	where, arr = AppendWhere(search.ShopName, "shop_name", where, arr) 
	where, arr = AppendWhere(search.ManagerAuthKey, "manager_auth_key", where, arr) 
	where, arr = AppendWhere(search.LastAccessTime, "last_access_time", where, arr) 
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
		var columns autodbdto.DBMShopDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.ShopId, 
			&columns.BusinessId, 
			&columns.ShopName, 
			&columns.ManagerAuthKey, 
			&columns.LastAccessTime, 
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
func InsertMShop(db *sql.Tx, search *autodbdto.DBMShopForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"shop_id", 
		"business_id", 
		"shop_name", 
		"manager_auth_key", 
		"last_access_time", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Sid, 
		search.ShopId, 
		search.BusinessId, 
		search.ShopName, 
		search.ManagerAuthKey, 
		search.LastAccessTime, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "INSERT INTO m_shop (" 
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
func UpdateMShop(db *sql.Tx, search *autodbdto.DBMShopForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"shop_id", 
		"business_id", 
		"shop_name", 
		"manager_auth_key", 
		"last_access_time", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Sid, 
		search.ShopId, 
		search.BusinessId, 
		search.ShopName, 
		search.ManagerAuthKey, 
		search.LastAccessTime, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "UPDATE m_shop SET " 
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
func BulkInsertMShop(db *sql.Tx, search []*autodbdto.DBMShopForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"shop_id", 
		"business_id", 
		"shop_name", 
		"manager_auth_key", 
		"last_access_time", 
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
					arr = append(arr, search[dataInd].ShopId) 
				case 2: 
					arr = append(arr, search[dataInd].BusinessId) 
				case 3: 
					arr = append(arr, search[dataInd].ShopName) 
				case 4: 
					arr = append(arr, search[dataInd].ManagerAuthKey) 
				case 5: 
					arr = append(arr, search[dataInd].LastAccessTime) 
				case 6: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 7: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 8: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 9: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 10: 
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
func BulkUpdateMShop(db *sql.Tx, search []*autodbdto.DBMShopForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"sid", 
		"shop_id", 
		"business_id", 
		"shop_name", 
		"manager_auth_key", 
		"last_access_time", 
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
					arr = append(arr, search[dataInd].ShopId) 
				case 2: 
					arr = append(arr, search[dataInd].BusinessId) 
				case 3: 
					arr = append(arr, search[dataInd].ShopName) 
				case 4: 
					arr = append(arr, search[dataInd].ManagerAuthKey) 
				case 5: 
					arr = append(arr, search[dataInd].LastAccessTime) 
				case 6: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 7: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 8: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 9: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 10: 
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
 
