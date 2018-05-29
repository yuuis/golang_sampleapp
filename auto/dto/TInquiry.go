package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBTInquiryDTO struct {
	Sid sql.NullInt64 
	MUserSid sql.NullInt64 
	InquiryDate NullTime 
	CategoryCode sql.NullString 
	Tel sql.NullString 
	Name sql.NullString 
	MailAddress sql.NullString 
	InquiryText sql.NullString 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBTInquiryForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	MUserSid string `validate:"required,max=20"` 
	InquiryDate string 
	CategoryCode string `validate:"max=20"` 
	Tel string `validate:"max=20"` 
	Name string `validate:"max=80"` 
	MailAddress string `validate:"max=255"` 
	InquiryText string `validate:"required,max=1024"` 
	DeleteFlag string `validate:"max=1"` 
	CreateUserId string `validate:"max=20"` 
	UpdateUserId string `validate:"max=20"` 
	CreateTime string 
	UpdateTime string 
}
type DBTInquiryResultDTO struct {
	Token string 
	Status int 
	Form *DBTInquiryForm 
	List []DBTInquiryForm 
	Errors []dto.ErrorForm 
}
func (cf *DBTInquiryForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.MUserSid:     "m_user_sid", 
		&cf.InquiryDate:     "inquiry_date", 
		&cf.CategoryCode:     "category_code", 
		&cf.Tel:     "tel", 
		&cf.Name:     "name", 
		&cf.MailAddress:     "mail_address", 
		&cf.InquiryText:     "inquiry_text", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFTInquiry(dto DBTInquiryDTO) DBTInquiryForm { 
	var form DBTInquiryForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.MUserSid = "" 
	if dto.MUserSid.Valid == true { 
		form.MUserSid = strconv.FormatInt(dto.MUserSid.Int64, 10) 
	} 
	form.InquiryDate = "" 
	if dto.InquiryDate.Valid == true { 
		form.InquiryDate = dto.InquiryDate.Time.Format("2006/01/02 15:04:05") 
	} 
	form.CategoryCode = "" 
	if dto.CategoryCode.Valid == true { 
		form.CategoryCode = dto.CategoryCode.String 
	} 
	form.Tel = "" 
	if dto.Tel.Valid == true { 
		form.Tel = dto.Tel.String 
	} 
	form.Name = "" 
	if dto.Name.Valid == true { 
		form.Name = dto.Name.String 
	} 
	form.MailAddress = "" 
	if dto.MailAddress.Valid == true { 
		form.MailAddress = dto.MailAddress.String 
	} 
	form.InquiryText = "" 
	if dto.InquiryText.Valid == true { 
		form.InquiryText = dto.InquiryText.String 
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
