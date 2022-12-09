package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/restapi-gin/controllers"
	"github.com/solnsumei/restapi-gin/models"
)

func main() {
	router := gin.Default()

	// Initialize DB
	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello Book API"})
	})

	booksRouter := router.Group("/books")
	{
		booksRouter.GET("/", controllers.FindBooks)
	}

	router.Run("localhost:3000")
}
