package main

import (
	"crud/controllers"
	"crud/initializers"
	"crud/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	// Posts routes batch
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	// Users routes batch
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	if err := r.Run(); err != nil {
		fmt.Println(err)
	}

}
