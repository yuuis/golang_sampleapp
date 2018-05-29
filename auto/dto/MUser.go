package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMUserDTO struct {
	Sid sql.NullInt64 
	UserId sql.NullString 
	Password sql.NullString 
	BusinessId sql.NullInt64 
	ActiveFlag sql.NullInt64 
	CorporateName sql.NullString 
	CorporateNameKana sql.NullString 
	DepartmentName sql.NullString 
	FamilyName sql.NullString 
	FirstName sql.NullString 
	FamilyNameKana sql.NullString 
	FirstNameKana sql.NullString 
	PostalCode sql.NullString 
	PrefectureName sql.NullString 
	Address sql.NullString 
	Address2 sql.NullString 
	Tel sql.NullString 
	MailAddress sql.NullString 
	ReminderOnetime sql.NullString 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
	StartTime NullTime 
	EndTime NullTime 
	NotClimFlag sql.NullInt64 
	Bpcode sql.NullString 
}
type DBMUserForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	UserId string `validate:"required,max=128"` 
	Password string `validate:"required,max=256"` 
	BusinessId string `validate:"required,max=20"` 
	ActiveFlag string `validate:"required,max=1"` 
	CorporateName string `validate:"required,max=80"` 
	CorporateNameKana string `validate:"required,max=80"` 
	DepartmentName string `validate:"required,max=80"` 
	FamilyName string `validate:"required,max=80"` 
	FirstName string `validate:"required,max=80"` 
	FamilyNameKana string `validate:"required,max=80"` 
	FirstNameKana string `validate:"required,max=80"` 
	PostalCode string `validate:"required,max=7"` 
	PrefectureName string `validate:"required,max=10"` 
	Address string `validate:"required,max=255"` 
	Address2 string `validate:"required,max=255"` 
	Tel string `validate:"required,max=20"` 
	MailAddress string `validate:"required,max=80"` 
	ReminderOnetime string `validate:"required,max=80"` 
	DeleteFlag string `validate:"required,max=1"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateTime string `validate:"required"` 
	UpdateTime string `validate:"required"` 
	StartTime string 
	EndTime string 
	NotClimFlag string `validate:"max=4"` 
	Bpcode string `validate:"max=50"` 
}
type DBMUserResultDTO struct {
	Token string 
	Status int 
	Form *DBMUserForm 
	List []DBMUserForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMUserForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.UserId:     "user_id", 
		&cf.Password:     "password", 
		&cf.BusinessId:     "business_id", 
		&cf.ActiveFlag:     "active_flag", 
		&cf.CorporateName:     "corporate_name", 
		&cf.CorporateNameKana:     "corporate_name_kana", 
		&cf.DepartmentName:     "department_name", 
		&cf.FamilyName:     "family_name", 
		&cf.FirstName:     "first_name", 
		&cf.FamilyNameKana:     "family_name_kana", 
		&cf.FirstNameKana:     "first_name_kana", 
		&cf.PostalCode:     "postal_code", 
		&cf.PrefectureName:     "prefecture_name", 
		&cf.Address:     "address", 
		&cf.Address2:     "address_2", 
		&cf.Tel:     "tel", 
		&cf.MailAddress:     "mail_address", 
		&cf.ReminderOnetime:     "reminder_onetime", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
		&cf.StartTime:     "start_time", 
		&cf.EndTime:     "end_time", 
		&cf.NotClimFlag:     "not_clim_flag", 
		&cf.Bpcode:     "bpcode", 
	} 
} 
func DTFMUser(dto DBMUserDTO) DBMUserForm { 
	var form DBMUserForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.UserId = "" 
	if dto.UserId.Valid == true { 
		form.UserId = dto.UserId.String 
	} 
	form.Password = "" 
	if dto.Password.Valid == true { 
		form.Password = dto.Password.String 
	} 
	form.BusinessId = "" 
	if dto.BusinessId.Valid == true { 
		form.BusinessId = strconv.FormatInt(dto.BusinessId.Int64, 10) 
	} 
	form.ActiveFlag = "" 
	if dto.ActiveFlag.Valid == true { 
		form.ActiveFlag = strconv.FormatInt(dto.ActiveFlag.Int64, 10) 
	} 
	form.CorporateName = "" 
	if dto.CorporateName.Valid == true { 
		form.CorporateName = dto.CorporateName.String 
	} 
	form.CorporateNameKana = "" 
	if dto.CorporateNameKana.Valid == true { 
		form.CorporateNameKana = dto.CorporateNameKana.String 
	} 
	form.DepartmentName = "" 
	if dto.DepartmentName.Valid == true { 
		form.DepartmentName = dto.DepartmentName.String 
	} 
	form.FamilyName = "" 
	if dto.FamilyName.Valid == true { 
		form.FamilyName = dto.FamilyName.String 
	} 
	form.FirstName = "" 
	if dto.FirstName.Valid == true { 
		form.FirstName = dto.FirstName.String 
	} 
	form.FamilyNameKana = "" 
	if dto.FamilyNameKana.Valid == true { 
		form.FamilyNameKana = dto.FamilyNameKana.String 
	} 
	form.FirstNameKana = "" 
	if dto.FirstNameKana.Valid == true { 
		form.FirstNameKana = dto.FirstNameKana.String 
	} 
	form.PostalCode = "" 
	if dto.PostalCode.Valid == true { 
		form.PostalCode = dto.PostalCode.String 
	} 
	form.PrefectureName = "" 
	if dto.PrefectureName.Valid == true { 
		form.PrefectureName = dto.PrefectureName.String 
	} 
	form.Address = "" 
	if dto.Address.Valid == true { 
		form.Address = dto.Address.String 
	} 
	form.Address2 = "" 
	if dto.Address2.Valid == true { 
		form.Address2 = dto.Address2.String 
	} 
	form.Tel = "" 
	if dto.Tel.Valid == true { 
		form.Tel = dto.Tel.String 
	} 
	form.MailAddress = "" 
	if dto.MailAddress.Valid == true { 
		form.MailAddress = dto.MailAddress.String 
	} 
	form.ReminderOnetime = "" 
	if dto.ReminderOnetime.Valid == true { 
		form.ReminderOnetime = dto.ReminderOnetime.String 
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
	form.StartTime = "" 
	if dto.StartTime.Valid == true { 
		form.StartTime = dto.StartTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.EndTime = "" 
	if dto.EndTime.Valid == true { 
		form.EndTime = dto.EndTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.NotClimFlag = "" 
	if dto.NotClimFlag.Valid == true { 
		form.NotClimFlag = strconv.FormatInt(dto.NotClimFlag.Int64, 10) 
	} 
	form.Bpcode = "" 
	if dto.Bpcode.Valid == true { 
		form.Bpcode = dto.Bpcode.String 
	} 
	return form 
} 
