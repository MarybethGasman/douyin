package controller

import (
	"douyin/src/cache"
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
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	VideoList  []Video `json:"video_list,omitempty"`
}

type Video struct {
	VideoId       int64  `json:"videoId,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"playUrl,omitempty"`
	CoverUrl      string `json:"coverUrl,omitempty"`
	FavoriteCount int64  `json:"favoriteCount,omitempty"`
	CommentCount  int64  `json:"commentCount,omitempty"`
	IsFavorite    bool   `json:"isFavorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type FavoriteController struct {
}

/**
点赞操作
*/
func (fc *FavoriteController) PostAction(ctx iris.Context) mvc.Result {
	/**
	取出客户端传回变量
	*/
	var token = ctx.URLParam("token")
	var userId = ctx.URLParam("user_id")
	var videoId = ctx.URLParam("video_id")
	// 动作类型 1-点赞 2-取消点赞
	var actionType = ctx.URLParam("action_type")
	// 先根据userId获得用户信息，并取得用户名
	username := ""
	row := DB.QueryRow("select name from tb_user where user_id = ?", userId)
	row.Scan(&username)
	if username == "" {
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "用户不存在，请重新登录",
			},
		}
	}
	// 鉴定token是否存在
	if cache.RCExists(token) {
		//更新用户token
		cache.RCSet(token, userId, time.Minute*30)
		// 已登录
		// action_type 1表示点赞
		if actionType == "1" {
			// 将点赞信息存入数据库(用户名，视频id，删除信息 1表示已删除，0表示未删除)
			_, err := DB.Exec(
				"insert into tb_favorite(username,video_id,isdeleted) values (?,?,?)",
				username, videoId, 0)
			if err != nil {
				panic("点赞视频失败！")
			}
			/**
			将数据存到Redis中
			键值对
			用户：点赞的多个视频
			*/
			cache.RCSAdd("user:liked:"+userId, videoId)
			/**
			返回点赞信息，点赞成功
			*/
			return mvc.Response{
				Object: favoriteResponse{
					Response: Response{
						StatusCode: 0,
						StatusMsg:  "点赞成功",
					},
				},
			}
		} else if actionType == "2" {
			// 2-取消点赞操作
			/**
			数据库层面
			修改数据库表中的信息，将isdeleted更新为1-表示该点赞已删除
			update操作
			根据user_id和video_id联合查询
			*/
			_, err := DB.Exec("update tb_favorite set isdeleted = 1 where username = ? and video_id = ?", username, videoId)
			if err != nil {
				log.Fatal(err)
				panic("取消点赞失败！")
			}
			/**
			Redis层面
			删除Redis键值对中的指定键值对信息
			*/
			cache.RCSRem("user:liked:"+userId, videoId)
			/**
			返回取消点赞信息，点赞失败
			*/
			return mvc.Response{
				Object: favoriteResponse{
					Response: Response{
						StatusCode: 0,
						StatusMsg:  "取消点赞成功",
					},
				},
			}
		} else {
			// action_type 除12外的情况也要考虑返回
			return mvc.Response{
				Object: favoriteResponse{
					Response: Response{
						StatusCode: -1,
						StatusMsg:  "action_type参数不正确",
					},
				},
			}
		}
	} else {
		// 未登录
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "用户未登录或注册",
			},
		}
	}
}

/**
点赞列表操作
*/
func (fc *FavoriteController) PostList(ctx iris.Context) mvc.Result {
	/**
	取出客户端传回变量
	*/
	var token = ctx.URLParam("token")
	var userId = ctx.URLParam("user_id")
	/**
	1、根据userId取出redis中对应的所有的videoId
	2、根据videoId取出数据库中
	*/
	// 取出Redis中对应的所有video_id并存放到result
	result, err := cache.RCSmembers("user:liked:" + userId).Result()
	if err != nil {
		panic("获取视频列表失败")
	}

	// 鉴权token，是否登录或者注册
	if token != "" && cache.RCExists(token) {
		//更新用户token
		cache.RCSet(token, userId, time.Minute*30)
		// 通过循环获取到的result逐个获取每个视频对应的视频信息
		// 定义需要返回的videolist切片
		video_list := make([]Video, 0, 50)
		for _, s := range result {
			// 获取到视频的相关信息，由video_id进行查询
			rows, err := DB.Query(
				"select video_id,author_name, play_url, cover_url, favorite_count, comment_count from tb_video where video_id = ?",
				s)
			if err != nil {
				log.Fatal(err)
			}
			for rows.Next() {
				var video Video
				err := rows.Scan(&video.VideoId, &video.Author.Name, &video.PlayUrl, &video.CoverUrl, &video.FavoriteCount, &video.CommentCount)
				if err != nil {
					return nil
				}
				// 根据author_name查询视频作者的相关具体信息
				query, err := DB.Query(
					"select user_id, follow_count, follower_count from tb_user where name = ?",
					video.Author.Name)
				if err != nil {
					log.Fatal(err)
				}

				for query.Next() {
					err1 := query.Scan(&video.Author.Id, &video.Author.FollowCount, &video.Author.FollowerCount)
					if err1 != nil {
						return nil
					}
					// 关注这块需要重新查询
					var isDeleted int
					todoRes, err := DB.Query(
						"select isdeleted from tb_relation where follower_id = ? and following_id = ?",
						userId, video.Author.Id)
					if err != nil {
						log.Fatal(err)
					}
					todoRes.Scan(&isDeleted)
					//已关注
					if isDeleted == 0 {
						video.Author.IsFollow = true
					} else if isDeleted == 1 {
						// 未关注
						video.Author.IsFollow = false
					} else {
						return mvc.Response{
							Object: favoriteListResponse{
								StatusCode: -1,
								StatusMsg:  "isDelete字段错误",
							},
						}
					}
				}
				video_list = append(video_list, video)
				defer rows.Close()
				defer query.Close()
			}
		}
		return mvc.Response{
			Object: favoriteListResponse{
				StatusCode: 0,
				StatusMsg:  "成功",
				VideoList:  video_list,
			},
		}
	} else {
		// 未登录
		return mvc.Response{
			Object: favoriteListResponse{
				StatusCode: -1,
				StatusMsg:  "用户未登录或注册",
			},
		}
	}
}
