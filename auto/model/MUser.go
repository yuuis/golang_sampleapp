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
func GetByPKMUser(sid string, r *http.Request) (autodbdto.DBMUserDTO, bool, bool) { 
	db := db.DbConn() 
 
	sql := "select * from m_user where sid = ? "
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMUserDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMUserDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMUserDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.UserId, 
			&columns.Password, 
			&columns.BusinessId, 
			&columns.ActiveFlag, 
			&columns.CorporateName, 
			&columns.CorporateNameKana, 
			&columns.DepartmentName, 
			&columns.FamilyName, 
			&columns.FirstName, 
			&columns.FamilyNameKana, 
			&columns.FirstNameKana, 
			&columns.PostalCode, 
			&columns.PrefectureName, 
			&columns.Address, 
			&columns.Address2, 
			&columns.Tel, 
			&columns.MailAddress, 
			&columns.ReminderOnetime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.StartTime, 
			&columns.EndTime, 
			&columns.NotClimFlag, 
			&columns.Bpcode, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMUserDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMUserDTO 
	return rdto, false, true 
} 
/* 
 * PKからデータを取得しつつロックをかける 
 * 取得できなかった場合は空の構造体とfalseを返す 
 */ 
func GetByPKMUserForUpdate(db *sql.Tx, sid string, r *http.Request) (autodbdto.DBMUserDTO, bool, bool) { 
 
	sql := "select * from m_user where sid = ? "
	sql = sql + " for update " 
	stmt, err := db.Prepare(sql) 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMUserDTO 
		return rv, false, false 
	} 
	defer stmt.Close() 
 
	rows, err := stmt.Query(sid) 
	common.WriteLog(config.DEBUG, sql, r) 
	common.WriteLog(config.DEBUG, sid, r) 
 
	if err != nil { 
		common.WriteErrorLog(config.FATAL, err, r) 
		var rv autodbdto.DBMUserDTO 
		return rv, false, false 
	} 
	defer rows.Close() 
	for rows.Next() { 
		var columns autodbdto.DBMUserDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.UserId, 
			&columns.Password, 
			&columns.BusinessId, 
			&columns.ActiveFlag, 
			&columns.CorporateName, 
			&columns.CorporateNameKana, 
			&columns.DepartmentName, 
			&columns.FamilyName, 
			&columns.FirstName, 
			&columns.FamilyNameKana, 
			&columns.FirstNameKana, 
			&columns.PostalCode, 
			&columns.PrefectureName, 
			&columns.Address, 
			&columns.Address2, 
			&columns.Tel, 
			&columns.MailAddress, 
			&columns.ReminderOnetime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.StartTime, 
			&columns.EndTime, 
			&columns.NotClimFlag, 
			&columns.Bpcode, 
		) 
 
		if err != nil { 
			common.WriteErrorLog(config.FATAL, err, r) 
			var rv autodbdto.DBMUserDTO 
			return rv, false, false 
		} 
		return columns, true, true 
	} 
 
	var rdto autodbdto.DBMUserDTO 
	return rdto, false, true 
} 
 
/* 
 * 指定したカラムに対してIN句を発行する 
 * 取得できなかった場合は空の配列を返す 
 */ 
