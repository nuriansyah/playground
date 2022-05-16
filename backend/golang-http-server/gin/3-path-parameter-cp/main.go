package main

import (
	"fmt"
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
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				fmt.Sprint("data not found"): err,
			})
			return
		}
		user, ok := data[userID]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				fmt.Sprint("data not found"): err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":    user.Name,
			"country": user.Country,
			"age":     user.Age,
		})
	} // TODO: replace this

}

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:user_id", ProfileHandler())
	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
