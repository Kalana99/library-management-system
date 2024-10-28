package services

import (
    "library-management-system/models"
    "library-management-system/repositories"
)

type BookService interface {
    CreateBook(book *models.Book) error
    GetBookByID(id int) (*models.Book, error)
    UpdateBook(book *models.Book) error
    DeleteBook(id int) error
}

type bookService struct {
    repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
    return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book *models.Book) error {
    return s.repo.Create(book)
}

func (s *bookService) GetBookByID(id int) (*models.Book, error) {
    return s.repo.GetByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
    return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id int) error {
    return s.repo.Delete(id)
}
