package config

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
)

import (
	"net/http"
	"net/url"
)

var Cos *cos.Client

func init() {
	log.Println("初始化cos配置...")
	u, _ := url.Parse(AppConfig.GetString("cos.url"))
	bucket := &cos.BaseURL{BucketURL: u}
	Cos = cos.NewClient(bucket, &http.Client{Transport: &cos.AuthorizationTransport{
		SecretID:  AppConfig.GetString("cos.SecretId"),
		SecretKey: AppConfig.GetString("cos.SecretKey"),
	}})
	log.Println("初始化cos配置完成")
}
