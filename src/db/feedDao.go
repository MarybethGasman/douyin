package db

import (
	"douyin/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"strconv"
	"time"
)

var (
	db  = NewDB()
	dns = config.AppConfig.Get("datasource.dataSourceName").(string) + "?parseTime=true" //gorm框架 连接mysql
)

type FeedDao struct {
}

type TbVideo struct {
	VideoId       int64     `gorm:"primaryKey"`
	AuthorName    string    `gorm:"default:"`
	PlayUrl       string    `gorm:"default:"`
	CoverUrl      string    `gorm:"default:"`
	FavoriteCount int64     `gorm:"default:"`
	CommentCount  int64     `gorm:"default:"`
	CreateDate    time.Time `gorm:"default:"`
	UpdateDate    time.Time `gorm:"default:"`
}

type TbUser struct {
	UserId        int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      int8
	Password      string
}

func NewDB() *gorm.DB {
	open, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		println("create gorm filed,err: " + err.Error())
		os.Exit(1)
	}
	return open
}

// Select30VideoByUpdate 根据时间戳获取最多30条视频记录
func (v *FeedDao) Select30VideoByUpdate(unix string) ([]TbVideo, error) {
	var tb []TbVideo
	date := "0000-00-00 00:00:00"
	if len(unix) > 0 {
		// 时间戳转换为时间
		parseInt, err := strconv.ParseInt(unix, 10, 64)
		if err != nil {
			return nil, err
		}
		date = time.UnixMilli(parseInt).Format("2006-01-02 15:04:05")
	}
	db.Order("update_date DESC").Where("update_date < ?", date).Limit(30).Find(&tb)
	return tb, nil
}

func (v *FeedDao) InsertVideo(table *TbVideo) bool {
	db.Create(table)
	return true
}

func (v *FeedDao) SelectUserById(id int64) *TbUser {
	user := &TbUser{}
	db.Where("user_id = ?", id).First(user)
	return user
}
