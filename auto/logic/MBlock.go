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
func SearchMBlock(form *autodbdto.DBMBlockForm, r *http.Request) []autodbdto.DBMBlockForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBMBlockForm 
 
	selectResult, _ := automodel.SelectMBlock(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFMBlock(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkMBlock(form *autodbdto.DBMBlockForm, r *http.Request) (autodbdto.DBMBlockForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKMBlock(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFMBlock(selectResult), true 
	} 
 
	var rform autodbdto.DBMBlockForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterMBlock(form *autodbdto.DBMBlockForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKMBlockForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateMBlock(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertMBlock(tx, form, r, []string{"sid"}) 
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
