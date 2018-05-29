package main

import (
	"../config"
	"../service/common"
	"../service/db"
	"database/sql"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const timelayout = "2006/01/02 15:04:05"

var outputPath = "/Users/yuuis/gowork/web/src/auto"
var configPath = "../../config"
var serviceCommonPath = "../../service/common"
var serviceDbPath = "../../service/db"
var serviceDtoPath = "../../service/dto"

type TableList struct {
	table_in sql.NullString
}
type ColumnList struct {
	field           sql.NullString
	fieldtype       sql.NullString
	fieldcollation  sql.NullString
	fieldnull       sql.NullString
	fieldkey        sql.NullString
	fielddefault    sql.NullString
	fieldextra      sql.NullString
	fieldprivileges sql.NullString
	fieldcomment    sql.NullString
}

func main() {
	_, err := db.DbInit()
	if err != nil {
		panic(err)
	}

	// 前回ファイルを全て削除する
	if err := os.RemoveAll(outputPath); err != nil {
		panic(err)
	}
	// テーブル一覧取得
	tableList := getTableList()

	urlList := ""
	contentsUrl := ""
	for ind, _ := range tableList {
		// テーブルの内容取得
		columnList := getColumnList(tableList[ind])
		// DTO作成
		makeDto(tableList[ind], columnList, ind)
		// Model作成
		makeModel(tableList[ind], columnList, ind)

		makeLogic(tableList[ind], columnList, ind)
		makeController(tableList[ind], columnList, ind)
		makeView(tableList[ind], columnList, ind)
		u, c := makeUrlList(tableList[ind], columnList, ind)
		urlList = urlList + u
		contentsUrl = contentsUrl + c
	}
	makeIndex(urlList)
	makeIndexHtml(contentsUrl)
	makeIndexController()
}

func makeIndex(urlList string) {
	dtoPath := outputPath + ""

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}

	contentString := ""
	contentString = contentString + "package auto\n"
	contentString = contentString + "\n"
	contentString = contentString + "import (\n"
	contentString = contentString + "	\"./controller\"\n"
	contentString = contentString + "	\"log\"\n"
	contentString = contentString + "	\"net/http\"\n"
	contentString = contentString + ")\n"
	contentString = contentString + "\n"
	contentString = contentString + "func autoAuthenticate(fn http.HandlerFunc) http.HandlerFunc {\n"
	contentString = contentString + "	return func(w http.ResponseWriter, r *http.Request) {\n"
	contentString = contentString + "		log.Println(\"before process\") // 処理の前の共通処理\n"
	contentString = contentString + "		fn(w, r)\n"
	contentString = contentString + "		log.Println(\"after process\") // 処理の後の共通処理\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "}\n"
	contentString = contentString + "\n"
	contentString = contentString + "func AutoControllerLoad() {\n"
	contentString = contentString + "	http.HandleFunc(\"/autosample\", autoAuthenticate(autocontroller.AutoIndexViewHandler))\n"
	contentString = contentString + "\n"
	contentString = contentString + urlList
	contentString = contentString + "\n"
	contentString = contentString + "}\n"

	ioutil.WriteFile(dtoPath+"/autocontroller.go", []byte(contentString), os.ModePerm)
}

func makeUrlList(table TableList, columns []ColumnList, cnt int) (string, string) {
	havePrimary := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			havePrimary = true
		}
	}
	if havePrimary == false {
		return "", ""
	}

	contentString := ""
	contentString = contentString + "	http.HandleFunc(\"/autosample/" + changeToGlobals(table.table_in.String) + "\", autoAuthenticate(autocontroller." + changeToGlobals(table.table_in.String) + "SearchViewHandler))\n"
	contentString = contentString + "	http.HandleFunc(\"/autosample/" + changeToGlobals(table.table_in.String) + "/register\", autoAuthenticate(autocontroller." + changeToGlobals(table.table_in.String) + "RegisterViewHandler))\n"

	contentUrl := ""
	contentUrl = contentUrl + "<div class=\"container\">\n"
	contentUrl = contentUrl + "  <div class=\"form-group row\">\n"
	contentUrl = contentUrl + "    <div class=\"col-form-label col-4 border border-dark m-1 pb-1\"><a href=\"/autosample/" + changeToGlobals(table.table_in.String) + "\">" + table.table_in.String + "</a></div>\n"
	contentUrl = contentUrl + "  </div>\n"
	contentUrl = contentUrl + "</div>\n"
	contentUrl = contentUrl + "\n"

	return contentString, contentUrl
}

func makeIndexHtml(contentsUrl string) {
	dtoPath := outputPath + "/view"

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}

	contentString := ""
	contentString = contentString + "{{template \"headerTpl\" dict \"Title\" \"Goサンプル\" \"SubTitle\" \"TInquriy\"}}\n"
	contentString = contentString + "<section>\n"
	contentString = contentString + "<h3>機能一覧</h3>\n"
	contentString = contentString + "\n"
	contentString = contentString + contentsUrl
	contentString = contentString + "</section>\n"
	contentString = contentString + "\n"
	contentString = contentString + "{{template \"footerTpl\" .}}\n"

	ioutil.WriteFile(dtoPath+"/index.html", []byte(contentString), os.ModePerm)
}

