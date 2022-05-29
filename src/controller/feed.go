package controller

import (
	"bytes"
	"douyin/src/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"os"
	"os/exec"
	"strings"
)

type FeedController struct {
}

func (fc *FeedController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/video/{videoName}", "GetFeedVideo")
	b.Handle("GET", "/", "Get")
}

func (fc *FeedController) Get(ctx iris.Context) {
	ctx.JSON(service.GetFeed(ctx.URLParam("latest_time"), ctx.URLParam("token")))
}

func (fc *FeedController) GetFeedVideo(b iris.Context) {
	videoName := b.Params().Get("videoName")

	stdout, _ := exec.Command("go", "env", "GOMOD").Output()
	path := string(bytes.TrimSpace(stdout))
	if path == "" {
		os.Exit(1)
	}
	ss := strings.Split(path, "\\")
	ss = ss[:len(ss)-1]
	path = strings.Join(ss, "\\") + "\\"

	b.SendFile(path+service.FilePath+videoName, videoName)
}
