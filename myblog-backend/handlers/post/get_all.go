package post

import (
	"encoding/json"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
)

func (p *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var postResponse = []models.PostResponse{}
	posts := p.postRepository.GetAll()
	for _, post := range posts {
		postResponse = append(postResponse, *models.PostToResponse(post))
	}
	json.NewEncoder(w).Encode(postResponse)
}
