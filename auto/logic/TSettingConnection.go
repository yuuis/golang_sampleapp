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
func SearchTSettingConnection(form *autodbdto.DBTSettingConnectionForm, r *http.Request) []autodbdto.DBTSettingConnectionForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTSettingConnectionForm 
 
	selectResult, _ := automodel.SelectTSettingConnection(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTSettingConnection(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTSettingConnection(form *autodbdto.DBTSettingConnectionForm, r *http.Request) (autodbdto.DBTSettingConnectionForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTSettingConnection(form.MSettingSid, r) 
 
	if hit { 
		return autodbdto.DTFTSettingConnection(selectResult), true 
	} 
 
	var rform autodbdto.DBTSettingConnectionForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTSettingConnection(form *autodbdto.DBTSettingConnectionForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.MSettingSid != ""  { 
		_, hit, _ = automodel.GetByPKTSettingConnectionForUpdate(tx,form.MSettingSid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTSettingConnection(tx, form, r, []string{"sid"}) 
	} else { 
		lid, writed = automodel.InsertTSettingConnection(tx, form, r, []string{"sid"}) 
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
