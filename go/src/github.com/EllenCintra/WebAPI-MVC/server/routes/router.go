package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/core/book"
	"github.com/hyperyuri/webapi-with-go/database"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	bookRepository := book.NewBookRepository(database.GetDB())
	bookService := book.NewBookService(bookRepository)
	bookController := book.NewControllerBook(bookService)

	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", bookController.ShowBook)
			books.GET("/", bookController.ShowAllBooks)
			books.POST("/", bookController.CreateBook)
			books.PUT("/:id", bookController.EditBook)
			books.DELETE("/:id", bookController.DeleteBook)

		}
	}

	return router
}
