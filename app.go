package main

import (
	"log"

	"example/auth-with-gin/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		hello := new(controllers.HelloController)
		v1.GET("/hello", hello.Default)
	}

	user := new(controllers.UserController)
	v1.POST("/signup", user.SignUp)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	router.Run(":5050")
}
