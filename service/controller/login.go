package controller

import (
  "net/http"
  // "html/template"
  "encoding/json"
  "log"
  "fmt"
  "../dto"
  "../data"
  "../common"
  "../logic"
  "../../config"
)

type Result struct {
  Result string `json:"result"`
  Error string `json:"error"`
}

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
//   common.WriteLog(config.DEBUG, "login", r)

//   if r.Method == "POST" {
//     // フォーム入力情報からUserForm型を作成
//     username := r.FormValue("name")
//     user := dto.UserForm{username, r.FormValue("password")}
//     page := new(Page)
//     var tmpl *template.Template
//     var err error 

//     // 認証処理
//     ok, err := logic.Authenticate(user)
//     if err != nil {
//       common.WriteErrorLog(config.DEBUG, err, nil)
//     }
//     if ok{
//       // ログイン成功
//       page.Title = "logined"
//       page.Message = "login success!"
//       page.Status = 1
//       data.SetStringSession(w, r, "user", username)
//       tmpl, err = common.ViewParses("view/login/logined.html")
//     } else {
//       // ログイン失敗
//       page.Title = "can't login"
//       page.Message = "login faild!"
//       page.Status = 2
//       tmpl, err = common.ViewParses("view/login/login.html")
//     } 

//     if err != nil {
//       common.WriteErrorLog(config.DEBUG, err, nil)
//     }
//     err = tmpl.Execute(w, page)
//     if err != nil {
//       common.WriteErrorLog(config.DEBUG, err, nil)
//     }
//   } else {
//     page := new(Page)
//     page.Title = "wrong access"
//     page.Status = 2

//     tmpl, err := common.ViewParses("view/login/login.html")
//     if err != nil {
//       common.WriteErrorLog(config.DEBUG, err, nil)
//     }
//     err = tmpl.Execute(w, page)
//     if err != nil {
//       common.WriteErrorLog(config.DEBUG, err, nil)
//     }
//   }
}

func LoginAsynchronousViewHandler(w http.ResponseWriter, r *http.Request) {
  common.WriteLog(config.DEBUG, "loginasynchronous", r)
  log.Println("name: " + r.FormValue("name"))
  log.Println("password: " + r.FormValue("password"))
  user := dto.UserForm{r.FormValue("name"), r.FormValue("password")}
  ok, errm := logic.Authenticate(user)

  if ok {
    // ログイン成功
    data.SetStringSession(w, r, "user", r.FormValue("name"))
    result := new(Result)
    result.Result = "success"
    output, _ := json.Marshal(&result)
    log.Println("output: " + string(output))
    fmt.Fprint(w, string(output))
  } else {
    // ログイン失敗
    result := new(Result)
    result.Result = "faild"
    result.Error = errm
    output, _ := json.Marshal(&result)
    log.Println("output: " + string(output))
    fmt.Fprint(w, string(output))
  } 
}

func LogoutViewHandler(w http.ResponseWriter, r *http.Request) {
  data.SetStringSession(w, r, "user", "")
  GetLoginViewHandler(w, r)
}