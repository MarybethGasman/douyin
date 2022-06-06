package controller

import (
	"database/sql"
	. "douyin/src/cache"
	. "douyin/src/common"
	"douyin/src/db"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"strconv"
)

type actionResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type listResponse struct {
	StatusCode string  `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	UserList   []User2 `json:"user_list,omitempty"`
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
	token := RCGet(context.URLParam("token"))
	if token.Err() == redis.Nil {
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 400,
				StatusMsg:  "登录超时",
			},
		}
	}
	userid := token.Val()
	touserid := context.URLParam("to_user_id")
	if userid == touserid {
		return mvc.Response{
			Object: actionResponse{StatusCode: 0, StatusMsg: "无法关注自己"},
		}
	}
	actiontype, _ := strconv.Atoi(context.URLParam("action_type"))

	row, err := db.DB.Query("select `isdeleted` from `tb_relation` where  `follower_id`=? and `following_id`=?", userid, touserid)
	if err != nil {
		log.Println("查询关注列表错误")
		return mvc.Response{
			Object: actionResponse{
				StatusCode: 500,
				StatusMsg:  "查询列表错误",
			},
		}
	}
	defer row.Close() //在执行完后关闭游标

	//开启事务
	tx, _ := db.DB.Begin()
	var result sql.Result
	//判断relation表中是否有对于关注关系
	exist := row.Next() //exist用于判断记录是否存在

	//TODO 防止关注/取消关注重复操作
	if exist {
		var curtype bool
		row.Scan(&curtype) //如果对应记录存在则匹配对应类型值
		chtype := (actiontype == 2)
		if curtype == chtype { //修改前后状态无变化
			return mvc.Response{
				Object: actionResponse{
					StatusCode: 400,
					StatusMsg:  "你已经关注/取关该用户，请勿重复操作",
				},
			}
		}
	}

	//根据是否存在记录是否存在决定修改/添加记录
	if !exist {
		result, _ = db.DB.Exec("insert into `tb_relation`(`follower_id`,`following_id`,`isdeleted`) "+
			"values(?,?,?)", userid, touserid, (actiontype+1)%2) // 2对应true执行关注操作(删除) 1对应true执行取消关注操作（不删除）
	} else {
		result, _ = db.DB.Exec("update `tb_relation` set isdeleted=? where `follower_id`=? "+
			"and `following_id`=?", (actiontype+1)%2, userid, touserid)
	}

	if actiontype == 1 {
		db.DB.Exec("update `tb_user` set `follow_count`=`follow_count`+1 where `user_id`=?", userid)       //关注数+1
		db.DB.Exec("update `tb_user` set `follower_count`=`follower_count`+1 where `user_id`=?", touserid) //粉丝数+1
	} else {
		db.DB.Exec("update `tb_user` set `follow_count`=`follow_count`-1 where `user_id`=?", userid)       //关注数-1
		db.DB.Exec("update `tb_user` set `follower_count`=`follower_count`-1 where `user_id`=?", touserid) //粉丝数-1
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
			StatusCode: 501,
			StatusMsg:  "发生错误，请稍后重试",
		},
	}
}

func (rc *RelationController) GetFollowList(context iris.Context) mvc.Result {
	token := RCGet(context.URLParam("token"))
	if token.Err() == redis.Nil {
		return mvc.Response{
			Object: listResponse{
				StatusCode: "100",
				StatusMsg:  "登录超时",
			},
		}
	}

	userid := token.Val() //根据token获取用户id
	log.Println("用户（ID）：" + userid + " token:" + context.URLParam("token"))
	rows, err := db.DB.Query("select `following_id` from `tb_relation` where `follower_id`=? and `isdeleted`=false", userid) //获取关注对象的id
	if err != nil {
		log.Printf("查询关注列表错误:%v\n", err)
		return mvc.Response{
			Object: listResponse{
				StatusCode: "500",
				StatusMsg:  "查询出错",
			},
		}
	}
	user_list := make([]User2, 0, 50)
	for rows.Next() {
		var user User2
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Printf("%v", err)
		}
		//获取关注用户信息
		row := db.DB.QueryRow("select `user_id`,`name`,`follow_count`,`follower_count` from `tb_user` where `user_id`=?", id)
		row.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount)
		user.IsFollow = true //从关注列表中只获取关注对象
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
