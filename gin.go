package main

import (
	"awesomeProject/conf"
	"awesomeProject/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//初始化数据库连接
	conf.Init()
	//启动分组路由
	r := server.NewRouter()
	//sc := model.TimmingMission()
	//在10000端口监听请求
	r.Run(":10000")
	//go c.Start() //启动定时任务
	//select {}
}
