package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/phamngocquang0072/ginpj1/initializers"
)

func init() {
	// Load .env file
	initializers.LoadEnvVariables()
}

func main() {

	//load port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":"+port)
}