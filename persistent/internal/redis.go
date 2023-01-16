package internal

import (
	"github.com/gomodule/redigo/redis"
	"handshake/conf"
)

var redisConn *redis.Pool

func RedisInit() {
	redisConf := conf.RedisConf()
	redisConn = &redis.Pool{
		MaxIdle:     redisConf.InitConn,
		MaxActive:   redisConf.MaxConn,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func CloseRedis() {
	redisConn.Close()
}

func RedisConn() *redis.Conn {
	conn := redisConn.Get()
	return &conn
}
