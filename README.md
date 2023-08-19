>这次打算使用gin框架，搭建一个小项目，实现任务清单的功能，可以对完成的任务进行确认，也可以删除，主要逻辑上实现对待办事项的增删改查功能。

## 前端
使用现成的Vue框架搭建的前端模板。首先去git hub上克隆一下前端代码：[Q1mi/bubble_frontend: bubble frontend base vue2.0 (github.com)](https://github.com/Q1mi/bubble_frontend)
克隆岛本地之后的项目目录：
![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/4b7398abbcf34d379967239f39d0c267~tplv-k3u1fbpfcp-watermark.image?)

之后，在目录下打开cmd
```bash
npm install
```
下载相关依赖，之后再使用
```bash
npm run build
```
对项目进行打包，生成dist目录,目录下的内容如下:

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/59641545b3674edca77bee03ccaf1890~tplv-k3u1fbpfcp-watermark.image?)
包含一个html文件和static文件夹，内含css和json和fonts

## 后端
创建项目，将前端打包的好号的dist目录下的文件拷贝到项目中

![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c0e806ad21cf4c4c81553aaa52c93846~tplv-k3u1fbpfcp-watermark.image?)
可以看到，index.html  static文件夹都被拷贝过来了。

之后，建立template文件夹，将前端代码放进去，编写main.go文件如下：
```go
package main  
  
import (  
"github.com/gin-gonic/gin"  
"net/http"  
)  
  
func main() {  
r := gin.Default()  
//告诉gin框架模板文件引用的静态文件去哪里找  
r.Static("/static", "static")  
// 告诉gin框架去哪里找模板文件  
r.LoadHTMLGlob("templates/*")  
r.GET("/", func(c *gin.Context) {  
c.HTML(http.StatusOK, "index.html", nil)  
})  
r.Run()  
}
```
运行项目，访问localhostL：8080便可以看到前端界面:

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/fd84342c2d494dd5a767090b30df3c30~tplv-k3u1fbpfcp-watermark.image?)

之后，简历结构体与前端进行对接：

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6a42a9a6d93048cabf54b0fd510a6a33~tplv-k3u1fbpfcp-watermark.image?)

## 实现的结果
具体实现的过程就不一一详述了，感兴趣的同学可以访问我的github项目地址：[yehowlong/bubble: Gin框架小项目再实践之任务清单增删改查 (github.com)](https://github.com/yehowlong/bubble)

下面是增删改查功能的展示：
增加一个任务：

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f946d6b6bedc4f3c8ce192397c837c6a~tplv-k3u1fbpfcp-watermark.image?)

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/709c4547f5604fcabfaa4aa496cf592f~tplv-k3u1fbpfcp-watermark.image?)

删除一个任务：

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/3bfdbc907e794395bbbf374384c54742~tplv-k3u1fbpfcp-watermark.image?)

修改一个任务的状态:

![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c5d523b3566349b7856ed10b05c4a13b~tplv-k3u1fbpfcp-watermark.image?)

查询所有任务：

访问主页自动查询所有任务
![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d5903e8a3909444c858333e7400e6e57~tplv-k3u1fbpfcp-watermark.image?)



## 参考
[lesson25_小清单项目启动_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1gJ411p7xC?p=25&vd_source=593aa969bb714309669e5f5ccafe096a)

[Q1mi/bubble_frontend: bubble frontend base vue2.0 (github.com)](https://github.com/Q1mi/bubble_frontend)

[yehowlong/bubble: Gin框架小项目再实践之任务清单增删改查 (github.com)](https://github.com/yehowlong/bubble)