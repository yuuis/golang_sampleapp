package dto

import (
	"github.com/mholt/binding"
	"net/http"
)

/*
 * https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
 */
type ListPostTestForm struct {
	MemcacheKey     string `validate:"required,max=10,min=1"`
	RedisKey        string `validate:"required,max=10,min=1"`
	MemcacheValue   string `validate:"required,max=10,min=1,hiragana"`
	RedisValue      string `validate:"required,max=10,min=1,katakana"`
	MemcachedValue  []string
	MemcachedRemark []string
	MemcachedDatas  []ListPostMemcachedData `validate:"required,dive,required"`
	RedisDatas      []ListPostRedisData
	Errors          []ErrorForm
}

type ListPostMemcachedData struct {
	MemcachedValue  string `validate:"required"`
	MemcachedRemark string `validate:"required,hiragana"`
}
type ListPostRedisData struct {
	RedisValue  string `validate:"required"`
	RedisRemark string `validate:"required"`
}

func (cf *ListPostTestForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.MemcacheKey:     "memcache_key",
		&cf.RedisKey:        "redis_key",
		&cf.MemcacheValue:   "memcache_value",
		&cf.RedisValue:      "redis_value",
		&cf.MemcachedValue:  "memcached_list_value[]",
		&cf.MemcachedRemark: "memcached_list_remark[]",
	}
}
