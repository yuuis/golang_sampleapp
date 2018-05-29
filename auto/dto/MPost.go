package autodbdto

import (
	"github.com/mholt/binding"
	"net/http"
	"../../service/dto"
	"database/sql"
	"strconv"
)

type DBMPostDTO struct {
	PostId sql.NullInt64 
	OldPostCode sql.NullString 
	PostCode sql.NullString 
	PrefectureNameKana sql.NullString 
	CityNameKana sql.NullString 
	AddressNameKana sql.NullString 
	PrefectureName sql.NullString 
	CityName sql.NullString 
	AddressName sql.NullString 
	Flag1 sql.NullInt64 
	Flag2 sql.NullInt64 
	Flag3 sql.NullInt64 
	Flag4 sql.NullInt64 
	Flag5 sql.NullInt64 
	Flag6 sql.NullInt64 
	CreateUserId sql.NullString 
	UpdateUserId sql.NullString 
	CreateDate NullTime 
	UpdateDate NullTime 
	DeleteFlag sql.NullInt64 
}
type DBMPostForm struct {
	Mode string 
	PostId string `validate:"required,max=11"` 
	OldPostCode string `validate:"required,max=5"` 
	PostCode string `validate:"required,max=7"` 
	PrefectureNameKana string `validate:"required,max=256"` 
	CityNameKana string `validate:"required,max=256"` 
	AddressNameKana string `validate:"required,max=256"` 
	PrefectureName string `validate:"required,max=256"` 
	CityName string `validate:"required,max=256"` 
	AddressName string `validate:"required,max=256"` 
	Flag1 string `validate:"required,max=11"` 
	Flag2 string `validate:"required,max=11"` 
	Flag3 string `validate:"required,max=11"` 
	Flag4 string `validate:"required,max=11"` 
	Flag5 string `validate:"required,max=11"` 
	Flag6 string `validate:"required,max=11"` 
	CreateUserId string `validate:"required,max=20"` 
	UpdateUserId string `validate:"required,max=20"` 
	CreateDate string `validate:"required"` 
	UpdateDate string `validate:"required"` 
	DeleteFlag string `validate:"required,max=11"` 
}
type DBMPostResultDTO struct {
	Token string 
	Status int 
	Form *DBMPostForm 
	List []DBMPostForm 
	Errors []dto.ErrorForm 
}
func (cf *DBMPostForm) FieldMap(req *http.Request) binding.FieldMap { 
	return binding.FieldMap{ 
		&cf.Mode: "mode",
		&cf.PostId:     "post_id", 
		&cf.OldPostCode:     "old_post_code", 
		&cf.PostCode:     "post_code", 
		&cf.PrefectureNameKana:     "prefecture_name_kana", 
		&cf.CityNameKana:     "city_name_kana", 
		&cf.AddressNameKana:     "address_name_kana", 
		&cf.PrefectureName:     "prefecture_name", 
		&cf.CityName:     "city_name", 
		&cf.AddressName:     "address_name", 
		&cf.Flag1:     "flag_1", 
		&cf.Flag2:     "flag_2", 
		&cf.Flag3:     "flag_3", 
		&cf.Flag4:     "flag_4", 
		&cf.Flag5:     "flag_5", 
		&cf.Flag6:     "flag_6", 
		&cf.CreateUserId:     "create_user_id", 
		&cf.UpdateUserId:     "update_user_id", 
		&cf.CreateDate:     "create_date", 
		&cf.UpdateDate:     "update_date", 
		&cf.DeleteFlag:     "delete_flag", 
	} 
} 
func DTFMPost(dto DBMPostDTO) DBMPostForm { 
	var form DBMPostForm 
	form.PostId = "" 
	if dto.PostId.Valid == true { 
		form.PostId = strconv.FormatInt(dto.PostId.Int64, 10) 
	} 
	form.OldPostCode = "" 
	if dto.OldPostCode.Valid == true { 
		form.OldPostCode = dto.OldPostCode.String 
	} 
	form.PostCode = "" 
	if dto.PostCode.Valid == true { 
		form.PostCode = dto.PostCode.String 
	} 
	form.PrefectureNameKana = "" 
	if dto.PrefectureNameKana.Valid == true { 
		form.PrefectureNameKana = dto.PrefectureNameKana.String 
	} 
	form.CityNameKana = "" 
	if dto.CityNameKana.Valid == true { 
		form.CityNameKana = dto.CityNameKana.String 
	} 
	form.AddressNameKana = "" 
	if dto.AddressNameKana.Valid == true { 
		form.AddressNameKana = dto.AddressNameKana.String 
	} 
	form.PrefectureName = "" 
	if dto.PrefectureName.Valid == true { 
		form.PrefectureName = dto.PrefectureName.String 
	} 
	form.CityName = "" 
	if dto.CityName.Valid == true { 
		form.CityName = dto.CityName.String 
	} 
	form.AddressName = "" 
	if dto.AddressName.Valid == true { 
		form.AddressName = dto.AddressName.String 
	} 
	form.Flag1 = "" 
	if dto.Flag1.Valid == true { 
		form.Flag1 = strconv.FormatInt(dto.Flag1.Int64, 10) 
	} 
	form.Flag2 = "" 
	if dto.Flag2.Valid == true { 
		form.Flag2 = strconv.FormatInt(dto.Flag2.Int64, 10) 
	} 
	form.Flag3 = "" 
	if dto.Flag3.Valid == true { 
		form.Flag3 = strconv.FormatInt(dto.Flag3.Int64, 10) 
	} 
	form.Flag4 = "" 
	if dto.Flag4.Valid == true { 
		form.Flag4 = strconv.FormatInt(dto.Flag4.Int64, 10) 
	} 
	form.Flag5 = "" 
	if dto.Flag5.Valid == true { 
		form.Flag5 = strconv.FormatInt(dto.Flag5.Int64, 10) 
	} 
	form.Flag6 = "" 
	if dto.Flag6.Valid == true { 
		form.Flag6 = strconv.FormatInt(dto.Flag6.Int64, 10) 
	} 
	form.CreateUserId = "" 
	if dto.CreateUserId.Valid == true { 
		form.CreateUserId = dto.CreateUserId.String 
	} 
	form.UpdateUserId = "" 
	if dto.UpdateUserId.Valid == true { 
		form.UpdateUserId = dto.UpdateUserId.String 
	} 
	form.CreateDate = "" 
	if dto.CreateDate.Valid == true { 
		form.CreateDate = dto.CreateDate.Time.Format("2006/01/02 15:04:05") 
	} 
	form.UpdateDate = "" 
	if dto.UpdateDate.Valid == true { 
		form.UpdateDate = dto.UpdateDate.Time.Format("2006/01/02 15:04:05") 
	} 
	form.DeleteFlag = "" 
	if dto.DeleteFlag.Valid == true { 
		form.DeleteFlag = strconv.FormatInt(dto.DeleteFlag.Int64, 10) 
	} 
	return form 
} 
