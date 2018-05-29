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
func SearchTInquiry(form *autodbdto.DBTInquiryForm, r *http.Request) []autodbdto.DBTInquiryForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTInquiryForm 
 
	selectResult, _ := automodel.SelectTInquiry(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTInquiry(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTInquiry(form *autodbdto.DBTInquiryForm, r *http.Request) (autodbdto.DBTInquiryForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTInquiry(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFTInquiry(selectResult), true 
	} 
 
	var rform autodbdto.DBTInquiryForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTInquiry(form *autodbdto.DBTInquiryForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKTInquiryForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTInquiry(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertTInquiry(tx, form, r, []string{"sid"}) 
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
