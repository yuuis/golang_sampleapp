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
func SearchRFirstConversion(form *autodbdto.DBRFirstConversionForm, r *http.Request) []autodbdto.DBRFirstConversionForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBRFirstConversionForm 
 
	selectResult, _ := automodel.SelectRFirstConversion(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFRFirstConversion(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkRFirstConversion(form *autodbdto.DBRFirstConversionForm, r *http.Request) (autodbdto.DBRFirstConversionForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKRFirstConversion(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFRFirstConversion(selectResult), true 
	} 
 
	var rform autodbdto.DBRFirstConversionForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterRFirstConversion(form *autodbdto.DBRFirstConversionForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKRFirstConversionForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateRFirstConversion(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertRFirstConversion(tx, form, r, []string{"sid"}) 
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
