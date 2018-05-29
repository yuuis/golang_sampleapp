package common

import (
	"../dto"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
	"unicode"
)

var validate *validator.Validate

/*
 * Validaterの初期設定、カスタムでエラーチェックを追加した場合はここで宣言する
 */
func InitValidater() {
	validate = validator.New()
	validate.RegisterValidation("hiragana", validateCustomFuncHiragana)
	validate.RegisterValidation("katakana", validateCustomFuncKatakana)
}

/*
 * Validaterを使用する際のロード
 */
func GetValidate() *validator.Validate {
	if validate == nil {
		InitValidater()
	}
	return validate
}

/*
 * エラー内容をチェックし、画面上取得できる形へ整形する
 */
func MakeErrorMessage(err error) []dto.ErrorForm {
	if validate == nil {
		InitValidater()
	}

	var errors []dto.ErrorForm

	for _, err := range err.(validator.ValidationErrors) {
		// 配列要素があった場合に、何番目の要素かを格納
		indStrings := strings.Split(err.Namespace(), "[")
		ind := "0"
		if indStrings != nil && len(indStrings) > 1 {
			ind = strings.Split(indStrings[1], "]")[0]
		}

		iind, _ := strconv.Atoi(ind)
		errors = append(errors, dto.ErrorForm{
			ErrorField:      err.Field(),
			ErrorFieldIndex: iind,
			ErrorMessage:    customErrorMessage(err.Tag(), err.Field(), err.Param()),
		})

	}

	return errors
}

func customErrorMessage(tag string, field string, param string) string {
	showFieldName := fieldMapping()[field] // フィールド名の変換
	if showFieldName == "" {
		showFieldName = field
	}

	if tag == "required" {
		return showFieldName + "は" + "必須です。"
	}
	if tag == "max" {
		return showFieldName + "は" + param + "以下で入力してください。"
	}
	if tag == "min" {
		return showFieldName + "は" + param + "以上で入力してください。"
	}
	if tag == "hiragana" {
		return showFieldName + "はひらがなで入力してください。"
	}
	if tag == "katakana" {
		return showFieldName + "はかたかなで入力してください。"
	}

	return "何らかのエラーが発生しました。"
}

func fieldMapping() map[string]string {
	return map[string]string{
		"RedisKey":      "Redisのキー",
		"RedisValue":    "Redisの値",
		"MemcacheKey":   "Memcachedのキー",
		"MemcacheValue": "Memcachedの値",
	}
}

/* 以下カスタムバリデータ ****************************************************************/

/*
 * カスタムValidate　ひらがな
 */
func validateCustomFuncHiragana(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	for _, r := range fl.Field().String() {
		if unicode.In(r, unicode.Hiragana) == false {
			return false
		}
	}
	return true
}

/*
 * カスタムValidate　カタカナ
 */
func validateCustomFuncKatakana(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	for _, r := range fl.Field().String() {
		if unicode.In(r, unicode.Katakana) == false {
			return false
		}
	}
	return true
}
