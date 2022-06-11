package db

import (
	. "douyin/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var gormDb *gorm.DB
var cnt = 0

func GetDBConnect() (conn *gorm.DB) {
	if gormDb != nil {
		return gormDb
	}

	cnt++
	dataSourceName := AppConfig.Get("datasource.dataSourceName").(string)
	Db, err := gorm.Open(mysql.Open(dataSourceName))
	log.Printf("数据库为mysql , 数据库链接为%s", dataSourceName)
	if err != nil {
		panic(err)
	}
	gormDb = Db
	return gormDb
}