func makeView(table TableList, columns []ColumnList, cnt int) {
	dtoPath := outputPath + "/view/" + table.table_in.String

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}

	tableColumns := 0
	for ind, _ := range columns {
		if columns[ind].fieldkey.String != "PRI" {
			tableColumns = tableColumns + 1
		}
	}
	havePrimary := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			havePrimary = true
		}
	}
	if havePrimary == false {
		return
	}

	contentString := ""
	contentString = contentString + "{{template \"headerTpl\" dict \"Title\" \"Goサンプル\" \"SubTitle\" \"" + changeToGlobals(table.table_in.String) + "\"}}\n"
	contentString = contentString + "<section>\n"
	contentString = contentString + "{{template \"datatablesTpl\" dict \"Id\" \"list\" \"Length\" \"" + strconv.Itoa(tableColumns) + "\" \"Order\" \"1\" \"AscDesc\" \"desc\"}}\n"
	contentString = contentString + "<script type=\"text/javascript\" charset=\"utf-8\">\n"
	contentString = contentString + "function funcGoRegister("
	addedflag := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			if addedflag {
				contentString = contentString + ","
			}
			contentString = contentString + "" + changeToGlobals(columns[ind].field.String) + ""
			addedflag = true
		}
	}
	contentString = contentString + ") {\n"
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			contentString = contentString + "	$(\"#" + changeToGlobals(columns[ind].field.String) + "ForRegister\").val(" + changeToGlobals(columns[ind].field.String) + ");\n"
		}
	}
	contentString = contentString + "	$(\"#" + changeToGlobals(table.table_in.String) + "RegisterForm\").submit();\n"
	contentString = contentString + "}\n"
	contentString = contentString + "</script>\n"
	contentString = contentString + "<h3>検索ページ</h3>\n"
	contentString = contentString + "<h5 class=\"bg-secondary text-white p-2\">検索条件</h5>\n"
	contentString = contentString + "<div class=\"container\">\n"
	contentString = contentString + "  {{template \"formTpl\" dict \"Action\" \"/autosample/" + changeToGlobals(table.table_in.String) + "\" \"Id\" \"" + changeToGlobals(table.table_in.String) + "Form\" \"Mode\" \"search\" \"Token\" .Token}}\n"

	for ind, _ := range columns {
		if columns[ind].fieldkey.String != "PRI" {
			contentString = contentString + "  <div class=\"form-group row\">\n"
			contentString = contentString + "    <div class=\"col-form-label col-4 border border-dark m-1 pb-1\">" + columns[ind].fieldcomment.String + "</div>\n"
			contentString = contentString + "    <div class=\"col-7 m-1\">\n"
			contentString = contentString + "      <input type=\"text\" name=\"" + columns[ind].field.String + "\" value=\"{{.Form." + changeToGlobals(columns[ind].field.String) + "}}\" id=\"" + changeToGlobals(columns[ind].field.String) + "\" class=\"form-control"
			if strings.Contains(columns[ind].field.String, "time") {
				contentString = contentString + " datetime"
			} else if !strings.Contains(columns[ind].field.String, "update") && strings.Contains(columns[ind].field.String, "date") {
				contentString = contentString + " date"
			}
			contentString = contentString + "\"/>\n"
			contentString = contentString + "      {{template \"errorsTpl\" dict \"Errors\" .Errors \"Target\" \"" + changeToGlobals(columns[ind].field.String) + "\" \"Marking\" \"" + changeToGlobals(columns[ind].field.String) + "\"}}\n"
			contentString = contentString + "    </div>\n"
			contentString = contentString + "  </div>\n"
		}
	}

	contentString = contentString + "  <div class=\"form-group row\">\n"
	contentString = contentString + "    <div class=\"offset-sm-2 col-sm-10\">\n"
	contentString = contentString + "      <button type=\"submit\" class=\"btn btn-outline-primary col-7 center-block\" id=\"searchButton\">検索</button>\n"
	contentString = contentString + "      <button type=\"button\" class=\"btn btn-outline-success col-3 center-block\" id=\"registerButton\" onclick=\"funcGoRegister('');\">新規作成</button>\n"
	contentString = contentString + "    </div>\n"
	contentString = contentString + "  </div>\n"
	contentString = contentString + "  </form>\n"
	contentString = contentString + "</div>\n"
	contentString = contentString + "<hr/>\n"
	contentString = contentString + "<table class=\"table\" id=\"list\">\n"
	contentString = contentString + "  <thead class=\"thead-light\">\n"
	contentString = contentString + "    <tr>\n"
	contentString = contentString + "      <th scope=\"col\">&nbsp;</th>\n"
	for ind, _ := range columns {
		if columns[ind].fieldkey.String != "PRI" {
			contentString = contentString + "      <th scope=\"col\">" + columns[ind].fieldcomment.String + "</th>\n"
		}
	}
	contentString = contentString + "    </tr>\n"
	contentString = contentString + "  </thead>\n"
	contentString = contentString + "  <tbody>\n"
	contentString = contentString + "    {{range .List}}\n"
	contentString = contentString + "    <tr>\n"
	contentString = contentString + "      <td><a href=\"javascript: onclick=funcGoRegister("
	addedflag = false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			if addedflag {
				contentString = contentString + ","
			}
			contentString = contentString + "'{{." + changeToGlobals(columns[ind].field.String) + "}}'"
			addedflag = true
		}
	}
	contentString = contentString + ");\">"
	addedflag = false
	for ind, _ := range columns {

		if columns[ind].fieldkey.String == "PRI" {
			if addedflag {
				contentString = contentString + ","
			}
			contentString = contentString + "{{." + changeToGlobals(columns[ind].field.String) + "}}"
			addedflag = true
		}
	}
	contentString = contentString + "</a></td>\n"
	for ind, _ := range columns {
		if columns[ind].fieldkey.String != "PRI" {
			contentString = contentString + "      <td>{{." + changeToGlobals(columns[ind].field.String) + "}}</td>\n"
		}
	}
	contentString = contentString + "    </tr>\n"
	contentString = contentString + "    {{end}}\n"
	contentString = contentString + "  </tbody>\n"
	contentString = contentString + "</table>\n"
	contentString = contentString + "<hr/>\n"
	contentString = contentString + "\n"
	contentString = contentString + "\n"
	contentString = contentString + "</section>\n"
	contentString = contentString + "\n"
	contentString = contentString + "{{template \"formTpl\" dict \"Action\" \"/autosample/" + changeToGlobals(table.table_in.String) + "/register\" \"Id\" \"" + changeToGlobals(table.table_in.String) + "RegisterForm\" \"Mode\" \"show\" \"Token\" .Token}}\n"
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			contentString = contentString + "<input type=\"hidden\" name=\"" + columns[ind].field.String + "\" id=\"" + changeToGlobals(columns[ind].field.String) + "ForRegister\"/>\n"
		}
	}
	contentString = contentString + "</form>\n"
	contentString = contentString + "\n"
	contentString = contentString + "{{template \"footerTpl\" .}}\n"
	contentString = contentString + "\n"
	ioutil.WriteFile(dtoPath+"/index.html", []byte(contentString), os.ModePerm)

	contentString = ""
	contentString = contentString + "{{template \"headerTpl\" dict \"Title\" \"Goサンプル\" \"SubTitle\" \"" + changeToGlobals(table.table_in.String) + "\"}}\n"
	contentString = contentString + "<section>\n"
	contentString = contentString + "<script type=\"text/javascript\" charset=\"utf-8\">\n"
	contentString = contentString + "function funcBack() {\n"
	contentString = contentString + "	$(\"#" + changeToGlobals(table.table_in.String) + "FormForBack\").submit();\n"
	contentString = contentString + "}\n"
	contentString = contentString + "</script>\n"
	contentString = contentString + "<h3>検索ページ</h3>\n"
	contentString = contentString + "<h5 class=\"bg-secondary text-white p-2\">登録変更</h5>\n"
	contentString = contentString + "<div class=\"container\">\n"
	contentString = contentString + "  {{template \"formTpl\" dict \"Action\" \"/autosample/" + changeToGlobals(table.table_in.String) + "/register\" \"Id\" \"" + changeToGlobals(table.table_in.String) + "FormRegister\" \"Mode\" \"register\" \"Token\" .Token}}\n"
	contentString = contentString + "  {{template \"completeTpl\" dict \"Status\" .Status}}\n"
	for ind, _ := range columns {
		if columns[ind].fieldextra.String == "auto_increment" {
			contentString = contentString + "  <input type=\"hidden\" name=\"" + columns[ind].field.String + "\" value=\"{{.Form." + changeToGlobals(columns[ind].field.String) + "}}\" />\n"
		} else {
			contentString = contentString + "  <div class=\"form-group row\">\n"
			contentString = contentString + "    <div class=\"col-form-label col-4 border border-dark m-1 pb-1\">" + columns[ind].fieldcomment.String + ""
			if columns[ind].fieldnull.String == "NO" {
				contentString = contentString + "<span class=\"text-danger\">※</span>"
			}
			contentString = contentString + "</div>\n"
			contentString = contentString + "    <div class=\"col-7 m-1\">\n"
			contentString = contentString + "      <input type=\"text\" name=\"" + columns[ind].field.String + "\" value=\"{{.Form." + changeToGlobals(columns[ind].field.String) + "}}\" id=\"" + changeToGlobals(columns[ind].field.String) + "\" class=\"form-control"
			if strings.Contains(columns[ind].field.String, "time") {
				contentString = contentString + " datetime"
			} else if !strings.Contains(columns[ind].field.String, "update") && strings.Contains(columns[ind].field.String, "date") {
				contentString = contentString + " date"
			}
			contentString = contentString + "\"/>\n"
			contentString = contentString + "      {{template \"errorsTpl\" dict \"Errors\" .Errors \"Target\" \"" + changeToGlobals(columns[ind].field.String) + "\" \"Marking\" \"" + changeToGlobals(columns[ind].field.String) + "\"}}\n"
			contentString = contentString + "    </div>\n"
			contentString = contentString + "  </div>\n"
		}
	}
	contentString = contentString + "  <div class=\"form-group row\">\n"
	contentString = contentString + "    <div class=\"offset-sm-2 col-sm-10\">\n"
	contentString = contentString + "      <button type=\"button\" class=\"btn btn-outline-secondary col-4 center-block\" id=\"backButton\" onclick=\"funcBack();\">戻る</button>\n"
	contentString = contentString + "      <button type=\"submit\" class=\"btn btn-outline-success col-4 center-block\" id=\"registerButton\">登録</button>\n"
	contentString = contentString + "    </div>\n"
	contentString = contentString + "  </div>\n"
	contentString = contentString + "  </form>\n"
	contentString = contentString + "</div>\n"
	contentString = contentString + "\n"
	contentString = contentString + "</section>\n"
	contentString = contentString + "\n"
	contentString = contentString + "{{template \"formTpl\" dict \"Action\" \"/autosample/" + changeToGlobals(table.table_in.String) + "\" \"Id\" \"" + changeToGlobals(table.table_in.String) + "FormForBack\" \"Mode\" \"back\" \"Token\" .Token}}\n"
	contentString = contentString + "</form>\n"
	contentString = contentString + "\n"
	contentString = contentString + "{{template \"footerTpl\" .}}\n"
	contentString = contentString + "\n"

	ioutil.WriteFile(dtoPath+"/register.html", []byte(contentString), os.ModePerm)
}

