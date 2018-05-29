package dto

import (
	"github.com/mholt/binding"
	"net/http"
)

/*
 * Controllerに直書きした場合は、Logicなどで使いまわせないので、機能単位でFORMを外部に出す
 */
type PostTestForm struct {
	MemcacheKey   string `validate:"required,max=10,min=1"`
	RedisKey      string `validate:"required,max=10,min=1"`
	MemcacheValue string `validate:"required,max=10,min=1,hiragana"`
	RedisValue    string `validate:"required,max=10,min=1,katakana"`
	Errors        []ErrorForm
}

func (cf *PostTestForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.MemcacheKey:   "memcache_key",
		&cf.RedisKey:      "redis_key",
		&cf.MemcacheValue: "memcache_value",
		&cf.RedisValue:    "redis_value",
	}
}
