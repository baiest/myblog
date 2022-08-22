package auth

import "github.com/baiest/myblog/myblog-backend/repository"

type AuthHandler struct {
	userRepository repository.IUserRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userRepository: repository.UserRepository{},
	}
}
