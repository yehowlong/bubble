>这是我的上一篇笔记：[Gin框架小项目再实践 ｜ 青训营 - 掘金 (juejin.cn)](https://juejin.cn/post/7268558840264917052)在上篇笔记中，我完成了一个任务清单的小项目，能够实现对任务的增删改查。但是项目的目录结构跟企业中实际开发的结构差别还是有点大，因此我决定学习并实践一下如何改造项目的目录结构。
## 目前的目录结构展示
使用tree /f命令查看目录结构如下：

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0b4650516ea64208ab4d21fdd28a24db~tplv-k3u1fbpfcp-watermark.image?)
接下来，将介绍各个层如何改造
## Controller
用来控制路由进来以后，执行哪些函数。

首先，简历controller文件夹下的controller.go文件，之后将各个处理器封装成函数放进去。

将请求首页的代码

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/28f99ed86cb94b40bf3960e36e987b7b~tplv-k3u1fbpfcp-watermark.image?)
改为：

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0e768b553e9d44a2a4b61c0e708c50c0~tplv-k3u1fbpfcp-watermark.image?)

在controller.controller里编写具体的函数：
![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/eb53cb0e35e14fdeb0a38cefeeae32bb~tplv-k3u1fbpfcp-watermark.image?)
同理，将post请求改为如下样式：

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/3ba3932853aa466b81c0cf3021d5ee5c~tplv-k3u1fbpfcp-watermark.image?)

其余几个结构均需要改名，封装，与上面的步骤相似，成果如下：

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/b0ff94eca4634ae1973d007173005132~tplv-k3u1fbpfcp-watermark.image?)

## dao层
即database access object 数据操作对象，我们将所有与数据库有关的代码移到这个目录下的mysql.go中：
![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/03202c2dc13645fda2977005dfaa071e~tplv-k3u1fbpfcp-watermark.image?)
main.go中初始化数据库的部分就变成了：
```go
err := dao.InitMySQL()
dao.DB.AutoMigrate(&Todo{})
defer dao.Close() //程序退出关闭数据库连接
```
## model
这里主要存放模型的定义，以及对这个模型的增删改查操作。

改造之后，应该把具体的增删改查的逻辑放在model层，而controller层应该只用调用model层的逻辑就可以了。如果逻辑比较复杂，controller层应该调用service层，service层再调用model层
即步骤应该是：
>url -->controller -->service -->model  
>请求来了-->控制器-->业务层-->模型层的增删改查

以controller层中的GetTodoList函数为例，原本controller中代码如下:
```go
func GetTodoList(c *gin.Context) {  
var todoList []Todo  
if err = DB.Find(&todoList).Error; err != nil {  
c.JSON(http.StatusOK, gin.H{"error": err.Error()})  
} else {  
c.JSON(http.StatusOK, todoList)  
}  
}
```

修改完代码如下：
```go
func GetTodoList(c *gin.Context) {  
todoList, err := models.GetAllTodo()  
if err != nil {  
c.JSON(  
http.StatusOK,  
gin.H{"error": err.Error()},  
)  
} else {  
c.JSON(http.StatusOK, todoList)  
}  
}
```
同时，models.GetAllTodo()函数的定义如下：
```go
func GetAllTodo() (todoList []*Todo, err error) {  
if err := dao.DB.Find(&todoList).Error; err != nil {  
return err  
  
}  
return nil, err  
}
```
同理，将controller层中的DeleteATodo函数改造成：
```go
func DeleteATodo(c *gin.Context) {  
id, ok := c.Params.Get("id")  
if !ok {  
c.JSON(http.StatusOK, gin.H{"error": "无效的id"})  
return  
}  
if err := models.DeleteATodo(id); err != nil {  
c.JSON(http.StatusOK, gin.H{"error": err.Error()})  
} else {  
c.JSON(http.StatusOK, gin.H{id: "deleted"})  
}  
}
```
modeler层中的DeleteATodo函数如下：
```go
func DeleteATodo(id string) (err error) {  
err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error  
}
```
## routers
main.go中还包含很多关于路由注册代码，其实这部分代码也是可以单独拿出来的。
此时main函数的代码如下：

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/14294c85fb4a486c8595c8d207b663ac~tplv-k3u1fbpfcp-watermark.image?)
非常的简洁
routers.go的代码如下：
![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0aa04d7160fe469aa6868f74dea934eb~tplv-k3u1fbpfcp-watermark.image?)

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0874c05431ce4a4aa7770b68f86ab35c~tplv-k3u1fbpfcp-watermark.image?)
## 改造完的目录结构展示

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d560914e09124600bce20446a9057c6a~tplv-k3u1fbpfcp-watermark.image?)

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/a8193a08b29248098985536a124bd572~tplv-k3u1fbpfcp-watermark.image?)

最后项目也推送至了github上面，感兴趣的同学可以下载下来研究一下:[yehowlong/bubble: Gin框架小项目再实践之任务清单增删改查 (github.com)](https://github.com/yehowlong/bubble)

`踩的一些小坑`

老是容易把函数的返回值(第二个括号里的参数)写到函数的参数列表里面去(第一个括号里的参数)


## 参考
[lesson27_企业级项目结构拆分_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1gJ411p7xC?p=27&spm_id_from=pageDriver&vd_source=593aa969bb714309669e5f5ccafe096a)