package controller

import (
	"../../config"
	"../common"
	"../dto"
	"../logic"
	"github.com/mholt/binding"
	"net/http"
)

func PostTestIndexViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "start", r)
	form := new(dto.PostTestForm)

	tmpl, err := common.ViewParses("./view/posttest/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, form)
	if err != nil {
		panic(err)
	}
}

func PostTestRegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "start", r)
	form := new(dto.PostTestForm)

	// POSTされてきたデータをFORMへ詰め込む
	if errs := binding.Bind(r, form); errs != nil {
		common.WriteErrorLog(config.ERROR, errs, r)
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	// デフォルト遷移先
	tmpl, err := common.ViewParses("./view/posttest/index.html")

	// Validaterをかける
	validaterErr := common.GetValidate().Struct(form)
	if validaterErr != nil {
		// バリデーションエラーがある場合の遷移先
		form.Errors = common.MakeErrorMessage(validaterErr)
	} else {
		// 成功の場合にロジックを実行
		logic.PostTestRegister(form)
		tmpl, err = common.ViewParses("./view/posttest/register.html")
	}

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, form)
	if err != nil {
		panic(err)
	}
}

func PostTestShowViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "start", r)
	form := new(dto.PostTestForm)

	// POSTされてきたデータをFORMへ詰め込む
	if errs := binding.Bind(r, form); errs != nil {
		common.WriteErrorLog(config.ERROR, errs, r)
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	// デフォルト遷移先
	tmpl, err := common.ViewParses("./view/posttest/index.html")

	// Validaterをかける
	validaterErr := common.GetValidate().StructExcept(form, "MemcacheValue", "RedisValue")
	if validaterErr != nil {
		// バリデーションエラーがある場合の遷移先
		form.Errors = common.MakeErrorMessage(validaterErr)
	} else {
		// 成功の場合にロジックを実行
		logic.PostTestShow(form)
		tmpl, err = common.ViewParses("./view/posttest/show.html")
	}

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, form)
	if err != nil {
		panic(err)
	}
}
