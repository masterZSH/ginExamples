package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	type LoginForm struct {
		User     string `form:"user" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		// 获取表单form
		if c.ShouldBind(&form) == nil {
			fmt.Print(form)
		}
	})

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// file upload
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})

	r.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {

			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path)

		}()
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
