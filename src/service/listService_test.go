package service

import (
	"douyin/src/common"
	db2 "douyin/src/db"
	"log"
	"testing"
)

func TestListService(t *testing.T) {
	db := db2.GetDBConnect()
	video := common.VideoList1{
		UserId:        2,
		PlayURL:       "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/a.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "this test",
	}
	re := db.Create(&video)
	log.Println("!!!", video.Id)
	log.Println("!!!", re.RowsAffected)

	var users []common.User

	re = db.Find(&users)
	n := re.RowsAffected
	var i int64
	for i = 0; i < n; i++ {
		log.Println("ID:", users[i].Id, "name:", users[i].Name)
	}
	log.Println("!!!")

	var videos []common.VideoList1
	re = db.Find(&videos)
	n = re.RowsAffected
	for i = 0; i < n; i++ {
		log.Println(videos[i])
	}
	log.Println("!!!!")
}
