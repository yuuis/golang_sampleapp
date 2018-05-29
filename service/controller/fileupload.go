package controller

import (
	"../../config"
	"../common"
	"../dto"
	"../logic"
	"encoding/json"
	"fmt"
	"net/http"
)

func FileUploadIndexViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "Start", r)

	tmpl, err := common.ViewParses("./view/fileupload/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
func FileUploadRegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "Start", r)

	//MultipartReaderを用いて受け取ったファイルを読み込み
	reader, err := r.MultipartReader()
	successFlag := 1
	message := ""
	//エラーが発生したら抜ける
	if err != nil {
		message = "アップロードされたファイルの読み込みに失敗しました。"
	} else {
		successFlag, message = logic.FileUploadRegister(r, reader)
	}

	var form = dto.FileUploadForm{
		Success: successFlag,
		Message: message,
	}

	// jsonエンコード
	outputJson, err := json.Marshal(&form)
	if err != nil {
		common.WriteErrorLog(config.INFO, err, r)
		panic(err)
	}

	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")

	// jsonデータを出力
	fmt.Fprint(w, string(outputJson))
}
