package service

import (
	"douyin/src/config"
	"douyin/src/db"
	"time"
)

var (
	videoUrl = config.AppConfig.GetString("video.videoUrl")
	imageUrl = config.AppConfig.GetString("video.imageUrl")
)

type FeedData struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	NextTime   int64       `json:"next_time"`
	VideoList  []VideoList `json:"video_list"`
}

type Author struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type VideoList struct {
	ID            int64  `json:"id"`
	Author        Author `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

func GetFeed(latestTime string, token string) *FeedData {
	data := &FeedData{}

	dao := &db.FeedDao{}
	videos, err := dao.Select30VideoByUpdate(latestTime)
	if err != nil {
		return &FeedData{
			StatusCode: 1,
			StatusMsg:  "查找视频失败,err: " + err.Error(),
		}
	}
	// 如果数据库视频没有数据 那么就加入下面这一条测试数据 防止客户端崩溃
	if len(videos) == 0 {
		data.VideoList = append(data.VideoList, VideoList{
			ID: 0, Author: Author{ID: 0, Name: "test"},
			PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		})
	}

	for _, v := range videos {
		data.VideoList = append(data.VideoList, VideoList{
			ID: v.VideoId, Author: Author{ID: 1, Name: v.AuthorName},
			PlayUrl:       getVideoUrl(v.PlayUrl),
			CoverUrl:      getImageUrl(v.CoverUrl),
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			//TODO 这块需要一个判断是否点赞
			IsFavorite: false,
		})
	}
	// 已经没有视频了 从头继续播放
	if len(videos) == 0 {
		data.NextTime = time.Now().UnixMilli()
	} else {
		data.NextTime = videos[len(videos)-1].UpdateDate.UnixMilli()
	}

	data.StatusCode = 0
	return data
}

func getVideoUrl(videoName string) string {
	return videoUrl + videoName
}

func getImageUrl(imageName string) string {
	return imageUrl + imageName
}
