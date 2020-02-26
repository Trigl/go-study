package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type ProcessorEvent struct {
	receiveTime int64
	headers     map[string]string
	body        []byte
	bodyLength  int
}

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	engine.Any("/haha", WebRoot)
	engine.Any("/hehe", WebRoot)
	// 绑定端口，然后启动应用
	engine.Run(":9205")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
 */
func WebRoot(context *gin.Context) {
	fmt.Println(context.Request.RequestURI)

	event := new(ProcessorEvent)
	event.receiveTime = time.Now().UnixNano() / 1e6
	event.headers = make(map[string]string, len(context.Request.Header))
	for k, v := range context.Request.Header {
		event.headers[k] = v[0]
	}

	bodybyte, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		fmt.Println("err")
	}
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodybyte))
	event.body = bodybyte
	event.bodyLength = len(bodybyte)

	context.String(http.StatusOK, fmt.Sprintf("header=%v,body=%s", event.headers, string(bodybyte)))
}
