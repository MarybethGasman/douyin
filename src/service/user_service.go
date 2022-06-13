package service

// 用户注册、登录、获取用户信息逻辑
// 登录状态由redis保持，当用户注册或登录后，将用户名和盐值MD5后作为key存入redis，value为用户id
// 该key的有效期为30分钟，用户每进行一次请求都会续期
// 常用的用户鉴权方法还有jwt，但简单的redis实现也可满足需求
// 登录注册传入的password都是经过加密的，以保证安全性
// author: 谭盟，张博思
import (
	. "douyin/cache"
	. "douyin/common"
	. "douyin/db"
	"douyin/utils"
	"time"
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
	//step1:判断用户名是否存在
	row := DB.QueryRow("select user_id from tb_user where name = ?", username)
	row.Scan(&userInDB.Id)
	userInDB.Password = password
	if userInDB.Exists() {
		return UserLoginAndRegisterResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		}
	} else {
		//step2:插入新记录
		result, err := DB.Exec(
			"insert into tb_user(name,password) values(?,?)",
			userInDB.Name, userInDB.Password)
		if err != nil {
			panic("新增数据错误")
		}
		userInDB.Id, err = result.LastInsertId() //新增数据的ID
		if err != nil {
			panic("获取新增数据ID错误")
		}
		//step3:保存登录状态信息
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
	rows.Scan(&userFormDB.Id, &userFormDB.Name, &userFormDB.Password)

	if userFormDB.IsCorrect(password) {
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

func Info(userId int64) User2 {
	var user User2
	row := DB.QueryRow(
		"select user_id,name,follow_count,follower_count,is_follow from tb_user where user_id = ?", userId)
	row.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow)
	return user
}
