package book

import "github.com/hyperyuri/webapi-with-go/database/entity"

type IBookService interface {
	GetBooks() (*[]entity.Book, error)
	GetBook(bookId int64) (*entity.Book, error)
	CreateBook(book *entity.Book) (*entity.Book, error)
	DeleteBook(book *entity.Book) error
	UpdateBook(book *entity.Book) (*entity.Book, error)
}

type ServiceBook struct {
	repository IBookRepository
}

func NewBookService(repository IBookRepository) IBookService {
	return &ServiceBook{repository: repository}
}

func (s *ServiceBook) GetBooks() (*[]entity.Book, error) {
	book, err := s.repository.GetBooks()
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *ServiceBook) GetBook(bookId int64) (*entity.Book, error) {
	book, err := s.repository.GetBook(bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *ServiceBook) CreateBook(book *entity.Book) (*entity.Book, error) {
	book, err := s.repository.CreateBook(book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *ServiceBook) DeleteBook(book *entity.Book) error {
	err := s.repository.DeleteBook(book)

	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceBook) UpdateBook(book *entity.Book) (*entity.Book, error) {
	book, err := s.repository.UpdateBook(book)

	if err != nil {
		return nil, err
	}

	return book, nil
}
