package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
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
