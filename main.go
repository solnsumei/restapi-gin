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

	// Create router group
	booksRouter := router.Group("/books")
	{
		booksRouter.GET("/", controllers.FindBooks)
		booksRouter.POST("/", controllers.CreateBook)
		booksRouter.GET("/:id", controllers.FindBook)
		booksRouter.PATCH("/:id", controllers.UpdateBook)
		booksRouter.DELETE("/:id", controllers.DeleteBook)
	}

	err := router.Run("localhost:3000")
	if err != nil {
		return
	}
}
