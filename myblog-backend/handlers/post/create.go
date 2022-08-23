package post

import (
	"encoding/json"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/baiest/myblog/myblog-backend/utils"
)

func (p *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data models.PostCreate
	var errorResponse models.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorResponse)
		return
	}

	data.IdUser = utils.GetIdUserRequest(r)
	post, err := p.postRepository.Create(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Error: "Post no created",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.PostResponse{
		Id:        post.Id,
		Title:     post.Content,
		Content:   post.Title,
		IdUser:    post.IdUser,
		CreatedAt: post.CreatedAt,
	})
}
