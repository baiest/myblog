package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/gorilla/mux"
)

func (p *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	var errorResponse models.ErrorResponse

	if err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		fmt.Println(err)
		return
	}

	post, err := p.postRepository.GetById(uint(id))
	if err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = p.postRepository.Delete(post.Id)

	if err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
