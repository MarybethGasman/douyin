package controller

import (
	"douyin/src/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type FeedController struct {
}

func (f *FeedController) Get(ctx iris.Context) mvc.Result {
	ctx.JSON(service.GetFeed(ctx.URLParam("latest_time")))
	return nil
}

func (f *FeedController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/video/{video}", "GetFeetVideo")
}

func (f *FeedController) GetFeetVideo(ctx iris.Context) mvc.Result {
	path := ctx.Params().Get("video")
	ctx.SendFile("./video/"+path, path)
	return nil
}
