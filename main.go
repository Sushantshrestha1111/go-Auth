package main

import (
	"fmt"
	"go-auth/controllers"
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
	r.POST("/signup", controllers.SignUp)

	r.Run() // listen and serve on 0.0.0.0:8080

}
