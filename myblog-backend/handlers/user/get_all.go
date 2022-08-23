package user

import (
	"encoding/json"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
)

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var usersResponse = []models.UserResponse{}
	users := u.userRepository.GetAll()
	for _, user := range users {
		usersResponse = append(usersResponse, *models.UserToResponse(user))
	}
	json.NewEncoder(w).Encode(usersResponse)
}
