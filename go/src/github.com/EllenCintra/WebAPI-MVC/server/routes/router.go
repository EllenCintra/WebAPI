package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/core/book"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", book.ShowBook)
			books.GET("/", book.ShowAllBooks)
			books.POST("/", book.CreateBook)
			books.PUT("/", book.EditBook)
			books.DELETE("/:id", book.DeleteBook)

		}
	}

	return router
}
