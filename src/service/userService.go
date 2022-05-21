package service

import (
	. "douyin/src/common"
	"douyin/src/utils"
	"time"
	."douyin/src/db"
	."douyin/src/cache"
)

type UserLoginAndRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(username string, password string) UserLoginAndRegisterResponse {
	var userInDB User
	userInDB.Id = 0
	userInDB.Name = username
	row := DB.QueryRow("select user_id from tb_user where name = ?", username)
	row.Scan(&userInDB.Id)
	userInDB.password = password
	if userInDB.exists() {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		}
	} else {
		//更新数据
		result, err := DB.Exec(
			"insert into tb_user(name,password) values(?,?)",
			userInDB.Name, userInDB.password)
		if err != nil {
			panic("新增数据错误")
		}
		userInDB.Id, err = result.LastInsertId() //新增数据的ID
		if err != nil {
			panic("获取新增数据ID错误")
		}

		token := utils.MD5WithSalt(userInDB.Name)
		RCSet(token, userInDB.Id, 30*time.Minute)

		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 0},
			UserId:   userInDB.Id,
			Token:    token,
		}
	}
}

func Login(username string, password string) UserLoginAndRegisterResponse {
	//查询数据库
	rows := DB.QueryRow(
		"select user_id,name,password from tb_user where name = ?",
		username)
	var userFormDB = User{}
	rows.Scan(&userFormDB.Id, &userFormDB.Name, &userFormDB.password)

	if userFormDB.isCorrect(password) {
		token := utils.MD5WithSalt(username)
		RCSet(token, userFormDB.Id, 30*time.Minute)
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 0},
			UserId:   userFormDB.Id,
			Token:    token,
		}
	} else {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
		}
	}
}

func Info(userId int64) User {
	var user User
	row := DB.QueryRow(
		"select user_id,naame,follow_count,follower_count,is_follow where user_id = ?", userId)
	row.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow)

	return user
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	password      string
}

func (user *User) isCorrect(password string) bool {
	return user.password == password
}
func (user *User) exists() bool {
	return user.Id > 0
}
func (user *User) equals(user2 User) bool {
	return user.Name == user2.Name &&
		user.password == user2.password
}
