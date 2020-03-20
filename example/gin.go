package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
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

	engine.Any("/log/*id", handler)
	// 动态加载 route
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("keke")
		engine.Any("/keke", handler)
	}()
	// 绑定端口，然后启动应用
	go endless.ListenAndServe(":9205", engine)

	time.Sleep(100000 * time.Millisecond)

	// 获取 routes 基本信息，如 path、method、handler
	routesInfo := engine.Routes()
	fmt.Printf("routes info: %v", routesInfo)
}

func handler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
