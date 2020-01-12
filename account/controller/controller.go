package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	// "strings"
)

// type Request struct {
// 	Method string //方法:POST,GET...
// 	URL *url.URL //URL结构体
// 	Proto      string // 协议："HTTP/1.0"
// 	ProtoMajor int    // 1
// 	ProtoMinor int    // 0
// 	Header Header    //头部信息
// 	Body io.ReadCloser //请求实体
// 	GetBody func() (io.ReadCloser, error) // Go 1.8
// 	ContentLength int64  //首部：Content-Length
// 	TransferEncoding []string
// 	Close bool           //是否已关闭
// 	Host string          //首部Host
// 	Form url.Values      //参数查询的数据
// 	PostForm url.Values // application/x-www-form-urlencoded类型的body解码后的数据
// 	MultipartForm *multipart.Form //文件上传时的数据
// 	Trailer Header
// 	RemoteAddr string          //请求地址
// 	RequestURI string          //请求的url地址
// 	TLS *tls.ConnectionState
// 	Cancel <-chan struct{} // 
// 	Response *Response //      响应数据
// }
type Controller struct {
    Data      map[interface{}]interface{}
    ChildName string
    TplNames  string
    Layout    []string
	TplExt    string
	Api map[string]func(http.ResponseWriter, *http.Request)
}

type Params struct {
	Name []string
	Password string
}
func (c *Controller) Check(w http.ResponseWriter, r *http.Request) {
	// application/json
	fmt.Println(r.Method)
	fmt.Println(r.Header.Get("Content-Type"))
	var params Params
	json.NewDecoder(r.Body).Decode(&params)
    fmt.Printf("username=%s, password=%s\n", params.Name[1], params.Password)

    fmt.Fprintf(w, `{"code":0}`)
}
