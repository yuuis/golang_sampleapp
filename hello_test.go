package main

import (
	"./service/common"
	"./service/controller"
	"./service/data"
	"./service/db"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
 * go test ./
 * http://blog.kaneshin.co/entry/2016/12/02/200108
 */
func TestHandler(t *testing.T) {
	_, err := db.DbInit() // DBコネクションをとっておく
	if err != nil {
		panic(err)
	}
	db.InitRedis()         // Redisコネクションをとっておく
	db.InitMemcache()      // Memcachedコネクションをとっておく
	data.InitSession()     // Sessionコネクションをとっておく
	common.InitValidater() // Valideterコネクションをとっておく
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))

	t.Run("controller.HelloViewHandler", func(t *testing.T) {
		t.Parallel()

		s := httptest.NewServer(http.HandlerFunc(controller.HelloViewHandler))
		defer s.Close()

		res, err := http.Get(s.URL)
		assert.NoError(t, err)
		assert.Equal(t, "text/plain", res.Header.Get("Content-Type"))
		assert.Equal(t, 200, res.StatusCode)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, "pong", string(body))
	})

}
