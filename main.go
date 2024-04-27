package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcus121neo/go-crud/controllers"
	"github.com/marcus121neo/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/", controllers.PostsCreate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostByID)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
