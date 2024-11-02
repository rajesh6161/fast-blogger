package userHandler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rajesh6161/fast-blogger/internal/app/services/userservice"
	"github.com/rajesh6161/fast-blogger/internal/app/validators"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

func CreateUser(c *fiber.Ctx) error {
	// Parse the JSON body into a User struct
	userInput := &models.UserCreate{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("Error parsing body: %s", err.Error()),
		})
	}

	// Validate the user data
	validationErr := validators.Validator(userInput)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErr)
	}
	user := &models.User{
		ID:       uuid.New(),
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password, // Hash before storing
	}

	// Attempt to create the user
	if err := userservice.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    fiber.StatusInternalServerError,
			"Message": fmt.Sprintf("Error creating user: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetAllUsers retrieves all users
func GetAllUsers(c *fiber.Ctx) error {
	users, err := userservice.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    fiber.StatusInternalServerError,
			"Message": fmt.Sprintf("Error fetching users: %s", err.Error()),
		})
	}
	return c.JSON(users)
}

// GetUserByID retrieves a user by their ID
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Parse the userID string to uuid.UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Invalid ID format",
		})
	}

	// Retrieve the user by ID
	user, err := userservice.GetUserByID(userID)
	if err != nil {
		log.Printf("Error getting user with ID: %v Error: %v", userID, err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Code":    fiber.StatusNotFound,
			"Message": fmt.Sprintf("Error: %v", err.Error()),
		})
	}

	return c.JSON(user)
}

// DeleteUser deletes a user by their ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Parse the userID string to uuid.UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Invalid ID format",
		})
	}

	// Attempt to delete the user
	if err := userservice.DeleteUser(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Code":    fiber.StatusNotFound,
			"Message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"Code":    fiber.StatusOK,
		"Message": "User successfully deleted",
	})
}

// UpdateUser updates a user's information
func UpdateUser(c *fiber.Ctx) error {
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": fmt.Sprintf("Error parsing body: %s", err.Error()),
		})
	}

	id := c.Params("id")

	// Parse the userID string to uuid.UUID
	userID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Invalid ID format",
		})
	}

	// Attempt to update the user
	updatedUser, err := userservice.UpdateUser(*user, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Code":    fiber.StatusInternalServerError,
			"Message": fmt.Sprintf("Error updating user: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}
