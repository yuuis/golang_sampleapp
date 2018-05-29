package data

import (
	"../../config"
	"github.com/bradfitz/gomemcache/memcache"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/gorilla/sessions"
	"net/http"
)

var store *gsm.MemcacheStore

func InitSession() {
	// Memcachedを使用する場合
	mc := memcache.New(config.MEMCACHEDHOSTNAME)
	mc.MaxIdleConns = config.SESSION_MAX_CONNECTION
	store = gsm.NewMemcacheStore(mc, "session_prefix_", []byte("secret-key-goes-here"))
	//store = sessions.NewCookieStore([]byte("something-very-secret"))
}

/*
 * Session情報登録
 *
 */
func SetStringSession(w http.ResponseWriter, r *http.Request, key string, value string) {
	//Session情報へ格納
	session, _ := store.Get(r, config.SESSION_NAME)
	session.Options = &sessions.Options{MaxAge: config.SESSION_MAX_AGE, Path: config.SESSION_PATH}

	session.Values[key] = value
	session.Save(r, w)
}

func SetIntSession(w http.ResponseWriter, r *http.Request, key string, value int) {
	//Session情報へ格納
	session, _ := store.Get(r, config.SESSION_NAME)
	session.Options = &sessions.Options{MaxAge: config.SESSION_MAX_AGE, Path: config.SESSION_PATH}

	session.Values[key] = value
	session.Save(r, w)
}

/*
 * Session情報取得
 *
 */
func GetStringSession(r *http.Request, key string) string {
	session, error := store.Get(r, config.SESSION_NAME)

	if error != nil {
		return ""
	}
	if session.Values[key] == nil {
		return ""
	}

	return session.Values[key].(string)
}
func GetIntSession(r *http.Request, key string) int {
	session, error := store.Get(r, config.SESSION_NAME)

	if error != nil {
		return 0
	}
	if session.Values[key] == nil {
		return 0
	}
	return session.Values[key].(int)
}

/*
 * Session情報破棄
 *
 */
func DeleteLoginUserSession(w http.ResponseWriter, r *http.Request) {
	//Sessionクリア
	session, _ := store.Get(r, config.SESSION_NAME)
	session.Options = &sessions.Options{MaxAge: -1, Path: config.SESSION_PATH}
	session.Save(r, w)
}
