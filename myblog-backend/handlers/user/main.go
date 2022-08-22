package user

import "github.com/baiest/myblog/myblog-backend/repository"

type UserHandler struct {
	userRepository repository.IUserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userRepository: repository.UserRepository{},
	}
}
