package main

import (
	"html/template"
	"net/http"
	"io"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"net/http/httputil"
	"github.com/julienschmidt/httprouter"
	 "Yq2/config"
)


type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}


func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//Request.Cookie("xxx")获取cookie
	cname, err1 := r.Cookie("username") //用户名
	sid, err2 := r.Cookie("session") //session-id

    if err1 != nil || err2 != nil {
		p := &HomePage{Name: "yangqiang"}
		//访客用户定位到首页
		t, e := template.ParseFiles("./templates/home.html")
		if e != nil {
			log.Printf("Parsing template home.html error: %s", e)
			return
		}
		//将数据渲染到模板
		t.Execute(w, p)
	    return
	}
	//登录用户
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		//重定向到/userhome页面（用户首页）
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	//获取表单元素username
	fname := r.FormValue("username")

	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}
	//渲染用户首页页面
	t, e := template.ParseFiles("./templates/userhome.html")
	if e != nil {
		log.Printf("Parsing userhome.html error: %s", e)
		return
	}

	t.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//如果请求方法不是POST，返回请求方法不正确
	log.Info("apiHandler ...")
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}
	if err := json.Unmarshal(res, apibody); err != nil {
		//body解析异常
		log.Info(err)
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	//根据apibody里面METHOD方法的不同进行相应转发
	request(apibody, w, r)
	defer r.Body.Close()
}

func proxyVideoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//代理到streamserver模块
	//proxy转发采用ELB
	u, _ := url.Parse("http://" + config.GetLbAddr() + ":9000/")
	//"net/http/httputil"提供类似NGINX功能的API
	//域名转换，数据原样转发
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func proxyUploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//代理到streamserver模块
	//proxy转发采用ELB
	u, _ := url.Parse("http://" + config.GetLbAddr() + ":9000/")
	//"net/http/httputil"提供类似NGINX功能的API
	//域名转换，数据原样转发
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

