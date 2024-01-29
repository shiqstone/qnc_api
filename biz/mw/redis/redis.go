package redis

import (
	"time"

	"github.com/go-redis/redis/v7"

	"qnc/pkg/constants"
)

var (
	expireTime = time.Hour * 1
	rdbQueue   *redis.Client
)

func InitRedis() {
	rdbQueue = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
}
