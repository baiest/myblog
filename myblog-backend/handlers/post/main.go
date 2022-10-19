package post

import "github.com/baiest/myblog/myblog-backend/repository"

type PostHandler struct {
	postRepository repository.IPostRepository
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		postRepository: repository.PostRepository{},
	}
}
