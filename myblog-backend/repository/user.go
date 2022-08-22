package repository

import (
	"github.com/baiest/myblog/myblog-backend/db"
	"github.com/baiest/myblog/myblog-backend/models"
)

type IUserRepository interface {
	Create(data models.UserCreate) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() []models.User
	// Update()
	// Delete()
}

type UserRepository struct{}

func (u UserRepository) Create(data models.UserCreate) (*models.User, error) {
	user := models.NewUser(data)
	created := db.DB.Create(&user)
	if err := created.Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (u UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	found := db.DB.First(&user, "email = ?", email)
	if err := found.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetAll() []models.User {
	var users []models.User
	db.DB.Find(&users)
	return users
}
