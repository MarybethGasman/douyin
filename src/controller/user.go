package controller

import (
	"douyin/src/common"
	userService "douyin/src/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
}
type UserResponse struct {
	common.Response
	User common.User2 `json:"user"`
}

func (uc *UserController) PostRegister(ctx iris.Context) mvc.Result {
	//参数接受校验
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")
	if username == "" || password == "" {
		return mvc.Response{
			Object: userService.UserLoginAndRegisterResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "参数不能为空"},
			},
		}
	}
	//password = utils.MD5(password)
	var response = userService.Register(username, password)

	return mvc.Response{
		Object: response,
	}
}

func (uc *UserController) PostLogin(ctx iris.Context) mvc.Response {
	//参数获取与校验
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")
	if username == "" || password == "" {
		return mvc.Response{
			Object: userService.UserLoginAndRegisterResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "参数不能为空"},
			},
		}
	}
	response := userService.Login(username, password)

	return mvc.Response{
		Object: response,
	}
}

func (uc *UserController) Get(ctx iris.Context) mvc.Response {
	var userId int64 = ctx.URLParamInt64Default("user_id", -1)
	if userId == -1 {
		return mvc.Response{
			Object: userService.UserLoginAndRegisterResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "参数不能为空"},
			},
		}
	}
	user := userService.Info(userId)
	return mvc.Response{
		Object: UserResponse{
			Response: common.Response{StatusCode: 0, StatusMsg: ""},
			User:     user,
		},
	}
}

//func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
//	a.Handle("POST", "/login/", "PostLogin")
//	a.Handle("POST", "/register/", "PostRegister")
//}
