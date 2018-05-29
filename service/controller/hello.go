package controller

import (
	"../../config"
	"../common"
	"../data"
	"../db"
	"../model"
	"net/http"
	"log"
)

type Page struct {
	Title string
	Count int
	Message string
	Status int
}

func HelloViewHandler(w http.ResponseWriter, r *http.Request) {
	//r.Header.Get("X-FORWARDED-FOR")
	common.WriteLog(config.INFO, "test", r)
	page := new(Page)
	page.Title = "Hello"

	page.Count = 1 + data.GetIntSession(r, "count")
	data.SetIntSession(w, r, "count", page.Count)

	log.Println("session: ", data.GetIntSession(r, "count") + 1)

	tmpl, err := common.ViewParses("./view/layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}

	model.TestSelectSQL(r)
	//	common.SendMail("mori@ibf.co.jp", "", "test@send.secm.jp", "テストメール", "テストメールです")

	db.SetMemcachedData("testKey", "あああああああ", 15*60)
	common.WriteLog(config.DEBUG, "memcached = "+db.GetMemcachedData("testKey"), r)

	db.SetRedisData("redisKey1", "れでぃすてすと", 5*60)
	common.WriteLog(config.DEBUG, "redis = "+db.GetRedisData("redisKey1"), r)
}
