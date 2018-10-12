package main 

import (
	"io"
	"os"
	"net/http"
	"html/template"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"time"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//根据模板HTML文件构造一个Template对象
	t, _ := template.ParseFiles("./videos/upload.html")
	//将Template和响应流关联起来
    t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("streamHandler ...")
	extend := ".mp4"
	//阿里云OSS文件支持H5流媒体播放
	log.Infoln("vid-id ==",p.ByName("vid-id"))
	targetUrl := "http://yq2-videos.oss-cn-beijing.aliyuncs.com/videos/" + p.ByName("vid-id") + extend
	//301 重定向
	http.Redirect(w, r, targetUrl, 301)
}

//不用OSS服务
func streamHandler_(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	video ,err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "open file error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	//流媒体播放
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}



func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infoln("uploadHandler ...")
	//限定最大body大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//根据设定文件大小解析Form文件
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		//HTTP 400
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		//HTTP 500
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return 
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		//HTTP 500
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	log.Infoln("vid-id ==",fn)
	extend := ".mp4"
	//将上传的文件保存在服务器
	err = ioutil.WriteFile(VIDEO_DIR + fn + extend, data, 0666) //权限不要设置的太大
	if err != nil {
		log.Printf("Write file error: %v", err)
		//HTTP 500
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	ossfn := "videos/" + fn + extend //上传到OSS中的文件名
	path := "./videos/" + fn + extend  //本地临时存储的文件名
	bn := "yq2-videos"
	//ossfn存储对象文件名， path路径，bn代表存储桶的名字
	ret := UploadToOss(ossfn, path, bn)
	if !ret {
		//HTTP 500
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	//移除文件
	rmerr := os.Remove(path)
	if rmerr != nil {
		log.Infoln("删除文件错误:",rmerr)
	}

	w.WriteHeader(http.StatusCreated) //HTTP 201
	io.WriteString(w, "Uploaded successfully")
}
