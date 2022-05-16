package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RequestMethodHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

func GetGinRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/get", RequestMethodHandler)
	router.POST("/post", RequestMethodHandler)
	router.PUT("/put", RequestMethodHandler)
	router.DELETE("/delete", RequestMethodHandler)
	router.PATCH("/patch", RequestMethodHandler)
	router.HEAD("/head", RequestMethodHandler)
	router.OPTIONS("/options", RequestMethodHandler)

	return router
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
