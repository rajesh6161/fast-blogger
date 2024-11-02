package postservice

import (
	"time"

	"github.com/google/uuid"
	"github.com/rajesh6161/fast-blogger/internal/db"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

func CreatePost(post *models.Post) error {
	db := db.GetDB()
	db.Create(&post)
	return nil
}

func GetAllPosts() ([]*models.Post, error) {
	db := db.GetDB()
	posts := []*models.Post{}
	res := db.Preload("Author").Preload("Likes").Preload("Comments").Find(&posts)
	if res.Error != nil {
		return posts, res.Error
	}
	return posts, nil
}

func GetPostByID(id uuid.UUID) (*models.Post, error) {
	db := db.GetDB()
	post := &models.Post{ID: id}
	res := db.Preload("Author").Preload("Likes").Preload("Comments").First(&post)
	if res.Error != nil {
		return post, res.Error
	}
	return post, nil
}

func UpdatePost(updatedPost models.Post, id uuid.UUID) (*models.Post, error) {
	db := db.GetDB()
	post, err := GetPostByID(id)
	if err != nil {
		return &models.Post{}, err
	}
	if len(updatedPost.Title) > 0 {
		post.Title = updatedPost.Title
	}
	if len(updatedPost.Body) > 0 {
		post.Body = updatedPost.Body
	}
	if len(updatedPost.ImageUrl) > 0 {
		post.ImageUrl = updatedPost.ImageUrl
	}
	post.Likes = append(post.Likes, updatedPost.Likes...)
	post.Comments = append(post.Comments, updatedPost.Comments...)
	post.DateUpdated = time.Now()

	res := db.Save(&post)
	if res.Error != nil {
		return post, res.Error
	}

	return post, nil
}

func DeletePost(id uuid.UUID) error {
	db := db.GetDB()
	// before deleting first check if that post exists or not
	post, err := GetPostByID(id)
	if err != nil {
		return err
	}

	res := db.Delete(&post)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
