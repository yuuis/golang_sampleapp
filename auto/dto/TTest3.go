package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBTTest3DTO struct {
	Number sql.NullInt64 
	InputDate NullTime 
	Floattest sql.NullFloat64 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBTTest3Form struct {
	Mode string 
	Number string `validate:"required,max=20"` 
	InputDate string `validate:"required"` 
	Floattest string `validate:"required"` 
	DeleteFlag string `validate:"required,max=1"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateTime string `validate:"required"` 
	UpdateTime string `validate:"required"` 
}
type DBTTest3ResultDTO struct {
	Token string 
	Status int 
	Form *DBTTest3Form 
	List []DBTTest3Form 
	Errors []dto.ErrorForm 
}
func (cf *DBTTest3Form) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Number:     "number", 
		&cf.InputDate:     "input_date", 
		&cf.Floattest:     "floattest", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFTTest3(dto DBTTest3DTO) DBTTest3Form { 
	var form DBTTest3Form 
	form.Number = "" 
	if dto.Number.Valid == true { 
		form.Number = strconv.FormatInt(dto.Number.Int64, 10) 
	} 
	form.InputDate = "" 
	if dto.InputDate.Valid == true { 
		form.InputDate = dto.InputDate.Time.Format("2006/01/02 15:04:05") 
	} 
	form.Floattest = "" 
	if dto.Floattest.Valid == true { 
		form.Floattest = strconv.FormatFloat(dto.Floattest.Float64, 'f', -1, 64) 
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
