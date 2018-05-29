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
func SearchTBeaconAccess(form *autodbdto.DBTBeaconAccessForm, r *http.Request) []autodbdto.DBTBeaconAccessForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBTBeaconAccessForm 
 
	selectResult, _ := automodel.SelectTBeaconAccess(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFTBeaconAccess(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkTBeaconAccess(form *autodbdto.DBTBeaconAccessForm, r *http.Request) (autodbdto.DBTBeaconAccessForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKTBeaconAccess(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFTBeaconAccess(selectResult), true 
	} 
 
	var rform autodbdto.DBTBeaconAccessForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterTBeaconAccess(form *autodbdto.DBTBeaconAccessForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKTBeaconAccessForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateTBeaconAccess(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertTBeaconAccess(tx, form, r, []string{"sid"}) 
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
