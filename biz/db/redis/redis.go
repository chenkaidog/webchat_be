package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"webchat_be/biz/config"
)

var rdbClient *redis.Client

func Init() {
	rdbClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetRedisConf().IP, config.GetRedisConf().Port),
		Password: config.GetRedisConf().Password,
		DB:       config.GetRedisConf().DB,
	})

	rdbClient.AddHook(new(loggerHook))
}

func GetRedisClient() *redis.Client {
	return rdbClient
}
