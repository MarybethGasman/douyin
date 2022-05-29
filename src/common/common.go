package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoList1 struct {
	Id            int64  `gorm:"column:video_id;primaryKey;autoIncrement:true" json:"id"`
	AuthorName    string `gorm:"column:author_name"`
	PlayURL       string `gorm:"column:play_url" json:"play_url"`
	CoverURL      string `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"  json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count"`
}

func (*VideoList1) TableName() string {
	return "tb_video"
}

type User2 struct {
	Id            int64  `gorm:"column:user_id;primaryKey;autoIncrement:true" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"`
	IsFollow      bool   `gorm:"column:is_follow" json:"is_follow"`
}

type VideoList2 struct {
	Id            int64  `gorm:"column:video_id;primaryKey;autoIncrement:true" json:"id"`
	Author        User2  `json:"author"`
	PlayURL       string `gorm:"column:play_url" json:"play_url"`
	CoverURL      string `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"  json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
type VideoListResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg,omitempty"`
	VideoLists []VideoList2 `json:"video_list"`
}

func (*VideoList2) TableName() string {
	return "tb_video"
}

type User struct {
	Id            int64  `gorm:"column:user_id;primaryKey;autoIncrement:true" json:"id,omitempty"`
	Name          string `gorm:"column:name" json:"name,omitempty"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"column:is_follow" json:"is_follow,omitempty"`
	Password      string `gorm:"column:password"`
}

func (*User) TableName() string {
	return "tb_user"
}

func (user *User) IsCorrect(password string) bool {
	return user.Password == password
}
func (user *User) Exists() bool {
	return user.Id > 0
}
