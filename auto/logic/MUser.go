package autologic 
 
import ( 
	"../../config" 
	"../../service/common" 
	"../../service/db" 
	"../dto" 
	"../model" 
	"net/http" 
	"strconv" 
) 
 
/* 
 * 検索 
 */ 
func SearchMUser(form *autodbdto.DBMUserForm, r *http.Request) []autodbdto.DBMUserForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBMUserForm 
 
	selectResult, _ := automodel.SelectMUser(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFMUser(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkMUser(form *autodbdto.DBMUserForm, r *http.Request) (autodbdto.DBMUserForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKMUser(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFMUser(selectResult), true 
	} 
 
	var rform autodbdto.DBMUserForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterMUser(form *autodbdto.DBMUserForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKMUserForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateMUser(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertMUser(tx, form, r, []string{"sid"}) 
		sid = strconv.FormatInt(lid, 10) 
	} 
 
	if writed { 
		tx.Commit() 
		return sid, true 
	} else { 
		tx.Rollback() 
		return "", false 
	} 
 
} 
