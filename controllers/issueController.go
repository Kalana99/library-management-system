package controllers

import (
    "library-management-system/models"
    "library-management-system/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type IssueController struct {
    service *services.IssueService
}

func NewIssueController(service *services.IssueService) *IssueController {
    return &IssueController{service: service}
}

func (c *IssueController) IssueBook(ctx *gin.Context) {
    var issue models.Issue
    if err := ctx.ShouldBindJSON(&issue); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.IssueBook(&issue); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue book"})
        return
    }
    ctx.JSON(http.StatusOK, issue)
}

func (c *IssueController) ReturnBook(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := c.service.ReturnBook(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}

func (c *IssueController) GetUserIssues(ctx *gin.Context) {
    userID, _ := strconv.Atoi(ctx.Param("user_id"))
    issues, err := c.service.GetIssuesByUserID(uint(userID))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve issues"})
        return
    }
    ctx.JSON(http.StatusOK, issues)
}
