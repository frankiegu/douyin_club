package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/Yq2/douyin_club/scheduler/dbops"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		log.Info("video id should not be empty")
		sendResponse(w, 400, "video id should not be empty")
		return 
	}
	//根据视频ID 删除DB里面的视频
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		log.Info("Internal server error")
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}