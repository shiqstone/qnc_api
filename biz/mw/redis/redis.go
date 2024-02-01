package redis

import (
	"time"

	"github.com/go-redis/redis/v7"

	"qnc/biz/mw/viper"
)

var (
	expireTime = time.Hour * 1
	rdbQueue   *redis.Client
	config     *viper.Redis
)

func InitRedis() {
	config = viper.Conf.Redis
	rdbQueue = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})
}
