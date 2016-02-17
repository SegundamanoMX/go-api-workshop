package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin" 
)


func main() {
    router := gin.Default()

    v1 := router.Group("/v1")
    {
        v1.GET("/hello/:name", getHello)
        v1.GET("/hello", getHello)
        v1.POST("/user/:id", postUser)
    }
    router.Run() // listen and server on 0.0.0.0:8080
}

func postUser(c *gin.Context) {
    id := c.Param("id") 
    other := c.DefaultQuery("page", "0") // shortcut for c.Request.URL.Query().Get("page")
    message := c.DefaultPostForm("message", "Nothing over here")
    
    str := fmt.Sprintf("Id: %s, Message: %s, Other: %s", id, message, other)
    c.JSON(http.StatusOK, gin.H{
        "message": str,
    })
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
