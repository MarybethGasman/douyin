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
	app.OnErrorCode(iris.StatusNotFound, notFound)
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
	addr := strconv.Itoa(AppConfig.Get("server.port").(int))
	app := newApp()
	app.UseGlobal(before)
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
