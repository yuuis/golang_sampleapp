package controller

import (
  "net/http"
  "html/template"
  "log"
  "../dto"
  "../data"
  "../common"
  "../logic"
  "../../config"
)


func GetLoginViewHandler(w http.ResponseWriter, r *http.Request) {
  common.WriteLog(config.DEBUG, "login", r)

  page := new(Page)
  page.Title = "login"

  tmpl, err := common.ViewParses("./view/login/login.html")
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }
}


func PostLoginViewHandler(w http.ResponseWriter, r *http.Request) {
  common.WriteLog(config.DEBUG, "login", r)
  r.ParseForm()

  if r.Method == "POST" {
    // フォーム入力情報からUserForm型を作成
    username := r.Form["username"][0]
    user := dto.UserForm{username, r.Form["password"][0]}
    page := new(Page)
    var tmpl *template.Template
    var err error 

    // 認証処理
    ok, err := logic.Authenticate(user)
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
    if ok{
      // ログイン成功
      page.Title = "logined"
      page.Message = "login success!"
      page.Status = 1
      data.SetStringSession(w, r, "user", username)
      tmpl, err = common.ViewParses("view/login/logined.html")
    } else {
      // ログイン失敗
      page.Title = "can't login"
      page.Message = "login faild!"
      page.Status = 2
      tmpl, err = common.ViewParses("view/login/login.html")
    } 

    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
    err = tmpl.Execute(w, page)
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
  } else {
    page := new(Page)
    page.Title = "wrong access"
    page.Status = 2

    tmpl, err := common.ViewParses("view/login/login.html")
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
    err = tmpl.Execute(w, page)
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
    }
  }
}

func LogoutViewHandler(w http.ResponseWriter, r *http.Request) {
  // user := data.GetStringSession(r, "user")
  // if user != "" {
    data.SetStringSession(w, r, "user", "")
    log.Println("hoge")
  // }
  // GetLoginViewHandler(w, r)
  page := new(Page)
  page.Title = "login"

  tmpl, err := common.ViewParses("./view/login/login.html")
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    common.WriteErrorLog(config.DEBUG, err, nil)
  }
}