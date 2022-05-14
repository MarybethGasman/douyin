package config

import (
	"bytes"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

var AppConfig = viper.New()

func init() {
	stdout, _ := exec.Command("go", "env", "GOMOD").Output()
	path := string(bytes.TrimSpace(stdout))
	if path == "" {
		os.Exit(1)
	}
	ss := strings.Split(path, "\\")
	ss = ss[:len(ss)-1]
	path = strings.Join(ss, "\\") + "\\"

	AppConfig.AddConfigPath(path)          //设置读取的文件路径
	AppConfig.SetConfigName("application") //设置读取的文件名
	AppConfig.SetConfigType("yaml")        //设置文件的类型
	//尝试进行配置读取
	if err := AppConfig.ReadInConfig(); err != nil {
		panic(err)
	}
}
