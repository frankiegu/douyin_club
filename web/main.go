package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	Logger "github.com/Yq2/douyin_club/web/logs"
)

var log = Logger.Log

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)

	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)

	router.POST("/userhome", userHomeHandler)

	router.POST("/api", apiHandler)   //转发服务模式之API模式

	router.GET("/videos/:vid-id", proxyVideoHandler)  //服务转发模式之proxy模式

	router.POST("/upload/:vid-id", proxyUploadHandler)   //服务转发模式之proxy模式
	//static静态文件路由
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

//api模式可以对原始数据增删改查，然后用网络请求库模拟HTTP请求
//proxy模式类似NGINX，Apache这些代理软件，只做域名转换，数据原样转发

func main() {
	log.Info("server [web] start listen 8080")
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}

