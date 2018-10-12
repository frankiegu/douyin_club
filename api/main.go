package main

import (
	"github.com/Yq2/douyin_club/api/session"
	"github.com/julienschmidt/httprouter"
	"net/http"
	Logger "github.com/Yq2/douyin_club/api/logs"
)

var log = Logger.Log

type middleWareHandler struct {
	r *httprouter.Router
}

//http.Handler是一个net/http标准库的接口
//middleWareHandler实现了这个接口里面的唯一方法ServeHTTP
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}

//Middware中间件
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	//在这里添加验证中间件
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:username", Login)  //rest api

	router.GET("/user/:username", GetUserInfo)

	router.POST("/user/:username/videos", AddNewVideo)

	router.GET("/user/:username/videos", ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	router.POST("/videos/:vid-id/comments", PostComment)

	router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}

func Prepare() {
	log.Info("Prepare ...")
	//将db中用户session加载到全局map中
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	//mh是实现了Handler接口的结构体类型，所以可以把mh当做一个接口类型传进去
	log.Info("server [api] start listen 8081")
	http.ListenAndServe(":8081", mh)
}

//listen -->RegisterHandlers --->handler
//golang中每个handler处理的时候都会开启一个goroutine


