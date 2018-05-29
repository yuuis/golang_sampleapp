package db

import (
	"../../config"
	"../common"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

/*
 * redisの初期処理を行う
 */
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.REDISHOSTNAME,
		PoolSize: config.REDISMAXCONNECTION,
		DB:       0, // use default DB
	})
}

/*
 * データセット
 */
func SetRedisData(key string, value string, secconds int) {
	err := RedisClient.Set(key, value, time.Duration(secconds)*time.Second).Err()
	if err != nil {
		common.WriteErrorLog(config.ERROR, err, nil)
	}
}

/*
 * データ取得
 */
func GetRedisData(key string) string {
	val, err := RedisClient.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		common.WriteErrorLog(config.ERROR, err, nil)
		return ""
	}
	return val
}
