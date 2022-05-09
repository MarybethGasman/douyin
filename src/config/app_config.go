package config

import "github.com/spf13/viper"

var AppConfig = viper.New()

func init() {
	AppConfig.AddConfigPath(".")           //设置读取的文件路径
	AppConfig.SetConfigName("application") //设置读取的文件名
	AppConfig.SetConfigType("yaml")        //设置文件的类型
	//尝试进行配置读取
	if err := AppConfig.ReadInConfig(); err != nil {
		panic(err)
	}
}
