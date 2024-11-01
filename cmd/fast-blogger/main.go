package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rajesh6161/fast-blogger/internal/app/handlers"
	"github.com/rajesh6161/fast-blogger/internal/db/datastore"
)

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	// Initialize sample posts from the datastore
	datastore.InitPosts()
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Get("/post/all", handlers.GetAllPosts)    // GET all posts
	app.Get("/post/:id", handlers.GetPostByID)    // GET a post by id
	app.Post("/post/create", handlers.CreatePost) // CREATE a new post
	app.Put("/post/:id", handlers.UpdatePost)     // UPDATE a post by id
	app.Delete("/post/:id", handlers.DeletePost)  // DELETE a post by id

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
