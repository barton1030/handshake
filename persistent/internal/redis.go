package internal

import (
	"github.com/gomodule/redigo/redis"
	"handshake/conf"
	"strconv"
	"time"
)

var redisConn *redis.Pool

func RedisInit() {
	redisConf := conf.RedisConf()
	host := redisConf.RedisHost()
	port := redisConf.RedisPort()
	stringPort := strconv.Itoa(port)
	addr := host + ":" + stringPort
	redisConn = &redis.Pool{
		MaxIdle:     redisConf.RedisInitConn(),
		MaxActive:   redisConf.RedisMaxConn(),
		IdleTimeout: 1 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}

func CloseRedis() {
	redisConn.Close()
}

func RedisConn() *redis.Pool {
	return redisConn
}
