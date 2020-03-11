package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()
	// 匹配  /user/john/   /user/john/give
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.Run(":8888")
}
