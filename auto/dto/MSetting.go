package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMSettingDTO struct {
	Sid sql.NullInt64 
	BusinessId sql.NullInt64 
	MUserSid sql.NullInt64 
	MShopSid sql.NullInt64 
	Key sql.NullString 
	Context sql.NullString 
	ActiveFlag sql.NullInt64 
	UrlMatchingPattern sql.NullInt64 
	Status sql.NullInt64 
	ScreenshotGetTime NullTime 
	RegularlyItemCode sql.NullString 
	SortOrder sql.NullInt64 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBMSettingForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	BusinessId string `validate:"required,max=20"` 
	MUserSid string `validate:"required,max=20"` 
	MShopSid string `validate:"max=20"` 
	Key string `validate:"required,max=255"` 
	Context string 
	ActiveFlag string `validate:"required,max=1"` 
	UrlMatchingPattern string `validate:"required,max=1"` 
	Status string `validate:"required,max=1"` 
	ScreenshotGetTime string 
	RegularlyItemCode string `validate:"max=255"` 
	SortOrder string `validate:"max=20"` 
	DeleteFlag string `validate:"required,max=1"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateTime string `validate:"required"` 
	UpdateTime string `validate:"required"` 
}
type DBMSettingResultDTO struct {
	Token string 
	Status int 
	Form *DBMSettingForm 
	List []DBMSettingForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMSettingForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.BusinessId:     "business_id", 
		&cf.MUserSid:     "m_user_sid", 
		&cf.MShopSid:     "m_shop_sid", 
		&cf.Key:     "key", 
		&cf.Context:     "context", 
		&cf.ActiveFlag:     "active_flag", 
		&cf.UrlMatchingPattern:     "url_matching_pattern", 
		&cf.Status:     "status", 
		&cf.ScreenshotGetTime:     "screenshot_get_time", 
		&cf.RegularlyItemCode:     "regularly_item_code", 
		&cf.SortOrder:     "sort_order", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFMSetting(dto DBMSettingDTO) DBMSettingForm { 
	var form DBMSettingForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.BusinessId = "" 
	if dto.BusinessId.Valid == true { 
		form.BusinessId = strconv.FormatInt(dto.BusinessId.Int64, 10) 
	} 
	form.MUserSid = "" 
	if dto.MUserSid.Valid == true { 
		form.MUserSid = strconv.FormatInt(dto.MUserSid.Int64, 10) 
	} 
	form.MShopSid = "" 
	if dto.MShopSid.Valid == true { 
		form.MShopSid = strconv.FormatInt(dto.MShopSid.Int64, 10) 
	} 
	form.Key = "" 
	if dto.Key.Valid == true { 
		form.Key = dto.Key.String 
	} 
	form.Context = "" 
	if dto.Context.Valid == true { 
		form.Context = dto.Context.String 
	} 
	form.ActiveFlag = "" 
	if dto.ActiveFlag.Valid == true { 
		form.ActiveFlag = strconv.FormatInt(dto.ActiveFlag.Int64, 10) 
	} 
	form.UrlMatchingPattern = "" 
	if dto.UrlMatchingPattern.Valid == true { 
		form.UrlMatchingPattern = strconv.FormatInt(dto.UrlMatchingPattern.Int64, 10) 
	} 
	form.Status = "" 
	if dto.Status.Valid == true { 
		form.Status = strconv.FormatInt(dto.Status.Int64, 10) 
	} 
	form.ScreenshotGetTime = "" 
	if dto.ScreenshotGetTime.Valid == true { 
		form.ScreenshotGetTime = dto.ScreenshotGetTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.RegularlyItemCode = "" 
	if dto.RegularlyItemCode.Valid == true { 
		form.RegularlyItemCode = dto.RegularlyItemCode.String 
	} 
	form.SortOrder = "" 
	if dto.SortOrder.Valid == true { 
		form.SortOrder = strconv.FormatInt(dto.SortOrder.Int64, 10) 
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
