package controller

import (
	db2 "douyin/src/db"
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
	token := ctx.URLParam("token")
	video_id := ctx.URLParam("video_id")
	//user_id := ctx.URLParam("user_id")
	if token == "" {
		return mvc.Response{
			Object: Response{StatusCode: 1, StatusMsg: "请先登录..."},
		}
	}
	//fmt.Println(token, video_id, user_id)
	sql := "select comment_id,comment_user_id,comment_content,comment_latest_time,user_name from comment left join user on comment_user_id=user_id and comment_video_id=? order by comment_latest_time desc;"
	rows, err := sqlSession.Query(sql, video_id)
	if err != nil {
		panic(err)
	}
	var commentList []Comment = make([]Comment, 0)
	var comment Comment
	for rows.Next() {
		//这里是查询video_id视频的所有评论，然后查询出来的全部封装到comment类中，再封装到CommentListResponse中返回到前端
		rows.Scan(&comment.Id, &comment.User.Id, &comment.Content, &comment.CreateDate, &comment.User.Name)
		commentList = append(commentList, comment)
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
	//若前端传过来的是一个json格式字符串的话下面方法可以获取并封装成一个actionRequest对象
	err := ctx.ReadJSON(&actionRequest)
	if err != nil {
		log.Fatal(err)
		return mvc.Response{Object: Response{StatusCode: 1, StatusMsg: "请求失败!!!"}}
	}
	if actionRequest.ActionType == 1 {
		//发布评论
		sql := "insert into comment values(?,?,?,?,?,?)"
		_, err := sqlSession.Exec(sql, nil, actionRequest.UserId, actionRequest.VideoId, actionRequest.CommentText, time.Now().Format("2006-01-02 15:04:05"), nil)
		if err != nil {
			return mvc.Response{Object: Response{StatusCode: 1, StatusMsg: "发表失败!!!"}}
		}
	} else {
		//删除评论
		sql := "delete from comment where comment_id=?"
		_, err := sqlSession.Exec(sql, actionRequest.CommentId)
		if err != nil {
			return mvc.Response{
				Object: Response{StatusCode: 1, StatusMsg: "删除失败,请重试!!!"},
			}
		}
	}
	return mvc.Response{
		Object: Response{StatusCode: 0, StatusMsg: "操作成功!!!"},
	}
}
