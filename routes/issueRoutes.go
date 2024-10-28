package routes

import (
    "library-management-system/controllers"
    "github.com/gin-gonic/gin"
)

func IssueRoutes(router *gin.Engine, issueController *controllers.IssueController) {
    issueRoutes := router.Group("/issues")
    {
        issueRoutes.POST("/", issueController.IssueBook)
        issueRoutes.PUT("/:id/return", issueController.ReturnBook)
        issueRoutes.GET("/user/:user_id", issueController.GetUserIssues)
    }
}
