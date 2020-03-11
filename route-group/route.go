package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()
	// /v1/login
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpointV1)
		v1.POST("/submit", submitEndpointV1)
		v1.POST("/read", readEndpointV1)
	}

	// /v2/login
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpointV2)
		v2.POST("/submit", submitEndpointV2)
		v2.POST("/read", readEndpointV2)
	}
	router.Run(":8888")
}

// v1 group
func loginEndpointV1(c *gin.Context) {
	c.String(http.StatusOK, "login v1")
}

func submitEndpointV1(c *gin.Context) {
	c.String(http.StatusOK, "submit v1")
}

func readEndpointV1(c *gin.Context) {
	c.String(http.StatusOK, "read v1")
}

// v2 group
func loginEndpointV2(c *gin.Context) {
	c.String(http.StatusOK, "login v2")
}

func submitEndpointV2(c *gin.Context) {
	c.String(http.StatusOK, "submit v2")
}

func readEndpointV2(c *gin.Context) {
	c.String(http.StatusOK, "read v2")
}
