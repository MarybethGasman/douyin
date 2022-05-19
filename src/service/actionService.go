package service

import (
	"douyin/src/cache"
	"douyin/src/db"
	"github.com/kataras/iris/v12"
	"io"
	"os"
)

const (
	FilePath = "upload/"
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
				"status_msg":  "MkdirAll filed,err: " + err.Error(),
			})
			return
		}
	}

	fw, err := os.Create(FilePath + head.Filename)

	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 2,
			"status_msg":  "创建文件失败,err: " + err.Error(),
		})
		return
	}
	defer fw.Close()
	_, err = io.Copy(fw, file)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status_code": 3,
			"status_msg":  "拷贝文件失败,err: " + err.Error(),
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"status_code": 0,
		"status_msg":  "保存文件成功",
	})
	// TODO 将信息插入到数据库
	dao := &db.FeedDao{}
	id, err := rcGet.Int64()
	if err != nil {
		panic("获取id失败,err: " + err.Error())
		return
	}

	user := dao.SelectUserById(id)
	dao.InsertVideo(&db.TbVideo{
		AuthorName: user.Name,
		PlayUrl:    head.Filename,
	})
}

func isHasDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}
