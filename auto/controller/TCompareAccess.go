package autocontroller

import (
	"../../config"
	"../../service/common"
	"../../service/data"
	"../dto"
	"../logic"
	"encoding/json"
	"github.com/mholt/binding"
	"net/http"
)

func TCompareAccessSearchViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "Start", r)
	// デフォルト遷移先
	tmpl, err := common.ViewParses("./auto/view/t_compare_access/index.html")

	var resultDTO autodbdto.DBTCompareAccessResultDTO
	resultDTO.Token = common.GenerateUID()

	// POSTされてきたデータをFORMへ詰め込む
	form := new(autodbdto.DBTCompareAccessForm)
	if errs := binding.Bind(r, form); errs != nil {
		common.WriteErrorLog(config.ERROR, errs, r)
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}
	resultDTO.Form = form // 結果用DTOへFormを格納

	if form.Mode == config.BUTTON_BACK {
		// セッションの内容をフォームに戻して再建策をかける
		jsonBytes := ([]byte)(data.GetStringSession(r, "TCompareAccessSearchViewHandlerKey"))
		json.Unmarshal(jsonBytes, resultDTO.Form)

		form.Mode = config.BUTTON_SEARCH
	}

	if form.Mode == config.BUTTON_SEARCH {
		common.WriteLog(config.DEBUG, "Mode = "+config.BUTTON_SEARCH, r)
		// Validaterをかける
		validaterErr := common.GetValidate().StructExcept(form, "Sid", "CreateUserId", "UpdateUserId", "CreateTime", "UpdateTime")
		if validaterErr != nil {
			// バリデーションエラーがある場合の遷移先
			common.WriteLog(config.DEBUG, "Validater Error", r)
			resultDTO.Errors = common.MakeErrorMessage(validaterErr)
		} else {
			// 成功の場合にロジックを実行
			autoResult := autologic.SearchTCompareAccess(form, r)
			resultDTO.List = autoResult // 結果用DTOへ検索結果を格納
			tmpl, err = common.ViewParses("./auto/view/t_compare_access/index.html")
			// 検索内容をセッションへ保持する
			jsonBytes, _ := json.Marshal(resultDTO.Form)
			data.SetStringSession(w, r, "TCompareAccessSearchViewHandlerKey", string(jsonBytes))
		}
	}

	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}

	err = tmpl.Execute(w, resultDTO)
	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}
}

func TCompareAccessRegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "Start", r)
	// デフォルト遷移先
	tmpl, err := common.ViewParses("./auto/view/t_compare_access/register.html")

	var resultDTO autodbdto.DBTCompareAccessResultDTO
	resultDTO.Token = common.GenerateUID()

	// POSTされてきたデータをFORMへ詰め込む
	form := new(autodbdto.DBTCompareAccessForm)
	if errs := binding.Bind(r, form); errs != nil {
		common.WriteErrorLog(config.ERROR, errs, r)
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}
	resultDTO.Form = form

	if form.Mode == config.BUTTON_REGISTER {
		common.WriteLog(config.DEBUG, "Mode = "+config.BUTTON_SEARCH, r)
		// Validaterをかける
		validaterErr := common.GetValidate().StructExcept(form, "Sid")
		if validaterErr != nil {
			// バリデーションエラーがある場合の遷移先
			common.WriteLog(config.DEBUG, "Validater Error", r)
			resultDTO.Errors = common.MakeErrorMessage(validaterErr)
			resultDTO.Status = config.RESULT_ERROR
		} else {
			// 成功の場合にロジックを実行
			sid, errFlag := autologic.RegisterTCompareAccess(form, r)
			if errFlag == false {
				resultDTO.Status = config.RESULT_ERROR
			} else {
				resultDTO.Status = config.RESULT_SUCCESS
				resultDTO.Form.Sid = sid
			}
			tmpl, err = common.ViewParses("./auto/view/t_compare_access/register.html")
		}
	} else {
		// 初期表示 SIDを指定されている場合でデータが取得できなかった場合は404画面へ遷移させる
		if form.Sid != ""  {
			resultForm, getted := autologic.GetByPkTCompareAccess(form, r) // 結果用DTOへFormを格納
			resultDTO.Form = &resultForm
			if getted == false {
				http.Redirect(w, r, "/notfound", 404)
			}
		}
	}

	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}

	err = tmpl.Execute(w, resultDTO)
	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}
}
