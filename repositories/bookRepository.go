package repositories

import (
    "library-management-system/models"
    "gorm.io/gorm"
)

type BookRepository interface {
    Create(book *models.Book) error
    GetByID(id int) (*models.Book, error)
    Update(book *models.Book) error
    Delete(id int) error
}

type bookRepository struct {
    db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
    return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *models.Book) error {
    return r.db.Create(book).Error
}

func (r *bookRepository) GetByID(id int) (*models.Book, error) {
    var book models.Book
    if err := r.db.First(&book, id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func (r *bookRepository) Update(book *models.Book) error {
    return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id int) error {
    return r.db.Delete(&models.Book{}, id).Error
}
