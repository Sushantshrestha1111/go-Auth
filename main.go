package main

import (
	"fmt"
	"go-auth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()

}

func main() {

	fmt.Println("hello3")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
