package main

import (
	"bubble/controller"
	"bubble/dao"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&Todo{})
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	r.Run()
	//v1版本代码

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一个代办事项
		//v1Group.GET("/todo/:id", func(c *gin.Context) {
		//
		//})
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

}
