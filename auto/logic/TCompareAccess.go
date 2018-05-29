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
func SearchTCompareAccess(form *autodbdto.DBTCompareAccessForm, r *http.Request) []autodbdto.DBTCompareAccessForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTCompareAccessForm 
 
	selectResult, _ := automodel.SelectTCompareAccess(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTCompareAccess(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTCompareAccess(form *autodbdto.DBTCompareAccessForm, r *http.Request) (autodbdto.DBTCompareAccessForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTCompareAccess(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFTCompareAccess(selectResult), true 
	} 
 
	var rform autodbdto.DBTCompareAccessForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTCompareAccess(form *autodbdto.DBTCompareAccessForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKTCompareAccessForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTCompareAccess(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertTCompareAccess(tx, form, r, []string{"sid"}) 
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
