package routes

import (
    "library-management-system/controllers"
    "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *controllers.UserController) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/", userController.CreateUser)
        userRoutes.GET("/:id", userController.GetUser)
        userRoutes.DELETE("/:id", userController.DeleteUser)
    }
}
