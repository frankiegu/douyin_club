package main 

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"Yq2/config"
)

//服务转发客户端
var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	log.Info("request ...")
	var resp *http.Response
	var err error
	log.Info("b.Url:",b.Url)
	u, _ := url.Parse(b.Url)  //生成一个URL对象
	u.Host = config.GetLbAddr() + ":" + u.Port()
	newUrl := u.String()
	log.Info("newUrl:",newUrl)
	switch b.Method {
	case http.MethodGet:  //GET
		req, _ := http.NewRequest("GET", newUrl, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req) //执行请求
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:  //POST
	    //bytes.NewBuffer([]byte(b.ReqBody)根据请求body生成一个Buffer指针类型，Buffer指针类型已经实现了io.Reader接口
		req, _ := http.NewRequest("POST", newUrl, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req) //执行请求
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(w, resp)
	case http.MethodDelete:  //DELETE
		req, _ := http.NewRequest("DELETE", newUrl, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)  //执行请求
		if err != nil {
			log.Println(err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest) //400
		io.WriteString(w, "Bad api request")
		return
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}