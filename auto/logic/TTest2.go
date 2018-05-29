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
func SearchTTest2(form *autodbdto.DBTTest2Form, r *http.Request) []autodbdto.DBTTest2Form { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTTest2Form 
 
	selectResult, _ := automodel.SelectTTest2(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTTest2(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTTest2(form *autodbdto.DBTTest2Form, r *http.Request) (autodbdto.DBTTest2Form, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTTest2(form.Code1, form.Code2, r) 
 
	if hit { 
		return autodbdto.DTFTTest2(selectResult), true 
	} 
 
	var rform autodbdto.DBTTest2Form 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTTest2(form *autodbdto.DBTTest2Form, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Code1 != ""  && form.Code2 != ""  { 
		_, hit, _ = automodel.GetByPKTTest2ForUpdate(tx,form.Code1, form.Code2, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTTest2(tx, form, r, []string{"sid"}) 
	} else { 
		lid, writed = automodel.InsertTTest2(tx, form, r, []string{"sid"}) 
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
