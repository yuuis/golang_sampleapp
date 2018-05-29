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
func SearchMSettingConnectionMessage(form *autodbdto.DBMSettingConnectionMessageForm, r *http.Request) []autodbdto.DBMSettingConnectionMessageForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBMSettingConnectionMessageForm 
 
	selectResult, _ := automodel.SelectMSettingConnectionMessage(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFMSettingConnectionMessage(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkMSettingConnectionMessage(form *autodbdto.DBMSettingConnectionMessageForm, r *http.Request) (autodbdto.DBMSettingConnectionMessageForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKMSettingConnectionMessage(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFMSettingConnectionMessage(selectResult), true 
	} 
 
	var rform autodbdto.DBMSettingConnectionMessageForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterMSettingConnectionMessage(form *autodbdto.DBMSettingConnectionMessageForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKMSettingConnectionMessageForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateMSettingConnectionMessage(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertMSettingConnectionMessage(tx, form, r, []string{"sid"}) 
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
