package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"time"
	"strconv"
	"database/sql/driver"
)

type NullTime struct { 
    Time  time.Time 
    Valid bool // Valid is true if Time is not NULL 
} 
 
// Scan implements the Scanner interface. 
func (nt *NullTime) Scan(value interface{}) error { 
    nt.Time, nt.Valid = value.(time.Time) 
    return nil 
} 
 
// Value implements the driver Valuer interface. 
func (nt NullTime) Value() (driver.Value, error) { 
    if !nt.Valid { 
        return nil, nil 
    } 
    return nt.Time, nil 
} 
type DBHMailLogDTO struct {
	Sid sql.NullInt64 
	Subject sql.NullString 
	FromAddress sql.NullString 
	ToAddress sql.NullString 
	RetrunPath sql.NullString 
	MailText sql.NullString 
	Description sql.NullString 
	SortOrder sql.NullInt64 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBHMailLogForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	Subject string `validate:"required,max=255"` 
	FromAddress string `validate:"max=255"` 
	ToAddress string `validate:"max=255"` 
	RetrunPath string `validate:"max=255"` 
	MailText string 
	Description string 
	SortOrder string `validate:"max=20"` 
	DeleteFlag string `validate:"max=1"` 
	CreateUserId string `validate:"max=20"` 
	UpdateUserId string `validate:"max=20"` 
	CreateTime string 
	UpdateTime string 
}
type DBHMailLogResultDTO struct {
	Token string 
	Status int 
	Form *DBHMailLogForm 
	List []DBHMailLogForm 
	Errors []dto.ErrorForm 
}
func (cf *DBHMailLogForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.Subject:     "subject", 
		&cf.FromAddress:     "from_address", 
		&cf.ToAddress:     "to_address", 
		&cf.RetrunPath:     "retrun_path", 
		&cf.MailText:     "mail_text", 
		&cf.Description:     "description", 
		&cf.SortOrder:     "sort_order", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFHMailLog(dto DBHMailLogDTO) DBHMailLogForm { 
	var form DBHMailLogForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.Subject = "" 
	if dto.Subject.Valid == true { 
		form.Subject = dto.Subject.String 
	} 
	form.FromAddress = "" 
	if dto.FromAddress.Valid == true { 
		form.FromAddress = dto.FromAddress.String 
	} 
	form.ToAddress = "" 
	if dto.ToAddress.Valid == true { 
		form.ToAddress = dto.ToAddress.String 
	} 
	form.RetrunPath = "" 
	if dto.RetrunPath.Valid == true { 
		form.RetrunPath = dto.RetrunPath.String 
	} 
	form.MailText = "" 
	if dto.MailText.Valid == true { 
		form.MailText = dto.MailText.String 
	} 
	form.Description = "" 
	if dto.Description.Valid == true { 
		form.Description = dto.Description.String 
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
