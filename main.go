package main

import (
    "log"
    "os"
    "library-management-system/config"
    "library-management-system/controllers"
    "library-management-system/models"
    "library-management-system/repositories"
    "library-management-system/routes"
    "library-management-system/services"
    "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	// Access environment variables
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Initialize the database connection
    config.InitDB(dbHost, dbPort, dbUser, dbPassword, dbName)

    // Automigrate models to create necessary tables
    err = config.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Issue{})
    if err != nil {
        log.Fatalf("Failed to migrate database models: %v", err)
    }
    log.Println("Database migration completed.")

    // Initialize repositories
    userRepo := repositories.NewUserRepository(config.DB)
    bookRepo := repositories.NewBookRepository(config.DB)
    issueRepo := repositories.NewIssueRepository(config.DB)

    // Initialize services
    userService := services.NewUserService(userRepo)
    bookService := services.NewBookService(bookRepo)
    issueService := services.NewIssueService(issueRepo)

    // Initialize controllers
    userController := controllers.NewUserController(userService)
    bookController := controllers.NewBookController(bookService)
    issueController := controllers.NewIssueController(issueService)

    // Set up Gin router
    router := gin.Default()

	// Set up a default route
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Set up routes
    routes.UserRoutes(router, userController)
    routes.BookRoutes(router, bookController)
    routes.IssueRoutes(router, issueController)

    // Start the server on port 8080
    log.Println("Server is running on http://localhost:8080")
    router.Run(":8080")
}
