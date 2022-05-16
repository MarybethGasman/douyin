package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

type FeedController struct {
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}
type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    true,
	},
}
var DemoUser = User{
	Id:            1,
	Name:          "xuzheng",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      true,
}

// Feed same demo video list for every request
func (cc *FeedController) Get(ctx iris.Context) mvc.Result {

	return mvc.Response{
		Object: FeedResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "请求成功"},
			VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		},
	}
}
