package dao

type Tb_relation struct {
	RelationId  int `gorm:"relation_id" json:"relation_id"`
	FollowerId  int `gorm:"follower_id" json:"follower_id"`
	FollowingId int `gorm:"following_id" json:"following_id"`
	Isdeleted   int `gorm:"isdeleted" json:"isdeleted"`
}

type Tb_user struct {
	UserId        int    `gorm:"user_id" json:"user_id"`
	Name          string `gorm:"name" json:"name"`
	FollowCount   int    `gorm:"follow_count" json:"follow_count"`
	FollowerCount int    `gorm:"follower_count" json:"follower_count"`
	IsFollow      int    `gorm:"is_follow" json:"is_follow"`
	Password      string `gorm:"password" json:"password"`
}

type Tb_video struct {
	VideoId       int    `gorm:"video_id" json:"video_id"`
	AuthorName    string `gorm:"author_name" json:"author_name"`
	PlayUrl       string `gorm:"play_url" json:"play_url"`
	CoverUrl      string `gorm:"cover_url" json:"cover_url"`
	FavoriteCount int    `gorm:"favorite_count" json:"favorite_count"`
	CommentCount  int    `gorm:"comment_count" json:"comment_count"`
}

type Tb_comment struct {
	CommentId int    `gorm:"comment_id" json:"comment_id"`
	Username  string `gorm:"username" json:"username"`
	Content   string `gorm:"content" json:"content"`
	Isdeleted int    `gorm:"isdeleted" json:"isdeleted"`
}

type Tb_favorite struct {
	FavoriteId int    `gorm:"favorite_id" json:"favorite_id"`
	Username   string `gorm:"username" json:"username"`
	VideoId    int    `gorm:"video_id" json:"video_id"`
	Isdeleted  int    `gorm:"isdeleted" json:"isdeleted"`
}
