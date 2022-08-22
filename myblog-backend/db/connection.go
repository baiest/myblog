package db

import (
	"fmt"
	"log"

	"github.com/baiest/myblog/myblog-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConnection(c models.DatabaseConfig) {
	DNS := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.User,
		c.Password,
		c.DbName,
		c.Port,
	)
	DBConnection, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Println("Database connected!")
	DB = DBConnection
}