func makeIndexController() {
	dtoPath := outputPath + "/controller"
	contentString := ""
	contentString = contentString + "package autocontroller\n"
	contentString = contentString + "\n"
	contentString = contentString + "import (\n"
	contentString = contentString + "	\"../../config\"\n"
	contentString = contentString + "	\"../../service/common\"\n"
	contentString = contentString + "	\"net/http\"\n"
	contentString = contentString + ")\n"
	contentString = contentString + "\n"
	contentString = contentString + "func AutoIndexViewHandler(w http.ResponseWriter, r *http.Request) {\n"
	contentString = contentString + "	common.WriteLog(config.INFO, \"Start\", r)\n"
	contentString = contentString + "	// デフォルト遷移先\n"
	contentString = contentString + "	tmpl, err := common.ViewParses(\"./auto/view/index.html\")\n"
	contentString = contentString + "\n"
	contentString = contentString + "	if err != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r)\n"
	contentString = contentString + "		panic(err)\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "\n"
	contentString = contentString + "	err = tmpl.Execute(w, nil)\n"
	contentString = contentString + "	if err != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r)\n"
	contentString = contentString + "		panic(err)\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "}\n"
	ioutil.WriteFile(dtoPath+"/index.go", []byte(contentString), os.ModePerm)

}
func makeController(table TableList, columns []ColumnList, cnt int) {
	dtoPath := outputPath + "/controller"

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}

	haveAutincliment := false
	for ind, _ := range columns {
		if columns[ind].fieldextra.String == "auto_increment" {
			haveAutincliment = true
		}
	}
	havePrimary := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			havePrimary = true
		}
	}
	if havePrimary == false {
		return
	}

	contentString := ""
	contentString = contentString + "package autocontroller\n"
	contentString = contentString + "\n"
	contentString = contentString + "import (\n"
	contentString = contentString + "	\"../../config\"\n"
	contentString = contentString + "	\"../../service/common\"\n"
	contentString = contentString + "	\"../../service/data\"\n"
	contentString = contentString + "	\"../dto\"\n"
	contentString = contentString + "	\"../logic\"\n"
	contentString = contentString + "	\"encoding/json\"\n"
	contentString = contentString + "	\"github.com/mholt/binding\"\n"
	contentString = contentString + "	\"net/http\"\n"
	contentString = contentString + ")\n"
	contentString = contentString + "\n"
	contentString = contentString + "func " + changeToGlobals(table.table_in.String) + "SearchViewHandler(w http.ResponseWriter, r *http.Request) {\n"
	contentString = contentString + "	common.WriteLog(config.INFO, \"Start\", r)\n"
	contentString = contentString + "	// デフォルト遷移先\n"
	contentString = contentString + "	tmpl, err := common.ViewParses(\"./auto/view/" + table.table_in.String + "/index.html\")\n"
	contentString = contentString + "\n"
	contentString = contentString + "	var resultDTO autodbdto.DB" + changeToGlobals(table.table_in.String) + "ResultDTO\n"
	contentString = contentString + "	resultDTO.Token = common.GenerateUID()\n"
	contentString = contentString + "\n"
	contentString = contentString + "	// POSTされてきたデータをFORMへ詰め込む\n"
	contentString = contentString + "	form := new(autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form)\n"
	contentString = contentString + "	if errs := binding.Bind(r, form); errs != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.ERROR, errs, r)\n"
	contentString = contentString + "		http.Error(w, errs.Error(), http.StatusBadRequest)\n"
	contentString = contentString + "		return\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "	resultDTO.Form = form // 結果用DTOへFormを格納\n"
	contentString = contentString + "\n"
	contentString = contentString + "	if form.Mode == config.BUTTON_BACK {\n"
	contentString = contentString + "		// セッションの内容をフォームに戻して再建策をかける\n"
	contentString = contentString + "		jsonBytes := ([]byte)(data.GetStringSession(r, \"" + changeToGlobals(table.table_in.String) + "SearchViewHandlerKey\"))\n"
	contentString = contentString + "		json.Unmarshal(jsonBytes, resultDTO.Form)\n"
	contentString = contentString + "\n"
	contentString = contentString + "		form.Mode = config.BUTTON_SEARCH\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "\n"
	contentString = contentString + "	if form.Mode == config.BUTTON_SEARCH {\n"
	contentString = contentString + "		common.WriteLog(config.DEBUG, \"Mode = \"+config.BUTTON_SEARCH, r)\n"
	contentString = contentString + "		// Validaterをかける\n"
	contentString = contentString + "		validaterErr := common.GetValidate().StructExcept(form"
	for ind, _ := range columns {
		if columns[ind].fieldnull.String == "NO" || columns[ind].field.String == "create_user_id" || columns[ind].field.String == "update_user_id" || columns[ind].field.String == "create_time" || columns[ind].field.String == "update_time" {
			contentString = contentString + ", \"" + changeToGlobals(columns[ind].field.String) + "\""
		}
	}
	contentString = contentString + ")\n"
	contentString = contentString + "		if validaterErr != nil {\n"
	contentString = contentString + "			// バリデーションエラーがある場合の遷移先\n"
	contentString = contentString + "			common.WriteLog(config.DEBUG, \"Validater Error\", r)\n"
	contentString = contentString + "			resultDTO.Errors = common.MakeErrorMessage(validaterErr)\n"
	contentString = contentString + "		} else {\n"
	contentString = contentString + "			// 成功の場合にロジックを実行\n"
	contentString = contentString + "			autoResult := autologic.Search" + changeToGlobals(table.table_in.String) + "(form, r)\n"
	contentString = contentString + "			resultDTO.List = autoResult // 結果用DTOへ検索結果を格納\n"
	contentString = contentString + "			tmpl, err = common.ViewParses(\"./auto/view/" + table.table_in.String + "/index.html\")\n"
	contentString = contentString + "			// 検索内容をセッションへ保持する\n"
	contentString = contentString + "			jsonBytes, _ := json.Marshal(resultDTO.Form)\n"
	contentString = contentString + "			data.SetStringSession(w, r, \"" + changeToGlobals(table.table_in.String) + "SearchViewHandlerKey\", string(jsonBytes))\n"
	contentString = contentString + "		}\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "\n"
	contentString = contentString + "	if err != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r)\n"
	contentString = contentString + "		panic(err)\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "\n"
	contentString = contentString + "	err = tmpl.Execute(w, resultDTO)\n"
	contentString = contentString + "	if err != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r)\n"
	contentString = contentString + "		panic(err)\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "}\n"
	contentString = contentString + "\n"
	contentString = contentString + "func " + changeToGlobals(table.table_in.String) + "RegisterViewHandler(w http.ResponseWriter, r *http.Request) {\n"
	contentString = contentString + "	common.WriteLog(config.INFO, \"Start\", r)\n"
	contentString = contentString + "	// デフォルト遷移先\n"
	contentString = contentString + "	tmpl, err := common.ViewParses(\"./auto/view/" + table.table_in.String + "/register.html\")\n"
	contentString = contentString + "\n"
	contentString = contentString + "	var resultDTO autodbdto.DB" + changeToGlobals(table.table_in.String) + "ResultDTO\n"
	contentString = contentString + "	resultDTO.Token = common.GenerateUID()\n"
	contentString = contentString + "\n"
	contentString = contentString + "	// POSTされてきたデータをFORMへ詰め込む\n"
	contentString = contentString + "	form := new(autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form)\n"
	contentString = contentString + "	if errs := binding.Bind(r, form); errs != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.ERROR, errs, r)\n"
	contentString = contentString + "		http.Error(w, errs.Error(), http.StatusBadRequest)\n"
	contentString = contentString + "		return\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "	resultDTO.Form = form\n"
	contentString = contentString + "\n"
	contentString = contentString + "	if form.Mode == config.BUTTON_REGISTER {\n"
	contentString = contentString + "		common.WriteLog(config.DEBUG, \"Mode = \"+config.BUTTON_SEARCH, r)\n"
	contentString = contentString + "		// Validaterをかける\n"
	contentString = contentString + "		validaterErr := common.GetValidate().StructExcept(form"
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			contentString = contentString + ", \"" + changeToGlobals(columns[ind].field.String) + "\""
		}
	}

	contentString = contentString + ")\n"
	contentString = contentString + "		if validaterErr != nil {\n"
	contentString = contentString + "			// バリデーションエラーがある場合の遷移先\n"
	contentString = contentString + "			common.WriteLog(config.DEBUG, \"Validater Error\", r)\n"
	contentString = contentString + "			resultDTO.Errors = common.MakeErrorMessage(validaterErr)\n"
	contentString = contentString + "			resultDTO.Status = config.RESULT_ERROR\n"
	contentString = contentString + "		} else {\n"
	contentString = contentString + "			// 成功の場合にロジックを実行\n"
	if haveAutincliment {
		contentString = contentString + "			sid, errFlag := autologic.Register" + changeToGlobals(table.table_in.String) + "(form, r)\n"
		contentString = contentString + "			if errFlag == false {\n"
		contentString = contentString + "				resultDTO.Status = config.RESULT_ERROR\n"
		contentString = contentString + "			} else {\n"
		contentString = contentString + "				resultDTO.Status = config.RESULT_SUCCESS\n"
		for ind, _ := range columns {
			if columns[ind].fieldextra.String == "auto_increment" {
				contentString = contentString + "				resultDTO.Form." + changeToGlobals(columns[ind].field.String) + " = sid\n"
			}
		}
		contentString = contentString + "			}\n"
	} else {
		contentString = contentString + "			_, errFlag := autologic.Register" + changeToGlobals(table.table_in.String) + "(form, r)\n"
		contentString = contentString + "			if errFlag == false {\n"
		contentString = contentString + "				resultDTO.Status = config.RESULT_ERROR\n"
		contentString = contentString + "			} else {\n"
		contentString = contentString + "				resultDTO.Status = config.RESULT_SUCCESS\n"
		contentString = contentString + "			}\n"
	}
	contentString = contentString + "			tmpl, err = common.ViewParses(\"./auto/view/" + table.table_in.String + "/register.html\")\n"
	contentString = contentString + "		}\n"
	contentString = contentString + "	} else {\n"
	contentString = contentString + "		// 初期表示 SIDを指定されている場合でデータが取得できなかった場合は404画面へ遷移させる\n"
	contentString = contentString + "		if "
	addflag := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			if addflag {
				contentString = contentString + " && "
			}
			contentString = contentString + "form." + changeToGlobals(columns[ind].field.String) + " != \"\" "
			addflag = true
		}
	}
	contentString = contentString + " {\n"
	contentString = contentString + "			resultForm, getted := autologic.GetByPk" + changeToGlobals(table.table_in.String) + "(form, r) // 結果用DTOへFormを格納\n"
	contentString = contentString + "			resultDTO.Form = &resultForm\n"
	contentString = contentString + "			if getted == false {\n"
	contentString = contentString + "				http.Redirect(w, r, \"/notfound\", 404)\n"
	contentString = contentString + "			}\n"
	contentString = contentString + "		}\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "\n"
	contentString = contentString + "	if err != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r)\n"
	contentString = contentString + "		panic(err)\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "\n"
	contentString = contentString + "	err = tmpl.Execute(w, resultDTO)\n"
	contentString = contentString + "	if err != nil {\n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r)\n"
	contentString = contentString + "		panic(err)\n"
	contentString = contentString + "	}\n"
	contentString = contentString + "}\n"

	ioutil.WriteFile(dtoPath+"/"+changeToGlobals(table.table_in.String)+".go", []byte(contentString), os.ModePerm)

}
func makeLogic(table TableList, columns []ColumnList, cnt int) {
	dtoPath := outputPath + "/logic"

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}
	havePrimary := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			havePrimary = true
		}
	}
	if havePrimary == false {
		return
	}

	contentString := ""
	contentString = contentString + "package autologic \n"
	contentString = contentString + " \n"
	contentString = contentString + "import ( \n"
	contentString = contentString + "	\"../../config\" \n"
	contentString = contentString + "	\"../../service/common\" \n"
	contentString = contentString + "	\"../../service/db\" \n"
	contentString = contentString + "	\"../dto\" \n"
	contentString = contentString + "	\"../model\" \n"
	contentString = contentString + "	\"net/http\" \n"
	contentString = contentString + "	\"strconv\" \n"
	contentString = contentString + ") \n"
	contentString = contentString + " \n"
	contentString = contentString + "/* \n"
	contentString = contentString + " * 検索 \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func Search" + changeToGlobals(table.table_in.String) + "(form *autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request) []autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form { \n"
	contentString = contentString + "	common.WriteLog(config.INFO, \"Start\", r) \n"
	contentString = contentString + "	var resultList []autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form \n"
	contentString = contentString + " \n"
	contentString = contentString + "	selectResult, _ := automodel.Select" + changeToGlobals(table.table_in.String) + "(form, r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	for ind, _ := range selectResult { \n"
	contentString = contentString + "		resultList = append(resultList, autodbdto.DTF" + changeToGlobals(table.table_in.String) + "(selectResult[ind])) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	return resultList \n"
	contentString = contentString + " \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"
	contentString = contentString + "/* \n"
	contentString = contentString + " * PKから情報取得 \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func GetByPk" + changeToGlobals(table.table_in.String) + "(form *autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request) (autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, bool) { \n"
	contentString = contentString + "	common.WriteLog(config.INFO, \"Start\", r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	selectResult, hit, _ := automodel.GetByPK" + changeToGlobals(table.table_in.String) + "("
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			contentString = contentString + "form." + changeToGlobals(columns[ind].field.String) + ", "
		}
	}
	contentString = contentString + "r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if hit { \n"
	contentString = contentString + "		return autodbdto.DTF" + changeToGlobals(table.table_in.String) + "(selectResult), true \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	var rform autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form \n"
	contentString = contentString + "	return rform, false \n"
	contentString = contentString + " \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"
	contentString = contentString + "/* \n"
	contentString = contentString + " * 登録 \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func Register" + changeToGlobals(table.table_in.String) + "(form *autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request) (string, bool) { \n"
	contentString = contentString + "	common.WriteLog(config.INFO, \"Start\", r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	db := db.DbConn() \n"
	contentString = contentString + "	tx, _ := db.Begin() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	hit := false \n"
	contentString = contentString + "	sid := \"\" \n"

	contentString = contentString + "	if "
	addflag := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			if addflag {
				contentString = contentString + " && "
			}
			contentString = contentString + "form." + changeToGlobals(columns[ind].field.String) + " != \"\" "
			addflag = true
		}
	}
	contentString = contentString + " { \n"

	contentString = contentString + "		_, hit, _ = automodel.GetByPK" + changeToGlobals(table.table_in.String) + "ForUpdate(tx,"
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			contentString = contentString + "form." + changeToGlobals(columns[ind].field.String) + ", "
		}
	}
	contentString = contentString + "r) \n"
	//	contentString = contentString + "		if hit == false { \n"
	//	contentString = contentString + "			return \"\", false \n"
	//	contentString = contentString + "		} \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	lid, writed := int64(0), false \n"
	contentString = contentString + "	if hit { \n"
	contentString = contentString + "		writed = automodel.Update" + changeToGlobals(table.table_in.String) + "(tx, form, r, []string{\"sid\"}) \n"
	for ind, _ := range columns {
		if columns[ind].fieldextra.String == "auto_increment" {
			contentString = contentString + "		sid = form." + changeToGlobals(columns[ind].field.String) + " \n"
		}
	}
	contentString = contentString + "	} else { \n"
	contentString = contentString + "		lid, writed = automodel.Insert" + changeToGlobals(table.table_in.String) + "(tx, form, r, []string{\"sid\"}) \n"
	contentString = contentString + "		sid = strconv.FormatInt(lid, 10) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if writed { \n"
	contentString = contentString + "		tx.Commit() \n"
	contentString = contentString + "		return sid, true \n"
	contentString = contentString + "	} else { \n"
	contentString = contentString + "		tx.Rollback() \n"
	contentString = contentString + "		return \"\", false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "} \n"

	ioutil.WriteFile(dtoPath+"/"+changeToGlobals(table.table_in.String)+".go", []byte(contentString), os.ModePerm)
}

