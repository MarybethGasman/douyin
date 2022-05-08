package controller

import (
	. "douyin/src/db"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
}

func (uc *UserController) PostRegister(ctx iris.Context) mvc.Result {
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")

	db := DB

	result, err := db.Exec(
		"insert into tb_user(username,password) values(?,?)",
		username, password)
	if err != nil {
		panic("新增数据错误")
	}
	newID, _ := result.LastInsertId() //新增数据的ID
	i, _ := result.RowsAffected()     //受影响行数

	iris.New().Logger().Infof("新增的数据ID：%d , 受影响行数：%d \n", newID, i)

	return mvc.Response{
		Object: map[string]interface{}{
			"token":       "",
			"user_id":     1,
			"status_msg":  "OK",
			"status_code": 0,
		},
	}
}

func (uc *UserController) PostLogin(ctx iris.Context) mvc.Result {
	var username = ctx.URLParam("username")
	var password = ctx.URLParam("password")

	db := DB

	rows := db.QueryRow(
		"select password from tb_user where username = ?",
		username)
	var passwordFormDB string
	rows.Scan(&passwordFormDB)

	if password == passwordFormDB {
		return mvc.Response{
			Object: map[string]interface{}{
				"token":       "",
				"user_id":     1,
				"status_msg":  "OK",
				"status_code": 0,
			},
		}
	} else {
		return mvc.Response{
			Object: map[string]interface{}{
				"token":       "",
				"user_id":     1,
				"status_msg":  "用户名或密码错误",
				"status_code": 1,
			},
		}
	}
}
