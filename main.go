package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	return DB.DB().Ping()

}
func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run()
	//v1版本代码

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			//3.返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
				//c.JSON(http.StatusOK,gin.H{
				//	"code":2000,
				//	"msg":"success",
				//	"data":todo,
				//})
			}

		})
		//查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		//查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}
}
