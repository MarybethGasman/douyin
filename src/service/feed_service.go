package service

import (
	"douyin/src/cache"
	"douyin/src/config"
	"douyin/src/db"
	"time"
)

var (
	videoUrl = config.AppConfig.GetString("video.videoUrl")
	imageUrl = config.AppConfig.GetString("video.imageUrl")
	dao      = &db.FeedDao{}
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
	Title         string `json:"title"`
}

func GetFeed(latestTime string, token string) *FeedData {
	data := &FeedData{}
	var userId int64
	userId = -1

	videos, err := dao.SelectVideoByUpdate(latestTime, 30)
	if err != nil {
		return &FeedData{
			StatusCode: 1,
			StatusMsg:  "查找视频失败,err: " + err.Error(),
		}
	}

	if len(token) > 0 {
		get := cache.RCGet(token)
		userId, err = get.Int64()
	}

	for _, v := range videos {
		data.VideoList = append(data.VideoList, VideoList{
			ID: v.VideoId,
			Author: Author{
				ID:            v.TbUser.UserId,
				Name:          v.TbUser.Name,
				FollowCount:   v.TbUser.FollowCount,
				FollowerCount: v.TbUser.FollowerCount,
				IsFollow:      isFollow(userId, v.TbUser.UserId),
			},
			PlayUrl:       getVideoUrl(v.PlayUrl),
			CoverUrl:      getImageUrl(v.CoverUrl),
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			Title:         v.Title,
			IsFavorite:    isFavorite(userId, v.VideoId),
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

func isFollow(uId int64, vId int64) bool {
	if uId == -1 {
		return false
	}
	// 查看该uId用户是否关注了该视频作者
	return dao.IsFollwer(uId, vId)
}

func isFavorite(uId int64, vId int64) bool {
	if uId == -1 {
		return false
	}
	// TODO 查看该uId用户是否点赞了vId该视频
	return dao.IsFavorite(uId, vId)
}
