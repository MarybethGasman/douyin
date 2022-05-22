package service

import (
	. "douyin/src/common"
	"douyin/src/controller"
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
func GetVideoListsById(id string) []controller.VideoList {
	var videoS []controller.VideoList
	var user User
	db := db2.GetDBConnect()
	result := db.Where("author_name = ?", id).Find(&videoS)
	db.Where("user_id = ?", id).First(&user)
	var i int64
	for i = 0; i < result.RowsAffected; i++ {
		videoS[i].Author = user
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
