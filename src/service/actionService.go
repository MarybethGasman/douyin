package service

import (
	"bytes"
	"douyin/src/cache"
	"douyin/src/config"
	"douyin/src/db"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/kataras/iris/v12"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

var (
	FilePath = config.AppConfig.GetString("video.filePath")
)

// Contribution 视频投稿
func Contribution(ctx iris.Context) {
	r := ctx.Request()

	text := r.FormValue("token")

	//  鉴权
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

	// 保存封面
	fileImage := GetSnapshot(FilePath+fileName, FilePath+strings.Trim(fileName, ".mp4"), 60)

	// 将信息插入到数据库
	dao := &db.FeedDao{}
	id, err := rcGet.Int64()
	if err != nil {
		panic("get id failed,err: " + err.Error())
	}

	dao.InsertVideo(&db.TbVideo{
		UserId:   id,
		PlayUrl:  fileName,
		CoverUrl: fileImage,
		Title:    title,
	})
}

func isHasDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	err = imaging.Save(img, snapshotPath+".jpeg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(snapshotPath, "/")
	snapshotName = names[len(names)-1] + ".jpeg"
	return
}
