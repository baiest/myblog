package repository

type IHelloRepository interface {
	Hello() map[string]interface{}
}

type HelloRepository struct{}

func (h *HelloRepository) Hello() map[string]interface{} {
	return map[string]interface{}{
		"message": "hello",
	}
}
