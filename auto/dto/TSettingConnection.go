package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBTSettingConnectionDTO struct {
	MSettingSid sql.NullInt64 
	LastComparedTime NullTime 
	LastXdShopTime NullTime 
	LastXdCartTime NullTime 
	LastConvertionTime NullTime 
	LastComparedTimeSp NullTime 
	LastXdShopTimeSp NullTime 
	LastXdCartTimeSp NullTime 
	LastConvertionTimeSp NullTime 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBTSettingConnectionForm struct {
	Mode string 
	MSettingSid string `validate:"required,max=20"` 
	LastComparedTime string 
	LastXdShopTime string 
	LastXdCartTime string 
	LastConvertionTime string 
	LastComparedTimeSp string 
	LastXdShopTimeSp string 
	LastXdCartTimeSp string 
	LastConvertionTimeSp string 
	DeleteFlag string `validate:"max=1"` 
	CreateUserId string `validate:"max=20"` 
	UpdateUserId string `validate:"max=20"` 
	CreateTime string 
	UpdateTime string 
}
type DBTSettingConnectionResultDTO struct {
	Token string 
	Status int 
	Form *DBTSettingConnectionForm 
	List []DBTSettingConnectionForm 
	Errors []dto.ErrorForm 
}
func (cf *DBTSettingConnectionForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.MSettingSid:     "m_setting_sid", 
		&cf.LastComparedTime:     "last_compared_time", 
		&cf.LastXdShopTime:     "last_xd_shop_time", 
		&cf.LastXdCartTime:     "last_xd_cart_time", 
		&cf.LastConvertionTime:     "last_convertion_time", 
		&cf.LastComparedTimeSp:     "last_compared_time_sp", 
		&cf.LastXdShopTimeSp:     "last_xd_shop_time_sp", 
		&cf.LastXdCartTimeSp:     "last_xd_cart_time_sp", 
		&cf.LastConvertionTimeSp:     "last_convertion_time_sp", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFTSettingConnection(dto DBTSettingConnectionDTO) DBTSettingConnectionForm { 
	var form DBTSettingConnectionForm 
	form.MSettingSid = "" 
	if dto.MSettingSid.Valid == true { 
		form.MSettingSid = strconv.FormatInt(dto.MSettingSid.Int64, 10) 
	} 
	form.LastComparedTime = "" 
	if dto.LastComparedTime.Valid == true { 
		form.LastComparedTime = dto.LastComparedTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastXdShopTime = "" 
	if dto.LastXdShopTime.Valid == true { 
		form.LastXdShopTime = dto.LastXdShopTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastXdCartTime = "" 
	if dto.LastXdCartTime.Valid == true { 
		form.LastXdCartTime = dto.LastXdCartTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastConvertionTime = "" 
	if dto.LastConvertionTime.Valid == true { 
		form.LastConvertionTime = dto.LastConvertionTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastComparedTimeSp = "" 
	if dto.LastComparedTimeSp.Valid == true { 
		form.LastComparedTimeSp = dto.LastComparedTimeSp.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastXdShopTimeSp = "" 
	if dto.LastXdShopTimeSp.Valid == true { 
		form.LastXdShopTimeSp = dto.LastXdShopTimeSp.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastXdCartTimeSp = "" 
	if dto.LastXdCartTimeSp.Valid == true { 
		form.LastXdCartTimeSp = dto.LastXdCartTimeSp.Time.Format("2006/01/02 15:04:05") 
	} 
	form.LastConvertionTimeSp = "" 
	if dto.LastConvertionTimeSp.Valid == true { 
		form.LastConvertionTimeSp = dto.LastConvertionTimeSp.Time.Format("2006/01/02 15:04:05") 
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
