package main

import (
	"github.com/gin-gonic/gin"
	"backend/controllers"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
