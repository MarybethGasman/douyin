package main

import (
	. "douyin/src/config"
	. "douyin/src/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
)

func newApp() *iris.Application {
	app := iris.New()
	//app.OnErrorCode(iris.StatusNotFound, notFound)
	mvc.Configure(app.Party("/douyin/user"), func(app *mvc.Application) {
		app.Handle(new(UserController))
	})
	// 获取视频流
	mvc.Configure(app.Party("/douyin/feed"), func(app *mvc.Application) {
		app.Handle(new(FeedController))
	})

	return app
}
func main() {
	addr := strconv.Itoa(AppConfig.Get("server.port").(int))
	app := newApp()
	err := app.Run(iris.Addr(":"+addr), iris.WithCharset("UTF-8"))
	if err != nil {
		return
	}
}

//func notFound(ctx iris.Context) {
//	code := ctx.GetStatusCode()
//	msg := "404 Not Found"
//	ctx.JSON(iris.Map{
//		"Message": msg,
//		"Code":    code,
//	})
//}
