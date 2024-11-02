package userservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rajesh6161/fast-blogger/internal/db/datastore"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

func CreateUser(user *models.User) error {
	user.ID = uuid.New()                            // Generate a new UUID for the user
	user.CreatedAt = time.Now()                     // Set creation time
	user.UpdatedAt = user.CreatedAt                 // Set update time as the creation time
	datastore.Users = append(datastore.Users, user) // Add the user to the data store
	return nil
}

// GetAllUsers returns all users in the data store
func GetAllUsers() []*models.User {
	return datastore.Users
}

// GetUserByID retrieves a user by ID from the data store
func GetUserByID(id uuid.UUID) (*models.User, error) {
	for _, user := range datastore.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// UpdateUser updates an existing user's information
func UpdateUser(updatedUser models.User, id uuid.UUID) (*models.User, error) {
	for _, user := range datastore.Users {
		if user.ID == id {
			if len(updatedUser.Name) > 0 {
				user.Name = updatedUser.Name
			}
			if len(updatedUser.Email) > 0 {
				user.Email = updatedUser.Email
			}
			user.Password = updatedUser.Password // Store hashed password
			user.UpdatedAt = time.Now()          // Update the last updated time
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// DeleteUser removes a user by ID from the data store
func DeleteUser(id uuid.UUID) error {
	for i, user := range datastore.Users {
		if user.ID == id {
			// Remove the user from the slice
			datastore.Users = append(datastore.Users[:i], datastore.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
