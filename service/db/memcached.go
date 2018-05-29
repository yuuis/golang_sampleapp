package db

import (
	"../../config"
	"../common"
	"github.com/bradfitz/gomemcache/memcache"
	//	"time"
)

var mc *memcache.Client

/**
 * Memcacheの初期化を行う
 */
func InitMemcache() {
	mc = memcache.New(config.MEMCACHEDHOSTNAME)
	mc.MaxIdleConns = config.MEMCACHEMAXCONNECTION
}

/*
 * Memcacheのデータを登録する
 */
func SetMemcachedData(key, data string, secconds int) {
	mc.Set(&memcache.Item{Key: key, Value: []byte(data), Expiration: int32(secconds)})
}

/*
 * Memcacheのデータを取得する
 */
func GetMemcachedData(key string) string {
	it, err := mc.Get(key)

	if err != nil {
		common.WriteErrorLog(config.DEBUG, err, nil)
		return ""
	} else {
		return string(it.Value)
	}
}
