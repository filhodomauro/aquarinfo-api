package main

import (
	"github.com/filhodomauro/aquarinfo-api/business"
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
	router.GET("business", business.List)
	router.Run()
}
