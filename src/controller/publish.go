package controller

import (
	"douyin/src/service"
	"github.com/kataras/iris/v12"
)

type PublishController struct {
}

func (pc *PublishController) Get() {

}

func (pc *PublishController) PostAction(ctx iris.Context) {
	service.Contribution(ctx)
}
