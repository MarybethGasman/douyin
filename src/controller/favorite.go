package controller

import (
	"douyin/src/cache"
	. "douyin/src/common"
	. "douyin/src/db"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"time"
)

/**
点赞返回结构体
*/
type favoriteResponse struct {
	Response
}

/**
点赞列表返回结构体
*/
type favoriteListResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg,omitempty"`
	VideoList  []VideoList2 `json:"video_list,omitempty"`
}

type FavoriteController struct {
}

/**
点赞操作
*/
func (fc *FavoriteController) PostAction(ctx iris.Context) mvc.Result {
	var token = ctx.URLParamDefault("token", "")
	var videoId = ctx.URLParamInt64Default("video_id", -1)
	// 动作类型 1-点赞 2-取消点赞
	var actionType = ctx.URLParamIntDefault("action_type", -1)

	if token == "" || videoId == -1 || actionType == -1 {
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "缺少参数或参数错误",
			},
		}
	}
	if !cache.RCExists(token) {
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "鉴权失败，检查登录状态",
			},
		}
	}
	var response mvc.Response
	userId, _ := cache.RCGet(token).Int64()
	cache.RCSet(token, userId, time.Minute*30)
	user := SelectUserById(userId)
	if actionType == 1 {
		tx, err := DB.Begin()
		_, err = tx.Exec(
			"insert into tb_favorite(user_id,video_id,is_deleted) values (?,?,?)",
			user.Id, videoId, 0)
		_, err = tx.Exec("update tb_video set favorite_count = favorite_count + 1 where video_id = ?", videoId)
		if err != nil {
			tx.Rollback()
			panic(err.Error())
		}
		tx.Commit()
		response = mvc.Response{
			Object: Response{
				StatusCode: 0,
				StatusMsg:  "点赞成功",
			},
		}
	} else if actionType == 2 {
		tx, err := DB.Begin()
		_, err = tx.Exec(
			"update tb_favorite set is_deleted = 1 where user_id = ? and video_id = ?",
			user.Id, videoId)
		_, err = tx.Exec("update tb_video set favorite_count = favorite_count - 1 where video_id = ?", videoId)
		if err != nil {
			tx.Rollback()
			panic(err.Error())
		}
		tx.Commit()
		response = mvc.Response{
			Object: Response{
				StatusCode: 0,
				StatusMsg:  "取消点赞成功",
			},
		}
	} else {
		response = mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "非法参数",
			},
		}
	}
	return response
}

/**
点赞列表操作
*/
func (fc *FavoriteController) GetList(ctx iris.Context) mvc.Result {
	/**
	取出客户端传回变量
	*/
	var token = ctx.URLParamDefault("token", "")
	var userId = ctx.URLParamInt64Default("user_id", -1)

	if token == "" || userId == -1 {
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "缺少参数或参数错误",
			},
		}
	}
	if !cache.RCExists(token) {
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "鉴权失败，检查登录状态",
			},
		}
	}
	var response mvc.Response
	userId, _ = cache.RCGet(token).Int64()
	cache.RCSet(token, userId, time.Minute*30)
	user := SelectUserById(userId)
	rows, err := DB.Query("select video_id, play_url, cover_url, favorite_count, comment_count, title,tu.user_id,tu.name,follow_count, follower_count from tb_video tv inner join tb_user tu on tv.user_id = tu.user_id where video_id in (select video_id from tb_favorite where favorite_id = ? and isdeleted=false)", user.Id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var videoListResponse VideoListResponse
	for rows.Next() {
		var favoriteVideo VideoList2
		err = rows.Scan(&favoriteVideo.Id, &favoriteVideo.PlayURL, &favoriteVideo.CoverURL, &favoriteVideo.FavoriteCount, &favoriteVideo.CommentCount, &favoriteVideo.Title, &favoriteVideo.Author.Id, &favoriteVideo.Author.Name, &favoriteVideo.Author.FollowCount, &favoriteVideo.Author.FollowerCount)
		videoListResponse.VideoLists = append(videoListResponse.VideoLists, favoriteVideo)
		if err != nil {
			log.Fatalln(err)
		}
	}
	response = mvc.Response{
		Object: videoListResponse,
	}
	return response
}
