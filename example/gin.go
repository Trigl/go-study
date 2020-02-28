package main

import (
	"fmt"
	"github.com/fvbock/endless"
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
	engine.Any("/haha", handler)
	engine.Any("/hehe", handler)

	// 动态加载 route
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("keke")
		engine.Any("/keke", handler)
	}()
	// 绑定端口，然后启动应用
	endless.ListenAndServe(":9205", engine)

}

func handler(ctx *gin.Context) {
	event := new(ProcessorEvent)
	event.receiveTime = time.Now().UnixNano() / 1e6
	header := make(map[string]string, len(ctx.Request.Header))
	for k, v := range ctx.Request.Header {
		header[k] = v[0]
	}
	event.headers = header

	switch ctx.Request.Method {
	case "GET":
		event.body = []byte(ctx.Request.URL.RawQuery)
	case "POST":
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Println("err")
			ctx.String(http.StatusBadRequest, "400")
			return
		}
		event.body = body
	}
	event.bodyLength = len(event.body)

	fmt.Println(event)
	ctx.String(http.StatusOK, fmt.Sprintf("header=%v,body=%s", event.headers, string(event.body)))
}
