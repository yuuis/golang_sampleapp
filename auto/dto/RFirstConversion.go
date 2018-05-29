package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBRFirstConversionDTO struct {
	Sid sql.NullInt64 
	MSettingSid sql.NullInt64 
	Url sql.NullString 
	UserKey sql.NullString 
	FirstTBeaconSid sql.NullInt64 
}
type DBRFirstConversionForm struct {
	Mode string 
	Sid string `validate:"required,max=20"` 
	MSettingSid string `validate:"required,max=20"` 
	Url string `validate:"max=255"` 
	UserKey string `validate:"required,max=255"` 
	FirstTBeaconSid string `validate:"required,max=20"` 
}
type DBRFirstConversionResultDTO struct {
	Token string 
	Status int 
	Form *DBRFirstConversionForm 
	List []DBRFirstConversionForm 
	Errors []dto.ErrorForm 
}
func (cf *DBRFirstConversionForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.Sid:     "sid", 
		&cf.MSettingSid:     "m_setting_sid", 
		&cf.Url:     "url", 
		&cf.UserKey:     "user_key", 
		&cf.FirstTBeaconSid:     "first_t_beacon_sid", 
	} 
} 
func DTFRFirstConversion(dto DBRFirstConversionDTO) DBRFirstConversionForm { 
	var form DBRFirstConversionForm 
	form.Sid = "" 
	if dto.Sid.Valid == true { 
		form.Sid = strconv.FormatInt(dto.Sid.Int64, 10) 
	} 
	form.MSettingSid = "" 
	if dto.MSettingSid.Valid == true { 
		form.MSettingSid = strconv.FormatInt(dto.MSettingSid.Int64, 10) 
	} 
	form.Url = "" 
	if dto.Url.Valid == true { 
		form.Url = dto.Url.String 
	} 
	form.UserKey = "" 
	if dto.UserKey.Valid == true { 
		form.UserKey = dto.UserKey.String 
	} 
	form.FirstTBeaconSid = "" 
	if dto.FirstTBeaconSid.Valid == true { 
		form.FirstTBeaconSid = strconv.FormatInt(dto.FirstTBeaconSid.Int64, 10) 
	} 
	return form 
} 
