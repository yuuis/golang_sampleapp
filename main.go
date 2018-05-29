package main

/*
 * go get github.com/gorilla/sessions
 * go get golang.org/x/text/encoding/japanese
 * go get gopkg.in/go-playground/validator.v9
 * go get github.com/mholt/binding
 * go get github.com/gorilla/sessions
 * go get github.com/bradleypeabody/gorilla-sessions-memcache
 * go get "github.com/go-sql-driver/mysql"
 * go get github.com/bradfitz/gomemcache/memcache
 * go get -u github.com/go-redis/redis
 * go get github.com/stretchr/testify/assert
 * validater document http://godoc.org/gopkg.in/go-playground/validator.v9
 *
 * go build ./src/server/main.go ./bin
 */
import (
	"./auto"
	"./service/common"
	"./service/controller"
	"./service/data"
	"./service/db"
	"log"
	"net/http"
)

func authenticate(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before process") // 処理の前の共通処理

		user := data.GetStringSession(r, "user")
		log.Println("pre_sessioon: " + user)
		log.Println(r.URL.Path)
		if user != "" || r.URL.Path == "/loginasynchronous" || r.URL.Path == "/login"{
			fn(w, r)
		} else {
			controller.GetLoginViewHandler(w, r)
		}

		log.Println("after process") // 処理の後の共通処理
		log.Println("session: " + data.GetStringSession(r, "user"))
	}
}

func main() {
	_, err := db.DbInit() // DBコネクションをとっておく
	if err != nil {
		panic(err)
	}
	db.InitRedis()         // Redisコネクションをとっておく
	db.InitMemcache()      // Memcachedコネクションをとっておく
	data.InitSession()     // Sessionコネクションをとっておく
	common.InitValidater() // Valideterコネクションをとっておく

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/notfound", authenticate(controller.NotFoundViewHandler))

	auto.AutoControllerLoad()

	// login
	http.HandleFunc("/login", authenticate(controller.GetLoginViewHandler))
  http.HandleFunc("/loginpost", authenticate(controller.PostLoginViewHandler))
  http.HandleFunc("/loginasynchronous", authenticate(controller.LoginAsynchronousViewHandler))
  http.HandleFunc("/logout", authenticate(controller.LogoutViewHandler))

	// fileupload
	http.HandleFunc("/fileupload", authenticate(controller.FileUploadIndexViewHandler))
	http.HandleFunc("/fileupload/register", authenticate(controller.FileUploadRegisterViewHandler))

	// postest
	http.HandleFunc("/posttest", authenticate(controller.PostTestIndexViewHandler))
	http.HandleFunc("/posttest/register", authenticate(controller.PostTestRegisterViewHandler))
	http.HandleFunc("/posttest/show", authenticate(controller.PostTestShowViewHandler))

	// listpostest
	http.HandleFunc("/listposttest", authenticate(controller.ListPostTestIndexViewHandler))
	http.HandleFunc("/listposttest/register", authenticate(controller.ListPostTestRegisterViewHandler))
	http.HandleFunc("/listposttest/show", authenticate(controller.ListPostTestShowViewHandler))

	// 最後に/をつけるとワイルドカードになる

	http.HandleFunc("/", authenticate(controller.HelloViewHandler))
	http.HandleFunc("/hello", authenticate(controller.HelloViewHandler))

	http.ListenAndServe(":3000", nil)
}
