package dal

import (
	"qnc/biz/dal/db"
	"qnc/biz/mw/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	redis.InitRedis()
}
