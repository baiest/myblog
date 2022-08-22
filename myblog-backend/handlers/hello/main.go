package hello

import (
	"github.com/baiest/myblog/myblog-backend/repository"
)

type HelloHandler struct {
	helloRepository repository.IHelloRepository
}

func NewByeHandler() *HelloHandler {
	return &HelloHandler{
		helloRepository: &repository.HelloRepository{},
	}
}
