# douyin_club

前后端解耦优势
1 解放生产力，提高合作效率
2 送耦合的架构更灵活，部署方便
3 性能的提升，可靠性的提升
前端可以部署多个服务，中间使用ELB NGINX
缺点：
1 工作量大  
2 前后端分离带来的成本
3 系统复杂度加大

后端服务 API
rest是一种设计风格，不是任何架构标准
当前restful api通常使用HTTP作为通信协议，JSON作为数据交换格式

rest api特点
1 同一接口
2 无状态
模式：
分层
CS模式 
REST API原则
1 以URL风格设计API
2 通过不同的METHOD（GET POST PUT DELETE）来区分对资源的CRUD
3 返回码（Status Code）符合HTTP资源描述的规定

api设计：用户
1 创建用户：
URL:/user	method:POST ,SC 201,400,500
2 用户登录：
URL：/user/:username method:POST,SC:200,400,500
3 获取用户基本信息：
URL：/user/:username method:GET ,SC :200,400,401,403,500

401表示没有通过验证
403表示通过了验证，但是没有操作某个资源的权限
4 用户注销：
URL：/user/:username method:DELETE,SC:204,400,401,403,500
204表示删除成功

api设计：用户资源
1 List all videos:列出一个用户下的视频
URL: /user/:username/videos  method:GET ,SC :200,400,500
2 Get one video :
URL: /user/:username/videos/:vid-id  method:GET,SC:200,400,500
3 Delete one video:
URL: /user/:username/videos/:vid-id  method:DELETE,SC:204,400,401,403,500


API 设计：评论
1 Show comments:显示一个视频下面的评论
URL: /videos/:vid-id/comments method:GET,SC:200,400,500
2 Post a comment:发表一个评论==在一个视频ID下面发表一个评论
URL: /videos/:vid-id/comments  method:POST,SC:201,400,500
3 Delete a comment:删除评论==删除一个视频ID下面的一个评论（ID）
URL: /videos/:vid-id/comment/:comment-id  method:DELETE ,SC:204,400,401,403,500


Streaming播放流处理，流控
静态视频，非RTMP （快手，抖音）
独立的服务，可独立部署
统一的api格式


Stream Server
Stream部分
Uploadfile部分

Scheduler任务调度器
1 REST FUL 的HTTP server
2 Timer
3 生产者、消费者模型下的task runner


架构预览
		    Producer/Dispatcher
                  |向下箭头
Timer           channel通信
                  |向下箭头
		   Consumer/Executor


前端服务
Go的模板引擎
模板引擎是将HTML解析和元素预置替换成最终的页面的工具
Go的模板有两种text/template和html/template
Go的模板采用动态生成的模式
		   Go模板引擎
		静态HTML
	        |
	   解析 |
	    | 
	模板       +      动态元素
	  |                   |
	  |	                  |
	  |---> 最终页面 <----|
处理流程：
静态HTML  --模板解析器--->模板 + 动态元素（数据） ----渲染----> 最终页面

web模块
127.0.0.1:8000/upload
127.0.0.1:9000/upload
这是两个不同的域
前端服务转发：
1 proxy 解决前端跨域访问问题
前端访问 127.0.0.1:8000/upload
服务转发工具：NGINX   Apache这些反向代理服务器
实际访问的是 127.0.0.1:9000/upload
2 api
前端发来请求
{
	url: "xxx",
	method: "POST",
	message: "123"
}
后端用httpclient处理

proxy是直接代理
某些api请求只能使用proxy代理，比如上传文件，用api方式处理起来很麻烦
如果有很多URL访问需要跨域那么直接用proxy会简便些
cross origin resource sharing
proxy方式和api方式的区别：
proxy类似NGINX，Apache这些反向代理软件
api方式相当于在后台用request这些网络访问库，自定义转发
proxy方式简单高效，不会对原始数据做改变，只是做域名转换，数据原样转发
api模式可以对原始请求数据进行修改，增删改查 然后用http库模拟HTTP请求（一些爬虫系统就是这么做的）


Cloud Native云原生
广义的云原生特点：
1 松耦合的架构（SOA/Microsoftservice）
2 无状态，伸缩性，冗余
3 平台无关性（不同的服务可以部署在不同的云上）

部署发布：
1 自动化部署
2 良好的迁移性
3 多云共生（把业务拆分到不同的云）

OSS对象存储
1 BUCKET存储桶
2 支持流媒体特性（图片显示，视频播放）
3 支持HTTP和HTTPS方式访问
4 支持加密访问
5 支持ELB
6 

OSS存储模式：
1 local ---> oss
2 client --> local --> policy,client+policy --> oss
3 oss --> callback --> local server

LB --> api1
LB --> api2

一个request进来  ----> LB --> 1,2
REDIS

cache ===> db强制同步
