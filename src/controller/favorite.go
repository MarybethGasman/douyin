package controller

// 点赞和点赞列表实现
// 直接对数据库进行操作
// 优点：
// 1.实现简单
// 2.数据准确，且实时性好
// 3.适用于业务规模较小，大概在十万级左右
// 缺点:
// 1.点赞属于高频操作，若请求过多将导致数据库压力过大
// 方案二：使用redis缓存实现
// 使用Set，key 为 favorite_{用户id}，value 为 该用户点赞的视频id集合
// 点赞操作只需向redis中加入一条数据，redis内数据定期刷回数据库
// 查询某用户点赞的所有视频：从数据库和redis取出该用户点赞的所有视频id集合，然后用此集合查询结果返回
// 判断某用户是否点赞某视频：先查redis，存在直接返回，不存在去数据库找

import (
	"douyin/src/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

import (
	"douyin/src/cache"
	. "douyin/src/common"
	. "douyin/src/db"
	"log"
	"time"
)

type favoriteResponse struct {
	Response
}

type favoriteListResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg,omitempty"`
	VideoList  []VideoList2 `json:"video_list,omitempty"`
}

type FavoriteController struct {
}

func (fc *FavoriteController) PostAction(ctx iris.Context) mvc.Result {
	var token = ctx.URLParamDefault("token", "")
	var videoId = ctx.URLParamInt64Default("video_id", -1)
	// 动作类型 1-点赞 2-取消点赞
	var actionType = ctx.URLParamIntDefault("action_type", -1)

	if token == "" || videoId == -1 || actionType < 1 || actionType > 2 {
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
	tx, err := DB.Begin()
	if RecordExists(userId, videoId) {
		//如果数据库中有记录，那么根据action_type改变is_deleted即可
		_, err = tx.Exec(
			"update tb_favorite set is_deleted = ? where user_id = ? and video_id = ?",
			actionType-1, userId, videoId)
		//更新tb_video的冗余字段，若为点赞操作则视频点赞数+1，若为取消点赞操作则-1
		number := 1
		if actionType == 2 {
			number = -1
		}
		_, err = tx.Exec("update tb_video set favorite_count = favorite_count + ? where video_id = ?",
			number, videoId)
		if err != nil {
			tx.Rollback()
			panic(err.Error())
		}
	} else {
		//若不存在记录，则只处理点赞操作
		if actionType == 1 {
			_, err = tx.Exec(
				"insert into tb_favorite(user_id,video_id,is_deleted) values (?,?,?)",
				user.Id, videoId, 0)
			_, err = tx.Exec("update tb_video set favorite_count = favorite_count + 1 where video_id = ?", videoId)
			if err != nil {
				tx.Rollback()
				panic(err.Error())
			}
		}
	}
	tx.Commit()
	response = mvc.Response{
		Object: Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
	}
	return response
}

func (fc *FavoriteController) GetList(ctx iris.Context) mvc.Result {
	var token = ctx.URLParamDefault("token", "")
	var authorId = ctx.URLParamInt64Default("user_id", -1)
	if token == "" || authorId == -1 {
		return mvc.Response{
			Object: Response{
				StatusCode: -1,
				StatusMsg:  "缺少参数或参数错误",
			},
		}
	}

	var response mvc.Response
	userId, _ := cache.RCGet(token).Int64()
	cache.RCSet(token, userId, time.Minute*30)
	rows, err := DB.Query("select video_id, play_url, cover_url, favorite_count, comment_count, title,tu.user_id,tu.name,follow_count, follower_count from tb_video tv inner join tb_user tu on tv.user_id = tu.user_id where video_id in (select video_id from tb_favorite where user_id = ? and is_deleted = 0)", authorId)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	baseUrl := config.AppConfig.GetString("video.imageUrl")
	var videoListResponse VideoListResponse
	for rows.Next() {
		var favoriteVideo VideoList2
		err = rows.Scan(&favoriteVideo.Id, &favoriteVideo.PlayURL, &favoriteVideo.CoverURL, &favoriteVideo.FavoriteCount, &favoriteVideo.CommentCount, &favoriteVideo.Title, &favoriteVideo.Author.Id, &favoriteVideo.Author.Name, &favoriteVideo.Author.FollowCount, &favoriteVideo.Author.FollowerCount)
		favoriteVideo.PlayURL = baseUrl + favoriteVideo.PlayURL
		favoriteVideo.CoverURL = baseUrl + favoriteVideo.CoverURL
		favoriteVideo.Author.IsFollow = isFollow(userId, favoriteVideo.Author.Id)
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

func isFollow(userId int64, authorId int64) bool {
	if userId == -1 {
		return false
	}
	row := DB.QueryRow("select relation_id from tb_relation where follower_id = ? and following_id = ? and isdeleted = 0", userId, authorId)
	relationId := -1
	row.Scan(&relationId)
	if relationId > 0 {
		return true
	}
	return false
}
