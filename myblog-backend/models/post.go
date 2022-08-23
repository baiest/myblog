package models

import "time"

type IdPost = uint

type Post struct {
	Id      IdPost `gorm:"primarykey"`
	Title   string `gorm:"type:varchar(100);not null"`
	Content string
	IdUser  IdUser `gorm:"not null"`
	DBModel
}
type PostResponse struct {
	Id        IdPost    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IdUser    IdUser    `json:"id_user"`
	CreatedAt time.Time `json:"created_at"`
}
type PostCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	IdUser  IdUser `json:"id_user"`
}

func NewPost(data PostCreate) *Post {
	return &Post{
		Title:   data.Title,
		Content: data.Content,
		IdUser:  data.IdUser,
	}
}
