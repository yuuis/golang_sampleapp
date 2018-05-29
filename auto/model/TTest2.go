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
func GetByPKTTest2(code1 string, code2 string, r *http.Request) (autodbdto.DBTTest2DTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from t_test2 where code1 = ? and code2 = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTTest2DTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(code1, code2) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, code1+code2, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTTest2DTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBTTest2DTO 
 
		err := rows.Scan( 
			&columns.Code1, 
			&columns.Code2, 
			&columns.Number, 
			&columns.InputDate, 
			&columns.Floattest, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBTTest2DTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBTTest2DTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKTTest2ForUpdate(db *sql.Tx, code1 string, code2 string, r *http.Request) (autodbdto.DBTTest2DTO, bool, bool) { 
 
	sql := "select * from t_test2 where code1 = ? and code2 = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTTest2DTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(code1, code2) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, code1+code2, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBTTest2DTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBTTest2DTO 
 
		err := rows.Scan( 
			&columns.Code1, 
			&columns.Code2, 
			&columns.Number, 
			&columns.InputDate, 
			&columns.Floattest, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBTTest2DTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBTTest2DTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInTTest2(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBTTest2DTO, bool) { 
	var rdto []autodbdto.DBTTest2DTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from t_test2 where " + targetColumn + " in (" 
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
		var columns autodbdto.DBTTest2DTO 
 
		err := rows.Scan( 
			&columns.Code1, 
			&columns.Code2, 
			&columns.Number, 
			&columns.InputDate, 
			&columns.Floattest, 
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
func SelectByInTTest2ForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBTTest2DTO, bool) { 
	var rdto []autodbdto.DBTTest2DTO 
 
 
	var arr []string 
 
	sql := "select * from t_test2 where " + targetColumn + " in (" 
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
		var columns autodbdto.DBTTest2DTO 
 
		err := rows.Scan( 
			&columns.Code1, 
			&columns.Code2, 
			&columns.Number, 
			&columns.InputDate, 
			&columns.Floattest, 
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
func SelectTTest2(search *autodbdto.DBTTest2Form, r *http.Request) ([]autodbdto.DBTTest2DTO, bool) { 
	var rdto []autodbdto.DBTTest2DTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from t_test2" 
	where := "" 
 
	where, arr = AppendWhere(search.Code1, "code1", where, arr) 
	where, arr = AppendWhere(search.Code2, "code2", where, arr) 
	where, arr = AppendWhere(search.Number, "number", where, arr) 
	where, arr = AppendWhere(search.InputDate, "input_date", where, arr) 
	where, arr = AppendWhere(search.Floattest, "floattest", where, arr) 
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
		var columns autodbdto.DBTTest2DTO 
 
		err := rows.Scan( 
			&columns.Code1, 
			&columns.Code2, 
			&columns.Number, 
			&columns.InputDate, 
			&columns.Floattest, 
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
func InsertTTest2(db *sql.Tx, search *autodbdto.DBTTest2Form, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"code1", 
		"code2", 
		"number", 
		"input_date", 
		"floattest", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Code1, 
		search.Code2, 
		search.Number, 
		search.InputDate, 
		search.Floattest, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "INSERT INTO t_test2 (" 
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
func UpdateTTest2(db *sql.Tx, search *autodbdto.DBTTest2Form, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"code1", 
		"code2", 
		"number", 
		"input_date", 
		"floattest", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
	} 
	values := []string{ 
		search.Code1, 
		search.Code2, 
		search.Number, 
		search.InputDate, 
		search.Floattest, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
	} 
	sql := "UPDATE t_test2 SET " 
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
	sql = sql + " where  code1 = ? and  code2 = ? " 
	arr = append(arr, search.Code1) 
	arr = append(arr, search.Code2) 
 
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
func BulkInsertTTest2(db *sql.Tx, search []*autodbdto.DBTTest2Form, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"code1", 
		"code2", 
		"number", 
		"input_date", 
		"floattest", 
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
					arr = append(arr, search[dataInd].Code1) 
				case 1: 
					arr = append(arr, search[dataInd].Code2) 
				case 2: 
					arr = append(arr, search[dataInd].Number) 
				case 3: 
					arr = append(arr, search[dataInd].InputDate) 
				case 4: 
					arr = append(arr, search[dataInd].Floattest) 
				case 5: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 6: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 7: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 8: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 9: 
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
func BulkUpdateTTest2(db *sql.Tx, search []*autodbdto.DBTTest2Form, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"code1", 
		"code2", 
		"number", 
		"input_date", 
		"floattest", 
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
					arr = append(arr, search[dataInd].Code1) 
				case 1: 
					arr = append(arr, search[dataInd].Code2) 
				case 2: 
					arr = append(arr, search[dataInd].Number) 
				case 3: 
					arr = append(arr, search[dataInd].InputDate) 
				case 4: 
					arr = append(arr, search[dataInd].Floattest) 
				case 5: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 6: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 7: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 8: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 9: 
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
 
