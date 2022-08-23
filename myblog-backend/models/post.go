package models

import "time"

type IdPost = uint

type Post struct {
	Id      IdPost `gorm:"primarykey"`
	Title   string `gorm:"type:varchar(100);not null"`
	Content string
	IdUser  IdUser `gorm:"not null;"`
	User    User   `gorm:"foreignKey:IdUser"`
	DBModel
}
type PostResponse struct {
	Id        IdPost       `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
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

func PostToResponse(data Post) *PostResponse {
	return &PostResponse{
		Id:        data.Id,
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		User:      *UserToResponse(data.User),
	}
}
