package persistent

import "handshake/persistent/internal"

func Init() {
	internal.DbInit()
	internal.RedisInit()
}

func Close() {
	internal.CloseDb()
	internal.CloseRedis()
}
