package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	mapping()
}

func mapping() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "fuck yeah!")
	})
	router.GET("/business", List)
	router.POST("/business", Save)
	router.Run()
}
