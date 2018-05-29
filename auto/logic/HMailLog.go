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
func SearchHMailLog(form *autodbdto.DBHMailLogForm, r *http.Request) []autodbdto.DBHMailLogForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBHMailLogForm 
 
	selectResult, _ := automodel.SelectHMailLog(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFHMailLog(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkHMailLog(form *autodbdto.DBHMailLogForm, r *http.Request) (autodbdto.DBHMailLogForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKHMailLog(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFHMailLog(selectResult), true 
	} 
 
	var rform autodbdto.DBHMailLogForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterHMailLog(form *autodbdto.DBHMailLogForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKHMailLogForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateHMailLog(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertHMailLog(tx, form, r, []string{"sid"}) 
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
