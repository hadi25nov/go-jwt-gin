package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.loadEnvVariables()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.run()

}
