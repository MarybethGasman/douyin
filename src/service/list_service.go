package service

import (
	. "douyin/common"
	db2 "douyin/db"
)

const (
	TableNameFavorite = "tb_favorite"
)

// Favorite mapped from table <tb_favorite>
type Favorite struct {
	FavoriteID int64  `gorm:"column:favorite_id;primaryKey;autoIncrement:true"`
	Username   string `gorm:"column:username"`
	VideoID    int64  `gorm:"column:video_id"`
	Isdeleted  int32  `gorm:"column:isdeleted"`
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}

/*
获取视频列表
*/
func GetVideoListsById(id int) []VideoList2 {

	var videos1 []VideoList1
	var user User2

	db := db2.GetDBConnect()
	result := db.Select("video_id", "user_id", "play_url", "cover_url", "favorite_count", "comment_count").Where("user_id = ?", id).Find(&videos1)
	db.Model(User{}).Where("user_id = ?", id).First(&user)
	n := result.RowsAffected
	videoS := make([]VideoList2, n)

	var i int64
	for i = 0; i < n; i++ {
		videoS[i].Id = videos1[i].Id
		//拼接得到视频地址
		//使用了feedService.go中的接口
		videoS[i].PlayURL = getVideoUrl(videos1[i].PlayURL)
		videoS[i].CommentCount = videos1[i].CommentCount
		videoS[i].FavoriteCount = videos1[i].FavoriteCount
		//我也不知道封面放在哪？
		videoS[i].CoverURL = getImageUrl(videos1[i].CoverURL)
		videoS[i].Author = user

		var favorite Favorite
		cnt := db.Where("user_id = ?", user.Id).Where("video_id = ?", videoS[i].Id).First(&favorite)
		if cnt.RowsAffected > 0 && favorite.Isdeleted != 1 {
			videoS[i].IsFavorite = true
		} else {
			videoS[i].IsFavorite = false
		}
		videoS[i].Title = videos1[i].Title
	}
	return videoS
}
