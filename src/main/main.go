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
	//配置404返回内容
	app.OnErrorCode(iris.StatusNotFound, notFound)
	//配置路由
	mvc.Configure(app.Party("/douyin/user"), func(app *mvc.Application) {
		app.Handle(new(UserController))
	})
	mvc.Configure(app.Party("/douyin/favorite"), func(app *mvc.Application) {
		app.Handle(new(FavoriteController))
	})
	mvc.Configure(app.Party("/douyin/comment"), func(app *mvc.Application) {
		app.Handle(new(CommentController))
	})
	mvc.Configure(app.Party("/douyin/publish"), func(app *mvc.Application) {
		app.Handle(new(PublishController))
	})
	mvc.Configure(app.Party("/douyin/relation"), func(app *mvc.Application) {
		app.Handle(new(RelationController))
	})
	mvc.Configure(app.Party("/douyin/feed"), func(app *mvc.Application) {
		app.Handle(new(FeedController))
	})
	return app
}

func main() {
	addr := strconv.Itoa(AppConfig.GetInt("server.port"))
	app := newApp()
	//每一次请求都打印请求路径，iris使用了责任链模式，很类似servlet里面的filter
	app.UseGlobal(before)
	//监听我们的服务端口，iris.WithoutPathCorrectionRedirection这个选项用于路径匹配，例如
	// /douyin/user/register 和 /douyin/user/register/ 这两个路径完全不同但却需要同一个请求处理器
	// 咱们的客户端发送的请求末尾都会带上一个分号，若未开启这个选项，请求会fail，都是教训啊(￣、￣)
	app.Run(iris.Addr(":"+addr), iris.WithCharset("UTF-8"), iris.WithoutPathCorrectionRedirection)
}

func before(ctx iris.Context) {
	iris.New().Logger().Info(ctx.Path())
	ctx.Next()
}

func notFound(ctx iris.Context) {
	code := ctx.GetStatusCode()
	msg := "404 Not Found"
	ctx.JSON(iris.Map{
		"Message": msg,
		"Code":    code,
	})
}
