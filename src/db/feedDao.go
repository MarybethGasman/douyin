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
	dns = config.AppConfig.GetString("datasource.dataSourceName") + "?parseTime=true" //gorm框架 连接mysql
)

type FeedDao struct {
}

type TbUser struct {
	UserId        int64  `gorm:"primaryKey"`
	Name          string `gorm:"default:"`
	FollowCount   int64  `gorm:"default:"`
	FollowerCount int64  `gorm:"default:"`
	IsFollow      int8   `gorm:"default:"`
	Password      string `gorm:"default:"`
}

type TbVideo struct {
	VideoId       int64 `gorm:"primaryKey"`
	UserId        int64
	PlayUrl       string    `gorm:"default:"`
	CoverUrl      string    `gorm:"default:"`
	FavoriteCount int64     `gorm:"default:"`
	CommentCount  int64     `gorm:"default:"`
	Title         string    `gorm:"default:"`
	CreateDate    time.Time `gorm:"default:"`
	UpdateDate    time.Time `gorm:"default:"`
	TbUser        TbUser    `gorm:"foreignKey:UserId;references:UserId"`
}

func (TbVideo) TableName() string {
	return "tb_video"
}

func (TbUser) TableName() string {
	return "tb_user"
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

// SelectVideoByUpdate 根据时间戳获取最多number条视频记录
func (v *FeedDao) SelectVideoByUpdate(unix string, number int) ([]TbVideo, error) {
	var tb []TbVideo
	if len(unix) > 0 {
		// 时间戳转换为时间
		parseInt, err := strconv.ParseInt(unix, 10, 64)
		if err != nil {
			return nil, err
		}
		date := time.UnixMilli(parseInt).Format("2006-01-02 15:04:05")
		db.Preload("TbUser").Limit(number).Where("update_date < ?", date).Order("update_date DESC").Find(&tb)
		return tb, nil
	}
	db.Preload("TbUser").Limit(number).Order("update_date DESC").Find(&tb)
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
