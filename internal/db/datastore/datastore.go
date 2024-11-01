package datastore

import (
	"github.com/rajesh6161/fast-blogger/internal/db/models"
	"github.com/rajesh6161/fast-blogger/internal/helpers"
)

var Posts []*models.Post

func InitPosts() {
	Posts = []*models.Post{
		{ID: helpers.UUIDParser("43512caa-bf57-49c8-8dfa-cc3aa3f315b8"), Title: "The Dawn of AI", Body: "Exploring the rise of AI technologies.", Author: "Jane Doe", ImageUrl: "https://example.com/image1.jpg", DateCreated: helpers.DateParser("2024-10-01"), DateUpdated: helpers.DateParser("2024-10-01")},
		{ID: helpers.UUIDParser("248016e0-fdd4-4d90-807a-e9cf4afc738a"), Title: "Blockchain Basics", Body: "Understanding the fundamentals of blockchain.", Author: "John Smith", ImageUrl: "https://example.com/image2.jpg", DateCreated: helpers.DateParser("2024-10-02"), DateUpdated: helpers.DateParser("2024-10-01")},
		{ID: helpers.UUIDParser("7b1cfb8d-1902-4e3a-a56c-860909010656"), Title: "Climate Change Insights", Body: "Analyzing the impacts of climate change.", Author: "Emma Brown", ImageUrl: "https://example.com/image3.jpg", DateCreated: helpers.DateParser("2024-10-03"), DateUpdated: helpers.DateParser("2024-10-01")},
	}
}
