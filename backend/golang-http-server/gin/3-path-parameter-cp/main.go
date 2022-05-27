package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, _ := strconv.Atoi(i)

		if _, ok := data[id]; ok {
			c.String(http.StatusOK, "Name: %v, Country: %v, Age: %v", data[id].Name, data[id].Country, data[id].Age)
		} else {
			c.String(http.StatusNotFound, "data not found")
		}
	} // TODO: replace this

}

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/profile/:id", ProfileHandler())
	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
