package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":1111",          // 地址端口
		Handler:        router,           // 绑定
		ReadTimeout:    10 * time.Second, // 超时时间
		WriteTimeout:   10 * time.Second, // 超时时间
		MaxHeaderBytes: 1 << 20,          // 1M
	}
	s.ListenAndServe()

}
