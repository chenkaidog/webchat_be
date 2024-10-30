package db

import (
	"webchat_be/biz/db/mysql"
	"webchat_be/biz/db/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
