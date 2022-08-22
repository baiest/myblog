package user

import (
	"encoding/json"
	"net/http"

	"github.com/baiest/myblog/myblog-backend/models"
)

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var usersResponse []models.UserResponse
	users := u.userRepository.GetAll()
	for _, user := range users {
		usersResponse = append(usersResponse, models.UserResponse{
			Id:       user.Id,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		})
	}
	json.NewEncoder(w).Encode(usersResponse)
}
