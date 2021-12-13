package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database/entity"
)

type ControllerBook struct {
	service IBookService
}

func NewControllerBook(service IBookService) ControllerBook {
	return ControllerBook{service: service}
}

func (s *ControllerBook) ShowAllBooks(c *gin.Context) {
	books, err := s.service.GetBooks()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, &books)
}

func (s *ControllerBook) ShowBook(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		return
	}

	book, err := s.service.GetBook(int64(newid))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find book by id: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (s *ControllerBook) CreateBook(c *gin.Context) {
	var book entity.Book
	errJson := c.ShouldBindJSON(&book)

	if errJson != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot bind JSON: " + errJson.Error(),
		})
		return
	}

	newBook, err := s.service.CreateBook(&book)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func (s *ControllerBook) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newid, integerErr := strconv.Atoi(id)

	if integerErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		return
	}

	book, getErr := s.service.GetBook(int64(newid))

	if getErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find book by id: " + getErr.Error(),
		})
		return
	}

	deleteErr := s.service.DeleteBook(book)

	if deleteErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot delete book: " + deleteErr.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *ControllerBook) EditBook(c *gin.Context) {
	id := c.Param("id")
	newid, integerErr := strconv.Atoi(id)

	if integerErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		return
	}

	book, getErr := s.service.GetBook(int64(newid))

	if getErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find book by id: " + getErr.Error(),
		})
		return
	}

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	bookSaved, err := s.service.UpdateBook(book)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot edit book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bookSaved)
}
