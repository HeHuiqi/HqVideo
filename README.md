## 项目说明

使用 "github.com/julienschmidt/httprouter"这个库来作为整个项目的路由库
### api层（独立服务）
这一层主要时多数据库操作,供web及客户端调用

```
$ go run api/main.go
Video:http://127.0.0.1:8000


```

### scheduler层（独立服务）
这一层主要时启动一个定时任务去删除用户的视频记录及文件

```
#运行
$ go run scheduler/main.go
Video-scheduler:http://127.0.0.1:9001


```

### streamserver层（独立服务）
这一层主要提供视频播放以及上传

```
# 运行，在9000端口
$ go run streamserver/main.go

Video-stream:http://127.0.0.1:9000


#上传在videos/upload.html

# 播放
http://127.0.0.1:9000/video/test.mp4

```

### web层（独立服务）
这一层就是页面相关的服务，主要是页面的显示,未完成

```
$ go run web/main.go
Video-web:http://127.0.0.1:8001

```
//百度网盘：Go语言实战流媒体视频网站

