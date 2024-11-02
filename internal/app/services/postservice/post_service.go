package postservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rajesh6161/fast-blogger/internal/db/datastore"
	"github.com/rajesh6161/fast-blogger/internal/db/models"
)

func CreatePost(post *models.Post) error {
	post.ID = uuid.New()
	post.DateCreated = time.Now()
	post.DateUpdated = time.Now()
	datastore.Posts = append(datastore.Posts, post)
	return nil
}

func GetAllPosts() []*models.Post {
	return datastore.Posts
}

func GetPostByID(id uuid.UUID) (*models.Post, error) {
	for _, p := range datastore.Posts {
		if p.ID == id {
			return p, nil
		}
	}
	return &models.Post{}, errors.New("post not found")
}

func UpdatePost(post models.Post, id uuid.UUID) (*models.Post, error) {
	old_post, err := GetPostByID(id)
	if err != nil {
		return &models.Post{}, err
	}
	if len(post.Title) > 0 {
		old_post.Title = post.Title
	}
	if len(post.Body) > 0 {
		old_post.Body = post.Body
	}
	if len(post.ImageUrl) > 0 {
		old_post.ImageUrl = post.ImageUrl
	}
	old_post.Likes = append(old_post.Likes, post.Likes...)
	old_post.Comments = append(old_post.Comments, post.Comments...)
	old_post.DateUpdated = time.Now()

	for i, p := range datastore.Posts {
		if p.ID == id {
			datastore.Posts[i] = old_post
			return old_post, nil
		}
	}
	return &models.Post{}, errors.New("failed to update post")
}

func DeletePost(id uuid.UUID) error {
	// before deleting first check if that post exists or not
	post, err := GetPostByID(id)
	if err != nil {
		return err
	}
	i := 0
	for _, p := range datastore.Posts {
		if p.ID != post.ID {
			datastore.Posts[i] = p
			i++
		}
	}
	datastore.Posts = datastore.Posts[:i]
	return nil
}
