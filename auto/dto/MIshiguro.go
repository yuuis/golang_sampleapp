package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMIshiguroDTO struct {
	Sid sql.NullInt64 
	Name sql.NullString 
	Code sql.NullString 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBMIshiguroForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	Name string `validate:"required,max=255"` 
	Code string `validate:"required,max=255"` 
	DeleteFlag string `validate:"required,max=1"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateTime string `validate:"required"` 
	UpdateTime string `validate:"required"` 
}
type DBMIshiguroResultDTO struct {
	Token string 
	Status int 
	Form *DBMIshiguroForm 
	List []DBMIshiguroForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMIshiguroForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.Name:     "name", 
		&cf.Code:     "code", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFMIshiguro(dto DBMIshiguroDTO) DBMIshiguroForm { 
	var form DBMIshiguroForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.Name = "" 
	if dto.Name.Valid == true { 
		form.Name = dto.Name.String 
	} 
	form.Code = "" 
	if dto.Code.Valid == true { 
		form.Code = dto.Code.String 
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
