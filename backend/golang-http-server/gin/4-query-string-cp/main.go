package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/movie", func(c *gin.Context) {
		genre := c.Query("genre")
		if genre == "" {
			genre = "general"
		}
		countryID := c.Query("country")
		if countryID == "" {
			c.String(http.StatusOK, "Here result of query of movie with genre %v", genre)
		} else {
			c.String(http.StatusOK, "Here result of query of movie with genre %v in country %v", genre, countryID)
		}
	})
	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
