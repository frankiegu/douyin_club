package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	Logger "github.com/Yq2/video_server/streamserver/logs"
)

var log = Logger.Log
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", streamHandler)

	router.POST("/upload/:vid-id", uploadHandler)

	router.GET("/testpage", testPageHandler)

	return router
}

//中间件处理
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//获取一张连接票
	if !m.l.GetConn() {
		//HTTP 429 表示连接过多
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	//处理逻辑
	m.r.ServeHTTP(w, r)
	//无论是否发生panic异常，都会归还一张票
	//go对每个连接都会新开一个goroutine来处理，有几个go就会执行几次本方法，就会从票池里面拿走多少张票
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	//票池里一共有10张票
	log.Info("stream limit 10")
	mh := NewMiddleWareHandler(r, 10)
	log.Info("server [streamserver] start listen 9000")
	http.ListenAndServe(":9000", mh)
}