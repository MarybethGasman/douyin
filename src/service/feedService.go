package service

import (
	"douyin/src/db"
	"time"
)

const (
	videoUrl = "http://10.196.62.4:8080/douyin/feed/video/"      // 哪个服务器存放着视频的目录
	imageUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/" // 哪个服务器存放着视频封面的目录
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
