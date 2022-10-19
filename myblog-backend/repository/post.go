package repository

import (
	"github.com/baiest/myblog/myblog-backend/db"
	"github.com/baiest/myblog/myblog-backend/models"
)

type IPostRepository interface {
	Create(data models.PostCreate) (*models.Post, error)
	GetAll() []models.Post
	GetById(id models.IdPost) (*models.Post, error)
	// Update()
	Delete(id models.IdPost) error
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
	db.DB.Preload("User").Find(&posts)
	return posts
}

func (p PostRepository) GetById(id models.IdPost) (*models.Post, error) {
	post := models.Post{
		Id: id,
	}
	found := db.DB.Preload("User").First(&post)
	if err := found.Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p PostRepository) Delete(id models.IdPost) error {
	return db.DB.Delete(models.Post{}, id).Error
}
