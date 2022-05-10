package db

import (
	"log"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	row := DB.QueryRow("select * from tb_user")
	var hh string
	row.Scan(&hh, &hh, &hh)
	log.Println("开始")
	log.Println(hh)
	log.Println("结束")
}
