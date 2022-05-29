package db

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

var user User

func Test_init(t *testing.T) {

}

func TestMysqlConnection(t *testing.T) {

	log.Println("连接数据库")
	row := DB.QueryRow("select * from tb_user")
	rows,err := DB.Query("select * from tb_user")
	defer rows.Close()
	if err!=nil {
		fmt.Printf("insert data error: %v\n", err)
	}

	var ll string
	for rows.Next() {
		e := rows.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow,&ll)
		if e == nil{
			buf,err := json.Marshal(user)
			if err != nil {
				panic(err)
			}
			log.Println(string(buf))
			log.Println(user.Id,user.Name,user.FollowerCount,user.IsFollow,user.IsFollow)
		}
	}
	log.Println("连接完成")
	var hh string
	row.Scan(&hh, &hh, &hh,&hh, &hh, &hh)
	log.Println("开始")
	log.Println(hh)
	log.Println("结束")

}
