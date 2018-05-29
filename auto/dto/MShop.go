package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMShopDTO struct {
	Sid sql.NullInt64 
	ShopId sql.NullString 
	BusinessId sql.NullInt64 
	ShopName sql.NullString 
	ManagerAuthKey sql.NullString 
	LastAccessTime NullTime 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBMShopForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	ShopId string `validate:"required,max=128"` 
	BusinessId string `validate:"required,max=20"` 
	ShopName string `validate:"max=255"` 
	ManagerAuthKey string `validate:"required,max=128"` 
	LastAccessTime string 
	DeleteFlag string `validate:"required,max=1"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateTime string `validate:"required"` 
	UpdateTime string `validate:"required"` 
}
type DBMShopResultDTO struct {
	Token string 
	Status int 
	Form *DBMShopForm 
	List []DBMShopForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMShopForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.ShopId:     "shop_id", 
		&cf.BusinessId:     "business_id", 
		&cf.ShopName:     "shop_name", 
		&cf.ManagerAuthKey:     "manager_auth_key", 
		&cf.LastAccessTime:     "last_access_time", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFMShop(dto DBMShopDTO) DBMShopForm { 
	var form DBMShopForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.ShopId = "" 
	if dto.ShopId.Valid == true { 
		form.ShopId = dto.ShopId.String 
	} 
	form.BusinessId = "" 
	if dto.BusinessId.Valid == true { 
		form.BusinessId = strconv.FormatInt(dto.BusinessId.Int64, 10) 
	} 
	form.ShopName = "" 
	if dto.ShopName.Valid == true { 
		form.ShopName = dto.ShopName.String 
	} 
	form.ManagerAuthKey = "" 
	if dto.ManagerAuthKey.Valid == true { 
		form.ManagerAuthKey = dto.ManagerAuthKey.String 
	} 
	form.LastAccessTime = "" 
	if dto.LastAccessTime.Valid == true { 
		form.LastAccessTime = dto.LastAccessTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.DeleteFlag = "" 
	if dto.DeleteFlag.Valid == true { 
		form.DeleteFlag = strconv.FormatInt(dto.DeleteFlag.Int64, 10) 
	} 
	form.CreateUserId = "" 
	if dto.CreateUserId.Valid == true { 
		form.CreateUserId = strconv.FormatInt(dto.CreateUserId.Int64, 10) 
	} 
	form.UpdateUserId = "" 
	if dto.UpdateUserId.Valid == true { 
		form.UpdateUserId = strconv.FormatInt(dto.UpdateUserId.Int64, 10) 
	} 
	form.CreateTime = "" 
	if dto.CreateTime.Valid == true { 
		form.CreateTime = dto.CreateTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.UpdateTime = "" 
	if dto.UpdateTime.Valid == true { 
		form.UpdateTime = dto.UpdateTime.Time.Format("2006/01/02 15:04:05") 
	} 
	return form 
} 
