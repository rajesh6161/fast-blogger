package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rajesh6161/fast-blogger/internal/app/services/postservice"
	"github.com/rajesh6161/fast-blogger/internal/app/validators"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

func CreatePost(c *fiber.Ctx) error {
	// Parse the JSON body into a Post struct
	post := &models.Post{}
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("Error parsing body: %s", err.Error()),
		})
	}
	// valdiating the post
	validation_err := validators.Validator(post)
	if validation_err != nil {
		return c.JSON(validation_err)
	}
	if err := postservice.CreatePost(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("Error creating post: %s", err.Error()),
		})
	}
	return c.JSON(post)
}
func GetAllPosts(c *fiber.Ctx) error {
	posts := postservice.GetAllPosts()
	return c.JSON(posts)
}

func GetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	// Parse the postID string to uuid.UUID
	postID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error parsing ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Invalid ID format",
		})
	}
	post, err := postservice.GetPostByID(postID)
	if err != nil {
		log.Printf("error getting post with ID: %v Error: %v", postID, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("error parsing ID: %v", err.Error()),
		})
	}
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	// Parse the postID string to uuid.UUID
	postID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error parsing ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Invalid ID format",
		})
	}
	err = postservice.DeletePost(postID)
	if err != nil {
		log.Printf("Error deleting: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Code":    fiber.StatusOK,
		"Message": "Post successfully deleted",
	})
}

func UpdatePost(c *fiber.Ctx) error {
	post := &models.Post{}
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("Error parsing body: %s", err.Error()),
		})
	}
	id := c.Params("id")
	// Parse the postID string to uuid.UUID
	postID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("error parsing ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Invalid ID format",
		})
	}

	updated_post, err := postservice.UpdatePost(*post, postID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("Error updating post: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(updated_post)
}
