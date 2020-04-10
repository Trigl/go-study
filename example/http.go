package main

import (
	"context"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var httpClient *http.Client

func main() {
	// init http client
	tr := &http.Transport{
		MaxIdleConns:        40,
		MaxIdleConnsPerHost: 40,
	}
	httpClient = &http.Client{Transport: tr}

	// start local http server
	engine := gin.Default()
	engine.Any("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	go endless.ListenAndServe(":8080", engine)

	for {
		sendPost()
		time.Sleep(time.Millisecond * 3)
	}
}

func sendPost() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	req, err := http.NewRequest("POST", "http://localhost:8080", strings.NewReader("test"))
	if err != nil {
		fmt.Println("http request error")
	}
	req = req.WithContext(ctx)
	resp, err := httpClient.Do(req)
	cancel()

	if err != nil {
		fmt.Printf("http Post error: %v", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("fail")
	}
}
