package controller

import "github.com/kataras/iris/v12/mvc"

type PublishController struct {
}

func (pc *PublishController) GetList() mvc.Response {
	return mvc.Response{
		Object: VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoVideos,
		},
	}
}