func makeModel(table TableList, columns []ColumnList, cnt int) {
	dtoPath := outputPath + "/model"

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}

	havePrimary := false
	for ind, _ := range columns {
		if columns[ind].fieldkey.String == "PRI" {
			havePrimary = true
		}
	}

	contentString := ""
	contentString = contentString + "package automodel \n"
	contentString = contentString + " \n"
	contentString = contentString + "import ( \n"
	contentString = contentString + "	\"../../config\" \n"
	contentString = contentString + "	\"../../service/common\" \n"
	contentString = contentString + "	\"../../service/db\" \n"
	contentString = contentString + "	\"../dto\" \n"
	contentString = contentString + "	\"net/http\" \n"
	contentString = contentString + "	\"strings\" \n"
	contentString = contentString + "	\"database/sql\" \n"
	contentString = contentString + ") \n"
	contentString = contentString + " \n"
	added := false
	if havePrimary {
		contentString = contentString + "/* \n"
		contentString = contentString + " * PKからデータを取得する \n"
		contentString = contentString + " * 取得できなかった場合は空の構造体とfalseを返す \n"
		contentString = contentString + " */ \n"
		contentString = contentString + "func GetByPK" + changeToGlobals(table.table_in.String) + "("

		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				contentString = contentString + columns[ind].field.String + " string, "
			}
		}
		contentString = contentString + "r *http.Request) (autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO, bool, bool) { \n"
		contentString = contentString + "	db := db.DbConn() \n"
		contentString = contentString + " \n"
		contentString = contentString + "	sql := \"select * from " + table.table_in.String + " where "
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + "and "
				}
				contentString = contentString + columns[ind].field.String + " = ? "
				added = true
			}
		}
		contentString = contentString + "\"\n"
		contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		var rv autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "		return rv, false, false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	defer stmt.Close() \n"
		contentString = contentString + " \n"
		contentString = contentString + "	rows, err := stmt.Query("
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + ", "
				}
				contentString = contentString + columns[ind].field.String + ""
				added = true
			}
		}
		contentString = contentString + ") \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, "
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + "+"
				}
				contentString = contentString + columns[ind].field.String + ""
				added = true
			}
		}
		contentString = contentString + ", r) \n"

		contentString = contentString + " \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		var rv autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "		return rv, false, false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	defer rows.Close() \n"
		contentString = contentString + "	for rows.Next() { \n"
		contentString = contentString + "		var columns autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + " \n"
		contentString = contentString + "		err := rows.Scan( \n"
		for ind, _ := range columns {
			contentString = contentString + "			&columns." + changeToGlobals(columns[ind].field.String) + ", \n"
		}
		contentString = contentString + "		) \n"
		contentString = contentString + " \n"
		contentString = contentString + "		if err != nil { \n"
		contentString = contentString + "			common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "			var rv autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "			return rv, false, false \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "		return columns, true, true \n"
		contentString = contentString + "	} \n"
		contentString = contentString + " \n"
		contentString = contentString + "	var rdto autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "	return rdto, false, true \n"
		contentString = contentString + "} \n"

		contentString = contentString + "/* \n"
		contentString = contentString + " * PKからデータを取得しつつロックをかける \n"
		contentString = contentString + " * 取得できなかった場合は空の構造体とfalseを返す \n"
		contentString = contentString + " */ \n"
		contentString = contentString + "func GetByPK" + changeToGlobals(table.table_in.String) + "ForUpdate(db *sql.Tx, "

		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				contentString = contentString + columns[ind].field.String + " string, "
			}
		}
		contentString = contentString + "r *http.Request) (autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO, bool, bool) { \n"
		//	contentString = contentString + "	db := db.DbConn() \n"
		contentString = contentString + " \n"
		contentString = contentString + "	sql := \"select * from " + table.table_in.String + " where "
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + "and "
				}
				contentString = contentString + columns[ind].field.String + " = ? "
				added = true
			}
		}
		contentString = contentString + "\"\n"
		contentString = contentString + "	sql = sql + \" for update \" \n"
		contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		var rv autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "		return rv, false, false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	defer stmt.Close() \n"
		contentString = contentString + " \n"
		contentString = contentString + "	rows, err := stmt.Query("
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + ", "
				}
				contentString = contentString + columns[ind].field.String + ""
				added = true
			}
		}
		contentString = contentString + ") \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, "
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + "+"
				}
				contentString = contentString + columns[ind].field.String + ""
				added = true
			}
		}
		contentString = contentString + ", r) \n"

		contentString = contentString + " \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		var rv autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "		return rv, false, false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	defer rows.Close() \n"
		contentString = contentString + "	for rows.Next() { \n"
		contentString = contentString + "		var columns autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + " \n"
		contentString = contentString + "		err := rows.Scan( \n"
		for ind, _ := range columns {
			contentString = contentString + "			&columns." + changeToGlobals(columns[ind].field.String) + ", \n"
		}
		contentString = contentString + "		) \n"
		contentString = contentString + " \n"
		contentString = contentString + "		if err != nil { \n"
		contentString = contentString + "			common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "			var rv autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "			return rv, false, false \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "		return columns, true, true \n"
		contentString = contentString + "	} \n"
		contentString = contentString + " \n"
		contentString = contentString + "	var rdto autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
		contentString = contentString + "	return rdto, false, true \n"
		contentString = contentString + "} \n"
	}
	contentString = contentString + " \n"
	contentString = contentString + "/* \n"
	contentString = contentString + " * 指定したカラムに対してIN句を発行する \n"
	contentString = contentString + " * 取得できなかった場合は空の配列を返す \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func SelectByIn" + changeToGlobals(table.table_in.String) + "(targetColumn string, in []string, r *http.Request) ([]autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO, bool) { \n"
	contentString = contentString + "	var rdto []autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
	contentString = contentString + " \n"
	contentString = contentString + "	db := db.DbConn() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	var arr []string \n"
	contentString = contentString + " \n"
	contentString = contentString + "	sql := \"select * from " + table.table_in.String + " where \" + targetColumn + \" in (\" \n"
	contentString = contentString + "	where := \"\" \n"
	contentString = contentString + "	for ind, _ := range in { \n"
	contentString = contentString + "		if where != \"\" { \n"
	contentString = contentString + "			where = where + \",\" \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		where = where + \"?\" \n"
	contentString = contentString + "		arr = append(arr, in[ind]) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + where + \")\" \n"
	contentString = contentString + "	if where == \"\" { \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer stmt.Close() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
	contentString = contentString + "	for i, _ := range arr { \n"
	contentString = contentString + "		dest[i] = &arr[i] \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	rows, err := stmt.Query(dest...) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer rows.Close() \n"
	contentString = contentString + "	for rows.Next() { \n"
	contentString = contentString + "		var columns autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
	contentString = contentString + " \n"
	contentString = contentString + "		err := rows.Scan( \n"
	for ind, _ := range columns {
		contentString = contentString + "			&columns." + changeToGlobals(columns[ind].field.String) + ", \n"
	}
	contentString = contentString + "		) \n"
	contentString = contentString + " \n"
	contentString = contentString + "		if err != nil { \n"
	contentString = contentString + "			common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "			return rdto, false \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		rdto = append(rdto, columns) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	return rdto, true \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"

	contentString = contentString + "/* \n"
	contentString = contentString + " * 指定したカラムに対してIN句を発行しつつロックをかける \n"
	contentString = contentString + " * 取得できなかった場合は空の配列を返す \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func SelectByIn" + changeToGlobals(table.table_in.String) + "ForUpdate(db *sql.Tx, targetColumn string, in []string, r *http.Request) ([]autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO, bool) { \n"
	contentString = contentString + "	var rdto []autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
	contentString = contentString + " \n"
	//contentString = contentString + "	db := db.DbConn() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	var arr []string \n"
	contentString = contentString + " \n"
	contentString = contentString + "	sql := \"select * from " + table.table_in.String + " where \" + targetColumn + \" in (\" \n"
	contentString = contentString + "	where := \"\" \n"
	contentString = contentString + "	for ind, _ := range in { \n"
	contentString = contentString + "		if where != \"\" { \n"
	contentString = contentString + "			where = where + \",\" \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		where = where + \"?\" \n"
	contentString = contentString + "		arr = append(arr, in[ind]) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + where + \")\" \n"
	contentString = contentString + "	if where == \"\" { \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + \" for update \" \n"
	contentString = contentString + " \n"
	contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer stmt.Close() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
	contentString = contentString + "	for i, _ := range arr { \n"
	contentString = contentString + "		dest[i] = &arr[i] \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	rows, err := stmt.Query(dest...) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer rows.Close() \n"
	contentString = contentString + "	for rows.Next() { \n"
	contentString = contentString + "		var columns autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
	contentString = contentString + " \n"
	contentString = contentString + "		err := rows.Scan( \n"
	for ind, _ := range columns {
		contentString = contentString + "			&columns." + changeToGlobals(columns[ind].field.String) + ", \n"
	}
	contentString = contentString + "		) \n"
	contentString = contentString + " \n"
	contentString = contentString + "		if err != nil { \n"
	contentString = contentString + "			common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "			return rdto, false \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		rdto = append(rdto, columns) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	return rdto, true \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"

	contentString = contentString + "/* \n"
	contentString = contentString + " * Formデータに含まれる情報を元にSQLを発行する \n"
	contentString = contentString + " * 空白の場合は検索対象としない \n"
	contentString = contentString + " * 取得できなかった場合は、空の配列を返す \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func Select" + changeToGlobals(table.table_in.String) + "(search *autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request) ([]autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO, bool) { \n"
	contentString = contentString + "	var rdto []autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
	contentString = contentString + " \n"
	contentString = contentString + "	db := db.DbConn() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	var arr []string \n"
	contentString = contentString + " \n"
	contentString = contentString + "	sql := \"select * from " + table.table_in.String + "\" \n"
	contentString = contentString + "	where := \"\" \n"
	contentString = contentString + " \n"
	for ind, _ := range columns {
		contentString = contentString + "	where, arr = AppendWhere(search." + changeToGlobals(columns[ind].field.String) + ", \"" + columns[ind].field.String + "\", where, arr) \n"
	}
	contentString = contentString + " \n"
	contentString = contentString + "	if where != \"\" { \n"
	contentString = contentString + "		sql = sql + \" where \" + where \n"
	contentString = contentString + " \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer stmt.Close() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
	contentString = contentString + "	for i, _ := range arr { \n"
	contentString = contentString + "		dest[i] = &arr[i] \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	rows, err := stmt.Query(dest...) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return rdto, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer rows.Close() \n"
	contentString = contentString + "	for rows.Next() { \n"
	contentString = contentString + "		var columns autodbdto.DB" + changeToGlobals(table.table_in.String) + "DTO \n"
	contentString = contentString + " \n"
	contentString = contentString + "		err := rows.Scan( \n"
	for ind, _ := range columns {
		contentString = contentString + "			&columns." + changeToGlobals(columns[ind].field.String) + ", \n"
	}
	contentString = contentString + "		) \n"
	contentString = contentString + " \n"
	contentString = contentString + "		if err != nil { \n"
	contentString = contentString + "			common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "			return rdto, false \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		rdto = append(rdto, columns) \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	return rdto, true \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"

	contentString = contentString + "/* \n"
	contentString = contentString + " * Formデータに含まれる情報を追加する \n"
	contentString = contentString + " * 除外したいカラムはexcludesへ配列で格納する \n"
	contentString = contentString + " * 成功時には追加したIDを返す \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func Insert" + changeToGlobals(table.table_in.String) + "(db *sql.Tx, search *autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request, excludes []string) (int64, bool) { \n"
	//	contentString = contentString + "	db := db.DbConn() \n"
	contentString = contentString + "	columnSql := \"\" \n"
	contentString = contentString + "	valueSql := \"\" \n"
	contentString = contentString + "	var arr []string \n"
	contentString = contentString + "	columns := []string{ \n"
	for ind, _ := range columns {
		contentString = contentString + "		\"" + columns[ind].field.String + "\", \n"
	}
	contentString = contentString + "	} \n"
	contentString = contentString + "	values := []string{ \n"
	for ind, _ := range columns {
		contentString = contentString + "		search." + changeToGlobals(columns[ind].field.String) + ", \n"
	}
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql := \"INSERT INTO " + table.table_in.String + " (\" \n"
	contentString = contentString + "	for ind, _ := range columns { \n"
	contentString = contentString + "		if IsColumnExcludes(columns[ind], excludes) == false { \n"
	contentString = contentString + "			if columnSql != \"\" { \n"
	contentString = contentString + "				columnSql = columnSql + \",\" \n"
	contentString = contentString + "			} \n"
	contentString = contentString + "			columnSql = columnSql + \"`\" + columns[ind] + \"`\" \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + columnSql + \") VALUES (\" \n"
	contentString = contentString + "	for ind, _ := range columns { \n"
	contentString = contentString + "		if IsColumnExcludes(columns[ind], excludes) == false { \n"
	contentString = contentString + "			if valueSql != \"\" { \n"
	contentString = contentString + "				valueSql = valueSql + \",\" \n"
	contentString = contentString + "			} \n"
	contentString = contentString + "			valueSql = valueSql + \"?\" \n"
	contentString = contentString + "			arr = append(arr, values[ind]) \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + valueSql + \")\" \n"
	contentString = contentString + " \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
	contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return 0, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer stmt.Close() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
	contentString = contentString + "	for i, _ := range arr { \n"
	contentString = contentString + "		dest[i] = &arr[i] \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	res, err := stmt.Exec(dest...) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return 0, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	id, _ := res.LastInsertId() \n"
	contentString = contentString + "	return id, true \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"

	if havePrimary {
		contentString = contentString + "/* \n"
		contentString = contentString + " * Formデータに含まれる情報を更新する \n"
		contentString = contentString + " * 除外したいカラムはexcludesへ配列で格納する \n"
		contentString = contentString + " */ \n"
		contentString = contentString + "func Update" + changeToGlobals(table.table_in.String) + "(db *sql.Tx, search *autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request, excludes []string) bool { \n"
		//contentString = contentString + "	db := db.DbConn() \n"
		contentString = contentString + "	columnSql := \"\" \n"
		contentString = contentString + "	var arr []string \n"
		contentString = contentString + "	columns := []string{ \n"
		for ind, _ := range columns {
			contentString = contentString + "		\"" + columns[ind].field.String + "\", \n"
		}
		contentString = contentString + "	} \n"
		contentString = contentString + "	values := []string{ \n"
		for ind, _ := range columns {
			contentString = contentString + "		search." + changeToGlobals(columns[ind].field.String) + ", \n"
		}
		contentString = contentString + "	} \n"
		contentString = contentString + "	sql := \"UPDATE " + table.table_in.String + " SET \" \n"
		contentString = contentString + "	for ind, _ := range columns { \n"
		contentString = contentString + "		if IsColumnExcludes(columns[ind], excludes) == false { \n"
		contentString = contentString + "			if columnSql != \"\" { \n"
		contentString = contentString + "				columnSql = columnSql + \",\" \n"
		contentString = contentString + "			} \n"
		contentString = contentString + "			columnSql = columnSql + \"`\" + columns[ind] + \"` = ?\" \n"
		contentString = contentString + "			arr = append(arr, values[ind]) \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	sql = sql + columnSql \n"
		contentString = contentString + "	sql = sql + \" where "
		added = false
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				if added {
					contentString = contentString + "and "
				}
				contentString = contentString + " " + columns[ind].field.String + " = ? "
				added = true
			}
		}
		contentString = contentString + "\" \n"
		for ind, _ := range columns {
			if columns[ind].fieldkey.String == "PRI" {
				contentString = contentString + "	arr = append(arr, search." + changeToGlobals(columns[ind].field.String) + ") \n"
			}
		}
		contentString = contentString + " \n"
		contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		return false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	defer stmt.Close() \n"
		contentString = contentString + " \n"
		contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
		contentString = contentString + "	for i, _ := range arr { \n"
		contentString = contentString + "		dest[i] = &arr[i] \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	stmt.Exec(dest...) \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
		contentString = contentString + " \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		return false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	return true \n"
		contentString = contentString + "} \n"
	}

	contentString = contentString + "/* \n"
	contentString = contentString + " * バルクインサート \n"
	contentString = contentString + " */ \n"
	contentString = contentString + "func BulkInsert" + changeToGlobals(table.table_in.String) + "(db *sql.Tx, search []*autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request, excludes []string) (int64, bool) { \n"
	//contentString = contentString + "	db := db.DbConn() \n"
	contentString = contentString + "	columnSql := \"\" \n"
	contentString = contentString + "	valueSql := \"\" \n"
	contentString = contentString + "	var arr []string \n"
	contentString = contentString + "	columns := []string{ \n"
	for ind, _ := range columns {
		contentString = contentString + "		\"" + columns[ind].field.String + "\", \n"
	}
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql := \"INSERT INTO t_inquiry (\" \n"
	contentString = contentString + "	for ind, _ := range columns { \n"
	contentString = contentString + "		if IsColumnExcludes(columns[ind], excludes) == false { \n"
	contentString = contentString + "			if columnSql != \"\" { \n"
	contentString = contentString + "				columnSql = columnSql + \",\" \n"
	contentString = contentString + "			} \n"
	contentString = contentString + "			columnSql = columnSql + \"`\" + columns[ind] + \"`\" \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + columnSql + \") VALUES \" \n"
	contentString = contentString + "	for dataInd, _ := range search { \n"
	contentString = contentString + "		if dataInd != 0 { \n"
	contentString = contentString + "			valueSql = valueSql + \",\" \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		valueSql = valueSql + \"(\" \n"
	contentString = contentString + "		valueSqlSub := \"\" \n"
	contentString = contentString + "		for ind, _ := range columns { \n"
	contentString = contentString + "			if IsColumnExcludes(columns[ind], excludes) == false { \n"
	contentString = contentString + "				if valueSqlSub != \"\" { \n"
	contentString = contentString + "					valueSqlSub = valueSqlSub + \",\" \n"
	contentString = contentString + "				} \n"
	contentString = contentString + "				valueSqlSub = valueSqlSub + \"?\" \n"
	contentString = contentString + "				switch ind { \n"
	for ind, _ := range columns {
		contentString = contentString + "				case " + strconv.Itoa(ind) + ": \n"
		contentString = contentString + "					arr = append(arr, search[dataInd]." + changeToGlobals(columns[ind].field.String) + ") \n"
	}
	contentString = contentString + "				} \n"
	contentString = contentString + "			} \n"
	contentString = contentString + "		} \n"
	contentString = contentString + "		valueSql = valueSql + valueSqlSub + \")\" \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	sql = sql + valueSql + \"\" \n"
	contentString = contentString + " \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
	contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
	contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return 0, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	defer stmt.Close() \n"
	contentString = contentString + " \n"
	contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
	contentString = contentString + "	for i, _ := range arr { \n"
	contentString = contentString + "		dest[i] = &arr[i] \n"
	contentString = contentString + "	} \n"
	contentString = contentString + "	res, err := stmt.Exec(dest...) \n"
	contentString = contentString + " \n"
	contentString = contentString + "	if err != nil { \n"
	contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
	contentString = contentString + "		return 0, false \n"
	contentString = contentString + "	} \n"
	contentString = contentString + " \n"
	contentString = contentString + "	id, _ := res.LastInsertId() \n"
	contentString = contentString + "	return id, true \n"
	contentString = contentString + "} \n"
	contentString = contentString + " \n"

	if havePrimary {
		contentString = contentString + "/* \n"
		contentString = contentString + " * バルクアップデート \n"
		contentString = contentString + " */ \n"
		contentString = contentString + "func BulkUpdate" + changeToGlobals(table.table_in.String) + "(db *sql.Tx, search []*autodbdto.DB" + changeToGlobals(table.table_in.String) + "Form, r *http.Request, excludes []string) (int64, bool) { \n"
		//contentString = contentString + "	db := db.DbConn() \n"
		contentString = contentString + "	columnSql := \"\" \n"
		contentString = contentString + "	valueSql := \"\" \n"
		contentString = contentString + "	duplicateKeys := \"\" \n"
		contentString = contentString + " \n"
		contentString = contentString + "	var arr []string \n"
		contentString = contentString + "	columns := []string{ \n"
		for ind, _ := range columns {
			contentString = contentString + "		\"" + columns[ind].field.String + "\", \n"
		}
		contentString = contentString + "	} \n"
		contentString = contentString + "	sql := \"INSERT INTO t_inquiry (\" \n"
		contentString = contentString + "	for ind, _ := range columns { \n"
		contentString = contentString + "		if IsColumnExcludes(columns[ind], excludes) == false { \n"
		contentString = contentString + "			if columnSql != \"\" { \n"
		contentString = contentString + "				columnSql = columnSql + \",\" \n"
		contentString = contentString + "			} \n"
		contentString = contentString + "			columnSql = columnSql + \"`\" + columns[ind] + \"`\" \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	sql = sql + columnSql + \") VALUES \" \n"
		contentString = contentString + "	for dataInd, _ := range search { \n"
		contentString = contentString + "		if dataInd != 0 { \n"
		contentString = contentString + "			valueSql = valueSql + \",\" \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "		valueSql = valueSql + \"(\" \n"
		contentString = contentString + "		valueSqlSub := \"\" \n"
		contentString = contentString + "		for ind, _ := range columns { \n"
		contentString = contentString + "			if IsColumnExcludes(columns[ind], excludes) == false { \n"
		contentString = contentString + "				if valueSqlSub != \"\" { \n"
		contentString = contentString + "					valueSqlSub = valueSqlSub + \",\" \n"
		contentString = contentString + "					if dataInd == 0 { \n"
		contentString = contentString + "						duplicateKeys = duplicateKeys + \",\" \n"
		contentString = contentString + "					} \n"
		contentString = contentString + "				} \n"
		contentString = contentString + "				valueSqlSub = valueSqlSub + \"?\" \n"
		contentString = contentString + "				if dataInd == 0 { \n"
		contentString = contentString + "					duplicateKeys = duplicateKeys + \"\" + columns[ind] + \" = VALUES(\" + columns[ind] + \")\" \n"
		contentString = contentString + "				} \n"
		contentString = contentString + "				switch ind { \n"
		for ind, _ := range columns {
			contentString = contentString + "				case " + strconv.Itoa(ind) + ": \n"
			contentString = contentString + "					arr = append(arr, search[dataInd]." + changeToGlobals(columns[ind].field.String) + ") \n"
		}
		contentString = contentString + "				} \n"
		contentString = contentString + "			} \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "		valueSql = valueSql + valueSqlSub + \")\" \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	sql = sql + valueSql + \"\" \n"
		contentString = contentString + "	sql = sql + \" ON DUPLICATE KEY UPDATE \" + duplicateKeys \n"
		contentString = contentString + " \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, sql, r) \n"
		contentString = contentString + "	common.WriteLog(config.DEBUG, strings.Join(arr, \",\"), r) \n"
		contentString = contentString + "	stmt, err := db.Prepare(sql) \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		return 0, false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	defer stmt.Close() \n"
		contentString = contentString + " \n"
		contentString = contentString + "	dest := make([]interface{}, len(arr)) \n"
		contentString = contentString + "	for i, _ := range arr { \n"
		contentString = contentString + "		dest[i] = &arr[i] \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	res, err := stmt.Exec(dest...) \n"
		contentString = contentString + " \n"
		contentString = contentString + "	if err != nil { \n"
		contentString = contentString + "		common.WriteErrorLog(config.FATAL, err, r) \n"
		contentString = contentString + "		return 0, false \n"
		contentString = contentString + "	} \n"
		contentString = contentString + " \n"
		contentString = contentString + "	id, _ := res.LastInsertId() \n"
		contentString = contentString + "	return id, true \n"
		contentString = contentString + "} \n"
		contentString = contentString + " \n"
	}

	if cnt == 1 {
		contentString = contentString + "func IsColumnExcludes(t string, excludes []string) bool { \n"
		contentString = contentString + "	for ind, _ := range excludes { \n"
		contentString = contentString + "		if excludes[ind] == t { \n"
		contentString = contentString + "			return true \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "	} \n"
		contentString = contentString + " \n"
		contentString = contentString + "	return false \n"
		contentString = contentString + "} \n"
		contentString = contentString + "func AppendWhere(s string, c string, w string, arr []string) (string, []string) { \n"
		contentString = contentString + "	if s != \"\" { \n"
		contentString = contentString + "		if w != \"\" { \n"
		contentString = contentString + "			w = w + \" and \" \n"
		contentString = contentString + "		} \n"
		contentString = contentString + "		w = w + \" \" + c + \" = ? \" \n"
		contentString = contentString + "		arr = append(arr, s) \n"
		contentString = contentString + "	} \n"
		contentString = contentString + "	return w, arr \n"
		contentString = contentString + "} \n"

	}

	ioutil.WriteFile(dtoPath+"/"+changeToGlobals(table.table_in.String)+".go", []byte(contentString), os.ModePerm)
}

