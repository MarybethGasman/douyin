package db

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestFeedDao_SelectUserById(t *testing.T) {
	dao := &FeedDao{}
	id := dao.SelectUserById(3)
	fmt.Println(id)
}

func TestFeedDao_SelectVideoByUpdate(t *testing.T) {
	dao := &FeedDao{}
	location, err := time.LoadLocation("Local")
	inLocation, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-05-22 18:53:42", location)
	formatInt := strconv.FormatInt(inLocation.UnixMilli(), 10)
	fmt.Println(formatInt)
	update, err := dao.SelectVideoByUpdate("", 2)
	if err != nil {
		return
	}
	for i, d := range update {
		fmt.Println(i, "--->", d)
	}
}

func TestFeedDao_InsertVideo(t *testing.T) {
	//db := &FeedDao{}
	//db.InsertVideo(&TbVideo{
	//	UserId:  30,
	//	PlayUrl: "head.Filename",
	//})
}

func TestFeedDao_PreLoad(t *testing.T) {
	var testTable []TbVideo
	db.Debug().Preload("TbUser").Find(&testTable)
	fmt.Println(testTable)
}
