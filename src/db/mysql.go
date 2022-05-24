package db

import (
	"database/sql"
	. "douyin/src/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB



func init() {

	driverName := AppConfig.Get("datasource.driverName").(string)
	dataSourceName := AppConfig.Get("datasource.dataSourceName").(string)
	//打印文件读取出来的内容:
	log.Printf("数据库为 %s, 数据库链接为%s", driverName, dataSourceName)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if db.Ping() != nil {
		panic("数据库连接错误")
	}
	DB = db
}
