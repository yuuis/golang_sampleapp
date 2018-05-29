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
func SearchTMemo(form *autodbdto.DBTMemoForm, r *http.Request) []autodbdto.DBTMemoForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTMemoForm 
 
	selectResult, _ := automodel.SelectTMemo(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTMemo(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTMemo(form *autodbdto.DBTMemoForm, r *http.Request) (autodbdto.DBTMemoForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTMemo(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFTMemo(selectResult), true 
	} 
 
	var rform autodbdto.DBTMemoForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTMemo(form *autodbdto.DBTMemoForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKTMemoForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTMemo(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertTMemo(tx, form, r, []string{"sid"}) 
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
