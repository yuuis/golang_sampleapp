package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBTMemoDTO struct {
	Sid sql.NullInt64 
	MSettingSid sql.NullInt64 
	InputTime NullTime 
	Memo sql.NullString 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBTMemoForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	MSettingSid string `validate:"required,max=20"` 
	InputTime string `validate:"required"` 
	Memo string `validate:"required"` 
	DeleteFlag string `validate:"max=1"` 
	CreateUserId string `validate:"max=20"` 
	UpdateUserId string `validate:"max=20"` 
	CreateTime string 
	UpdateTime string 
}
type DBTMemoResultDTO struct {
	Token string 
	Status int 
	Form *DBTMemoForm 
	List []DBTMemoForm 
	Errors []dto.ErrorForm 
}
func (cf *DBTMemoForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.MSettingSid:     "m_setting_sid", 
		&cf.InputTime:     "input_time", 
		&cf.Memo:     "memo", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFTMemo(dto DBTMemoDTO) DBTMemoForm { 
	var form DBTMemoForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.MSettingSid = "" 
	if dto.MSettingSid.Valid == true { 
		form.MSettingSid = strconv.FormatInt(dto.MSettingSid.Int64, 10) 
	} 
	form.InputTime = "" 
	if dto.InputTime.Valid == true { 
		form.InputTime = dto.InputTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.Memo = "" 
	if dto.Memo.Valid == true { 
		form.Memo = dto.Memo.String 
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
