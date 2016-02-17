package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/hello/:name", getHello)
		v1.GET("/hello", getHello)
		v1.PUT("/user/:id", putUser)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/user", postUser)
	}
	router.Run() // listen and server on 0.0.0.0:8080
}

func putUser(c *gin.Context) {
	id := c.Param("id")
	other := c.DefaultQuery("page", "0") // shortcut for c.Request.URL.Query().Get("page")
	message := c.DefaultPostForm("message", "Nothing over here")

	str := fmt.Sprintf("Id: %s, Message: %s, Other: %s", id, message, other)
	c.JSON(http.StatusOK, gin.H{
		"message": str,
	})
}

func postUser(c *gin.Context) {
	var user User

	if c.BindJSON(&user) == nil {
		fmt.Println(user)
		if user.Name == "sergio" && user.Password == "pizza" {
			c.JSON(http.StatusOK, gin.H{"status": "you are italian"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}

func getHello(c *gin.Context) {
	var str string
	name := c.Param("name")
	if name != "" {
		str = fmt.Sprint("Hello, ", name)
	} else {
		str = fmt.Sprint("Hello, World!")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": str,
	})
}
