package post

import "github.com/baiest/myblog/myblog-backend/repository"

type PostHandler struct {
	postRepository repository.IPostRepository
}

func NewUserHandler() *PostHandler {
	return &PostHandler{
		postRepository: repository.PostRepository{},
	}
}
