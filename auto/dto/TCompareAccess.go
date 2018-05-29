package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBTCompareAccessDTO struct {
	Sid sql.NullInt64 
	Key sql.NullString 
	Name sql.NullString 
	Url sql.NullString 
	Percent sql.NullInt64 
	Encode sql.NullString 
	Forward sql.NullInt64 
	RequestUrl sql.NullString 
	RefererUrl sql.NullString 
	AccessTime NullTime 
	SessionId sql.NullString 
	CookiesKey sql.NullString 
	Price sql.NullInt64 
	IpAddress sql.NullString 
	DeleteFlag sql.NullInt64 
	CreateUserId sql.NullInt64 
	UpdateUserId sql.NullInt64 
	CreateTime NullTime 
	UpdateTime NullTime 
}
type DBTCompareAccessForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	Key string `validate:"max=50"` 
	Name string `validate:"max=255"` 
	Url string `validate:"max=255"` 
	Percent string `validate:"max=3"` 
	Encode string `validate:"max=10"` 
	Forward string `validate:"max=1"` 
	RequestUrl string `validate:"max=255"` 
	RefererUrl string `validate:"max=255"` 
	AccessTime string 
	SessionId string `validate:"max=50"` 
	CookiesKey string `validate:"max=50"` 
	Price string `validate:"max=11"` 
	IpAddress string `validate:"max=20"` 
	DeleteFlag string `validate:"max=1"` 
	CreateUserId string `validate:"max=20"` 
	UpdateUserId string `validate:"max=20"` 
	CreateTime string 
	UpdateTime string 
}
type DBTCompareAccessResultDTO struct {
	Token string 
	Status int 
	Form *DBTCompareAccessForm 
	List []DBTCompareAccessForm 
	Errors []dto.ErrorForm 
}
func (cf *DBTCompareAccessForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.Key:     "key", 
		&cf.Name:     "name", 
		&cf.Url:     "url", 
		&cf.Percent:     "percent", 
		&cf.Encode:     "encode", 
		&cf.Forward:     "forward", 
		&cf.RequestUrl:     "request_url", 
		&cf.RefererUrl:     "referer_url", 
		&cf.AccessTime:     "access_time", 
		&cf.SessionId:     "session_id", 
		&cf.CookiesKey:     "cookies_key", 
		&cf.Price:     "price", 
		&cf.IpAddress:     "ip_address", 
		&cf.DeleteFlag:     "delete_flag", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateTime:     "create_time", 
		&cf.UpdateTime:     "update_time", 
	} 
} 
func DTFTCompareAccess(dto DBTCompareAccessDTO) DBTCompareAccessForm { 
	var form DBTCompareAccessForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.Key = "" 
	if dto.Key.Valid == true { 
		form.Key = dto.Key.String 
	} 
	form.Name = "" 
	if dto.Name.Valid == true { 
		form.Name = dto.Name.String 
	} 
	form.Url = "" 
	if dto.Url.Valid == true { 
		form.Url = dto.Url.String 
	} 
	form.Percent = "" 
	if dto.Percent.Valid == true { 
		form.Percent = strconv.FormatInt(dto.Percent.Int64, 10) 
	} 
	form.Encode = "" 
	if dto.Encode.Valid == true { 
		form.Encode = dto.Encode.String 
	} 
	form.Forward = "" 
	if dto.Forward.Valid == true { 
		form.Forward = strconv.FormatInt(dto.Forward.Int64, 10) 
	} 
	form.RequestUrl = "" 
	if dto.RequestUrl.Valid == true { 
		form.RequestUrl = dto.RequestUrl.String 
	} 
	form.RefererUrl = "" 
	if dto.RefererUrl.Valid == true { 
		form.RefererUrl = dto.RefererUrl.String 
	} 
	form.AccessTime = "" 
	if dto.AccessTime.Valid == true { 
		form.AccessTime = dto.AccessTime.Time.Format("2006/01/02 15:04:05") 
	} 
	form.SessionId = "" 
	if dto.SessionId.Valid == true { 
		form.SessionId = dto.SessionId.String 
	} 
	form.CookiesKey = "" 
	if dto.CookiesKey.Valid == true { 
		form.CookiesKey = dto.CookiesKey.String 
	} 
	form.Price = "" 
	if dto.Price.Valid == true { 
		form.Price = strconv.FormatInt(dto.Price.Int64, 10) 
	} 
	form.IpAddress = "" 
	if dto.IpAddress.Valid == true { 
		form.IpAddress = dto.IpAddress.String 
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
