package utils

import (
	. "douyin/src/config"
	"douyin/src/db"
	"fmt"
	"github.com/gohouse/converter"
	"gorm.io/gen"
)

/*

只创建结构体

包为："github.com/gohouse/converter"

SavePath为输出结果存放文件

DSN用于配置数据库的信息

TagKey用于配置Tag，此处写为gorm

EnableJsonTag用于确认是否输出时有json标签

Table用于指定表，如果不写则输出库中的所有表
*/
func GetGen() {
	dataSourceName := AppConfig.Get("datasource.dataSourceName").(string)
	err := converter.NewTable2Struct().
		SavePath("./dao.go").
		Dsn(dataSourceName).
		TagKey("gorm").
		EnableJsonTag(true).
		//Table().
		Run()
	fmt.Println(err)
}


/*
这个连查询语句都给你创建好
包为：	"gorm.io/gen"
*/
//
// 想对已有的model生成crud等基础方法可以直接指定model struct ，例如model.User{}
// 如果是想直接生成表的model和crud方法，则可以指定表名称，例如g.GenerateModel("company")
// 想自定义某个表生成特性，比如struct的名称/字段类型/tag等，可以指定opt，例如g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
//g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
//这句有bug，别用


func GetGen2() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../model", // 生成代码的输出路径
	})
	g.UseDB(db.GetDBConnect())
	//g.ApplyBasic(g.GenerateModel("user")) // 绑定表结构
	g.GenerateModelAs("tb_comment","Comment")
	g.GenerateModelAs("tb_favorite","Favorite")
	g.GenerateModelAs("tb_relation","Relation")
	g.GenerateModelAs("tb_user","User")
	g.GenerateModelAs("tb_video","Video")

	//g.ApplyBasic(model.Comment{})
	//g.ApplyBasic(g.GenerateModelAs("tb_comment","Comment"),
	//g.GenerateModelAs("tb_favorite","Favorite"),
	//g.GenerateModelAs("tb_relation","Relation"),
	//g.GenerateModelAs("tb_user","User"),
	//g.GenerateModelAs("tb_video","Video"),
	//) // 绑定表结构

	//g.ApplyInterface(func( model.UserMethod) {},g.GenerateModel("tb_user"))
	g.Execute()
}