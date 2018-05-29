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
func SearchMShop(form *autodbdto.DBMShopForm, r *http.Request) []autodbdto.DBMShopForm { 
	common.WriteLog(config.INFO, "Start", r) 
	var resultList []autodbdto.DBMShopForm 
 
	selectResult, _ := automodel.SelectMShop(form, r) 
 
	for ind, _ := range selectResult { 
		resultList = append(resultList, autodbdto.DTFMShop(selectResult[ind])) 
	} 
 
	return resultList 
 
} 
 
/* 
 * PKから情報取得 
 */ 
func GetByPkMShop(form *autodbdto.DBMShopForm, r *http.Request) (autodbdto.DBMShopForm, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	selectResult, hit, _ := automodel.GetByPKMShop(form.Sid, r) 
 
	if hit { 
		return autodbdto.DTFMShop(selectResult), true 
	} 
 
	var rform autodbdto.DBMShopForm 
	return rform, false 
 
} 
 
/* 
 * 登録 
 */ 
func RegisterMShop(form *autodbdto.DBMShopForm, r *http.Request) (string, bool) { 
	common.WriteLog(config.INFO, "Start", r) 
 
	db := db.DbConn() 
	tx, _ := db.Begin() 
 
	hit := false 
	sid := "" 
	if form.Sid != ""  { 
		_, hit, _ = automodel.GetByPKMShopForUpdate(tx,form.Sid, r) 
	} 
 
	lid, writed := int64(0), false 
	if hit { 
		writed = automodel.UpdateMShop(tx, form, r, []string{"sid"}) 
		sid = form.Sid 
	} else { 
		lid, writed = automodel.InsertMShop(tx, form, r, []string{"sid"}) 
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
