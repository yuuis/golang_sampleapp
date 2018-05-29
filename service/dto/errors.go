package dto

/*
 * エラー共通DTO
 */
type ErrorForm struct {
	ErrorField      string
	ErrorFieldIndex int
	ErrorMessage    string
}
