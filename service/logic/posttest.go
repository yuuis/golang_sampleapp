package logic

import (
	"../db"
	"../dto"
)

func PostTestRegister(form *dto.PostTestForm) {
	db.SetMemcachedData(form.MemcacheKey, form.MemcacheValue, 1*60) // 入力されたKeyとValueをTTL1分で登録
	db.SetRedisData(form.RedisKey, form.RedisValue, 1*60)           // 入力されたKeyとValueをTTL1分で登録
}

func PostTestShow(form *dto.PostTestForm) {
	form.MemcacheValue = db.GetMemcachedData(form.MemcacheKey) // 入力されたKeyを使用してValueを取得
	form.RedisValue = db.GetRedisData(form.RedisKey)           // 入力されたKeyを使用してValueを取得
}
