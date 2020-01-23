package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
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

	// 日志
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.GET("/testForm", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)
			// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
		} else if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, `the body should be formB`)
		} else {
			c.String(http.StatusOK, `other`)
		}
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
