package post

import (
	"encoding/json"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
)

func (u *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var postResponse = []models.PostResponse{}
	users := u.postRepository.GetAll()
	for _, post := range users {
		postResponse = append(postResponse, models.PostResponse{
			Id:        post.Id,
			Title:     post.Content,
			Content:   post.Title,
			IdUser:    post.IdUser,
			CreatedAt: post.CreatedAt,
		})
	}
	json.NewEncoder(w).Encode(postResponse)
}
