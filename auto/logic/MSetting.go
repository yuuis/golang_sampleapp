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
func SearchMSetting(form *autodbdto.DBMSettingForm, r *http.Request) []autodbdto.DBMSettingForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBMSettingForm 
 
	selectResult, _ := automodel.SelectMSetting(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFMSetting(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkMSetting(form *autodbdto.DBMSettingForm, r *http.Request) (autodbdto.DBMSettingForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKMSetting(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFMSetting(selectResult), true 
	} 
 
	var rform autodbdto.DBMSettingForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterMSetting(form *autodbdto.DBMSettingForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKMSettingForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateMSetting(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertMSetting(tx, form, r, []string{"sid"}) 
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
