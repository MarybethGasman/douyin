package controller

import (
	"douyin/src/cache"
	. "douyin/src/common"
	"douyin/src/service"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"log"
)

type PublishController struct {
}

func (pc *PublishController) Get() {

}

func (pc *PublishController) GetList(ctx iris.Context) {

	request := ctx.Request()
	//获取参数
	token := request.FormValue("token")
	userID := request.FormValue("user_id")
	userId := cache.RCGet(token).Val()

	//查看输出
	log.Println("token:"+token, "userId:"+userID, "userid2:"+userId)

	if userId == "" {
		_, err := ctx.JSON(VideoListResponse{
			StatusCode: 301,
			StatusMsg:  "鉴权失败，请检测是否登录",
			VideoLists: []VideoList2{},
		})
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	//获取视频列表
	videoLists := service.GetVideoListsById(userId)

	//不知道为什么videoLists不能为空，等我解决吧
	//已解决：videoLists应为是数组，所以是空是[]，而不是nil
	_, err := ctx.JSON(VideoListResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoLists: videoLists,
	})
	log.Println(json.Marshal(videoLists))
	if err != nil {
		log.Println(err.Error())
	}
}

func (pc *PublishController) PostAction(ctx iris.Context) {
	service.Contribution(ctx)
}
