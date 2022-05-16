package controller

import "github.com/kataras/iris/v12/mvc"

type FavoriteController struct {
}
type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

func (fc *FavoriteController) GetList() mvc.Response {
	return mvc.Response{
		Object: VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoVideos,
		},
	}
}
