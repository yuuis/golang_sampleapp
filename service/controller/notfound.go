package controller

import (
	"../../config"
	"../common"
	"net/http"
)

func NotFoundViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, " 404 ", r)
	page := new(Page)

	tmpl, err := common.ViewParses("./view/common/404.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}
