package controller

import (
	"../../config"
	"../common"
	"../dto"
	"../logic"
	"github.com/mholt/binding"
	"net/http"
)

func ListPostTestIndexViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "start", r)
	form := new(dto.ListPostTestForm)

	tmpl, err := common.ViewParses("./view/listposttest/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, form)
	if err != nil {
		panic(err)
	}
}

func ListPostTestRegisterViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "start", r)
	form := new(dto.ListPostTestForm)

	// POSTされてきたデータをFORMへ詰め込む
	if errs := binding.Bind(r, form); errs != nil {
		common.WriteErrorLog(config.ERROR, errs, r)
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}
	// バリデート＆表示用のリストへ詰め替える
	for ind, _ := range form.MemcachedValue {
		memcachedData := dto.ListPostMemcachedData{
			MemcachedValue:  form.MemcachedValue[ind],
			MemcachedRemark: form.MemcachedRemark[ind],
		}
		form.MemcachedDatas = append(form.MemcachedDatas, memcachedData)
	}

	// デフォルト遷移先
	tmpl, err := common.ViewParses("./view/listposttest/index.html")

	// Validaterをかける
	validaterErr := common.GetValidate().Struct(form)
	if validaterErr != nil {
		// バリデーションエラーがある場合の遷移先
		form.Errors = common.MakeErrorMessage(validaterErr)
	} else {
		// 成功の場合にロジックを実行
		logic.ListPostTestRegister(form)
		tmpl, err = common.ViewParses("./view/listposttest/register.html")
	}

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, form)
	if err != nil {
		panic(err)
	}
}

func ListPostTestShowViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "start", r)
	form := new(dto.ListPostTestForm)

	// POSTされてきたデータをFORMへ詰め込む
	if errs := binding.Bind(r, form); errs != nil {
		common.WriteErrorLog(config.ERROR, errs, r)
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	// デフォルト遷移先
	tmpl, err := common.ViewParses("./view/listposttest/index.html")

	// Validaterをかける
	validaterErr := common.GetValidate().StructExcept(form, "MemcacheValue", "RedisValue")
	if validaterErr != nil {
		// バリデーションエラーがある場合の遷移先
		form.Errors = common.MakeErrorMessage(validaterErr)
	} else {
		// 成功の場合にロジックを実行
		logic.ListPostTestShow(form)
		tmpl, err = common.ViewParses("./view/listposttest/show.html")
	}

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, form)
	if err != nil {
		panic(err)
	}
}
