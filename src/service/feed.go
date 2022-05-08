package service

import "time"

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

func GetFeed(latestTime string) *FeedData {
	data := &FeedData{}

	data.VideoList = append(data.VideoList, VideoList{
		ID: 1, Author: Author{ID: 1, Name: "zhangshan"},
		PlayUrl:    "http://10.196.62.4:8080/douyin/feed/video/movie.mp4",
		CoverUrl:   "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		IsFavorite: false,
	})

	data.VideoList = append(data.VideoList, VideoList{
		ID: 2, Author: Author{ID: 2, Name: "lisi"},
		PlayUrl:    "http://10.196.62.4:8080/douyin/feed/video/big_buck_bunny.mp4",
		CoverUrl:   "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		IsFavorite: true,
	})

	data.VideoList = append(data.VideoList, VideoList{
		ID: 3, Author: Author{ID: 2, Name: "lisi"},
		PlayUrl:    "http://10.196.62.4:8080/douyin/feed/video/oceans.mp4",
		CoverUrl:   "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		IsFavorite: true,
	})
	data.NextTime = time.Now().Unix()

	return data
}
