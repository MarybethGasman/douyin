package main

// 抖音项目
// 三带一队:
// @陈鹤中 @张博思 @简懿豪 @谭盟 @徐政 @杨彬烜 @解城文 @梁明栩
// 使用iris框架作为web框架，部分接口使用gorm操作数据库，部分接口使用原生接口操作数据库
// 使用viper作为项目application.yaml配置管理，数据存储使用mysql和redis
// 登录，注册接口：谭盟
// 视频流接口和投稿接口 简懿豪
// 用户信息和粉丝列表 张博思
// 发布列表 梁明栩
// 赞操作和点赞列表 杨彬烜
// 评论操作和评论列表 徐政
// 关注操作和关注列表 陈鹤中
import (
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
