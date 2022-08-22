package auth

import (
	"encoding/json"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthHandler) SingIn(w http.ResponseWriter, r *http.Request) {
	var data models.UserCreate
	var errorResponse models.ErrorResponse
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errorResponse.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errorResponse)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Password = string(hashedPass)
	user, err := a.userRepository.Create(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Error: "User no created",
		})
		return
	}

	json.NewEncoder(w).Encode(&models.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
	})
}
