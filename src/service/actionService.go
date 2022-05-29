package service

import (
	"douyin/src/cache"
	"douyin/src/config"
	"douyin/src/db"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"strconv"
	"time"
)

var (
	FilePath = config.AppConfig.GetString("video.filePath")
)

// Contribution 视频投稿
func Contribution(ctx iris.Context) {
	r := ctx.Request()

	text := r.FormValue("token")

	// TODO 鉴权
	rcGet := cache.RCGet(text)
	if rcGet == nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 4,
			"status_msg":  "该用户没用权限",
		})
		return
	}

	title := r.FormValue("title")

	file, head, err := r.FormFile("data")
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 1,
			"status_msg":  "没有找到data参数,err: " + err.Error(),
		})
		return
	}
	defer file.Close()

	if b, _ := isHasDir(FilePath); !b {
		err = os.MkdirAll(FilePath, 0777)
		if err != nil {
			ctx.JSON(map[string]interface{}{
				"status_code": 5,
				"status_msg":  "create folder failed,err: " + err.Error(),
			})
			return
		}
	}

	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + head.Filename

	fw, err := os.Create(FilePath + fileName)

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 2,
			"status_msg":  "create file failed,err: " + err.Error(),
		})
		return
	}
	defer fw.Close()
	_, err = io.Copy(fw, file)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 3,
			"status_msg":  "copy file failed,err: " + err.Error(),
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"status_code": 0,
		"status_msg":  "save file success",
	})
	// TODO 将信息插入到数据库
	dao := &db.FeedDao{}
	id, err := rcGet.Int64()
	if err != nil {
		panic("get id failed,err: " + err.Error())
		return
	}

	dao.InsertVideo(&db.TbVideo{
		UserId:  id,
		PlayUrl: fileName,
		Title:   title,
	})
}

func isHasDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}
