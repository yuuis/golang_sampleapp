package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMBlockDTO struct {
	Sid sql.NullInt64 
	BusinessId sql.NullInt64 
	MUserSid sql.NullInt64 
	TargetNo sql.NullInt64 
	IpAddress sql.NullString 
	SortOrder sql.NullInt64 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBMBlockForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	BusinessId string `validate:"required,max=20"` 
	MUserSid string `validate:"required,max=20"` 
	TargetNo string `validate:"required,max=20"` 
	IpAddress string `validate:"required,max=255"` 
	SortOrder string `validate:"max=20"` 
	DeleteFlag string `validate:"required,max=1"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateTime string `validate:"required"` 
	UpdateTime string `validate:"required"` 
}
type DBMBlockResultDTO struct {
	Token string 
	Status int 
	Form *DBMBlockForm 
	List []DBMBlockForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMBlockForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.BusinessId:     "business_id", 
		&cf.MUserSid:     "m_user_sid", 
		&cf.TargetNo:     "target_no", 
		&cf.IpAddress:     "ip_address", 
		&cf.SortOrder:     "sort_order", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFMBlock(dto DBMBlockDTO) DBMBlockForm { 
	var form DBMBlockForm 
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
	form.TargetNo = "" 
	if dto.TargetNo.Valid == true { 
		form.TargetNo = strconv.FormatInt(dto.TargetNo.Int64, 10) 
	} 
	form.IpAddress = "" 
	if dto.IpAddress.Valid == true { 
		form.IpAddress = dto.IpAddress.String 
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
