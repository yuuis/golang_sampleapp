package autocontroller

import (
	"../../config"
	"../../service/common"
	"net/http"
)

func AutoIndexViewHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteLog(config.INFO, "Start", r)
	// デフォルト遷移先
	tmpl, err := common.ViewParses("./auto/view/index.html")

	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		common.WriteErrorLog(config.FATAL, err, r)
		panic(err)
	}
}