func makeDto(table TableList, columns []ColumnList, cnt int) {
	dtoPath := outputPath + "/dto"

	if err := os.MkdirAll(dtoPath, 0777); err != nil {
		panic(err)
	}

	contentString := ""
	if cnt == 0 {
		contentString = contentString + "type NullTime struct { \n"
		contentString = contentString + "    Time  time.Time \n"
		contentString = contentString + "    Valid bool // Valid is true if Time is not NULL \n"
		contentString = contentString + "} \n"
		contentString = contentString + " \n"
		contentString = contentString + "// Scan implements the Scanner interface. \n"
		contentString = contentString + "func (nt *NullTime) Scan(value interface{}) error { \n"
		contentString = contentString + "    nt.Time, nt.Valid = value.(time.Time) \n"
		contentString = contentString + "    return nil \n"
		contentString = contentString + "} \n"
		contentString = contentString + " \n"
		contentString = contentString + "// Value implements the driver Valuer interface. \n"
		contentString = contentString + "func (nt NullTime) Value() (driver.Value, error) { \n"
		contentString = contentString + "    if !nt.Valid { \n"
		contentString = contentString + "        return nil, nil \n"
		contentString = contentString + "    } \n"
		contentString = contentString + "    return nt.Time, nil \n"
		contentString = contentString + "} \n"
	}
	contentString = contentString + "type DB" + changeToGlobals(table.table_in.String) + "DTO struct {\n"
	for ind, _ := range columns {
		contentString = contentString + "\t" + changeToStructType(columns[ind], false) + " \n"
	}
	contentString = contentString + "}\n"

	contentString = contentString + "type DB" + changeToGlobals(table.table_in.String) + "Form struct {\n"
	contentString = contentString + "\tMode string \n"
	for ind, _ := range columns {
		contentString = contentString + "\t" + changeToStructType(columns[ind], true) + " \n"
	}
	contentString = contentString + "}\n"

	contentString = contentString + "type DB" + changeToGlobals(table.table_in.String) + "ResultDTO struct {\n"
	contentString = contentString + "\tToken string \n"
	contentString = contentString + "\tStatus int \n"
	contentString = contentString + "\tForm *DB" + changeToGlobals(table.table_in.String) + "Form \n"
	contentString = contentString + "\tList []DB" + changeToGlobals(table.table_in.String) + "Form \n"
	contentString = contentString + "\tErrors []dto.ErrorForm \n"
	contentString = contentString + "}\n"

	contentString = contentString + "func (cf *" + "DB" + changeToGlobals(table.table_in.String) + "Form) FieldMap(req *http.Request) binding.FieldMap { \n"
	contentString = contentString + "	return binding.FieldMap{ \n"
	contentString = contentString + "		&cf.Mode: \"mode\",\n"
	for ind, _ := range columns {
		contentString = contentString + "		&cf." + changeToGlobals(columns[ind].field.String) + ":     \"" + columns[ind].field.String + "\", \n"
	}
	contentString = contentString + "	} \n"
	contentString = contentString + "} \n"

	contentString = contentString + "func DTF" + changeToGlobals(table.table_in.String) + "(dto DB" + changeToGlobals(table.table_in.String) + "DTO) DB" + changeToGlobals(table.table_in.String) + "Form { \n"
	contentString = contentString + "	var form DB" + changeToGlobals(table.table_in.String) + "Form \n"
	for ind, _ := range columns {
		contentString = contentString + "	form." + changeToGlobals(columns[ind].field.String) + " = \"\" \n"
		if strings.Contains(columns[ind].fieldtype.String, "char") {
			contentString = contentString + "	if dto." + changeToGlobals(columns[ind].field.String) + ".Valid == true { \n"
			contentString = contentString + "		form." + changeToGlobals(columns[ind].field.String) + " = dto." + changeToGlobals(columns[ind].field.String) + ".String \n"
			contentString = contentString + "	} \n"
		}
		if strings.Contains(columns[ind].fieldtype.String, "text") {
			contentString = contentString + "	if dto." + changeToGlobals(columns[ind].field.String) + ".Valid == true { \n"
			contentString = contentString + "		form." + changeToGlobals(columns[ind].field.String) + " = dto." + changeToGlobals(columns[ind].field.String) + ".String \n"
			contentString = contentString + "	} \n"
		}
		if strings.Contains(columns[ind].fieldtype.String, "int") {
			contentString = contentString + "	if dto." + changeToGlobals(columns[ind].field.String) + ".Valid == true { \n"
			contentString = contentString + "		form." + changeToGlobals(columns[ind].field.String) + " = strconv.FormatInt(dto." + changeToGlobals(columns[ind].field.String) + ".Int64, 10) \n"
			contentString = contentString + "	} \n"
		}
		if strings.Contains(columns[ind].fieldtype.String, "float") {
			contentString = contentString + "	if dto." + changeToGlobals(columns[ind].field.String) + ".Valid == true { \n"
			contentString = contentString + "		form." + changeToGlobals(columns[ind].field.String) + " = strconv.FormatFloat(dto." + changeToGlobals(columns[ind].field.String) + ".Float64, 'f', -1, 64) \n"
			contentString = contentString + "	} \n"
		}
		if strings.Contains(columns[ind].fieldtype.String, "bool") {
			contentString = contentString + "	if dto." + changeToGlobals(columns[ind].field.String) + ".Valid == true { \n"
			contentString = contentString + "		form." + changeToGlobals(columns[ind].field.String) + " = strconv.FormatBool(dto." + changeToGlobals(columns[ind].field.String) + ".Bool, 10) \n"
			contentString = contentString + "	} \n"
		}
		if strings.Contains(columns[ind].fieldtype.String, "time") {
			contentString = contentString + "	if dto." + changeToGlobals(columns[ind].field.String) + ".Valid == true { \n"
			contentString = contentString + "		form." + changeToGlobals(columns[ind].field.String) + " = dto." + changeToGlobals(columns[ind].field.String) + ".Time.Format(\"" + timelayout + "\") \n"
			contentString = contentString + "	} \n"
		}
	}
	contentString = contentString + "	return form \n"
	contentString = contentString + "} \n"

	imports := "package autodbdto\n"
	imports = imports + "\n"
	imports = imports + "import (\n"
	imports = imports + "	\"github.com/mholt/binding\"\n"
	imports = imports + "	\"net/http\"\n"
	imports = imports + "	\"" + serviceDtoPath + "\"\n"

	if strings.Contains(contentString, "sql.") {
		imports = imports + "	\"database/sql\"\n"
	}
	if strings.Contains(contentString, "time.Time") {
		imports = imports + "	\"time\"\n"
	}
	if strings.Contains(contentString, "strconv.") {
		imports = imports + "	\"strconv\"\n"
	}
	if strings.Contains(contentString, "driver.") {
		imports = imports + "	\"database/sql/driver\"\n"
	}
	imports = imports + ")\n"
	imports = imports + "\n"

	ioutil.WriteFile(dtoPath+"/"+changeToGlobals(table.table_in.String)+".go", []byte(imports+contentString), os.ModePerm)
}

