package controller

import (
	"database/sql"
	"douyin/src/cache"
	"douyin/src/db"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

type actionResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type listResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type Relation struct {
	RelationId  int64
	FollowerId  int64
	FollowingId int64
	IsDeleted   bool
}
type RelationController struct {
}

func (rc *RelationController) PostAction(context iris.Context) mvc.Result {
	token := context.Params().Get("token")
	if !cache.RCExists(token) {
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 400,
				StatusMsg:  "登录超时",
			},
		}
	}
	userid, _ := context.Params().GetInt("user_id")
	touserid, _ := context.Params().GetInt("to_user_id")
	actiontype, _ := context.Params().GetInt("action_type")
	cache.RCSet(token, userid, time.Minute*30) //更新用户token
	rows, err := db.DB.Query("select `follower_id`,`following_id` from `tb_relation`")
	if err != nil {
		fmt.Errorf("查询关注列表错误")
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 500,
				StatusMsg:  "查询列表错误",
			},
		}
	}
	defer rows.Close() //在执行完后关闭游标
	var result sql.Result
	if rows == nil {
		result, err = db.DB.Exec("insert into `tb_relation`(`follower_id`,`following_id`,`isdeleted`) "+
			"values(?,?,?)", userid, touserid, actiontype-1)
	} else {
		result, err = db.DB.Exec("update `tb_relation` set isdeleted=? where `follower_id`=? "+
			"and `following_id`=?", actiontype-1, userid, touserid)
	}
	ar, _ := result.RowsAffected()
	var stm string
	switch actiontype {
	case 1:
		{
			stm = "关注"
		}
	case 2:
		{
			stm = "取消关注"
		}
	}
	if ar != 0 {
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 0,
				StatusMsg:  stm + "成功",
			},
		}
	}
	return mvc.Response{
		Object: actionResponse{
			StatusCode: 300,
			StatusMsg:  "发生错误，请稍后重试",
		},
	}
}

func (rc *RelationController) GetList() mvc.Result {

}
