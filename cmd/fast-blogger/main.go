package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	postHandler "github.com/rajesh6161/fast-blogger/internal/app/handlers/post"
	userHandler "github.com/rajesh6161/fast-blogger/internal/app/handlers/user"
	"github.com/rajesh6161/fast-blogger/internal/db"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "fast-blogger-db"
)

/*
RUN this query in Postgres
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();
*/

// for custom error handling
type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	// Set up the DSN for the database connection
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	// Initialize the database connection
	db := db.Initialize(dsn)
	if db == nil {
		log.Fatal("Failed to initialize database connection")
	}

	// Migrate models
	if err := db.AutoMigrate(&models.Post{}, &models.User{}, &models.Like{}, &models.Comment{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Get("/post/all", postHandler.GetAllPosts)    // GET all posts
	app.Get("/post/:id", postHandler.GetPostByID)    // GET a post by id
	app.Post("/post/create", postHandler.CreatePost) // CREATE a new post
	app.Put("/post/:id", postHandler.UpdatePost)     // UPDATE a post by id
	app.Delete("/post/:id", postHandler.DeletePost)  // DELETE a post by id

	app.Get("/user/all", userHandler.GetAllUsers)    // GET all users
	app.Get("/user/:id", userHandler.GetUserByID)    // GET an user by id
	app.Post("/user/create", userHandler.CreateUser) // CREATE an new user
	app.Put("/user/:id", userHandler.UpdateUser)     // UPDATE an user by id
	app.Delete("/user/:id", userHandler.DeleteUser)  // DELETE an user by id

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
