package internal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"handshake/conf"
)

var dbConnHandler *gorm.DB

// DbInit -- 初始化数据库连接
func DbInit() {
	dbConf := conf.DbConf()
	dbAddr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=true&loc=Local", dbConf.User, dbConf.Pwd, dbConf.Host, dbConf.Port, dbConf.DbName)
	fmt.Println(dbAddr)
	dbConn, err := gorm.Open(dbConf.Type, dbAddr)
	if err != nil {
		panic(err)
	}
	dbConn.DB().SetMaxIdleConns(dbConf.InitConn)
	dbConn.DB().SetMaxOpenConns(dbConf.MaxConn)
	dbConnHandler = dbConn
}

func CloseDb() {
	dbConnHandler.Close()
}

// DbConn -- 通过对外function提供服务
func DbConn() *gorm.DB {
	return dbConnHandler
}
