package controller

import (
	"douyin/src/cache"
	. "douyin/src/common"
	"douyin/src/service"
	"github.com/kataras/iris/v12"
	"log"
)

type VideoList struct {
	Id            int64  `gorm:"column:video_id;primaryKey;autoIncrement:true" json:"id"`
	Author        User   `json:"author"`
	AuthorName    string `gorm:"column:author_name"`
	PlayURL       string `gorm:"column:play_url" json:"play_url"`
	CoverURL      string `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"  json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
type VideoListResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg,omitempty"`
	VideoLists []VideoList `json:"video_list,omitempty"`
}

func (*VideoList) TableName() string {
	return "tb_video"
}

type PublishController struct {
}

func (pc *PublishController) Get() {

}

func (pc *PublishController) GetList(ctx iris.Context) {
	request := ctx.Request()
	//获取参数
	token := request.FormValue("token")
	//userID := request.FormValue("user_id")
	userId := cache.RCGet(token)

	if userId == nil {
		_, err := ctx.JSON(VideoListResponse{
			StatusCode: 0,
			StatusMsg:  "鉴权失败，请检测是否登录",
			VideoLists: nil,
		})
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	//获取视频列表
	videoLists := service.GetVideoListsById(userId.Val())

	_, err := ctx.JSON(VideoListResponse{
		StatusCode: 1,
		StatusMsg:  "成功",
		VideoLists: videoLists,
	})
	if err != nil {
		log.Println(err.Error())
	}
}

func (pc *PublishController) PostAction(ctx iris.Context) {
	service.Contribution(ctx)
}
