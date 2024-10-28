package controllers

import (
    "library-management-system/models"
    "library-management-system/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type BookController struct {
    BookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
    return &BookController{BookService: bookService}
}

func (bc *BookController) AddBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := bc.BookService.CreateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
        return
    }

    c.JSON(http.StatusCreated, book)
}

func (bc *BookController) GetBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    book, err := bc.BookService.GetBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, book)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book.ID = uint(id)
    if err := bc.BookService.UpdateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
        return
    }

    c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    if err := bc.BookService.DeleteBook(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
