package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/Yq2/video_server/scheduler/taskrunner"
	 Logger "github.com/Yq2/video_server/scheduler/logs"
)
var log = Logger.Log

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {

	log.Info("start goroutine taskrunner Start ...")
	go taskrunner.Start()
	r := RegisterHandlers()
	log.Info("server [scheduler] listen 9001")
	http.ListenAndServe(":9001", r)
}