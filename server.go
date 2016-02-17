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

type WorkerConf struct {
	Port string
}

func main() {
	// Run an instance though a goroutine
	conf1 := WorkerConf{":3001"}
	go mainApiWorker(conf1)

	// Run an instance in the main thread
	conf2 := WorkerConf{":3002"}
	altApiWorker(conf2)
}

func mainApiWorker(conf WorkerConf) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/hello/:name", getHello)
		v1.GET("/hello", getHello)
		v1.PUT("/user/:id", putUser)
	}
	router.Run(conf.Port) // listen and server on 0.0.0.0:3001
}

func altApiWorker(conf WorkerConf) {
	router := gin.Default()

	admin := router.Group("/admin")
	{
		admin.GET("/hello", getHelloAdmin)
		admin.POST("/user", postUser)
	}
	router.Run(conf.Port) // listen and server on 0.0.0.0:3002
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

func getHelloAdmin(c *gin.Context) {
	var str string
	str = fmt.Sprint("Hello, Admin!")

	c.JSON(http.StatusOK, gin.H{
		"message": str,
	})
}
