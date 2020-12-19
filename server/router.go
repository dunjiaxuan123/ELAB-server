package server

import (
	"awesomeProject/api"
	"awesomeProject/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//分组路由初始化
	r := gin.Default()
	//中间件方法，
	r.Use(middleware.Cors())

	//分组路由
	v1 := r.Group("/api")
	{
		//v1.POST("/form", api.All)
		//之后的接口就在这里写就可以，相关的用法参考文档资料即可
		v1.POST("/get", api.GetTime)
	}

	return r
}
