package db

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func TestGormMysql(t *testing.T) {
	var db *gorm.DB
	for i := 0; i < 5; i++ {
		db = GetDBConnect()
		fmt.Println("连接次数",cnt)
	}
	fmt.Println(db)

}
