package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/core/book"
	"github.com/hyperyuri/webapi-with-go/core/user"
	"github.com/hyperyuri/webapi-with-go/database"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	bookRepository := book.NewBookRepository(database.GetDB())
	bookService := book.NewBookService(bookRepository)
	bookController := book.NewControllerBook(bookService)

	userRepository := user.NewUserRepository(database.GetDB())
	userService := user.NewServiceUser(userRepository)
	userController := user.NewControllerUser(userService)

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

		users := main.Group("users")
		{
			users.GET("/:id", userController.GetUser)
			users.GET("/", userController.GetUsers)
			users.POST("/", userController.CreateUser)
			users.PUT("/:id", userController.EditUser)
			users.DELETE("/:id", userController.DeleteUser)

		}
	}

	return router
}
