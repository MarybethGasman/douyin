package controller

import (
	"database/sql"
	"douyin/src/cache"
	. "douyin/src/common"
	"douyin/src/db"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"strconv"
	"time"
)

type actionResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type listResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list,omitempty"`
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
	token := context.FormValue("token")
	if "" == token && !cache.RCExists(token) {
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 400,
				StatusMsg:  "登录超时",
			},
		}
	}
	userid := context.FormValue("user_id")
	touserid := context.FormValue("to_user_id")
	actiontype, _ := strconv.Atoi(context.FormValue("action_type"))
	cache.RCSet(token, userid, time.Minute*30) //更新用户token

	rows, err := db.DB.Query("select `follower_id`,`following_id` from `tb_relation` where  `follower_id`=? and `following_id`=?", userid, touserid)
	if err != nil {
		log.Println("查询关注列表错误")
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 500,
				StatusMsg:  "查询列表错误",
			},
		}
	}
	defer rows.Close() //在执行完后关闭游标

	//开启事务
	tx, _ := db.DB.Begin()
	var result sql.Result
	if !rows.Next() {
		result, _ = db.DB.Exec("insert into `tb_relation`(`follower_id`,`following_id`,`isdeleted`) "+
			"values(?,?,?)", userid, touserid, (actiontype+1)%2) // 2对应true执行关注操作(删除) 1对应true执行取消关注操作（不删除）
	} else {
		result, _ = db.DB.Exec("update `tb_relation` set isdeleted=? where `follower_id`=? "+
			"and `following_id`=?", (actiontype+1)%2, userid, touserid)
	}
	if actiontype == 1 {
		db.DB.Exec("update `tb_user` set `follow_count`=`follow_count`+1 where `user_id`=?", touserid)   //粉丝数+1
		db.DB.Exec("update `tb_user` set `follower_count`=`follower_count`+1 where `user_id`=?", userid) //关注数+1
	} else {
		db.DB.Exec("update `tb_user` set `follow_count`=`follow_count`-1 where `user_id`=?", touserid)   //粉丝数-1
		db.DB.Exec("update `tb_user` set `follower_count`=`follower_count`-1 where `user_id`=?", userid) //关注数-1
	}
	tx.Commit() //提交事务
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
	log.Println("用户：" + userid + "执行了" + stm + "操作")
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

func (rc *RelationController) GetList(context iris.Context) mvc.Result {
	userid := context.FormValue("user_id")
	token := context.FormValue("token")
	if token == "" || !cache.RCExists(token) {
		return mvc.Response{
			Object: listResponse{
				StatusCode: "100",
				StatusMsg:  "登录超时",
			},
		}
	}
	log.Println(token)
	rows, err := db.DB.Query("select `follower_id` from `tb_relation` where `follower_id`=? and `isdeleted`=false", userid)
	if err != nil {
		log.Printf("查询关注列表错误:%v\n", err)
	}
	user_list := make([]User, 0, 50)
	for rows.Next() {
		var user User
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Printf("%v", err)
		}
		//获取关注用户信息
		row := db.DB.QueryRow("select `user_id`,`name`,`follow_count`,`follower_count` from `tb_user` where `user_id`=?", id)
		row.Scan(&user.Id, &user.Name, &user.FollowerCount, &user.FollowCount)
		user.IsFollow = true
		user_list = append(user_list, user)
	}
	defer rows.Close()
	return mvc.Response{
		Object: listResponse{
			StatusCode: "0",
			StatusMsg:  "请求完成",
			UserList:   user_list,
		},
	}
}
