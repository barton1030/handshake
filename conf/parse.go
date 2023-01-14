package conf

import "fmt"
import "github.com/BurntSushi/toml"

func Init() {
	_, err := toml.DecodeFile("./conf/config.toml", &c)
	if err != nil {
		panic(fmt.Sprintf("config parse fail: %v", err))
	}
	fmt.Println("config:", c)
}

var c Config

type Config struct {
	ServerConfig Server `toml:"server"`
	DbConfig     Db     `toml:"database"`
	RedisConfig  Redis  `toml:"redis"`
}

type Server struct {
	ServerName string `toml:"server_name"`
	ServerPort int    `toml:"server_port"`
}

type Db struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Pwd      string `toml:"password"`
	Type     string `toml:"type"`
	DbName   string `toml:"db_name"`
	InitConn int    `toml:"init_conn"`
	MaxConn  int    `toml:"max_conn"`
}

type Redis struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Pwd      string `toml:"password"`
	InitConn int    `toml:"init_conn"`
	MaxConn  int    `toml:"max_conn"`
}

func ServerConf() (serverConf Server) {
	return c.ServerConfig
}

func DbConf() (dbConf Db) {
	return c.DbConfig
}

func RedisConf() (redisConf Redis) {
	return c.RedisConfig
}
