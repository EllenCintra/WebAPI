package book

import (
	"github.com/hyperyuri/webapi-with-go/database/entity"
	"gorm.io/gorm"
)

type IBookRepository interface {
	GetBooks() (*[]entity.Book, error)
	GetBook(bookId int64) (*entity.Book, error)
	CreateBook(book *entity.Book) (*entity.Book, error)
	DeleteBook(book *entity.Book) error
	UpdateBook(book *entity.Book) (*entity.Book, error)
}

type RepositoryBook struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &RepositoryBook{db: db}
}

func (r *RepositoryBook) GetBooks() (*[]entity.Book, error) {
	var book []entity.Book
	err := r.db.Find(&book).Error

	return &book, err
}

func (r *RepositoryBook) GetBook(bookId int64) (*entity.Book, error) {
	var book entity.Book
	err := r.db.Where("id = ?", bookId).Find(&book).Error

	return &book, err
}

func (r *RepositoryBook) CreateBook(book *entity.Book) (*entity.Book, error) {
	return book, r.db.Create(&book).Error
}

func (r *RepositoryBook) DeleteBook(book *entity.Book) error {
	err := r.db.Delete(book).Error

	return err
}

func (r *RepositoryBook) UpdateBook(book *entity.Book) (*entity.Book, error) {
	return book, r.db.Save(&book).Error
}
