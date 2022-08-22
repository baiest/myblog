package bye

import "github.com/baiest/myblog/myblog-backend/repository"

type ByeHandler struct {
	byeRepository repository.IByeRepository
}

func NewByeHandler() *ByeHandler {
	return &ByeHandler{
		byeRepository: &repository.ByeRepository{},
	}
}
