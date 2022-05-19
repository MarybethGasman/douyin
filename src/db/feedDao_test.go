package db

import (
	"fmt"
	"testing"
)

func TestFeedDao_SelectUserById(t *testing.T) {
	dao := &FeedDao{}
	id := dao.SelectUserById(3)
	fmt.Println(id)
}

func TestFeedDao_InsertVideo(t *testing.T) {
	dao := &FeedDao{}
	dao.InsertVideo(&TbVideo{
		AuthorName: "123456",
		PlayUrl:    "head.Filename",
	})
}
