package controller

import (
	"douyin/src/cache"
	. "douyin/src/db"
	"douyin/src/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserController struct {
}

func (uc *UserController) PostRegister(ctx iris.Context) mvc.Result {
	var username = ctx.URLParam("username")

	userId := 0
	row := DB.QueryRow("select user_id from tb_user where name = ?", username)
	row.Scan(&userId)
	if userId > 0 {
		return mvc.Response{
			Object: Response{1, "User already exist"},
		}
	}
	var password = ctx.URLParam("password")

	//password = utils.MD5(password)

	result, err := DB.Exec(
		"insert into tb_user(name,password) values(?,?)",
		username, password)
	if err != nil {
		panic("新增数据错误")
	}
	newID, _ := result.LastInsertId() //新增数据的ID

	token := utils.MD5WithSalt(username)
	cache.RCSet(token, newID, 30*time.Minute)

	return mvc.Response{
		Object: UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newID,
			Token:    token,
		},
	}
}

func (uc *UserController) PostLogin(ctx iris.Context) mvc.Result {
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")

	rows := DB.QueryRow(
		"select user_id,password from tb_user where name = ?",
		username)

	var passwordFormDB string
	var userId int64
	rows.Scan(&userId, &passwordFormDB)
	token := utils.MD5WithSalt(username)
	cache.RCSet(token, userId, 30*time.Minute)

	if passwordFormDB == password {
		return mvc.Response{
			Object: UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userId,
				Token:    token,
			},
		}
	} else {
		return mvc.Response{
			Object: UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
			},
		}
	}
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

var user User = User{
	Id:            1,
	Name:          "liry",
	FollowCount:   100,
	FollowerCount: 500,
	IsFollow:      true,
}

//用户信息接口
func (uc *UserController) Get(ctx iris.Context) mvc.Response {
	//fmt.Println(ctx.URLParams())
	//user_id := ctx.URLParam("user_id")
	return mvc.Response{
		Object: UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		},
	}
}

func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("POST", "/login/", "PostLogin")
	a.Handle("POST", "/register/", "PostRegister")
}
