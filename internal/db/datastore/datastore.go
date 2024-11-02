package datastore

import (
	"time"

	"github.com/rajesh6161/fast-blogger/internal/db/models"
	"github.com/rajesh6161/fast-blogger/internal/helpers"
)

var Posts []*models.Post
var Users []*models.User

func InitPosts() {
	Posts = []*models.Post{
		{ID: helpers.UUIDParser("43512caa-bf57-49c8-8dfa-cc3aa3f315b8"), Title: "The Dawn of AI", Body: "Exploring the rise of AI technologies.", Author: Users[0], ImageUrl: "https://example.com/image1.jpg", Likes: []*models.Like{
			{ID: helpers.UUIDParser("b6574153-f940-4965-89c9-97a20b040c11"), User: Users[0]},
			{ID: helpers.UUIDParser("827d9912-a534-4b48-8b1b-a3559a59a886"), User: Users[1]},
		}, Comments: []*models.Comment{
			{ID: helpers.UUIDParser("87c4ca16-e9de-456f-b13a-e1944bcb3ec0"), Content: "This is a comment!", By: Users[2]},
		}, DateCreated: helpers.DateParser("2024-10-01"), DateUpdated: helpers.DateParser("2024-10-01")},

		{ID: helpers.UUIDParser("248016e0-fdd4-4d90-807a-e9cf4afc738a"), Title: "Blockchain Basics", Body: "Understanding the fundamentals of blockchain.", Author: Users[1], ImageUrl: "https://example.com/image2.jpg", Likes: []*models.Like{
			{ID: helpers.UUIDParser("b6574153-f940-4965-89c9-97a20b040c11"), User: Users[0]},
		}, Comments: []*models.Comment{
			{ID: helpers.UUIDParser("74904e07-2e1f-4991-8444-fe7ea13fe700"), Content: "This is a comment 2!", By: Users[0]},
		}, DateCreated: helpers.DateParser("2024-10-02"), DateUpdated: helpers.DateParser("2024-10-01")},

		{ID: helpers.UUIDParser("7b1cfb8d-1902-4e3a-a56c-860909010656"), Title: "Climate Change Insights", Body: "Analyzing the impacts of climate change.", Author: Users[2], ImageUrl: "https://example.com/image3.jpg", Likes: []*models.Like{}, Comments: []*models.Comment{}, DateCreated: helpers.DateParser("2024-10-03"), DateUpdated: helpers.DateParser("2024-10-01")},
	}
}

func InitUsers() {
	Users = []*models.User{
		{
			ID:        helpers.UUIDParser("d57cb00c-fe27-4a56-ad69-223b38e20403"),
			Name:      "Alice Johnson",
			Email:     "alice.johnson@example.com",
			Password:  "hashedpassword1",             // In a real app, store hashed passwords
			CreatedAt: time.Now().AddDate(0, 0, -10), // 10 days ago
			UpdatedAt: time.Now().AddDate(0, 0, -5),  // 5 days ago
		},
		{
			ID:        helpers.UUIDParser("901c2cc4-7bc2-4f17-b4da-cc7870fe1b62"),
			Name:      "Bob Smith",
			Email:     "bob.smith@example.com",
			Password:  "hashedpassword2",
			CreatedAt: time.Now().AddDate(0, 0, -20), // 20 days ago
			UpdatedAt: time.Now().AddDate(0, 0, -15), // 15 days ago
		},
		{
			ID:        helpers.UUIDParser("7f5552cd-0513-4014-a37f-e14d18dcfecb"),
			Name:      "Carol Williams",
			Email:     "carol.williams@example.com",
			Password:  "hashedpassword3",
			CreatedAt: time.Now().AddDate(0, 0, -30), // 30 days ago
			UpdatedAt: time.Now().AddDate(0, 0, -25), // 25 days ago
		},
	}
}
