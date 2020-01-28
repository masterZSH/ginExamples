package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("test", "zsh")
		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		test := c.MustGet("test").(string)

		// 打印："12345"
		log.Println(test)
		c.JSON(http.StatusOK, gin.H{
			"test": test,
		})
	})
	r.Run(":8080")
}
