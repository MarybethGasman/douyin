package controller

import (
	"douyin/src/cache"
	. "douyin/src/common"
	db2 "douyin/src/db"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentListRequest struct {
	UserId  int64  `json:"user_id"`
	Token   string `json:"token"`
	VideoId int64  `json:"video_id"`
}

type CommentListResponse struct {
	//Response Response
	StatusCode  int32     `json:"status_code"` //0成功,其他值失败
	StatusMsg   string    `json:"status_msg"`
	CommentList []Comment `json:"comment_list"`
}

type CommentActionRequest struct {
	UserId      int64  `json:"user_id"`
	Token       string `json:"token"`
	VideoId     int64  `json:"video_id"`
	ActionType  int64  `json:"action_type"` //1-发布评论，2-删除评论
	CommentText string `json:"comment_text"`
	CommentId   int64  `json:"comment_id"` //需要删除评论的id
}
type CommentActionResponse struct {
	StatusCode int32  `json:"status_code"` //0表示成功，其他值表示失败
	StatusMsg  string `json:"status_msg"`
}

var sqlSession = db2.DB

type CommentController struct {
}

/**

"/douyin/comment/list/?user_id=&token=&video_id=1"
获取所有评论

*/

func (cc *CommentController) GetList(ctx iris.Context) mvc.Result {
	video_id := ctx.URLParam("video_id")
	sql := "select comment_id,tb_comment.user_id,content,create_time,name from tb_comment inner join tb_user on tb_comment.user_id=tb_user.user_id and video_id=? order by create_time desc;"
	rows, err := sqlSession.Query(sql, video_id)
	if err != nil {
		panic(err)
	}
	//延时资源关闭
	defer rows.Close()

	var commentList []Comment = make([]Comment, 0)
	var comment Comment
	for rows.Next() {
		//这里是查询video_id视频的所有评论，然后查询出来的全部封装到comment类中，再封装到CommentListResponse中返回到前端
		rows.Scan(&comment.Id, &comment.User.Id, &comment.Content, &comment.CreateDate, &comment.User.Name)
		commentList = append(commentList, comment)
	}
	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}

	return mvc.Response{
		Object: CommentListResponse{StatusCode: 0, StatusMsg: "查询成功!", CommentList: commentList},
	}
}

/**
"/douyin/comment/action
发表评论或者删除评论，根据
*/

func (cc *CommentController) PostAction(ctx iris.Context) mvc.Result {
	var actionRequest CommentActionRequest

	fmt.Println(ctx.URLParams())

	//actionRequest.UserId, _ = ctx.URLParamInt64("user_id")
	actionRequest.Token = ctx.URLParam("token")
	actionRequest.VideoId, _ = ctx.URLParamInt64("video_id")
	actionRequest.ActionType, _ = ctx.URLParamInt64("action_type")
	actionRequest.CommentText = ctx.URLParam("comment_text")
	actionRequest.CommentId, _ = ctx.URLParamInt64("comment_id")
	actionRequest.UserId, _ = cache.RCGet(actionRequest.Token).Int64()

	if actionRequest.ActionType == 1 {
		//发布评论
		sql := "insert into tb_comment values(?,?,?,?,?)"
		_, err := sqlSession.Exec(sql, nil, actionRequest.UserId, actionRequest.VideoId, actionRequest.CommentText, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			return mvc.Response{Object: Response{StatusCode: 1, StatusMsg: "发表失败!!!"}}
		}
		
		rows, _ := sqlSession.Query("select count(*) from tb_comment where video_id=?", actionRequest.VideoId)
		var count int
		if rows.Next() {
			rows.Scan(&count)
		}
		sql1 := "update tb_video set comment_count=? where video_id=?"
		sqlSession.Exec(sql1, count, actionRequest.VideoId)
	} else {
		//删除评论,得保证是本人删除，也就是删除的是当前用户的评论
		sql := "delete from tb_comment where comment_id=? and user_id=?"
		_, err := sqlSession.Exec(sql, actionRequest.CommentId, actionRequest.UserId)
		if err != nil {
			return mvc.Response{
				Object: Response{StatusCode: 1, StatusMsg: "删除失败,请重试!!!"},
			}
		}

	
		rows, _ := sqlSession.Query("select count(*) from tb_comment where video_id=?", actionRequest.VideoId)
		var count int64
		if rows.Next() {
			rows.Scan(&count)
		}
		sqlSession.Exec("update tb_video set comment_count=? where video_id=?", count)
	}
	return mvc.Response{
		Object: Response{StatusCode: 0, StatusMsg: "操作成功!!!"},
	}
}
