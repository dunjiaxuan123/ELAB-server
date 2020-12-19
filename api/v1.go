package api

import (
	"awesomeProject/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//func All(c *gin.Context) {
//	var user User
//
//	user.Username = c.PostForm("username")
//	user.Password = c.PostForm("password")
//
//	c.JSON(http.StatusOK, gin.H{
//		"username": user.Username,
//		"password": user.Password,
//	})
//	fmt.Println(user.Username)
//	fmt.Println(user.Password)
//}

func GetTime(c *gin.Context) {
	//获取查询的结果
	result := model.GetAllDB()
	//获取人数
	numbers := len(result)
	//arr := QuickSort(result, 0, numbers-1)
	for k, _ := range result {
		fmt.Println(result[k])
	}
	c.JSON(http.StatusOK, gin.H{
		//分别是人员总数和查询结果，result的格式是 map[int]map[string]string
		"total": numbers,
		//不必在意命名^-^
		"name": result,
	})
}
