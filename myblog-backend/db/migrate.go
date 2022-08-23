package db

import "github.com/baiest/myblog/myblog-backend/models"

func Migrate() {
	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.Post{})
}
