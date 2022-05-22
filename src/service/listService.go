package service

import (
	. "douyin/src/common"
	db2 "douyin/src/db"
)

const TableNameFavorite = "tb_favorite"

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
func GetVideoListsById(id string) []VideoList2 {
	var videos1 []VideoList1
	var user User
	db := db2.GetDBConnect()
	result := db.Select("video_id", "author_name", "play_url", "cover_url", "favorite_count", "comment_count").Where("author_name = ?", id).Find(&videos1)
	db.Where("user_id = ?", id).First(&user)
	n := result.RowsAffected
	videoS := make([]VideoList2, n)

	//给一条数据给他，要不然会崩溃
	if n == 0 {
		var vil VideoList2
		vil.Id = 4
		vil.Title = "?>"
		vil.PlayURL = "https://www.w3schools.com/html/movie.mp4"
		vil.CoverURL = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
		vil.IsFavorite = true
		vil.Author.Id = 4
		vil.FavoriteCount = 0
		vil.CommentCount = 0
		vil.Author.Name = "da"
		vil.Author.FollowerCount = 0
		vil.Author.IsFollow = true
		vil.Author.FollowCount = 0
		videoS = append(videoS, vil)
	}

	var i int64
	for i = 0; i < n; i++ {
		videoS[i].Id = videos1[i].Id
		videoS[i].PlayURL = videos1[i].PlayURL
		videoS[i].CommentCount = videos1[i].CommentCount
		videoS[i].FavoriteCount = videos1[i].FavoriteCount
		videoS[i].CoverURL = videos1[i].CoverURL
		videoS[i].Author.Id = user.Id
		videoS[i].Author.Name = user.Name
		videoS[i].Author.IsFollow = user.IsFollow
		videoS[i].Author.FollowerCount = user.FollowCount
		videoS[i].Author.FollowCount = user.FollowerCount

		var favorite Favorite
		cnt := db.Where("username = ?", user.Id).Where("video_id = ?", videoS[i].Id).First(&favorite)
		if cnt.RowsAffected > 0 && favorite.Isdeleted != 1 {
			videoS[i].IsFavorite = true
		} else {
			videoS[i].IsFavorite = false
		}
		videoS[i].Title = "数据库没有这个字段"
	}
	return videoS
}
