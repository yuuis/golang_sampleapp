package common

import (
	"errors"
	"html/template"
	"path/filepath"
	"strings"
)

/*
 * 呼び出し先テンプレートの共通化処理
 * define は{{define "name"}}～{{end}}で囲われている範囲がdefineとして定義されるため、1ファイルにまとめても良い
 * dict は呼び出し元で複数の引数を渡すために使用される
 */
func ViewParses(t string) (*template.Template, error) {
	tname := filepath.Base(t)
	var tmpl = template.Must(template.New(tname).Funcs(template.FuncMap{
		// Pass values ​​to children
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		// Replaces newlines with <br>
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
		// Skips sanitation on the parameter.  Do not use with dynamic data.
		"raw": func(text string) template.HTML {
			return template.HTML(text)
		},
	}).ParseFiles(
		t, // メインは先頭でロードする
		"./view/common/templates.html",
	))
	return tmpl, nil
}
