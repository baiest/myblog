package repository

import (
	"github.com/baiest/myblog/myblog-backend/db"
	"github.com/baiest/myblog/myblog-backend/models"
)

type IPostRepository interface {
	Create(data models.PostCreate) (*models.Post, error)
	GetAll() []models.Post
	// GetByEmail(email string) (*models.User, error)
	// Update()
	// Delete()
}

type PostRepository struct{}

func (p PostRepository) Create(data models.PostCreate) (*models.Post, error) {
	post := models.NewPost(data)
	created := db.DB.Create(&post)
	if err := created.Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (p PostRepository) GetAll() []models.Post {
	var posts []models.Post
	db.DB.Find(&posts)
	return posts
}
