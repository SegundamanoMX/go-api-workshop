package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()

    v1 := router.Group("/v1")
    {
        v1.GET("/hello", getHello)
    }
    router.Run() // listen and server on 0.0.0.0:8080
}

func getHello(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
}