func changeToStructType(column ColumnList, addValidate bool) string {
	rs := changeToGlobals(column.field.String)
	indStrings := strings.Split(column.fieldtype.String, "(")
	ind := "0"
	if indStrings != nil && len(indStrings) > 1 {
		ind = strings.Split(indStrings[1], ")")[0]
	}

	requred := ""
	if column.fieldnull.String == "NO" {
		requred = "required"
	}

	if addValidate == false {
		if strings.Contains(column.fieldtype.String, "char") {
			rs = rs + " sql.NullString"
		}
		if strings.Contains(column.fieldtype.String, "text") {
			rs = rs + " sql.NullString"
		}
		if strings.Contains(column.fieldtype.String, "int") {
			rs = rs + " sql.NullInt64"
		}
		if strings.Contains(column.fieldtype.String, "float") {
			rs = rs + " sql.NullFloat64"

		}
		if strings.Contains(column.fieldtype.String, "bool") {
			rs = rs + " sql.NullBool"
		}
		if strings.Contains(column.fieldtype.String, "time") {
			rs = rs + " NullTime"
		}
	} else {
		validate := ""
		if requred != "" {
			validate = validate + "required"
		}
		if ind != "" && ind != "0" {
			if validate != "" {
				validate = validate + ","
			}
			validate = validate + "max=" + ind
		}
		rs = rs + " string"
		if validate != "" {
			rs = rs + "" + " `validate:\"" + validate + "\"`"
		}
	}
	return rs
}
func changeToGlobals(s string) string {
	ss := strings.Split(s, "_")
	rs := ""
	for ind, _ := range ss {
		rs = rs + strings.ToUpper(ss[ind][:1])
		rs = rs + ss[ind][1:]
	}

	return rs
}

func getTableList() []TableList {
	db := db.DbConn()
	sql := "SHOW TABLES FROM " + config.DBNAME
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	common.WriteLog(config.DEBUG, sql, nil)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tableList []TableList
	for rows.Next() {
		var columns TableList

		err := rows.Scan(
			&columns.table_in,
		)

		if err != nil {
			panic(err)
		} else {
			tableList = append(tableList, columns)
		}
	}
	return tableList
}

func getColumnList(table TableList) []ColumnList {
	db := db.DbConn()
	var columnList []ColumnList

	sql := "show FULL columns from " + table.table_in.String
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	common.WriteLog(config.DEBUG, sql, nil)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var columns ColumnList

		err := rows.Scan(
			&columns.field,
			&columns.fieldtype,
			&columns.fieldcollation,
			&columns.fieldnull,
			&columns.fieldkey,
			&columns.fielddefault,
			&columns.fieldextra,
			&columns.fieldprivileges,
			&columns.fieldcomment,
		)

		if err != nil {
			panic(err)
		} else {
			columnList = append(columnList, columns)
		}
	}

	return columnList
}
