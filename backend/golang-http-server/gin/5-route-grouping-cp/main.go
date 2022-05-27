package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	if _, ok := movies[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"data": movies[id],
		}) // TODO: replace this
	} else {
		c.String(http.StatusNotFound, "data not found")
	}
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	movie := router.Group("/movie")
	{
		movie.GET("/get/:id", MovieGetHandler)
		movie.GET("/list", MovieListHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
