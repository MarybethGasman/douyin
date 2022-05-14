package controller

import (
	"douyin/src/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type FeedController struct {
}

func (fc *FeedController) Get(ctx iris.Context) mvc.Result {
	ctx.JSON(service.GetFeed(ctx.URLParam("latest_time")))
	return nil
}

func (fc *FeedController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/feed/video/{videoName}", "GetFeedVideo")
}

func (fc *FeedController) GetFeedVideo(b iris.Context) {
	videoName := b.URLParam("videoName")
	b.SendFile("E:\\GolandProjects\\simple-demo\\video\\"+videoName, videoName)
}
