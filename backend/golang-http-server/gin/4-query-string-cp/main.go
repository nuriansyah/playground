package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/movie", func(ctx *gin.Context) {
		genre := ctx.Query("genre")

		ctx.String(http.StatusOK, "Here result of query of movie with genre %s", genre)
	})
	router.GET("/movie/:genre/:country", func(ctx *gin.Context) {
		country := ctx.Param("country ID")
		genre := ctx.Param("action")
		ctx.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", genre, country)
	})
	/* router.GET("/movie/:genre/:country", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Here result of query of movie with genre %s in country %s", ctx.Param("genre"), ctx.Param("country"))
	})
	router.GET("/movie/:genre", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Here result of query of movie with genre %s", ctx.Param("genre"))
	}) */

	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
