package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var DB *sql.DB

func init() {
	config := viper.New()
	config.AddConfigPath("../../")      //设置读取的文件路径
	config.SetConfigName("application") //设置读取的文件名
	config.SetConfigType("yaml")        //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	driverName := config.Get("datasource.driverName").(string)
	dataSourceName := config.Get("datasource.dataSourceName").(string)
	//打印文件读取出来的内容:
	log.Printf("数据库为 %s, 数据库链接为%s", driverName, dataSourceName)
	_db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if _db.Ping() != nil {
		panic("数据库连接错误")
	}
	DB = _db
}
