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
func SearchMIshiguro(form *autodbdto.DBMIshiguroForm, r *http.Request) []autodbdto.DBMIshiguroForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBMIshiguroForm 
 
	selectResult, _ := automodel.SelectMIshiguro(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFMIshiguro(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkMIshiguro(form *autodbdto.DBMIshiguroForm, r *http.Request) (autodbdto.DBMIshiguroForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKMIshiguro(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFMIshiguro(selectResult), true 
	} 
 
	var rform autodbdto.DBMIshiguroForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterMIshiguro(form *autodbdto.DBMIshiguroForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKMIshiguroForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateMIshiguro(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertMIshiguro(tx, form, r, []string{"sid"}) 
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
