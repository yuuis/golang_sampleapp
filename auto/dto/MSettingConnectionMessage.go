package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMSettingConnectionMessageDTO struct {
	Sid sql.NullInt64 
	LastComparedTimeFlag sql.NullInt64 
	LastXdShopTimeFlag sql.NullInt64 
	LastXdCartTimeFlag sql.NullInt64 
	LastConvertionTimeFlag sql.NullInt64 
	StatusMessage sql.NullString 
	DetailMessage sql.NullString 
	RemarkMessage sql.NullString 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBMSettingConnectionMessageForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	LastComparedTimeFlag string `validate:"max=1"` 
	LastXdShopTimeFlag string `validate:"max=1"` 
	LastXdCartTimeFlag string `validate:"max=1"` 
	LastConvertionTimeFlag string `validate:"max=1"` 
	StatusMessage string `validate:"max=255"` 
	DetailMessage string `validate:"max=1024"` 
	RemarkMessage string `validate:"max=1024"` 
	DeleteFlag string `validate:"max=1"` 
	CreateUserId string `validate:"max=20"` 
	UpdateUserId string `validate:"max=20"` 
	CreateTime string 
	UpdateTime string 
}
type DBMSettingConnectionMessageResultDTO struct {
	Token string 
	Status int 
	Form *DBMSettingConnectionMessageForm 
	List []DBMSettingConnectionMessageForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMSettingConnectionMessageForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.LastComparedTimeFlag:     "last_compared_time_flag", 
		&cf.LastXdShopTimeFlag:     "last_xd_shop_time_flag", 
		&cf.LastXdCartTimeFlag:     "last_xd_cart_time_flag", 
		&cf.LastConvertionTimeFlag:     "last_convertion_time_flag", 
		&cf.StatusMessage:     "status_message", 
		&cf.DetailMessage:     "detail_message", 
		&cf.RemarkMessage:     "remark_message", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFMSettingConnectionMessage(dto DBMSettingConnectionMessageDTO) DBMSettingConnectionMessageForm { 
	var form DBMSettingConnectionMessageForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.LastComparedTimeFlag = "" 
	if dto.LastComparedTimeFlag.Valid == true { 
		form.LastComparedTimeFlag = strconv.FormatInt(dto.LastComparedTimeFlag.Int64, 10) 
	} 
	form.LastXdShopTimeFlag = "" 
	if dto.LastXdShopTimeFlag.Valid == true { 
		form.LastXdShopTimeFlag = strconv.FormatInt(dto.LastXdShopTimeFlag.Int64, 10) 
	} 
	form.LastXdCartTimeFlag = "" 
	if dto.LastXdCartTimeFlag.Valid == true { 
		form.LastXdCartTimeFlag = strconv.FormatInt(dto.LastXdCartTimeFlag.Int64, 10) 
	} 
	form.LastConvertionTimeFlag = "" 
	if dto.LastConvertionTimeFlag.Valid == true { 
		form.LastConvertionTimeFlag = strconv.FormatInt(dto.LastConvertionTimeFlag.Int64, 10) 
	} 
	form.StatusMessage = "" 
	if dto.StatusMessage.Valid == true { 
		form.StatusMessage = dto.StatusMessage.String 
	} 
	form.DetailMessage = "" 
	if dto.DetailMessage.Valid == true { 
		form.DetailMessage = dto.DetailMessage.String 
	} 
	form.RemarkMessage = "" 
	if dto.RemarkMessage.Valid == true { 
		form.RemarkMessage = dto.RemarkMessage.String 
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
