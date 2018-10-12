package main 

//api代理模式
type ApiBody struct {
	Url string `json:"url"`  //需要访问的URL
	Method string `json:"method"`  //请求方法
	ReqBody string `json:"req_body"` //body
}

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized = Err{Error: "api not recognized, bad resquest", ErrorCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults = Err{Error: "internal service error", ErrorCode: "003"}
)
