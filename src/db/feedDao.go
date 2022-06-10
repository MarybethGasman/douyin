package db

import (
	"douyin/src/config"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

type TbRelation struct {
	RelationId  int64  `gorm:"primaryKey"`
	FollowerId  string `gorm:"default:"`
	FollowingId int64  `gorm:"default:"`
	Isdeleted   int8   `gorm:"default:"`
}

type TbFavorite struct {
	FavoriteId int64 `gorm:"primaryKey"`
	UserId     int64 `gorm:"default:"`
	VideoId    int64 `gorm:"default:"`
	IsDeleted  int8  `gorm:"default:"`
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

// 根据用户名和视频Id查询是否有点赞
func (v *FeedDao) IsFavorite(userID int64, videoID int64) bool {
	var count int64
	db.Model(&TbFavorite{}).Where("`user_id` = ? and `video_id` = ? and `is_deleted` = 0", userID, videoID).Count(&count)
	return count > 0
}

// 根据用户名id查询是否有关注 true 表示关注
func (v *FeedDao) IsFollwer(originUserID int64, targetUserID int64) bool {
	var size int64
	db.Model(&TbRelation{}).Where("`follower_id` = ? and `following_id` = ? and `isdeleted` = 0", originUserID, targetUserID).Count(&size)
	return size > 0
}
