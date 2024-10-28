package controllers

import (
    "library-management-system/models"
    "library-management-system/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type UserController struct {
    service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
    return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
    var user models.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.CreateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    user, err := c.service.GetUserByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := c.service.DeleteUser(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
