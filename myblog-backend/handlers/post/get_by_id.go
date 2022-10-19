package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/gorilla/mux"
)

func (p *PostHandler) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.ParseUint(params["id"], 10, 64)
	var errorResponse models.ErrorResponse

	post, err := p.postRepository.GetById(uint(id))
	if err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	json.NewEncoder(w).Encode(models.PostToResponse(*post))
}
