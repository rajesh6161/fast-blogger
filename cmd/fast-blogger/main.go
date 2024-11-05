package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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

	engine := html.New("./internal/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Static("/static", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "partials/layout")
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{}, "partials/layout")
	})
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{}, "partials/layout")
	})

	app.Get("/api/post/all", postHandler.GetAllPosts)    // GET all posts
	app.Get("/api/post/:id", postHandler.GetPostByID)    // GET a post by id
	app.Post("/api/post/create", postHandler.CreatePost) // CREATE a new post
	app.Put("/api/post/:id", postHandler.UpdatePost)     // UPDATE a post by id
	app.Delete("/api/post/:id", postHandler.DeletePost)  // DELETE a post by id

	app.Get("/api/user/all", userHandler.GetAllUsers)    // GET all users
	app.Get("/api/user/:id", userHandler.GetUserByID)    // GET an user by id
	app.Post("/api/user/create", userHandler.CreateUser) // CREATE an new user
	app.Put("/api/user/:id", userHandler.UpdateUser)     // UPDATE an user by id
	app.Delete("/api/user/:id", userHandler.DeleteUser)  // DELETE an user by id

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
