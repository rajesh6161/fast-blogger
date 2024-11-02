package userservice

import (
	"github.com/google/uuid"
	"github.com/rajesh6161/fast-blogger/internal/db"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

func CreateUser(user *models.User) error {
	db := db.GetDB()
	res := db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// GetAllUsers returns all users in the data store
func GetAllUsers() ([]*models.User, error) {
	db := db.GetDB()
	users := []*models.User{}
	res := db.Find(&users)
	if res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

// GetUserByID retrieves a user by ID from the data store
func GetUserByID(id uuid.UUID) (*models.User, error) {
	db := db.GetDB()
	user := &models.User{ID: id}
	res := db.First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// UpdateUser updates an existing user's information
func UpdateUser(updatedUser models.User, id uuid.UUID) (*models.User, error) {
	db := db.GetDB()
	user, err := GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if len(updatedUser.Name) > 0 {
		user.Name = updatedUser.Name
	}
	if len(updatedUser.Email) > 0 {
		user.Email = updatedUser.Email
	}
	if len(updatedUser.Password) > 0 {
		user.Password = updatedUser.Password // Hash before storing
	}
	res := db.Save(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// DeleteUser removes a user by ID from the data store
func DeleteUser(id uuid.UUID) error {
	db := db.GetDB()
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}
	res := db.Delete(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
