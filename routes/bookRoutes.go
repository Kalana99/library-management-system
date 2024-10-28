package routes

import (
    "library-management-system/controllers"
    "github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine, bookController *controllers.BookController) {
    bookRoutes := router.Group("/books")
    {
        bookRoutes.POST("/", bookController.AddBook)
        bookRoutes.GET("/:id", bookController.GetBook)
        bookRoutes.PUT("/:id", bookController.UpdateBook)
        bookRoutes.DELETE("/:id", bookController.DeleteBook)
    }
}
