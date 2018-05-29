package data

import (
	"net/http"
)

/*
 * 事前処理
 *
 * 1.ログインチェック
 *
 */
func Before(w http.ResponseWriter, r *http.Request) bool {
	var checkLogin bool
	//不正な閲覧（権限のないクライアントへの閲覧試行）
	checkLogin = true
	return checkLogin
}
