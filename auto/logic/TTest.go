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
func SearchTTest(form *autodbdto.DBTTestForm, r *http.Request) []autodbdto.DBTTestForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTTestForm 
 
	selectResult, _ := automodel.SelectTTest(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTTest(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTTest(form *autodbdto.DBTTestForm, r *http.Request) (autodbdto.DBTTestForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTTest(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFTTest(selectResult), true 
	} 
 
	var rform autodbdto.DBTTestForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTTest(form *autodbdto.DBTTestForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKTTestForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTTest(tx, form, r, []string{"sid"}) 
	} else { 
		lid, writed = automodel.InsertTTest(tx, form, r, []string{"sid"}) 
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
