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
		postResponse = append(postResponse, *models.PostToResponse(post))
	}
	json.NewEncoder(w).Encode(postResponse)
}
