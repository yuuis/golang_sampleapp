package logic

import (
	"../../config"
	"../common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

/*
 * 成否(1成功）及び、失敗の場合にエラーメッセージを返却
 */
func FileUploadRegister(r *http.Request, reader *multipart.Reader) (int, string) {
	message := ""

	//forで複数ファイルがある場合に、すべてのファイルが終わるまで読み込む
	if message == "" {
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}

			//ファイル名がない場合はスキップする
			if part.FileName() == "" {
				continue
			}

			//uploadedfileディレクトリに受け取ったファイル名でファイルを作成
			uploadedFile, err := os.Create("/tmp/" + part.FileName())
			if err != nil {
				message = "ファイルの書き込みに失敗しました。"
				common.WriteErrorLog(config.INFO, err, r)
				break
			}

			//作ったファイルに読み込んだファイルの内容を丸ごとコピー
			_, err = io.Copy(uploadedFile, part)
			if err != nil {
				message = "ファイルの書き込みに失敗しました。"
				common.WriteErrorLog(config.INFO, err, r)
				break
			}
		}
	}

	successFlag := 1
	if message != "" {
		successFlag = 0
	}

	return successFlag, message
}