func SelectByInMUser(targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMUserDTO, bool) { 
	var rdto []autodbdto.DBMUserDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_user where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMUserDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.UserId, 
			&columns.Password, 
			&columns.BusinessId, 
			&columns.ActiveFlag, 
			&columns.CorporateName, 
			&columns.CorporateNameKana, 
			&columns.DepartmentName, 
			&columns.FamilyName, 
			&columns.FirstName, 
			&columns.FamilyNameKana, 
			&columns.FirstNameKana, 
			&columns.PostalCode, 
			&columns.PrefectureName, 
			&columns.Address, 
			&columns.Address2, 
			&columns.Tel, 
			&columns.MailAddress, 
			&columns.ReminderOnetime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.StartTime, 
			&columns.EndTime, 
			&columns.NotClimFlag, 
			&columns.Bpcode, 
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
func SelectByInMUserForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DBMUserDTO, bool) { 
	var rdto []autodbdto.DBMUserDTO 
 
 
	var arr []string 
 
	sql := "select * from m_user where " + targetColumn + " in (" 
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
		var columns autodbdto.DBMUserDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.UserId, 
			&columns.Password, 
			&columns.BusinessId, 
			&columns.ActiveFlag, 
			&columns.CorporateName, 
			&columns.CorporateNameKana, 
			&columns.DepartmentName, 
			&columns.FamilyName, 
			&columns.FirstName, 
			&columns.FamilyNameKana, 
			&columns.FirstNameKana, 
			&columns.PostalCode, 
			&columns.PrefectureName, 
			&columns.Address, 
			&columns.Address2, 
			&columns.Tel, 
			&columns.MailAddress, 
			&columns.ReminderOnetime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.StartTime, 
			&columns.EndTime, 
			&columns.NotClimFlag, 
			&columns.Bpcode, 
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
func SelectMUser(search *autodbdto.DBMUserForm, r *http.Request) ([]autodbdto.DBMUserDTO, bool) { 
	var rdto []autodbdto.DBMUserDTO 
 
	db := db.DbConn() 
 
	var arr []string 
 
	sql := "select * from m_user" 
	where := "" 
 
	where, arr = AppendWhere(search.Sid, "sid", where, arr) 
	where, arr = AppendWhere(search.UserId, "user_id", where, arr) 
	where, arr = AppendWhere(search.Password, "password", where, arr) 
	where, arr = AppendWhere(search.BusinessId, "business_id", where, arr) 
	where, arr = AppendWhere(search.ActiveFlag, "active_flag", where, arr) 
	where, arr = AppendWhere(search.CorporateName, "corporate_name", where, arr) 
	where, arr = AppendWhere(search.CorporateNameKana, "corporate_name_kana", where, arr) 
	where, arr = AppendWhere(search.DepartmentName, "department_name", where, arr) 
	where, arr = AppendWhere(search.FamilyName, "family_name", where, arr) 
	where, arr = AppendWhere(search.FirstName, "first_name", where, arr) 
	where, arr = AppendWhere(search.FamilyNameKana, "family_name_kana", where, arr) 
	where, arr = AppendWhere(search.FirstNameKana, "first_name_kana", where, arr) 
	where, arr = AppendWhere(search.PostalCode, "postal_code", where, arr) 
	where, arr = AppendWhere(search.PrefectureName, "prefecture_name", where, arr) 
	where, arr = AppendWhere(search.Address, "address", where, arr) 
	where, arr = AppendWhere(search.Address2, "address_2", where, arr) 
	where, arr = AppendWhere(search.Tel, "tel", where, arr) 
	where, arr = AppendWhere(search.MailAddress, "mail_address", where, arr) 
	where, arr = AppendWhere(search.ReminderOnetime, "reminder_onetime", where, arr) 
	where, arr = AppendWhere(search.DeleteFlag, "delete_flag", where, arr) 
	where, arr = AppendWhere(search.CreateUserId, "create_user_id", where, arr) 
	where, arr = AppendWhere(search.UpdateUserId, "update_user_id", where, arr) 
	where, arr = AppendWhere(search.CreateTime, "create_time", where, arr) 
	where, arr = AppendWhere(search.UpdateTime, "update_time", where, arr) 
	where, arr = AppendWhere(search.StartTime, "start_time", where, arr) 
	where, arr = AppendWhere(search.EndTime, "end_time", where, arr) 
	where, arr = AppendWhere(search.NotClimFlag, "not_clim_flag", where, arr) 
	where, arr = AppendWhere(search.Bpcode, "bpcode", where, arr) 
 
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
		var columns autodbdto.DBMUserDTO 
 
		err := rows.Scan( 
			&columns.Sid, 
			&columns.UserId, 
			&columns.Password, 
			&columns.BusinessId, 
			&columns.ActiveFlag, 
			&columns.CorporateName, 
			&columns.CorporateNameKana, 
			&columns.DepartmentName, 
			&columns.FamilyName, 
			&columns.FirstName, 
			&columns.FamilyNameKana, 
			&columns.FirstNameKana, 
			&columns.PostalCode, 
			&columns.PrefectureName, 
			&columns.Address, 
			&columns.Address2, 
			&columns.Tel, 
			&columns.MailAddress, 
			&columns.ReminderOnetime, 
			&columns.DeleteFlag, 
			&columns.CreateUserId, 
			&columns.UpdateUserId, 
			&columns.CreateTime, 
			&columns.UpdateTime, 
			&columns.StartTime, 
			&columns.EndTime, 
			&columns.NotClimFlag, 
			&columns.Bpcode, 
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
func InsertMUser(db *sql.Tx, search *autodbdto.DBMUserForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"user_id", 
		"password", 
		"business_id", 
		"active_flag", 
		"corporate_name", 
		"corporate_name_kana", 
		"department_name", 
		"family_name", 
		"first_name", 
		"family_name_kana", 
		"first_name_kana", 
		"postal_code", 
		"prefecture_name", 
		"address", 
		"address_2", 
		"tel", 
		"mail_address", 
		"reminder_onetime", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"start_time", 
		"end_time", 
		"not_clim_flag", 
		"bpcode", 
	} 
	values := []string{ 
		search.Sid, 
		search.UserId, 
		search.Password, 
		search.BusinessId, 
		search.ActiveFlag, 
		search.CorporateName, 
		search.CorporateNameKana, 
		search.DepartmentName, 
		search.FamilyName, 
		search.FirstName, 
		search.FamilyNameKana, 
		search.FirstNameKana, 
		search.PostalCode, 
		search.PrefectureName, 
		search.Address, 
		search.Address2, 
		search.Tel, 
		search.MailAddress, 
		search.ReminderOnetime, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
		search.StartTime, 
		search.EndTime, 
		search.NotClimFlag, 
		search.Bpcode, 
	} 
	sql := "INSERT INTO m_user (" 
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
func UpdateMUser(db *sql.Tx, search *autodbdto.DBMUserForm, r *http.Request, excludes []string) bool { 
	columnSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"user_id", 
		"password", 
		"business_id", 
		"active_flag", 
		"corporate_name", 
		"corporate_name_kana", 
		"department_name", 
		"family_name", 
		"first_name", 
		"family_name_kana", 
		"first_name_kana", 
		"postal_code", 
		"prefecture_name", 
		"address", 
		"address_2", 
		"tel", 
		"mail_address", 
		"reminder_onetime", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"start_time", 
		"end_time", 
		"not_clim_flag", 
		"bpcode", 
	} 
	values := []string{ 
		search.Sid, 
		search.UserId, 
		search.Password, 
		search.BusinessId, 
		search.ActiveFlag, 
		search.CorporateName, 
		search.CorporateNameKana, 
		search.DepartmentName, 
		search.FamilyName, 
		search.FirstName, 
		search.FamilyNameKana, 
		search.FirstNameKana, 
		search.PostalCode, 
		search.PrefectureName, 
		search.Address, 
		search.Address2, 
		search.Tel, 
		search.MailAddress, 
		search.ReminderOnetime, 
		search.DeleteFlag, 
		search.CreateUserId, 
		search.UpdateUserId, 
		search.CreateTime, 
		search.UpdateTime, 
		search.StartTime, 
		search.EndTime, 
		search.NotClimFlag, 
		search.Bpcode, 
	} 
	sql := "UPDATE m_user SET " 
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
func BulkInsertMUser(db *sql.Tx, search []*autodbdto.DBMUserForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	var arr []string 
	columns := []string{ 
		"sid", 
		"user_id", 
		"password", 
		"business_id", 
		"active_flag", 
		"corporate_name", 
		"corporate_name_kana", 
		"department_name", 
		"family_name", 
		"first_name", 
		"family_name_kana", 
		"first_name_kana", 
		"postal_code", 
		"prefecture_name", 
		"address", 
		"address_2", 
		"tel", 
		"mail_address", 
		"reminder_onetime", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"start_time", 
		"end_time", 
		"not_clim_flag", 
		"bpcode", 
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
					arr = append(arr, search[dataInd].UserId) 
				case 2: 
					arr = append(arr, search[dataInd].Password) 
				case 3: 
					arr = append(arr, search[dataInd].BusinessId) 
				case 4: 
					arr = append(arr, search[dataInd].ActiveFlag) 
				case 5: 
					arr = append(arr, search[dataInd].CorporateName) 
				case 6: 
					arr = append(arr, search[dataInd].CorporateNameKana) 
				case 7: 
					arr = append(arr, search[dataInd].DepartmentName) 
				case 8: 
					arr = append(arr, search[dataInd].FamilyName) 
				case 9: 
					arr = append(arr, search[dataInd].FirstName) 
				case 10: 
					arr = append(arr, search[dataInd].FamilyNameKana) 
				case 11: 
					arr = append(arr, search[dataInd].FirstNameKana) 
				case 12: 
					arr = append(arr, search[dataInd].PostalCode) 
				case 13: 
					arr = append(arr, search[dataInd].PrefectureName) 
				case 14: 
					arr = append(arr, search[dataInd].Address) 
				case 15: 
					arr = append(arr, search[dataInd].Address2) 
				case 16: 
					arr = append(arr, search[dataInd].Tel) 
				case 17: 
					arr = append(arr, search[dataInd].MailAddress) 
				case 18: 
					arr = append(arr, search[dataInd].ReminderOnetime) 
				case 19: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 20: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 21: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 22: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 23: 
					arr = append(arr, search[dataInd].UpdateTime) 
				case 24: 
					arr = append(arr, search[dataInd].StartTime) 
				case 25: 
					arr = append(arr, search[dataInd].EndTime) 
				case 26: 
					arr = append(arr, search[dataInd].NotClimFlag) 
				case 27: 
					arr = append(arr, search[dataInd].Bpcode) 
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
func BulkUpdateMUser(db *sql.Tx, search []*autodbdto.DBMUserForm, r *http.Request, excludes []string) (int64, bool) { 
	columnSql := "" 
	valueSql := "" 
	duplicateKeys := "" 
 
	var arr []string 
	columns := []string{ 
		"sid", 
		"user_id", 
		"password", 
		"business_id", 
		"active_flag", 
		"corporate_name", 
		"corporate_name_kana", 
		"department_name", 
		"family_name", 
		"first_name", 
		"family_name_kana", 
		"first_name_kana", 
		"postal_code", 
		"prefecture_name", 
		"address", 
		"address_2", 
		"tel", 
		"mail_address", 
		"reminder_onetime", 
		"delete_flag", 
		"create_user_id", 
		"update_user_id", 
		"create_time", 
		"update_time", 
		"start_time", 
		"end_time", 
		"not_clim_flag", 
		"bpcode", 
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
					arr = append(arr, search[dataInd].UserId) 
				case 2: 
					arr = append(arr, search[dataInd].Password) 
				case 3: 
					arr = append(arr, search[dataInd].BusinessId) 
				case 4: 
					arr = append(arr, search[dataInd].ActiveFlag) 
				case 5: 
					arr = append(arr, search[dataInd].CorporateName) 
				case 6: 
					arr = append(arr, search[dataInd].CorporateNameKana) 
				case 7: 
					arr = append(arr, search[dataInd].DepartmentName) 
				case 8: 
					arr = append(arr, search[dataInd].FamilyName) 
				case 9: 
					arr = append(arr, search[dataInd].FirstName) 
				case 10: 
					arr = append(arr, search[dataInd].FamilyNameKana) 
				case 11: 
					arr = append(arr, search[dataInd].FirstNameKana) 
				case 12: 
					arr = append(arr, search[dataInd].PostalCode) 
				case 13: 
					arr = append(arr, search[dataInd].PrefectureName) 
				case 14: 
					arr = append(arr, search[dataInd].Address) 
				case 15: 
					arr = append(arr, search[dataInd].Address2) 
				case 16: 
					arr = append(arr, search[dataInd].Tel) 
				case 17: 
					arr = append(arr, search[dataInd].MailAddress) 
				case 18: 
					arr = append(arr, search[dataInd].ReminderOnetime) 
				case 19: 
					arr = append(arr, search[dataInd].DeleteFlag) 
				case 20: 
					arr = append(arr, search[dataInd].CreateUserId) 
				case 21: 
					arr = append(arr, search[dataInd].UpdateUserId) 
				case 22: 
					arr = append(arr, search[dataInd].CreateTime) 
				case 23: 
					arr = append(arr, search[dataInd].UpdateTime) 
				case 24: 
					arr = append(arr, search[dataInd].StartTime) 
				case 25: 
					arr = append(arr, search[dataInd].EndTime) 
				case 26: 
					arr = append(arr, search[dataInd].NotClimFlag) 
				case 27: 
					arr = append(arr, search[dataInd].Bpcode) 
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
 
